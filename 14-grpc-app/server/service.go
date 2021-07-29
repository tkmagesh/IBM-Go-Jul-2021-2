package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	resp := &proto.AddResponse{Result: result}
	return resp, nil
}

func (s *server) Average(stream proto.AppService_AverageServer) error {
	var sum int64
	var count int64

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result := sum / count
			resp := &proto.AverageResponse{Result: result}
			fmt.Println(resp)
			stream.SendAndClose(resp)
			break
		}
		if err != nil {
			return err
		}
		sum += req.GetX()
		count++
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, &server{})
	e := grpcServer.Serve(listener)
	if e != nil {
		log.Fatalln(e)
	}
}
