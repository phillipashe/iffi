package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/phillipashe/iffi/internal/image_handler"
)

func ServeHTML() {
	fmt.Printf("Starting server at port 8080\n")
	http.Handle("/", http.FileServer(http.Dir("./web")))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	go ServeHTML()
	image_handler.InitializeImageHandler()
}
