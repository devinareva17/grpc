package main


import (
   "flag"
   "fmt"
   "context"
   "log"
   "net"


   pb "github.com/devinareva17/grpc/calculator/calculatorpb"
   "google.golang.org/grpc"
   "google.golang.org/grpc/reflection"
   "google.golang.org/grpc/codes"
)


var port = flag.Int("port", 50053, "Port server")


type server struct{
   pb.UnimplementedCalculatorServer
}


// mendefinisikan fungsi add
func (s *server) Add(ctx context.Context, in *pb.Request) (*pb.AddResponse, error) {
   //print N1 dan N2 dari client
   log.Printf("Received request with N1=%.2f and N2=%.2f\n", in.N1, in.N2)
   return &pb.AddResponse{N1: in.N1 + in.N2}, nil
}


// mendefinisikan fungsi substract
func (s *server) Subtract(ctx context.Context, in *pb.Request) (*pb.SubstractResponse, error) {
   return &pb.SubstractResponse{N1: in.N1 - in.N2}, nil
}


// mendefinisikan fungsi Multiply
func (s *server) Multiply(ctx context.Context, in *pb.Request) (*pb.MultiplyResponse, error) {
   return &pb.MultiplyResponse{N1: in.N1 * in.N2}, nil
}


// mendefinisikan fungsi Divide
func (s *server) Divide(ctx context.Context, in *pb.Request) (*pb.DivideResponse, error) {
   if in.N2 == 0 {
       return nil, grpc.Errorf(codes.InvalidArgument, "Cannot divide by zero")
   }
   return &pb.DivideResponse{N1: in.N1 / in.N2}, nil
}


func main() {
   flag.Parse()
   lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
   if err != nil {
       log.Fatalf("failed to listen: %s", err)
   }
   s := grpc.NewServer()
   pb.RegisterCalculatorServer(s, &server{})
   reflection.Register(s)
   fmt.Printf("listening on port %d...\n", *port)
   if err := s.Serve(lis); err != nil {
       log.Fatalf("failed to serve: %v", err)
   }
}
