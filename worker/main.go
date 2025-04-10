package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"

	"temporalapp/listworkflow"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "hello", worker.Options{OnFatalError: func(err error) {
		log.Fatalf("workflow error: %s", err)
	}})

	w.RegisterWorkflow(listworkflow.WorkflowOne)
	w.RegisterWorkflow(listworkflow.WorkflowTwo)
	w.RegisterActivity(listworkflow.HelloActivity)
	w.RegisterActivity(listworkflow.ByeActivity)

	w.RegisterWorkflowWithOptions(listworkflow.BuildSampleFlow(listworkflow.FactorySample), workflow.RegisterOptions{Name: listworkflow.SampleFlowWorkflowName})

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
