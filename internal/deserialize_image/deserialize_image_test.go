package deserialize_image

import (
	pb "iffi/proto/image"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestDeserializeImage(t *testing.T) {

	iguana_b64 := "abcd"

	// load and serialize test image
	iguana_img := &pb.Image{
		B64: iguana_b64,
	}
	iguana_pbuf, err := proto.Marshal(iguana_img)

	// deserialize the test image
	res := DeserializeImage(iguana_pbuf)

	if err != nil {
		t.Fatalf("failed to call deserialize_image: %v", err)
	}
	if res != "abcd" {
		t.Fatalf("This image isn't of an iguana!  Got:%v\nExpected:%v", res, iguana_b64)
	}
}
