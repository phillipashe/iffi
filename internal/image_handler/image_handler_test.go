package image_handler

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/phillipashe/iffi/proto/image"
	"google.golang.org/grpc"
)

func SetupEndpoint() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterDecodeImageServer(srv, &server{})

	log.Println("Starting gRPC server on port 50051...")
	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}

func TestDecode(t *testing.T) {

	SetupEndpoint()

	// Create a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a client for the gRPC service
	client := pb.NewDecodeImageClient(conn)

	// Create a context for the RPC
	ctx := context.Background()

	// Prepare the request
	request := &pb.Image{
		// Set image fields
	}

	// Make the RPC call
	response, err := client.Decode(ctx, request)
	if err != nil {
		t.Fatalf("Failed to call Decode: %v", err)
	}

	// Verify the response
	expectedMessage := "Hello world"
	if response.Decoded != expectedMessage {
		t.Errorf("Unexpected response. Expected: %s, Got: %s", expectedMessage, response.Decoded)
	}
}
