package main

import (
	"context"
	"fmt"
	"log"
	"time"

	todoPb "starktrix/grpc/src/proto/todo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func todo() {
	fmt.Println("=== Starting todo client ==========")
	// conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), /*grpc.WithBlock()*/)
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := todoPb.NewTodoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.Add(ctx, &todoPb.Todo{Id: "clientId", Title: "client_title", Body: "client_body"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.GetMessage())
}