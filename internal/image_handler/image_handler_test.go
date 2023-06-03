package image_handler

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"testing"

	pb "github.com/phillipashe/iffi/proto/image"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func MakeProto() *pb.Image {
	iguana_file, err := os.Open("../../testing/iguana_with_exif.b64")
	if err != nil {
		log.Fatalf("failed to retrieve iguana image from disk")
	}
	defer iguana_file.Close()

	scanner := bufio.NewScanner(iguana_file)
	scanner.Scan()
	iguana_b64 := scanner.Text()

	// load and serialize test image
	iguana_img := &pb.Image{
		B64: iguana_b64,
	}
	return iguana_img
}

func SetupEndpoint() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterDecodeImageServer(srv, &server{})

	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}

func TestDecode(t *testing.T) {

	SetupEndpoint()

	// Create a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a client for the gRPC service
	client := pb.NewDecodeImageClient(conn)

	// Create a context for the RPC
	ctx := context.Background()

	// Prepare the request
	request := MakeProto()

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
