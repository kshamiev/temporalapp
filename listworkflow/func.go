package listworkflow

import (
	"strings"

	"go.temporal.io/sdk/workflow"
)

func WorkflowID(ctx workflow.Context) string {
	parts := strings.Split(workflow.GetInfo(ctx).WorkflowExecution.ID, "/")
	if len(parts) < 2 {
		return workflow.GetInfo(ctx).WorkflowExecution.ID
	}
	return parts[1]
}

// BuildSampleFlow converts a SampleFlow workflow struct into a valid workflow function
func BuildSampleFlow(ctor func(workflow.Context, *SampleFlowWorkflowInput) (SampleFlowWorkflow, error)) func(workflow.Context, *FlowRequest) error {
	return func(ctx workflow.Context, req *FlowRequest) error {
		input := &SampleFlowWorkflowInput{
			Req: req,
			DeleteProfile: &DeleteProfileSignal{
				Channel: workflow.GetSignalChannel(ctx, DeleteProfileSignalName),
			},
		}
		wf, err := ctor(ctx, input)
		if err != nil {
			return err
		}
		if err := workflow.SetQueryHandler(ctx, GetProfileQueryName, wf.GetProfile); err != nil {
			return err
		}
		{
			opts := workflow.UpdateHandlerOptions{}
			if err := workflow.SetUpdateHandlerWithOptions(ctx, UpdateProfileUpdateName, wf.UpdateProfile, opts); err != nil {
				return err
			}
		}
		return wf.Execute(ctx)
	}
}
