package listworkflow

import (
	"fmt"
	"path"
	"runtime"
	"time"

	"go.temporal.io/sdk/workflow"
)

func FactorySample(ctx workflow.Context, in *SampleFlowWorkflowInput) (SampleFlowWorkflow, error) {
	fmt.Println("FactorySample", workflow.GetInfo(ctx).WorkflowExecution.ID, workflow.GetInfo(ctx).WorkflowExecution.RunID)
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
	fmt.Println("EXECUTE", workflow.GetInfo(ctx).WorkflowExecution.ID, workflow.GetInfo(ctx).WorkflowExecution.RunID)

	// Создаем логер
	logger := workflow.GetLogger(ctx)

	// Задаем стандартные настройки для activity
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{StartToCloseTimeout: time.Second})

	// Выполняем HelloActivity
	var helloResult string
	if err := workflow.ExecuteActivity(ctx, HelloActivity, "HelloActivity").Get(ctx, &helloResult); err != nil {
		logger.Error("HelloActivity failed.", "Error", err)
		return err
	}
	logger.Info("HelloActivity activity completed.", "result", helloResult)

	// Выполняем ByeActivity
	var byeResult string
	if err := workflow.ExecuteActivity(ctx, ByeActivity, "ByeActivity").Get(ctx, &byeResult); err != nil {
		logger.Error("ByeActivity failed.", "Error", err)
		return err
	}
	logger.Info("ByeActivity activity completed.", "result", byeResult)

	w.DeleteProfile.Receive(ctx)
	return workflow.ErrCanceled
}

func (w *Sample) GetProfile() (*Profile, error) {
	return w.Req, nil
}

func (w *Sample) UpdateProfile(ctx workflow.Context, in *Profile) (*Profile, error) {
	w.Req.Id = in.Id
	w.Req.Name = in.Name
	w.Req.Phone = in.Phone

	// Изменение с использованием версий
	// switch version(ctx, 1) {
	// case workflow.DefaultVersion: // старая версия кода
	// case 1: // новая версия кода
	// 	encodedValue := workflow.SideEffect(ctx, func(ctx workflow.Context) interface{} {
	// 		return uuid.NewString()
	// 	})
	// 	if err := encodedValue.Get(&w.Req.Address); err != nil {
	// 		return nil, err
	// 	}
	// }

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
