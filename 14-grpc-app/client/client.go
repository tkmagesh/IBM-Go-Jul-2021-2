package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := proto.NewAppServiceClient(conn)
	//doRequestResponse(client)
	//doClientStreaming(client)
	//doServerStreaming(client)
	doBiDirectionalStreaming(client)
}

func doBiDirectionalStreaming(client proto.AppServiceClient) {
	stream, err := client.Greet(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	requestData := []*proto.GreetRequest{
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Magesh",
				LastName:  "Kuppan",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Suresh",
				LastName:  "Kannan",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Ramesh",
				LastName:  "Jayaraman",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Rajesh",
				LastName:  "Pandit",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Naresh",
				LastName:  "Kumar",
			},
		},
	}

	go func() {
		for _, req := range requestData {
			stream.Send(req)
		}
		stream.CloseSend()
	}()
	/* wg := &sync.WaitGroup{}
	wg.Add(1) */
	done := make(chan bool)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("Thats all folks!")
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("Greet Result:", res.GetGreetMessage())
		}
		done <- true
	}()
	<-done
}
func doServerStreaming(client proto.AppServiceClient) {
	primeReq := &proto.PrimeRequest{RangeStart: 10, RangeEnd: 100}
	stream, e := client.Prime(context.Background(), primeReq)
	if e != nil {
		log.Fatalln(e)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Thats all folks!")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Prime Result:", res.GetNo())
	}
}

func doClientStreaming(client proto.AppServiceClient) {
	nos := []int64{10, 20, 30, 40, 50}
	stream, err := client.Average(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		req := &proto.AverageRequest{
			X: no,
		}
		fmt.Println("Sending : ", no)
		time.Sleep(500 * time.Millisecond)
		stream.Send(req)
	}
	response, e := stream.CloseAndRecv()
	if e != nil {
		log.Fatalln(e)
	}

	log.Println("Average Result:", response.GetResult())
}

func doRequestResponse(client proto.AppServiceClient) {
	req := &proto.AddRequest{X: 100, Y: 200}
	resp, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Add Result:", resp.GetResult())
}
