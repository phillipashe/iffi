package main

import "github.com/phillipashe/iffi/internal/image_handler"

func HelloWorld() string {
	return "Hello world"
}

func main() {
	HelloWorld()
	image_handler.HandleImage()
}
