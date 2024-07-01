package main

import (
	"fmt"
	"net"
	"starktrix/grpc/src/service"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50052")
	if err != nil {
		fmt.Println("error establishing tcp conn")
	}
	// todoServer := service.RegisterTodoService()
	// fmt.Println(" ==== Start GRPC server ====")
	// if err := todoServer.Serve(lis); err != nil {
	// 	e := fmt.Errorf("failed to serve: %v", err)
	// 	fmt.Println(e)
	// }
	streamServer := service.RegisterStreamService()
	fmt.Println(" ==== Start GRPC server ====")
	if err := streamServer.Serve(lis); err != nil {
		e := fmt.Errorf("failed to serve: %v", err)
		fmt.Println(e)
	}
}