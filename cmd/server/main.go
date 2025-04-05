package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"

	"temporalapp/internal/generated/server"
	"temporalapp/internal/generated/temporal"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type srv struct {
	server.CustomerServer
	tcl temporal.CustomerClient
}

func (s *srv) GetProfile(ctx context.Context, request *server.GetProfileRequest) (*temporal.Profile, error) {
	return s.tcl.GetProfile(
		ctx,
		evalCustomerWorkflowID(request),
		"",
	)
}

func (s *srv) UpdateProfile(ctx context.Context, request *server.UpdateProfileRequest) (*temporal.Profile, error) {
	return s.tcl.UpdateProfile(
		ctx,
		evalCustomerWorkflowID(request),
		"",
		&temporal.UpdateProfileRequest{
			Name: request.Name,
		})
}

func (s *srv) DeleteProfile(ctx context.Context, request *server.DeleteProfileRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{},
		s.tcl.DeleteProfile(
			ctx,
			evalCustomerWorkflowID(request),
			"",
		)
}

func (s *srv) SetAddress(ctx context.Context, request *server.SetAddressRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{},
		s.tcl.SetAddress(
			ctx,
			evalCustomerWorkflowID(request),
			"",
			&temporal.SetAddressRequest{Address: &temporal.Address{
				Title: request.Title,
				Lat:   request.Lat,
				Long:  request.Long,
			}})
}

func (s *srv) GetCart(ctx context.Context, request *server.GetCartRequest) (*temporal.Cart, error) {
	return s.tcl.GetCart(
		ctx,
		evalCustomerWorkflowID(request),
		"",
	)
}

func (s *srv) UpdateCart(ctx context.Context, request *server.UpdateCartRequest) (*temporal.Cart, error) {
	return s.tcl.UpdateCart(
		ctx,
		evalCustomerWorkflowID(request),
		"",
		&temporal.UpdateCartRequest{
			Products: request.Products,
		})
}

func (s *srv) DeleteCart(ctx context.Context, request *server.DeleteCartRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{},
		s.tcl.DeleteCart(
			ctx,
			evalCustomerWorkflowID(request),
			"",
		)
}

func (s *srv) GetOrder(ctx context.Context, request *server.GetOrderRequest) (*temporal.Order, error) {
	// TODO здесь мы будем получать конкретный заказ. Мы это рассматрим позже
	panic("implement me")
}

func (s *srv) GetOrders(ctx context.Context, request *server.GetOrdersRequest) (*server.GetOrdersResponse, error) {
	// TODO здесь мы будем получать список заказов. Мы это тоже рассматрим позже
	panic("implement me")
}

func (s *srv) Checkout(ctx context.Context, request *server.CheckoutRequest) (*temporal.Order, error) {
	return s.tcl.Checkout(
		ctx,
		evalCustomerWorkflowID(request),
		"",
		&temporal.CheckoutRequest{
			PaymentType: request.PaymentType,
		})
}

func (s *srv) NewCustomer(ctx context.Context, in *server.NewCustomerRequest) (*temporal.Profile, error) {
	run, err := s.tcl.CustomerFlowAsync(ctx, &temporal.CustomerFlowRequest{
		Name:  in.Name,
		Phone: in.Phone,
	})
	if err != nil {
		return nil, err
	}
	profile, err := s.tcl.GetProfile(ctx, run.ID(), run.RunID())
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func evalCustomerWorkflowID(msg interface{ ProtoReflect() protoreflect.Message }) string {
	workflowID, _ := expression.EvalExpression(temporal.CustomerFlowIdexpression, msg.ProtoReflect())
	return workflowID
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	tcl := temporal.NewCustomerClient(c)
	s := grpc.NewServer()
	server.RegisterCustomerServer(s, &srv{
		tcl: tcl,
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
