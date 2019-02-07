// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entroq.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DepType int32

const (
	DepType_CLAIM  DepType = 0
	DepType_DELETE DepType = 1
	DepType_CHANGE DepType = 2
	DepType_DEPEND DepType = 3
)

var DepType_name = map[int32]string{
	0: "CLAIM",
	1: "DELETE",
	2: "CHANGE",
	3: "DEPEND",
}

var DepType_value = map[string]int32{
	"CLAIM":  0,
	"DELETE": 1,
	"CHANGE": 2,
	"DEPEND": 3,
}

func (x DepType) String() string {
	return proto.EnumName(DepType_name, int32(x))
}

func (DepType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{0}
}

// TaskID contains the ID and version of a task. Together these make a unique
// identifier for that task.
type TaskID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version              int32    `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskID) Reset()         { *m = TaskID{} }
func (m *TaskID) String() string { return proto.CompactTextString(m) }
func (*TaskID) ProtoMessage()    {}
func (*TaskID) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{0}
}

func (m *TaskID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskID.Unmarshal(m, b)
}
func (m *TaskID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskID.Marshal(b, m, deterministic)
}
func (m *TaskID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskID.Merge(m, src)
}
func (m *TaskID) XXX_Size() int {
	return xxx_messageInfo_TaskID.Size(m)
}
func (m *TaskID) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskID.DiscardUnknown(m)
}

var xxx_messageInfo_TaskID proto.InternalMessageInfo

func (m *TaskID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskID) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

// TaskData contains only the data portion of a task. Useful for insertion.
type TaskData struct {
	// The name of the queue for this task.
	Queue string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	// The epoch time in millis when this task becomes available.
	AtMs int64 `protobuf:"varint,2,opt,name=at_ms,json=atMs,proto3" json:"at_ms,omitempty"`
	// The task's opaque payload.
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskData) Reset()         { *m = TaskData{} }
func (m *TaskData) String() string { return proto.CompactTextString(m) }
func (*TaskData) ProtoMessage()    {}
func (*TaskData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{1}
}

func (m *TaskData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskData.Unmarshal(m, b)
}
func (m *TaskData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskData.Marshal(b, m, deterministic)
}
func (m *TaskData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskData.Merge(m, src)
}
func (m *TaskData) XXX_Size() int {
	return xxx_messageInfo_TaskData.Size(m)
}
func (m *TaskData) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskData.DiscardUnknown(m)
}

var xxx_messageInfo_TaskData proto.InternalMessageInfo

func (m *TaskData) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *TaskData) GetAtMs() int64 {
	if m != nil {
		return m.AtMs
	}
	return 0
}

func (m *TaskData) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// TaskChange identifies a task by ID and specifies the new data it should contain.
// All fields should be filled in. Empty fields result in deleting data from that field.
type TaskChange struct {
	OldId                *TaskID   `protobuf:"bytes,1,opt,name=old_id,json=oldId,proto3" json:"old_id,omitempty"`
	NewData              *TaskData `protobuf:"bytes,2,opt,name=new_data,json=newData,proto3" json:"new_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TaskChange) Reset()         { *m = TaskChange{} }
func (m *TaskChange) String() string { return proto.CompactTextString(m) }
func (*TaskChange) ProtoMessage()    {}
func (*TaskChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{2}
}

func (m *TaskChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskChange.Unmarshal(m, b)
}
func (m *TaskChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskChange.Marshal(b, m, deterministic)
}
func (m *TaskChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskChange.Merge(m, src)
}
func (m *TaskChange) XXX_Size() int {
	return xxx_messageInfo_TaskChange.Size(m)
}
func (m *TaskChange) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskChange.DiscardUnknown(m)
}

var xxx_messageInfo_TaskChange proto.InternalMessageInfo

func (m *TaskChange) GetOldId() *TaskID {
	if m != nil {
		return m.OldId
	}
	return nil
}

func (m *TaskChange) GetNewData() *TaskData {
	if m != nil {
		return m.NewData
	}
	return nil
}

