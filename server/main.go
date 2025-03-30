package main

import (
	"context"
	"log"
	"net"

	"github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"temporalapp/generated/server"
	"temporalapp/generated/temporal"
)

func main() {
	// Запускаем простейший GRPC-сервер
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	tcl := temporal.NewCustomerClient(c)
	s := grpc.NewServer()
	server.RegisterCustomerServer(s, &srv{tcl: tcl})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type srv struct {
	server.CustomerServer
	tcl temporal.CustomerClient
}

func (s *srv) Create(ctx context.Context, in *temporal.CreateRequest) (*temporal.Profile, error) {
	run, err := s.tcl.CreateAsync(ctx, in)
	if err != nil {
		return nil, err
	}
	order, err := s.tcl.Read(ctx, run.ID(), run.RunID())
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *srv) Read(ctx context.Context, in *server.ReadRequest) (*temporal.Profile, error) {
	id, err := expression.EvalExpression(temporal.CreateIdexpression, in.ProtoReflect())
	if err != nil {
		return nil, err
	}
	read, err := s.tcl.Read(ctx, id, "")
	if err != nil {
		return nil, err
	}
	return read, nil
}

func (s *srv) Update(ctx context.Context, in *server.UpdateRequest) (*temporal.Profile, error) {
	id, err := expression.EvalExpression(temporal.CreateIdexpression, in.ProtoReflect())
	if err != nil {
		return nil, err
	}
	update, err := s.tcl.Update(ctx, id, "", &temporal.UpdateRequest{
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}
	return update, nil
}

func (s *srv) Delete(ctx context.Context, in *server.DeleteRequest) (*emptypb.Empty, error) {
	id, err := expression.EvalExpression(temporal.CreateIdexpression, in.ProtoReflect())
	if err != nil {
		return nil, err
	}
	if err := s.tcl.Delete(ctx, id, ""); err != nil {
		return nil, err
	}
	return nil, nil
}
