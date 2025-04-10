package listworkflow

import (
	"time"

	"github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	"go.temporal.io/sdk/workflow"
)

// Это основной workflow
//
// workflow details: (name: "pbworkflow.Sample.SampleFlow", id: "sample/${! sampleId.or(id.or(uuid_v4())) }")
type SampleFlowWorkflow interface {
	// Execute defines the entrypoint to a(n) pbworkflow.Sample.SampleFlow workflow
	Execute(ctx workflow.Context) error

	// Получение профиля из запущенного workflow
	// https://docs.temporal.io/encyclopedia/workflow-message-passing#writing-query-handlers
	GetProfile() (*Profile, error)

	// Обновление профиля в запущенном workflow
	// https://docs.temporal.io/encyclopedia/workflow-message-passing#writing-query-handlers
	UpdateProfile(workflow.Context, *Profile) (*Profile, error)
}

//

type SampleFlowWorkflowInput struct {
	Req           *FlowRequest
	DeleteProfile *DeleteProfileSignal
}

type FlowRequest struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
}

type FlowResponse struct {
	Id    int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Price float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
}

type Profile struct {
	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone   string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

// ////

// SampleTaskQueue is the default task-queue for a pbworkflow.Sample worker
const SampleTaskQueue = "general"

// pbworkflow.Sample workflow names
const (
	SampleFlowWorkflowName = "pbworkflow.Sample.SampleFlow"
)

// pbworkflow.Sample workflow id expressions
var (
	SampleFlowIdexpression = expression.MustParseExpression("sample/${! sampleId.or(id.or(uuid_v4())) }")
)

// pbworkflow.Sample query names
const (
	GetProfileQueryName = "pbworkflow.Sample.GetProfile"
)

// pbworkflow.Sample signal names
const (
	DeleteProfileSignalName = "pbworkflow.Sample.DeleteProfile"
)

// pbworkflow.Sample update names
const (
	UpdateProfileUpdateName = "pbworkflow.Sample.UpdateProfile"
)

// ////

// DeleteProfileSignal describes a(n) pbworkflow.Sample.DeleteProfile signal
type DeleteProfileSignal struct {
	Channel workflow.ReceiveChannel
}

// NewDeleteProfileSignal initializes a new pbworkflow.Sample.DeleteProfile signal wrapper
func NewDeleteProfileSignal(ctx workflow.Context) *DeleteProfileSignal {
	return &DeleteProfileSignal{Channel: workflow.GetSignalChannel(ctx, DeleteProfileSignalName)}
}

// Receive blocks until a(n) pbworkflow.Sample.DeleteProfile signal is received
func (s *DeleteProfileSignal) Receive(ctx workflow.Context) (*Profile, bool) {
	var resp Profile
	more := s.Channel.Receive(ctx, &resp)
	return &resp, more
}

// ReceiveAsync checks for a pbworkflow.Sample.DeleteProfile signal without blocking
func (s *DeleteProfileSignal) ReceiveAsync() *Profile {
	var resp Profile
	if ok := s.Channel.ReceiveAsync(&resp); !ok {
		return nil
	}
	return &resp
}

// ReceiveWithTimeout blocks until a(n) pbworkflow.Sample.DeleteProfile signal is received or timeout expires.
// Returns more value of false when Channel is closed.
// Returns ok value of false when no value was found in the channel for the duration of timeout or the ctx was canceled.
// resp will be nil if ok is false.
func (s *DeleteProfileSignal) ReceiveWithTimeout(ctx workflow.Context, timeout time.Duration) (resp *Profile, ok bool, more bool) {
	resp = &Profile{}
	if ok, more = s.Channel.ReceiveWithTimeout(ctx, timeout, &resp); !ok {
		return nil, false, more
	}
	return
}

// Select checks for a(n) pbworkflow.Sample.DeleteProfile signal without blocking
func (s *DeleteProfileSignal) Select(sel workflow.Selector, fn func(*Profile)) workflow.Selector {
	return sel.AddReceive(s.Channel, func(workflow.ReceiveChannel, bool) {
		req := s.ReceiveAsync()
		if fn != nil {
			fn(req)
		}
	})
}