// Task is a complete task object, containing IDs, data, and metadata.
type Task struct {
	// The name of the queue for this task.
	Queue   string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Version int32  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	// The epoch time in millis when this task becomes available.
	AtMs int64 `protobuf:"varint,4,opt,name=at_ms,json=atMs,proto3" json:"at_ms,omitempty"`
	// The UUID representing the claimant (owner) for this task.
	ClaimantId string `protobuf:"bytes,5,opt,name=claimant_id,json=claimantId,proto3" json:"claimant_id,omitempty"`
	// The task's opaque payload.
	Value []byte `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	// Epoch times in millis for creation and update of this task.
	CreatedMs            int64    `protobuf:"varint,7,opt,name=created_ms,json=createdMs,proto3" json:"created_ms,omitempty"`
	ModifiedMs           int64    `protobuf:"varint,8,opt,name=modified_ms,json=modifiedMs,proto3" json:"modified_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{3}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Task) GetAtMs() int64 {
	if m != nil {
		return m.AtMs
	}
	return 0
}

func (m *Task) GetClaimantId() string {
	if m != nil {
		return m.ClaimantId
	}
	return ""
}

func (m *Task) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Task) GetCreatedMs() int64 {
	if m != nil {
		return m.CreatedMs
	}
	return 0
}

func (m *Task) GetModifiedMs() int64 {
	if m != nil {
		return m.ModifiedMs
	}
	return 0
}

// QueueStats contains the name of the queue and the number of tasks within it.
type QueueStats struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NumTasks             int32    `protobuf:"varint,2,opt,name=num_tasks,json=numTasks,proto3" json:"num_tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueStats) Reset()         { *m = QueueStats{} }
func (m *QueueStats) String() string { return proto.CompactTextString(m) }
func (*QueueStats) ProtoMessage()    {}
func (*QueueStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{4}
}

func (m *QueueStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueStats.Unmarshal(m, b)
}
func (m *QueueStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueStats.Marshal(b, m, deterministic)
}
func (m *QueueStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueStats.Merge(m, src)
}
func (m *QueueStats) XXX_Size() int {
	return xxx_messageInfo_QueueStats.Size(m)
}
func (m *QueueStats) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueStats.DiscardUnknown(m)
}

var xxx_messageInfo_QueueStats proto.InternalMessageInfo

func (m *QueueStats) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QueueStats) GetNumTasks() int32 {
	if m != nil {
		return m.NumTasks
	}
	return 0
}

// ClaimRequest is sent to attempt to claim a task from a queue. The claimant ID
// should be unique to the requesting worker (e.g., if multiple workers are in
// the same process, they should all have different claimant IDs assigned).
type ClaimRequest struct {
	// The party requesting a task.
	ClaimantId string `protobuf:"bytes,1,opt,name=claimant_id,json=claimantId,proto3" json:"claimant_id,omitempty"`
	// The queue name to claim a task from.
	Queue string `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	// The duration of the claim, if/once successful.
	DurationMs           int64    `protobuf:"varint,3,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClaimRequest) Reset()         { *m = ClaimRequest{} }
func (m *ClaimRequest) String() string { return proto.CompactTextString(m) }
func (*ClaimRequest) ProtoMessage()    {}
func (*ClaimRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{5}
}

func (m *ClaimRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClaimRequest.Unmarshal(m, b)
}
func (m *ClaimRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClaimRequest.Marshal(b, m, deterministic)
}
func (m *ClaimRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRequest.Merge(m, src)
}
func (m *ClaimRequest) XXX_Size() int {
	return xxx_messageInfo_ClaimRequest.Size(m)
}
func (m *ClaimRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRequest proto.InternalMessageInfo

func (m *ClaimRequest) GetClaimantId() string {
	if m != nil {
		return m.ClaimantId
	}
	return ""
}

func (m *ClaimRequest) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *ClaimRequest) GetDurationMs() int64 {
	if m != nil {
		return m.DurationMs
	}
	return 0
}

// ClaimResponse is returned when a claim is fulfilled or becomes obviously impossible.
//
// For TryClaim calls, if no task is available, the task will be empty and the
// status code sent to the client will be OK.
//
// For Claim calls, the lack of a task is an error, usually a timeout or
// cancelation.
type ClaimResponse struct {
	Task                 *Task    `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClaimResponse) Reset()         { *m = ClaimResponse{} }
