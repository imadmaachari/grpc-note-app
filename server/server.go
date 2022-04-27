package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/imadmaachari/grpc-note-app/prototype/note/pb"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedNoteServiceServer
	Notes []pb.Note
}

func (s *server) CreateNoteList(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	s.Notes = append(s.Notes, *req.Note)
	return &pb.CreateNoteResponse{Note: req.Note}, nil
}

func (s *server) GetNoteList(ctx context.Context, req *pb.GetNoteListRequest) (*pb.GetNoteListResponse, error) {
	res := []*pb.Note{}
	for _, n := range s.Notes {
		u := pb.Note{
			StudentUuid: n.StudentUuid,
			Instructor:  n.Instructor,
			Subject:     n.Subject,
			Language:    n.Language,
		}
		res = append(res, &u)
	}

	ans := pb.GetNoteListResponse{Notes: res}
	return &ans, nil
}

func main() {
	fmt.Println("Welcome to Note gRPC server")
	lis, err := net.Listen("tcp", "0.0.0.0:50056")
	if err != nil {
		errors.Wrap(err, "Failed tolisten the port")
	}

	s := grpc.NewServer()
	pb.RegisterNoteServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		errors.Wrap(err, "Failed to server the listener")
	}
}
