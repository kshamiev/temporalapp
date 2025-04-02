package main

import (
	"log"

	"gitlab.tn.ru/golang/app"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

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

	// err = w.Run(worker.InterruptCh())
	err = w.Start()
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
	app.Lock(nil)
}
