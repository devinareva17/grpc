package main

import (
	"context"
	"log"
	"net"

	pb "github.com/devinareva17/grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/codes"
)

const (
	protocol = "tcp"
	port     = ":50053"
)

type server struct{pb.UnimplementedCalculatorServer}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddReply, error) {
    log.Printf("Received request with N1=%d and N2=%d\n", in.N1, in.N2)
	return &pb.AddReply{N1: in.N1 + in.N2}, nil
}

func (s *server) Subtract(ctx context.Context, in *pb.SubtractRequest) (*pb.SubtractReply, error) {
	return &pb.SubtractReply{N1: in.N1 - in.N2}, nil
}

func (s *server) Multiply(ctx context.Context, in *pb.MultiplyRequest) (*pb.MultiplyReply, error) {
	return &pb.MultiplyReply{N1: in.N1 * in.N2}, nil
}


func (s *server) Divide(ctx context.Context, in *pb.DivideRequest) (*pb.DivideReply, error) {
    if in.N2 == 0 {
        return nil, grpc.Errorf(codes.InvalidArgument, "Cannot divide by zero")
    }
    return &pb.DivideReply{N1: in.N1 / in.N2}, nil
}

func main() {
	lis, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}