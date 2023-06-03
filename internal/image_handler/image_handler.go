package image_handler

import (
	"context"
	"log"
	"net"

	"github.com/phillipashe/iffi/internal/decode_image"
	pb "github.com/phillipashe/iffi/proto/image"
	"google.golang.org/grpc"
)

type server struct {
	// Embed the unimplemented server
	pb.DecodeImageServer
}

func (s *server) Decode(ctx context.Context, req *pb.Image) (*pb.DecodedImage, error) {
	// message := "Hello world"
	gps := decode_image.GetExif(req.B64)
	response := &pb.DecodedImage{Decoded: gps}
	return response, nil
}

func InitializeImageHandler() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterDecodeImageServer(srv, &server{})

	log.Println("Starting gRPC server on port 50051...")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
