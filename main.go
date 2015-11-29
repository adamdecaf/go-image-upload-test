package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/adamdecaf/go-image-upload-test/routes"
)

func main() {
	// set routes
	http.Handle("/", http.FileServer(http.Dir("./html/")))
	http.HandleFunc("/ping", routes.Ping)
	http.HandleFunc("/upload", routes.Upload)

	fmt.Printf("Starting http server on :8080\n")

	err := http.ListenAndServe(":8080", nil) // set listen port

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
