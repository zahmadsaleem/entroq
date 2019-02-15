// Package qtest contains standard testing routines for exercising various backends in similar ways.
package qtest

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/shiblon/entroq"
	grpcbackend "github.com/shiblon/entroq/grpc"
	"github.com/shiblon/entroq/qsvc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/shiblon/entroq/proto"
	hpb "google.golang.org/grpc/health/grpc_health_v1"
)

const bufSize = 1 << 20

// Dialer returns a net connection.
type Dialer func() (net.Conn, error)

// ClientService starts an in-memory gRPC network service via StartService,
// then creates an EntroQ client that connects to it. It returns the client and
// a function that can be deferred for cleanup.
//
// The opener is used by the service to connect to storage. The client always
// uses a grpc opener.
func ClientService(ctx context.Context, opener entroq.BackendOpener) (client *entroq.EntroQ, stop func(), err error) {
	s, dial, err := StartService(ctx, opener)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		if err != nil {
			s.Stop()
		}
	}()

	client, err = entroq.New(ctx, grpcbackend.Opener("bufnet",
		grpcbackend.WithNiladicDialer(dial),
		grpcbackend.WithInsecure()))
	if err != nil {
		return nil, nil, fmt.Errorf("start client on in-memory service: %v", err)
	}

	return client, func() {
		client.Close()
		s.Stop()
	}, nil
}

// StartService starts an in-memory gRPC network service and returns a function for creating client connections to it.
func StartService(ctx context.Context, opener entroq.BackendOpener) (*grpc.Server, Dialer, error) {
	lis := bufconn.Listen(bufSize)
	svc, err := qsvc.New(ctx, opener)
	if err != nil {
		return nil, nil, fmt.Errorf("start service: %v", err)
	}
	s := grpc.NewServer()
	hpb.RegisterHealthServer(s, health.NewServer())
	pb.RegisterEntroQServer(s, svc)
	go s.Serve(lis)

	return s, lis.Dial, nil
}

// SimpleWorker tests basic worker functionality while tasks are coming in and
// being waited on.
func SimpleWorker(ctx context.Context, t *testing.T, client *entroq.EntroQ, qPrefix string) {
	t.Helper()

	queue := path.Join(qPrefix, "simple_worker")

	w := client.NewWorker(queue)

	var consumed []*entroq.Task
	ctx, cancel := context.WithCancel(ctx)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		for w.Next(ctx) {
			if _, _, err := w.DoModify(func(ctx context.Context) ([]entroq.ModifyArg, error) {
				consumed = append(consumed, w.Task())
				return []entroq.ModifyArg{w.Task().AsDeletion()}, nil
			}); err != nil {
				return fmt.Errorf("worker loop error: %v", err)
			}
		}
		return w.Err()
	})

	select {
	case <-time.After(2 * time.Second):
	case <-ctx.Done():
		t.Fatalf("Sleep: %v", ctx.Err())
	}

	var inserted []*entroq.Task
	for i := 0; i < 10; i++ {
		ins, _, err := client.Modify(ctx, entroq.InsertingInto(queue))
		if err != nil {
			t.Fatalf("Failed to insert task: %v", err)
		}
		inserted = append(inserted, ins...)
		select {
		case <-time.After(1 * time.Second):
		case <-ctx.Done():
			t.Fatalf("Canceled while inserting: %v", ctx.Err())
		}
	}

	for {
		empty, err := client.QueuesEmpty(ctx, entroq.MatchExact(queue))
		if err != nil {
			t.Fatalf("Error checking for empty queue: %v", err)
		}
		if empty {
			break
		}
		select {
		case <-ctx.Done():
			t.Fatalf("Context error waiting for queues to empty: %v", err)
		case <-time.After(2 * time.Second):
		}
	}

	cancel()
	if err := g.Wait(); err != nil && status.Code(err) != codes.Canceled {
		t.Fatalf("Worker exit error: %v", err)
	}

	if diff := EqualAllTasksVersionIncr(inserted, consumed, 1); diff != "" {
		t.Errorf("Tasks inserted not the same as tasks consumed:\n%v", diff)
	}
}

