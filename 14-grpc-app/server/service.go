package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	time.Sleep(5 * time.Second)
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

func isPrime(n int64) bool {
	var i int64
	for i = 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func (s *server) Prime(req *proto.PrimeRequest, stream proto.AppService_PrimeServer) error {
	start := req.GetRangeStart()
	end := req.GetRangeEnd()
	for no := start; no <= end; no++ {
		if isPrime(no) {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Sending : ", no)
			resp := &proto.PrimeResponse{No: no}
			stream.Send(resp)
		}
	}
	fmt.Println("End of generating prime nos")
	return nil
}

func (s *server) Greet(stream proto.AppService_GreetServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		greeting := req.GetGreeting()
		first_name := greeting.GetFirstName()
		last_name := greeting.GetLastName()
		greetMsg := fmt.Sprintf("Hi %s %s, Have a nice day!", first_name, last_name)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Sending : ", greetMsg)
		resp := &proto.GreetResponse{GreetMessage: greetMsg}
		stream.Send(resp)
	}
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
