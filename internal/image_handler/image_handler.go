package image_handler

import (
	"context"
	"log"
	"math"
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
	response := decode_image.GetExif(req.ImageData)

	// TODO add error handling
	return response, nil
}

func InitializeImageHandler() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer(
		// set max received size to 20mb
		grpc.MaxRecvMsgSize(2 * int(math.Pow(10, 7))),
	)
	pb.RegisterDecodeImageServer(srv, &server{})

	log.Println("Starting gRPC server on port 50051...")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
