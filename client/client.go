package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/imadmaachari/grpc-note-app/prototype/note/pb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Welcome to Note gRPC client")
	conn, err := grpc.Dial("localhost:50056", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer conn.Close()
	c := pb.NewNoteServiceClient(conn)

	//CREATE NOTE
	fmt.Println("Creating Notes :")
	createNoteRequest := &pb.CreateNoteRequest{
		Note: &pb.Note{
			StudentUuid: "XUI-123-43F",
			Instructor:  "DSF-4RF-5SS",
			Mark:        15,
			Language:    "Go gRPC",
		},
	}
	createNoteResponse, err := c.CreateNoteList(context.Background(), createNoteRequest)
	if err != nil {
		log.Fatalf("Error while creating Note %v", err)
	}
	fmt.Println(createNoteResponse)
	createNoteRequest2 := &pb.CreateNoteRequest{
		Note: &pb.Note{
			StudentUuid: "XUI-123-43F",
			Instructor:  "DSF-4RF-5SS",
			Mark:        15,
			Language:    "Go gRPC",
		},
	}
	createNoteResponse2, err := c.CreateNoteList(context.Background(), createNoteRequest2)
	if err != nil {
		log.Fatalf("Error while creating Note %v", err)
	}
	fmt.Println(createNoteResponse2)
	//LIST NOTES
	fmt.Println("Listing Notes :")
	getNoteListRequest := &pb.GetNoteListRequest{}
	getNoteListResponse, err := c.GetNoteList(context.Background(), getNoteListRequest)
	if err != nil {
		log.Fatalf("Error while fetching Notes %v", err)
	}
	fmt.Println(getNoteListResponse)
	//DELETE NOTE
}
