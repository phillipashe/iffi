package deserialize_image

import (
	"bufio"
	pb "iffi/proto/image"
	"os"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestDeserializeImage(t *testing.T) {

	iguana_file, err := os.Open("../../testing/iguana_with_exif.b64")
	if err != nil {
		t.Errorf("failed to retrieve iguana image from disk")
		return
	}
	defer iguana_file.Close()

	scanner := bufio.NewScanner(iguana_file)
	scanner.Scan()
	iguana_b64 := scanner.Text()
	// iguana_b64 := "abcd"

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
	if res != iguana_b64 {
		t.Fatalf("This image isn't of an iguana!  Got:%v\nExpected:%v", res, iguana_b64)
	}
}
