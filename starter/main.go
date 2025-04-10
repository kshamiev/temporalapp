package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"

	"temporalapp/listworkflow"
)

// c.QueryWorkflow()
// c.ExecuteWorkflow()
// c.UpdateWorkflow()
// c.SignalWorkflow()

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "hello_world_workflowID",
		TaskQueue: "hello",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, listworkflow.WorkflowOne, "Temporal1")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
	err = we.Get(context.Background(), nil)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	we, err = c.ExecuteWorkflow(context.Background(), workflowOptions, listworkflow.WorkflowTwo, "Temporal2")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
	err = we.Get(context.Background(), nil)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// STEP 1 CREATE
	workflowOptions = client.StartWorkflowOptions{
		ID:        "sample/" + uuid.NewString(),
		TaskQueue: "hello",
	}
	run, err := c.ExecuteWorkflow(context.Background(), workflowOptions, listworkflow.SampleFlowWorkflowName, &listworkflow.FlowRequest{Name: "pusik"})
	if err != nil {
		log.Fatalln(err)
	}
	if run == nil {
		log.Fatalln("execute workflow returned nil run")
	}
	fmt.Println("STEP 1 CREATE", "WorkflowID", run.GetID(), "RunID", run.GetRunID())

	// STEP 2 UPDATE
	handle, err := c.UpdateWorkflow(context.Background(), client.UpdateWorkflowOptions{
		UpdateID:   "",
		WorkflowID: run.GetID(),
		RunID:      run.GetRunID(),
		UpdateName: listworkflow.UpdateProfileUpdateName,
		Args: []any{&listworkflow.Profile{
			Id:      "11111111",
			Name:    "Vasiliy Pupkin",
			Phone:   "2222222",
			Address: "Деревня гадюкино",
		}},
		WaitForStage:        client.WorkflowUpdateStageCompleted,
		FirstExecutionRunID: "",
	})
	if err != nil {
		log.Fatalln(err)
	}
	var resp listworkflow.Profile
	err = handle.Get(context.Background(), &resp)
	fmt.Println("STEP 2 UPDATE", resp)

	// STEP 3 GET
	if val, err := c.QueryWorkflow(context.Background(), run.GetID(), run.GetRunID(), listworkflow.GetProfileQueryName); err != nil {
		log.Fatalln(err)
	} else if err = val.Get(&resp); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("STEP 3 GET", resp)

	// STEP 4 DELETE
	signal := &listworkflow.Profile{
		Id:      "99999999",
		Name:    "Хаджа Нассредин",
		Phone:   "44444444444",
		Address: "Село хрюкино",
	}
	err = c.SignalWorkflow(context.Background(), run.GetID(), run.GetRunID(), listworkflow.DeleteProfileSignalName, signal)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("STEP 4 DELETE")

	log.Println("Workflows completed")
}
