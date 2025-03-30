package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporalapp/generated/temporal"
	"temporalapp/listworkflow"
)

func main() {
	// Инициализируем Worker сразу из кодгена с помощью cli
	app, err := temporal.NewCustomerCli(
		temporal.NewCustomerCliOptions().WithWorker(
			func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
				w := worker.New(c, temporal.CustomerTaskQueue, worker.Options{})
				temporal.RegisterCreateWorkflow(w, listworkflow.Register)
				return w, nil
			}),
	)
	if err != nil {
		log.Fatalf("error initializing example cli: %v", err)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
