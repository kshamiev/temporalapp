package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporalapp/internal/actions"
	"temporalapp/internal/generated/temporal"
	"temporalapp/internal/services/assortment"
	"temporalapp/internal/services/payment"
	"temporalapp/internal/services/vendors"
	"temporalapp/internal/workflows/checkout"
	ch "temporalapp/internal/workflows/checkout/activities"
	"temporalapp/internal/workflows/customer"
	ca "temporalapp/internal/workflows/customer/activities"
	"temporalapp/internal/workflows/processing"
	"temporalapp/internal/workflows/processing/activities"
)

func main() {
	app, err := temporal.NewCustomerCli(
		temporal.NewCustomerCliOptions().WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
			w := worker.New(c, temporal.CustomerTaskQueue, worker.Options{})
			w.RegisterActivity(actions.HelloActivity)
			w.RegisterActivity(actions.ByeActivity)
			temporal.RegisterCustomerFlowWorkflow(w, customer.Register)
			temporal.RegisterCheckoutFlowWorkflow(w, checkout.Register)
			temporal.RegisterProcessingFlowWorkflow(w, processing.Register)
			temporal.RegisterCustomerActivities(w,
				ca.Register(
					assortment.New(),
				),
			)
			temporal.RegisterCheckoutActivities(w,
				ch.Register(
					assortment.New(),
					payment.New(),
				),
			)
			temporal.RegisterProcessingActivities(w,
				activities.Register(
					payment.New(),
					vendors.New(),
				),
			)
			return w, nil
		}),
	)
	if err != nil {
		log.Fatalf("error initializing example cli: %v", err)
	}

	// run cli
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
