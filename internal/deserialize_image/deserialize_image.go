package deserialize_image

import (
	"fmt"
	pb "iffi/proto/image"

	"google.golang.org/protobuf/proto"
)

func DeserializeImage(img []byte) string {
	image := &pb.Image{}
	err := proto.Unmarshal(img, image)
	if err != nil {
		fmt.Printf("error deserializing the image protobuf: %v", err)
	}
	return image.B64
}