func (m *ClaimResponse) String() string { return proto.CompactTextString(m) }
func (*ClaimResponse) ProtoMessage()    {}
func (*ClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{6}
}

func (m *ClaimResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClaimResponse.Unmarshal(m, b)
}
func (m *ClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClaimResponse.Marshal(b, m, deterministic)
}
func (m *ClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimResponse.Merge(m, src)
}
func (m *ClaimResponse) XXX_Size() int {
	return xxx_messageInfo_ClaimResponse.Size(m)
}
func (m *ClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimResponse proto.InternalMessageInfo

func (m *ClaimResponse) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

// ModifyRequest sends a request to modify a set of tasks with given
// dependencies. It is performed in a transaction, in which either all
// suggested modifications succeed and all dependencies are satisfied, or
// nothing is committed at all. A failure due to dependencies (in any
// of changes, deletes, or inserts) will be permanent.
//
// All successful changes will cause the requester to become the claimant.
type ModifyRequest struct {
	ClaimantId           string        `protobuf:"bytes,1,opt,name=claimant_id,json=claimantId,proto3" json:"claimant_id,omitempty"`
	Inserts              []*TaskData   `protobuf:"bytes,2,rep,name=inserts,proto3" json:"inserts,omitempty"`
	Changes              []*TaskChange `protobuf:"bytes,3,rep,name=changes,proto3" json:"changes,omitempty"`
	Deletes              []*TaskID     `protobuf:"bytes,4,rep,name=deletes,proto3" json:"deletes,omitempty"`
	Depends              []*TaskID     `protobuf:"bytes,5,rep,name=depends,proto3" json:"depends,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ModifyRequest) Reset()         { *m = ModifyRequest{} }
func (m *ModifyRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyRequest) ProtoMessage()    {}
func (*ModifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{7}
}

func (m *ModifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyRequest.Unmarshal(m, b)
}
func (m *ModifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyRequest.Marshal(b, m, deterministic)
}
func (m *ModifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyRequest.Merge(m, src)
}
func (m *ModifyRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyRequest.Size(m)
}
func (m *ModifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyRequest proto.InternalMessageInfo

func (m *ModifyRequest) GetClaimantId() string {
	if m != nil {
		return m.ClaimantId
	}
	return ""
}

func (m *ModifyRequest) GetInserts() []*TaskData {
	if m != nil {
		return m.Inserts
	}
	return nil
}

func (m *ModifyRequest) GetChanges() []*TaskChange {
	if m != nil {
		return m.Changes
	}
	return nil
}

func (m *ModifyRequest) GetDeletes() []*TaskID {
	if m != nil {
		return m.Deletes
	}
	return nil
}

func (m *ModifyRequest) GetDepends() []*TaskID {
	if m != nil {
		return m.Depends
	}
	return nil
}

// ModifyResponse returns inserted and updated tasks when successful.
// A dependency error (which is permanent) comes through as gRPC's NOT_FOUND code.
type ModifyResponse struct {
	Inserted             []*Task  `protobuf:"bytes,1,rep,name=inserted,proto3" json:"inserted,omitempty"`
	Changed              []*Task  `protobuf:"bytes,2,rep,name=changed,proto3" json:"changed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyResponse) Reset()         { *m = ModifyResponse{} }
func (m *ModifyResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyResponse) ProtoMessage()    {}
func (*ModifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{8}
}

func (m *ModifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyResponse.Unmarshal(m, b)
}
func (m *ModifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyResponse.Marshal(b, m, deterministic)
}
func (m *ModifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyResponse.Merge(m, src)
}
func (m *ModifyResponse) XXX_Size() int {
	return xxx_messageInfo_ModifyResponse.Size(m)
}
func (m *ModifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyResponse proto.InternalMessageInfo

func (m *ModifyResponse) GetInserted() []*Task {
	if m != nil {
		return m.Inserted
	}
	return nil
}

func (m *ModifyResponse) GetChanged() []*Task {
	if m != nil {
		return m.Changed
	}
	return nil
}

// ModifyDep can be returned with a gRPC NotFound status indicating which
// dependencies failed. This is done via the gRPC error return, not directly
// in the response proto.
type ModifyDep struct {
	Type                 DepType  `protobuf:"varint,1,opt,name=type,proto3,enum=proto.DepType" json:"type,omitempty"`
	Id                   *TaskID  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyDep) Reset()         { *m = ModifyDep{} }
func (m *ModifyDep) String() string { return proto.CompactTextString(m) }
func (*ModifyDep) ProtoMessage()    {}
func (*ModifyDep) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{9}
}

func (m *ModifyDep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyDep.Unmarshal(m, b)
}
func (m *ModifyDep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyDep.Marshal(b, m, deterministic)
}
func (m *ModifyDep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyDep.Merge(m, src)
}
func (m *ModifyDep) XXX_Size() int {
	return xxx_messageInfo_ModifyDep.Size(m)
}
func (m *ModifyDep) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyDep.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyDep proto.InternalMessageInfo

func (m *ModifyDep) GetType() DepType {
	if m != nil {
		return m.Type
	}
	return DepType_CLAIM
}

func (m *ModifyDep) GetId() *TaskID {
	if m != nil {
		return m.Id
	}
	return nil
}

// TasksRequest is sent to request a complete listing of tasks for the
// given queue. If claimant_id is empty, all tasks (not just expired
// or owned tasks) are returned.
type TasksRequest struct {
	ClaimantId           string   `protobuf:"bytes,1,opt,name=claimant_id,json=claimantId,proto3" json:"claimant_id,omitempty"`
	Queue                string   `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TasksRequest) Reset()         { *m = TasksRequest{} }
func (m *TasksRequest) String() string { return proto.CompactTextString(m) }
func (*TasksRequest) ProtoMessage()    {}
func (*TasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{10}
}

func (m *TasksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TasksRequest.Unmarshal(m, b)
}
func (m *TasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TasksRequest.Marshal(b, m, deterministic)
}
func (m *TasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TasksRequest.Merge(m, src)
}
func (m *TasksRequest) XXX_Size() int {
	return xxx_messageInfo_TasksRequest.Size(m)
}
func (m *TasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TasksRequest proto.InternalMessageInfo

func (m *TasksRequest) GetClaimantId() string {
	if m != nil {
		return m.ClaimantId
	}
	return ""
}

func (m *TasksRequest) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *TasksRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// TasksReqponse contains the tasks requested.
type TasksResponse struct {
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TasksResponse) Reset()         { *m = TasksResponse{} }
func (m *TasksResponse) String() string { return proto.CompactTextString(m) }
func (*TasksResponse) ProtoMessage()    {}
func (*TasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{11}
}

func (m *TasksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TasksResponse.Unmarshal(m, b)
}
func (m *TasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TasksResponse.Marshal(b, m, deterministic)
}
func (m *TasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TasksResponse.Merge(m, src)
}
func (m *TasksResponse) XXX_Size() int {
	return xxx_messageInfo_TasksResponse.Size(m)
}
func (m *TasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TasksResponse proto.InternalMessageInfo

func (m *TasksResponse) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

// QueuesRequest is sent to request a listing of all known queues.
type QueuesRequest struct {
	MatchPrefix          []string `protobuf:"bytes,1,rep,name=match_prefix,json=matchPrefix,proto3" json:"match_prefix,omitempty"`
	MatchExact           []string `protobuf:"bytes,2,rep,name=match_exact,json=matchExact,proto3" json:"match_exact,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueuesRequest) Reset()         { *m = QueuesRequest{} }
func (m *QueuesRequest) String() string { return proto.CompactTextString(m) }
func (*QueuesRequest) ProtoMessage()    {}
func (*QueuesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{12}
}

func (m *QueuesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueuesRequest.Unmarshal(m, b)
}
func (m *QueuesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueuesRequest.Marshal(b, m, deterministic)
}
func (m *QueuesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueuesRequest.Merge(m, src)
}
func (m *QueuesRequest) XXX_Size() int {
	return xxx_messageInfo_QueuesRequest.Size(m)
}
func (m *QueuesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueuesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueuesRequest proto.InternalMessageInfo

func (m *QueuesRequest) GetMatchPrefix() []string {
	if m != nil {
		return m.MatchPrefix
	}
	return nil
}

func (m *QueuesRequest) GetMatchExact() []string {
	if m != nil {
		return m.MatchExact
	}
	return nil
}

func (m *QueuesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// QueuesResponse contains the requested list of queue statistics.
type QueuesResponse struct {
	Queues               []*QueueStats `protobuf:"bytes,1,rep,name=queues,proto3" json:"queues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *QueuesResponse) Reset()         { *m = QueuesResponse{} }
func (m *QueuesResponse) String() string { return proto.CompactTextString(m) }
func (*QueuesResponse) ProtoMessage()    {}
func (*QueuesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6477f1161d6b143, []int{13}
}

func (m *QueuesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueuesResponse.Unmarshal(m, b)
}
func (m *QueuesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueuesResponse.Marshal(b, m, deterministic)
}
func (m *QueuesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueuesResponse.Merge(m, src)
}
func (m *QueuesResponse) XXX_Size() int {
	return xxx_messageInfo_QueuesResponse.Size(m)
}
func (m *QueuesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueuesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueuesResponse proto.InternalMessageInfo

func (m *QueuesResponse) GetQueues() []*QueueStats {
	if m != nil {
		return m.Queues
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.DepType", DepType_name, DepType_value)
	proto.RegisterType((*TaskID)(nil), "proto.TaskID")
	proto.RegisterType((*TaskData)(nil), "proto.TaskData")
	proto.RegisterType((*TaskChange)(nil), "proto.TaskChange")
	proto.RegisterType((*Task)(nil), "proto.Task")
	proto.RegisterType((*QueueStats)(nil), "proto.QueueStats")
	proto.RegisterType((*ClaimRequest)(nil), "proto.ClaimRequest")
	proto.RegisterType((*ClaimResponse)(nil), "proto.ClaimResponse")
	proto.RegisterType((*ModifyRequest)(nil), "proto.ModifyRequest")
	proto.RegisterType((*ModifyResponse)(nil), "proto.ModifyResponse")
	proto.RegisterType((*ModifyDep)(nil), "proto.ModifyDep")
	proto.RegisterType((*TasksRequest)(nil), "proto.TasksRequest")
	proto.RegisterType((*TasksResponse)(nil), "proto.TasksResponse")
	proto.RegisterType((*QueuesRequest)(nil), "proto.QueuesRequest")
	proto.RegisterType((*QueuesResponse)(nil), "proto.QueuesResponse")
}

func init() { proto.RegisterFile("entroq.proto", fileDescriptor_c6477f1161d6b143) }

var fileDescriptor_c6477f1161d6b143 = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5b, 0x6b, 0xdb, 0x4c,
	0x10, 0xfd, 0x24, 0x59, 0xb2, 0x3d, 0xbe, 0x7c, 0xee, 0x26, 0x05, 0x93, 0x12, 0x92, 0x88, 0x96,
	0x5c, 0x0a, 0xa1, 0xb8, 0x04, 0x0a, 0xa5, 0x0f, 0xc1, 0x36, 0xad, 0x69, 0x1c, 0x12, 0xd5, 0x6f,
	0x85, 0xba, 0x5b, 0xef, 0xa6, 0x11, 0xb1, 0x2e, 0xd1, 0xae, 0x93, 0xf8, 0x87, 0xf4, 0x8f, 0xf5,
	0xb1, 0xbf, 0xa6, 0xec, 0xcd, 0x96, 0x12, 0x07, 0x5a, 0xfa, 0x64, 0xed, 0x99, 0x99, 0x9d, 0x73,
	0x66, 0xce, 0x1a, 0xea, 0x34, 0xe6, 0x59, 0x72, 0x7d, 0x98, 0x66, 0x09, 0x4f, 0x90, 0x2b, 0x7f,
	0xfc, 0x0e, 0x78, 0x23, 0xcc, 0xae, 0x06, 0x3d, 0xd4, 0x04, 0x3b, 0x24, 0x6d, 0x6b, 0xdb, 0xda,
	0xab, 0x06, 0x76, 0x48, 0x50, 0x1b, 0xca, 0x37, 0x34, 0x63, 0x61, 0x12, 0xb7, 0xed, 0x6d, 0x6b,
	0xcf, 0x0d, 0xcc, 0xd1, 0xff, 0x08, 0x15, 0x51, 0xd3, 0xc3, 0x1c, 0xa3, 0x75, 0x70, 0xaf, 0x67,
	0x74, 0x46, 0x75, 0xa1, 0x3a, 0xa0, 0x35, 0x70, 0x31, 0x1f, 0x47, 0x4c, 0x56, 0x3a, 0x41, 0x09,
	0xf3, 0x21, 0x13, 0xa9, 0x37, 0x78, 0x3a, 0xa3, 0x6d, 0x67, 0xdb, 0xda, 0xab, 0x07, 0xea, 0xe0,
	0x7f, 0x01, 0x10, 0x97, 0x75, 0x2f, 0x71, 0xfc, 0x9d, 0xa2, 0xe7, 0xe0, 0x25, 0x53, 0x32, 0xd6,
	0x44, 0x6a, 0x9d, 0x86, 0x62, 0x7b, 0xa8, 0x38, 0x06, 0x6e, 0x32, 0x25, 0x03, 0x82, 0x0e, 0xa0,
	0x12, 0xd3, 0xdb, 0x31, 0xc1, 0x1c, 0xcb, 0x0e, 0xb5, 0xce, 0xff, 0xb9, 0x3c, 0xc1, 0x2b, 0x28,
	0xc7, 0xf4, 0x56, 0x7c, 0xf8, 0x3f, 0x2d, 0x28, 0x09, 0xf4, 0x11, 0xa6, 0x4a, 0xb5, 0xbd, 0x4a,
	0xb5, 0x53, 0x50, 0xbd, 0xd4, 0x54, 0xca, 0x69, 0xda, 0x82, 0xda, 0x64, 0x8a, 0xc3, 0x08, 0xc7,
	0x5c, 0x90, 0x76, 0xe5, 0x3d, 0x60, 0xa0, 0x01, 0x59, 0x8a, 0xf6, 0x72, 0xa2, 0xd1, 0x26, 0xc0,
	0x24, 0xa3, 0x98, 0x53, 0x22, 0x2e, 0x2c, 0xcb, 0x0b, 0xab, 0x1a, 0x51, 0xb7, 0x46, 0x09, 0x09,
	0x2f, 0x42, 0x15, 0xaf, 0xc8, 0x38, 0x18, 0x68, 0xc8, 0xfc, 0x77, 0x00, 0xe7, 0x82, 0xfe, 0x27,
	0x8e, 0x39, 0x43, 0x08, 0x4a, 0x31, 0x8e, 0x8c, 0x30, 0xf9, 0x8d, 0x9e, 0x41, 0x35, 0x9e, 0x45,
	0x63, 0x8e, 0xd9, 0x15, 0xd3, 0xfb, 0xab, 0xc4, 0xb3, 0x48, 0x4c, 0x82, 0xf9, 0x17, 0x50, 0xef,
	0x0a, 0x8a, 0x01, 0xbd, 0x9e, 0x51, 0xc6, 0xef, 0xab, 0xb0, 0x56, 0xa9, 0x50, 0xb3, 0xb3, 0xf3,
	0xb3, 0xdb, 0x82, 0x1a, 0x99, 0x65, 0x98, 0x87, 0x49, 0x2c, 0x68, 0x3a, 0x8a, 0xa6, 0x81, 0x86,
	0xcc, 0x7f, 0x05, 0x0d, 0xdd, 0x87, 0xa5, 0x49, 0xcc, 0x44, 0x45, 0x49, 0x30, 0xd2, 0xcb, 0xad,
	0xe5, 0x96, 0x16, 0xc8, 0x80, 0xff, 0xcb, 0x82, 0xc6, 0x50, 0xe8, 0x9c, 0xff, 0x31, 0xb7, 0x7d,
	0x28, 0x87, 0x31, 0xa3, 0x19, 0x17, 0x3a, 0x9d, 0x95, 0x5e, 0xd0, 0x71, 0xf4, 0x12, 0xca, 0x13,
	0xe9, 0x33, 0x41, 0x56, 0xa4, 0x3e, 0xc9, 0xa5, 0x2a, 0x07, 0x06, 0x26, 0x03, 0xed, 0x42, 0x99,
	0xd0, 0x29, 0xe5, 0x54, 0x6c, 0xdc, 0x79, 0xe8, 0x45, 0x13, 0x55, 0x89, 0x29, 0x8d, 0x09, 0x6b,
	0xbb, 0x8f, 0x24, 0xca, 0xa8, 0xff, 0x15, 0x9a, 0x46, 0x9b, 0x9e, 0xc7, 0x2e, 0x54, 0x14, 0x37,
	0x2a, 0x94, 0x39, 0xf7, 0x67, 0xb2, 0x08, 0xa2, 0x17, 0x86, 0x39, 0xd1, 0x22, 0x0b, 0x79, 0x26,
	0xe6, 0x9f, 0x42, 0x55, 0x75, 0xe8, 0xd1, 0x14, 0xf9, 0x50, 0xe2, 0xf3, 0x54, 0xd9, 0xa2, 0xd9,
	0x69, 0xea, 0x82, 0x1e, 0x4d, 0x47, 0xf3, 0x94, 0x06, 0x32, 0x86, 0x36, 0x17, 0xf6, 0x7f, 0x40,
	0xdb, 0x0e, 0x89, 0xff, 0x19, 0xea, 0xd2, 0x31, 0xff, 0x68, 0x94, 0x75, 0x70, 0xa7, 0x61, 0x14,
	0x72, 0xfd, 0xa4, 0xd4, 0xc1, 0xef, 0x40, 0x43, 0x5f, 0xae, 0xa7, 0xb1, 0x03, 0xae, 0xf2, 0xeb,
	0x8a, 0x51, 0xa8, 0x88, 0x1f, 0x42, 0x43, 0x1a, 0x7f, 0xc1, 0x68, 0x07, 0xea, 0x11, 0xe6, 0x93,
	0xcb, 0x71, 0x9a, 0xd1, 0x8b, 0xf0, 0x4e, 0x96, 0x56, 0x83, 0x9a, 0xc4, 0xce, 0x24, 0x24, 0x5f,
	0x93, 0x4c, 0xa1, 0x77, 0x78, 0xc2, 0xe5, 0xfc, 0xaa, 0x01, 0x48, 0xa8, 0x2f, 0x90, 0x47, 0xe8,
	0xbd, 0x85, 0xa6, 0x69, 0xa5, 0xf9, 0xed, 0x83, 0x27, 0xf5, 0x18, 0x82, 0xc6, 0x3d, 0xcb, 0xa7,
	0x18, 0xe8, 0x84, 0x83, 0x37, 0x50, 0xd6, 0x83, 0x46, 0x55, 0x70, 0xbb, 0x27, 0xc7, 0x83, 0x61,
	0xeb, 0x3f, 0x04, 0xe0, 0xf5, 0xfa, 0x27, 0xfd, 0x51, 0xbf, 0x65, 0x89, 0xef, 0xee, 0x87, 0xe3,
	0xd3, 0xf7, 0xfd, 0x96, 0xad, 0xf0, 0xb3, 0xfe, 0x69, 0xaf, 0xe5, 0x74, 0x7e, 0xd8, 0xe0, 0xf5,
	0xc5, 0x1f, 0xf5, 0x39, 0x3a, 0x82, 0xca, 0x28, 0x9b, 0xcb, 0x17, 0x84, 0xd6, 0x74, 0xaf, 0xfc,
	0xbb, 0xdd, 0x58, 0x2f, 0x82, 0x9a, 0x66, 0x07, 0xdc, 0xbf, 0xae, 0x39, 0x02, 0x4f, 0x19, 0x07,
	0x99, 0x78, 0xe1, 0x15, 0x6e, 0x3c, 0xbd, 0x87, 0x2e, 0x5b, 0xc9, 0x15, 0x2e, 0x5a, 0xe5, 0xdd,
	0xb2, 0x68, 0x55, 0xdc, 0xf2, 0x11, 0x78, 0x6a, 0xae, 0x8b, 0x56, 0x85, 0x8d, 0x2e, 0x5a, 0x15,
	0x87, 0xff, 0xcd, 0x93, 0xe8, 0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x18, 0xac, 0x5a, 0x59,
	0xc6, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EntroQClient is the client API for EntroQ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EntroQClient interface {
	TryClaim(ctx context.Context, in *ClaimRequest, opts ...grpc.CallOption) (*ClaimResponse, error)
	Claim(ctx context.Context, in *ClaimRequest, opts ...grpc.CallOption) (*ClaimResponse, error)
	Modify(ctx context.Context, in *ModifyRequest, opts ...grpc.CallOption) (*ModifyResponse, error)
	Tasks(ctx context.Context, in *TasksRequest, opts ...grpc.CallOption) (*TasksResponse, error)
	Queues(ctx context.Context, in *QueuesRequest, opts ...grpc.CallOption) (*QueuesResponse, error)
}

type entroQClient struct {
	cc *grpc.ClientConn
}

func NewEntroQClient(cc *grpc.ClientConn) EntroQClient {
	return &entroQClient{cc}
}

func (c *entroQClient) TryClaim(ctx context.Context, in *ClaimRequest, opts ...grpc.CallOption) (*ClaimResponse, error) {
	out := new(ClaimResponse)
	err := c.cc.Invoke(ctx, "/proto.EntroQ/TryClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entroQClient) Claim(ctx context.Context, in *ClaimRequest, opts ...grpc.CallOption) (*ClaimResponse, error) {
	out := new(ClaimResponse)
	err := c.cc.Invoke(ctx, "/proto.EntroQ/Claim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entroQClient) Modify(ctx context.Context, in *ModifyRequest, opts ...grpc.CallOption) (*ModifyResponse, error) {
	out := new(ModifyResponse)
	err := c.cc.Invoke(ctx, "/proto.EntroQ/Modify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entroQClient) Tasks(ctx context.Context, in *TasksRequest, opts ...grpc.CallOption) (*TasksResponse, error) {
	out := new(TasksResponse)
	err := c.cc.Invoke(ctx, "/proto.EntroQ/Tasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entroQClient) Queues(ctx context.Context, in *QueuesRequest, opts ...grpc.CallOption) (*QueuesResponse, error) {
	out := new(QueuesResponse)
	err := c.cc.Invoke(ctx, "/proto.EntroQ/Queues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntroQServer is the server API for EntroQ service.
type EntroQServer interface {
	TryClaim(context.Context, *ClaimRequest) (*ClaimResponse, error)
	Claim(context.Context, *ClaimRequest) (*ClaimResponse, error)
	Modify(context.Context, *ModifyRequest) (*ModifyResponse, error)
	Tasks(context.Context, *TasksRequest) (*TasksResponse, error)
	Queues(context.Context, *QueuesRequest) (*QueuesResponse, error)
}

func RegisterEntroQServer(s *grpc.Server, srv EntroQServer) {
	s.RegisterService(&_EntroQ_serviceDesc, srv)
}

func _EntroQ_TryClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntroQServer).TryClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EntroQ/TryClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntroQServer).TryClaim(ctx, req.(*ClaimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntroQ_Claim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntroQServer).Claim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EntroQ/Claim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntroQServer).Claim(ctx, req.(*ClaimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntroQ_Modify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntroQServer).Modify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EntroQ/Modify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntroQServer).Modify(ctx, req.(*ModifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntroQ_Tasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntroQServer).Tasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EntroQ/Tasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntroQServer).Tasks(ctx, req.(*TasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntroQ_Queues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntroQServer).Queues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EntroQ/Queues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntroQServer).Queues(ctx, req.(*QueuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EntroQ_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EntroQ",
	HandlerType: (*EntroQServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TryClaim",
			Handler:    _EntroQ_TryClaim_Handler,
		},
		{
			MethodName: "Claim",
			Handler:    _EntroQ_Claim_Handler,
		},
		{
			MethodName: "Modify",
			Handler:    _EntroQ_Modify_Handler,
		},
		{
			MethodName: "Tasks",
			Handler:    _EntroQ_Tasks_Handler,
		},
		{
			MethodName: "Queues",
			Handler:    _EntroQ_Queues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entroq.proto",
}