// SimpleSequence tests some basic functionality of a task manager, over gRPC.
func SimpleSequence(ctx context.Context, t *testing.T, client *entroq.EntroQ, qPrefix string) {
	t.Helper()

	now := time.Now()

	sleep := func(d time.Duration) {
		select {
		case <-time.After(d):
		case <-ctx.Done():
			t.Fatalf("Context canceled: %v", ctx.Err())
		}
	}

	queue := path.Join(qPrefix, "simple_sequence")

	// Claim from empty queue.
	task, err := client.TryClaim(ctx, queue, 100*time.Millisecond)
	if err != nil {
		t.Fatalf("Got unexpected error for claiming from an empty queue: %v", err)
	}
	if task != nil {
		t.Fatalf("Got unexpected non-nil claim response from empty queue:\n%s", task)
	}

	insWant := []*entroq.Task{
		{
			Queue:    queue,
			At:       now,
			Value:    []byte("hello"),
			Claimant: client.ID(),
		},
		{
			Queue:    queue,
			At:       now.Add(100 * time.Millisecond),
			Value:    []byte("there"),
			Claimant: client.ID(),
		},
	}
	var insData []*entroq.TaskData
	for _, task := range insWant {
		insData = append(insData, task.Data())
	}

	inserted, changed, err := client.Modify(ctx, entroq.Inserting(insData...))
	if err != nil {
		t.Fatalf("Got unexpected error inserting two tasks: %v", err)
	}
	if changed != nil {
		t.Fatalf("Got unexpected changes during insertion: %v", err)
	}
	if diff := EqualAllTasks(insWant, inserted); diff != "" {
		t.Fatalf("Modify tasks unexpected result, ignoring ID and time fields (-want +got):\n%v", diff)
	}
	// Also check that their arrival times are 100 ms apart as expected:
	if diff := inserted[1].At.Sub(inserted[0].At); diff != 100*time.Millisecond {
		t.Fatalf("Wanted At difference to be %v, got %v", 100*time.Millisecond, diff)
	}

	// Get queues.
	queuesWant := map[string]int{queue: 2}
	queuesGot, err := client.Queues(ctx)
	if err != nil {
		t.Fatalf("Getting queues failed: %v", err)
	}
	if diff := cmp.Diff(queuesWant, queuesGot); diff != "" {
		t.Fatalf("Queues (-want +got):\n%v", diff)
	}

	// Get all tasks.
	tasksGot, err := client.Tasks(ctx, queue)
	if err != nil {
		t.Fatalf("Tasks call failed after insertions: %v", err)
	}
	if diff := EqualAllTasks(insWant, tasksGot); diff != "" {
		t.Fatalf("Tasks unexpected return, ignoring ID and time fields (-want +got):\n%v", diff)
	}

	// Claim ready task.
	claimCtx, _ := context.WithTimeout(ctx, 5*time.Second)
	claimed, err := client.Claim(claimCtx, queue, 10*time.Second)

	if err != nil {
		t.Fatalf("Got unexpected error for claiming from a queue with one ready task: %v", err)
	}
	if claimed == nil {
		t.Fatalf("Unexpected nil result from blocking Claim")
	}
	if diff := EqualTasksVersionIncr(insWant[0], claimed, 1); diff != "" {
		t.Fatalf("Claim tasks differ, ignoring ID and times:\n%v", diff)
	}
	if got, lower, upper := claimed.At, now.Add(9*time.Second), now.Add(11*time.Second); got.Before(lower) || got.After(upper) {
		t.Fatalf("Claimed arrival time not in time bounds [%v, %v]: %v", lower, upper, claimed.At)
	}

	// TryClaim not ready task.
	tryclaimed, err := client.TryClaim(ctx, queue, 10*time.Second)
	if err != nil {
		t.Fatalf("Got unexpected error for claiming from a queue with no ready tasks: %v", err)
	}
	if tryclaimed != nil {
		t.Fatalf("Got unexpected non-nil claim response from a queue with no ready tasks:\n%s", tryclaimed)
	}

	// Make sure the next claim will work.
	sleep(100 * time.Millisecond)
	tryclaimed, err = client.TryClaim(ctx, queue, 5*time.Second)
	if err != nil {
		t.Fatalf("Got unexpected error for claiming from a queue with one ready task: %v", err)
	}
	if diff := EqualTasksVersionIncr(insWant[1], tryclaimed, 1); diff != "" {
		t.Fatalf("TryClaim got unexpected task, ignoring ID and time fields (-want +got):\n%v", diff)
	}
	if got, lower, upper := tryclaimed.At, time.Now().Add(4*time.Second), time.Now().Add(6*time.Second); got.Before(lower) || got.After(upper) {
		t.Fatalf("TryClaimed arrival time not in time bounds [%v, %v]: %v", lower, upper, tryclaimed.At)
	}
}

