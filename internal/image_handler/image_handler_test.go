package image_handler

import (
	// "context"
	"testing"
	// "time"
	// "google.golang.org/grpc"
)

func TestGetMessage(t *testing.T) {

}
func TestHandleImage(t *testing.T) {

	expectedResponse := "placeholder"

	res := HandleImage()
	if res != expectedResponse {
		t.Fatalf("Response was invalid")
	}

	// // Wait for the server to start
	// time.Sleep(time.Second)

	// // Set up a connection to the server
	// conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	// if err != nil {
	// 	t.Fatalf("failed to dial server: %v", err)
	// }
	// defer conn.Close()

	// // Create a client for the Greeter service
	// client := pb.NewGreeterClient(conn)

	// // Call the Hello endpoint with a test message
	// resp, err := client.Hello(context.Background(), &pb.HelloRequest{Name: "Alice"})
	// if err != nil {
	// 	t.Fatalf("failed to call Hello: %v", err)
	// }

	// // Verify the response message
	// if resp.Message != "Hello, Alice!" {
	// 	t.Errorf("unexpected response: %s", resp.Message)
	// }
}
