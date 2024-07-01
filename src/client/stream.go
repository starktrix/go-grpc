package main

import (
	"context"
	"fmt"
	"io"

	// "io"
	"log"
	"time"

	pb "starktrix/grpc/src/proto/stream"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func stream() {
	fmt.Println("=== Starting stream client ====")
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
	c := pb.NewStreamServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
    defer cancel()
	
	// CLIENT STREAM
	// clientStream(ctx, c)

	// SERVER STREAM
	serverStream(ctx, c)
	
}


func clientStream(ctx context.Context, c pb.StreamServiceClient) {
	stream, err := c.ClientStream(ctx)
	if err != nil {
		log.Fatalf("%v.ClientStream(_) = _, %v", c, err)
	  }
	  for i :=0; i < 10; i++ {
		err := stream.Send(
			&pb.StreamValue{
				Id: int32(i), 
				Data: "data", 
				Src: "client stream", 
				Dst: "server",
				},
			)
			if err != nil {
				// return errors.New("error sending client stream")
				fmt.Println("error sending stream")
				return
			}
	  }
	  reply, err := stream.CloseAndRecv()
	  if err == io.EOF {return}
		if err != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
		 }
		 log.Printf("Route summary: %v", reply)
		//  return
}

func serverStream(ctx context.Context, c pb.StreamServiceClient) {
	stream, err := c.ServerStream(ctx, &pb.ServerStreamStart{Serverstart: "server started"})
	if err != nil {
		log.Fatalf("%v.ServerStream(_) = _, %v", c, err)
	  }
	  for {
		val, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ServerStream(_) = _, %v", c, err)
		}
		log.Println(val)
	}
}



