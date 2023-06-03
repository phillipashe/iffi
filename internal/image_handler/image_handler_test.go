package image_handler

import (
	"context"
	"log"
	"math"
	"net"
	"os"
	"testing"

	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/phillipashe/iffi/proto/image"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// helper func to check that two floats are close enough to consider "equal"
func FloatDifferenceWithinTolerance(f0 float64, f1 float64, tolerance float64) bool {
	return math.Abs(f0-f1) > tolerance
}

// Create a proto to send to the gRPC endpoint
func MakeProto(imageFileName string) *pb.Image {
	imageData, err := os.ReadFile("../../testing/" + imageFileName)
	if err != nil {
		log.Fatalf("failed to retrieve iguana image from disk")
	}

	// load and serialize test image
	pb_img := &pb.Image{
		ImageData: imageData,
	}
	return pb_img
}

// set up the gRPC endpoint for unit testing
func SetupEndpoint() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer(
		grpc.MaxRecvMsgSize(2 * int(math.Pow(10, 7))),
	)
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
	request := MakeProto("olympus.jpeg")

	// Make the RPC call
	response, err := client.Decode(ctx, request)
	if err != nil {
		t.Fatalf("Failed to call Decode: %v", err)
	}

	// Verify the response
	expectedLatitude := 50.819053
	expectedLongitude := -0.136792
	expectedDatetime := &timestamp.Timestamp{Seconds: 1390507038, Nanos: 0}
	if FloatDifferenceWithinTolerance(response.GetLatitude(), expectedLatitude, 0.001) {
		t.Errorf("Unexpected response. Expected: %f, Got: %f", expectedLatitude, response.Latitude)
	}
	if FloatDifferenceWithinTolerance(response.GetLongitude(), expectedLongitude, 0.001) {
		t.Errorf("Unexpected response. Expected: %f, Got: %f", expectedLongitude, response.Longitude)
	}
	if response.GetDatetime().GetSeconds() != expectedDatetime.GetSeconds() {
		t.Errorf("Unexpected response. Expected: %s, Got: %s", expectedDatetime, response.Datetime)
	}
}
