package service

import (
	"context"
	"errors"
	"fmt"
	todoPb "starktrix/grpc/src/proto/todo"

	"google.golang.org/grpc"
)
type todoServer struct {
	// todoPb.UnimplementedTodoServiceServer
	todoPb.TodoServiceServer
}



func (srv *todoServer) Add(ctx context.Context, req *todoPb.Todo) (*todoPb.TodoMessage, error) {
	todo := fmt.Sprintf("%s-%s-%s\n", req.GetId(), req.GetTitle(), req.GetBody())
	fmt.Print(todo)
	todoMsg := &todoPb.TodoMessage{
		Message: todo,
	}
	return todoMsg, nil
	// return nil, errors.New("err: no data found")
}
func (srv *todoServer) GetAll(ctx context.Context, req *todoPb.Empty) (*todoPb.TodoList, error) {
	return nil, errors.New("err: no data found")
}
func (srv *todoServer) GetTodo(ctx context.Context, req *todoPb.TodoMessage) (*todoPb.Todo, error) {
	return nil, errors.New("err: no data found")
}
func (srv *todoServer) Edit(ctx context.Context, req *todoPb.TodoId) (*todoPb.TodoMessage, error) {
	return nil, errors.New("err: no data found")
}


// Register Server

func RegisterTodoService() *grpc.Server {
	s := grpc.NewServer()
	todoPb.RegisterTodoServiceServer(s, &todoServer{})
	return s
}