package main

import (
	"context"
	"grpc-app/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := proto.NewAppServiceClient(conn)
	req := &proto.AddRequest{X: 100, Y: 200}
	resp, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Add Result:", resp.GetResult())
}
