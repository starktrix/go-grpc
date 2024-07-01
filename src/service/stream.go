package service

import (
	"errors"
	"fmt"
	"io"
	pb "starktrix/grpc/src/proto/stream"

	"google.golang.org/grpc"
)


type streamServer struct {
	pb.UnimplementedStreamServiceServer
}

func (srv *streamServer) ClientStream(stream pb.StreamService_ClientStreamServer) error {
	// Recv() & SendAndClose()
	var count int
	for {
		val, err := stream.Recv()
		if err == io.EOF {return nil}
		if err != nil {
			return errors.New("err reading client stream")
		}
		fmt.Println("value received: ", val)
		count++
		if count > 5 {
			break
		}
		
	}
	return stream.SendAndClose(&pb.ClientStreamEnd{Clientend: "client end"})
	// return nil
	// return errors.New("err")
}
func (srv *streamServer) ServerStream(start *pb.ServerStreamStart, stream pb.StreamService_ServerStreamServer) error {
	// Send()
	fmt.Println("===  server stream started ===")
	fmt.Println("start :", start)
	var count int
	for {
		err := stream.Send(
			&pb.StreamValue{
			Id: 5, 
			Data: "data", 
			Src: "server", 
			Dst: "client",
			},
		)
		if err != nil {
			return errors.New("err server sending stream")
		}
		count++
		if count > 10 {
			break
		}
	}
	return nil
	// return errors.New("err")
}
func (srv *streamServer) BiStream(stream pb.StreamService_BiStreamServer) error {
	// Send() & Recv()
	var count int
	for {
		err := stream.Send(
			&pb.StreamValue{
			Id: 5, 
			Data: "data", 
			Src: "server", 
			Dst: "client",
			},
		)
		if err != nil {
			return errors.New("err server sending bi stream")
		}
		val, err := stream.Recv()
		if err == io.EOF {return nil}
		if err != nil {
			return errors.New("err server receving bi stream")
		}
		fmt.Println("value received: ", val)
		count++
		if count > 10 {
			break
		}
	}
	return errors.New("err")

}


func RegisterStreamService() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, &streamServer{})
	return s
}