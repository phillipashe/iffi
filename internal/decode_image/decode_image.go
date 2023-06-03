package decode_image

import (
	"bytes"
	"log"

	pb "github.com/phillipashe/iffi/proto/image"
	"github.com/rwcarlsen/goexif/exif"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetExif(image []byte) *pb.DecodedImage {

	// Decode the EXIF data
	x, err := exif.Decode(bytes.NewReader(image))
	if err != nil {
		log.Fatal(err)
	}

	// Get latitude, longitude, and timestamp
	// TODO: add error handling
	lat, long, _ := x.LatLong()
	tm, _ := x.DateTime()

	// return these values inside the response proto
	return &pb.DecodedImage{
		Longitude: long,
		Latitude:  lat,
		Datetime:  timestamppb.New(tm),
	}
}
