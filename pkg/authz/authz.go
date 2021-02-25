// Package authz contains standard data structures for representing permissions and authorization requests / responses.
package authz // import "entrogo.com/entroq/pkg/authz"

import (
	"context"
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Authorizer is an abstraction over Rego policy. Provide one of these to
// manage policy files and changes. The query is expected to return a nil error
// when authorized, and a non-nil error when not authorized (or smoething else
// goes wrong). If the non-nil error is an AuthzError, it can be unpacked for
// information about which queues and actions were disallowed.
type Authorizer interface {
	// Authorize sends a request with context to see if something is allowed.
	// The request contains information about the entity making the request and
	// a nil error indicates that the request is permitted. A non-nil error can
	// be returned for system errors, or for permission denied reasons. The
	// latter will be of type AuthzError and can be detected in standard
	// errors.Is/As ways..
	Authorize(context.Context, *Request) error

	// Close cleans up any resources, or policy watchdogs, that the authorizer
	// might need in order to do its work.
	Close() error
}

// Action is an authorization-style action that can be requested (and
// allowed or denied) for a particular queue spec.
type Action string

const (
	Claim  Action = "CLAIM"
	Delete Action = "DELETE"
	Change Action = "CHANGE"
	Insert Action = "INSERT"
	Read   Action = "READ"
	All    Action = "*"
)

// Request conatins an authorization request to send to OPA.
type Request struct {
	// Authz contains information that came in with the request (headers).
	Authz *Authorization `json:"authz"`
	// Queues contains information about what is desired: what queues to
	// operate on, and what should be done to them.
	Queues []*Queue `json:"queues"`
}

// NewYAMLRequest creates a request from YAML/JSON.
func NewYAMLRequest(y string) (*Request, error) {
	req := new(Request)
	if err := yaml.Unmarshal([]byte(y), req); err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	return req, nil
}

// AUthorization represents per-request authz information. It can ideally come in many
// forms. The first supported form is a "token", such as from an Authorization
// header.
type Authorization struct {
	// An HTTP Authorization header is split into its type and credentials and
	// included here when available.
	Type        string `json:"type"`
	Credentials string `json:"credentials"`

	// Never use this in practice. This allows the user to be set directly for testing.
	// The code checks for this and creates an error if present, unless
	// specifically allowed for testing.
	TestUser string `json:"testuser"`
}

// NewHeaderAuthorization creates an authorization structure from a header value.
func NewHeaderAuthorization(val string) *Authorization {
	pieces := strings.SplitN(val, " ", 2)
	az := new(Authorization)
	switch len(pieces) {
	case 1:
		az.Credentials = pieces[0]
	case 2:
		az.Type, az.Credentials = pieces[0], pieces[1]
	}
	return az
}

// String returns the authoriation header value.
func (a *Authorization) String() string {
	return strings.TrimSpace(strings.Join([]string{a.Type, a.Credentials}, " "))
}

// Queue contains information about a single queue (it is expected that
// only one match string will be specified. Behavior of multiple specifications
// is not necessarily well defined, and depends on policy execution order.
type Queue struct {
	// An exact name to match.
	Exact string `yaml: ",omitempty" json:"exact,omitempty"`
	// The kind of matching to do (default exact)
	Prefix string `yaml: ",omitempty" json:"prefix,omitempty"`
	// Actions contains the desired things to be done with this queue.
	Actions []Action `yaml:",flow" json:"actions"`
}

// AuthzError contains the reply from OPA, if non-empty. An empty UnmatchedQueues field implies
// that the action is allowed.
type AuthzError struct {
	Allow bool `json:"allow"`

	User string `json:"user"`
	// Failed contains the queue information for things that were not
	// found to be allowed by the policy. It will only contain the actions that
	// were not matched. If multiple actions were desired for a single queue,
	// only those disallowed are expected to be given back in the response.
	Failed []*Queue `json:"failed"`

	// For other kinds of errors.
	Err string `json:"err"`
}

// Success returns whether this error represents success (instead of an auth error).
func (e *AuthzError) Success() bool {
	return e == nil || (e.Allow && e.Err == "" && len(e.Failed) == 0)
}

// Error satisfies the error interface, producing a string error that contains
// unmatched queue/action information.
func (e *AuthzError) Error() string {
	y, err := yaml.Marshal(e.Failed)
	if err != nil {
		return fmt.Sprintf("user %q not authorized, failed to get data with reasons: %v", e.User, err)
	}
	if e.Err != "" {
		return fmt.Sprintf("auth error: %v", err)
	}
	return fmt.Sprintf("user %q not authorized, missing queue/actions:\n%s", e.User, string(y))
}
