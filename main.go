package main

import (
	"context"
	"grpc-1/invoicer"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("listener error: %v", err)
	}

	serviceRegistrar := grpc.NewServer()
	server := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serviceRegistrar, server)

	err = serviceRegistrar.Serve(listener)
	if err != nil {
		log.Fatalf("serviceRegistrar err: %v", err)
	}
}
