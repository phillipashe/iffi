package decode_image

import (
	"bytes"
	"encoding/base64"
	"log"
	"strconv"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

func ConvertToBuffer(b64 string) *bytes.Reader {

	// Remove the data URI prefix
	data := strings.TrimPrefix(b64, "data:image/png;base64,")

	// Decode the Base64 string
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Fatal(err)
	}

	// Create a buffer from the decoded data
	buffer := bytes.NewReader(decoded)

	return buffer
}

func GetExif(b64 string) string {
	// Open the image file
	file := ConvertToBuffer(b64)

	// Decode the EXIF data
	x, err := exif.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	lat, long, _ := x.LatLong()

	return strconv.FormatFloat(lat, 'f', -1, 64) + strconv.FormatFloat(long, 'f', -1, 64)
}
