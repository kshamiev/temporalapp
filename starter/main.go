package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"temporalapp/listworkflow"
)

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

	log.Println("Workflows completed")
}