// QueueMatch tests various queue matching functions against a client.
func QueueMatch(ctx context.Context, t *testing.T, client *entroq.EntroQ, qPrefix string) {
	t.Helper()

	queue1 := path.Join(qPrefix, "queue-1")
	queue2 := path.Join(qPrefix, "queue-2")
	queue3 := path.Join(qPrefix, "queue-3")
	quirkyQueue := path.Join(qPrefix, "quirky=queue")

	wantQueues := map[string]int{
		queue1:      1,
		queue2:      2,
		queue3:      3,
		quirkyQueue: 1,
	}

	// Add tasks so that queues have a certain number of things in them, as above.
	var toInsert []entroq.ModifyArg
	for q, n := range wantQueues {
		for i := 0; i < n; i++ {
			toInsert = append(toInsert, entroq.InsertingInto(q))
		}
	}
	inserted, _, err := client.Modify(ctx, toInsert...)
	if err != nil {
		t.Fatalf("in QueueMatch - inserting empty tasks: %v", err)
	}

	// Check that we got everything inserted.
	if want, got := len(inserted), len(toInsert); want != got {
		t.Fatalf("in QueueMatch - want %d inserted, got %d", want, got)
	}

	// Check that we can get exact numbers for all of the above using MatchExact.
	for q, n := range wantQueues {
		qs, err := client.Queues(ctx, entroq.MatchExact(q))
		if err != nil {
			t.Fatalf("QueueMatch single - getting queue: %v", err)
		}
		if len(qs) != 1 {
			t.Errorf("QueueMatch single - expected 1 entry, got %d", len(qs))
		}
		if want, got := n, qs[q]; want != got {
			t.Errorf("QueueMatch single - expected %d values in queue %q, got %d", want, q, got)
		}
	}

	// Check that passing multiple exact matches works properly.
	multiExactCases := []struct {
		q1 string
		q2 string
	}{
		{queue1, queue2},
		{queue1, queue3},
		{quirkyQueue, queue2},
		{"bogus", queue3},
	}

	for _, c := range multiExactCases {
		qs, err := client.Queues(ctx, entroq.MatchExact(c.q1), entroq.MatchExact(c.q2))
		if err != nil {
			t.Fatalf("QueueMatch multi - getting multiple queues: %v", err)
		}
		if len(qs) > 2 {
			t.Errorf("QueueMatch multi - expected no more than 2 entries, got %d", len(qs))
		}
		want1, want2 := wantQueues[c.q1], wantQueues[c.q2]
		if got1, got2 := qs[c.q1], qs[c.q2]; want1 != got1 || want2 != got2 {
			t.Errorf("QueueMatch multi - wanted %q:%d, %q:%d, got %q:%d, %q:%d", c.q1, want1, c.q2, want2, c.q1, got1, c.q2, got2)
		}
	}

	// Check prefix matching.
	prefixCases := []struct {
		prefix string
		qn     int
		n      int
	}{
		{path.Join(qPrefix, "queue-"), 3, 6},
		{path.Join(qPrefix, "qu"), 4, 7},
		{path.Join(qPrefix, "qui"), 1, 1},
	}

	for _, c := range prefixCases {
		qs, err := client.Queues(ctx, entroq.MatchPrefix(c.prefix))
		if err != nil {
			t.Fatalf("QueueMatch prefix - queues error: %v", err)
		}
		if want, got := c.qn, len(qs); want != got {
			t.Errorf("QueueMatch prefix - want %d queues, got %d", want, got)
		}
		tot := 0
		for _, n := range qs {
			tot += n
		}
		if want, got := c.n, tot; want != got {
			t.Errorf("QueueMatch prefix - want %d total items, got %d", want, got)
		}
	}
}

// EqualAllTasks returns a string diff if any of the tasks in the lists are unequal.
func EqualAllTasks(want, got []*entroq.Task) string {
	if len(want) != len(got) {
		return cmp.Diff(want, got)
	}
	if len(want) == 0 {
		return ""
	}
	var diffs []string
	for i, w := range want {
		g := got[i]
		if w.Queue != g.Queue || w.Claimant != g.Claimant || !bytes.Equal(w.Value, g.Value) {
			diffs = append(diffs, cmp.Diff(w, g))
		}
	}
	if len(diffs) != 0 {
		return strings.Join(diffs, "\n")
	}
	return ""
}

// EqualAllTasksVersionIncr returns a non-empty diff if any of the tasks are
// unequal, taking a version increment into account for the 'got' tasks.
func EqualAllTasksVersionIncr(want, got []*entroq.Task, versionBump int) string {
	if diff := EqualAllTasks(want, got); diff != "" {
		return diff
	}
	var diffs []string
	for i, w := range want {
		g := got[i]
		if w.Version+int32(versionBump) != g.Version {
			diffs = append(diffs, cmp.Diff(g, w))
		}
	}
	if len(diffs) != 0 {
		return strings.Join(diffs, "\n")
	}
	return ""
}

// Checks for task equality, ignoring Version and all time fields.
func EqualTasks(want, got *entroq.Task) string {
	return EqualAllTasks([]*entroq.Task{want}, []*entroq.Task{got})
}

// EqualTasksVersionIncr checks for equality, allowing a version increment.
func EqualTasksVersionIncr(want, got *entroq.Task, versionBump int) string {
	return EqualAllTasksVersionIncr([]*entroq.Task{want}, []*entroq.Task{got}, versionBump)
}
