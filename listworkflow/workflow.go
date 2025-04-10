package listworkflow

import (
	"fmt"
	"path"
	"runtime"

	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"
)

func FactorySample(ctx workflow.Context, in *SampleFlowWorkflowInput) (SampleFlowWorkflow, error) {
	fmt.Println("FactorySample")
	return &Sample{
		Req: &Profile{
			Id:    WorkflowID(ctx),
			Name:  in.Req.Name,
			Phone: in.Req.Phone,
		},
		DeleteProfile: in.DeleteProfile,
	}, nil
}

type Sample struct {
	Req           *Profile
	DeleteProfile *DeleteProfileSignal
}

func (w *Sample) Execute(ctx workflow.Context) error {
	fmt.Println("EXECUTE")
	w.DeleteProfile.Receive(ctx)
	return workflow.ErrCanceled
}

func (w *Sample) GetProfile() (*Profile, error) {
	return w.Req, nil
}

func (w *Sample) UpdateProfile(ctx workflow.Context, in *Profile) (*Profile, error) {
	w.Req.Name = in.Name
	w.Req.Phone = in.Phone

	// Изменение с использованием версий
	switch version(ctx, 1) {
	case workflow.DefaultVersion: // старая версия кода
	case 1: // новая версия кода
		encodedValue := workflow.SideEffect(ctx, func(ctx workflow.Context) interface{} {
			return uuid.NewString()
		})
		if err := encodedValue.Get(&w.Req.Address); err != nil {
			return nil, err
		}
	}

	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println("workflow ID:", workflow.GetInfo(ctx).WorkflowExecution.ID)
	fmt.Println("workflow instance ID:", workflow.GetInfo(ctx).WorkflowExecution.RunID)
	fmt.Println("workflow instance method ID:", workflow.GetCurrentUpdateInfo(ctx).ID)
	fmt.Println()

	return w.Req, nil
}

func version(ctx workflow.Context, newVersion int) workflow.Version {
	method := "method"
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		if fn := runtime.FuncForPC(pc); fn != nil {
			method = path.Base(fn.Name())
		}
	}
	return workflow.GetVersion(ctx, method+"-"+workflow.GetCurrentUpdateInfo(ctx).ID, workflow.DefaultVersion, workflow.Version(newVersion))
}
