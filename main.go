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
	http.Handle("/i/", http.StripPrefix("/i/", http.FileServer(http.Dir("./tmp"))))
	http.HandleFunc("/upload", routes.Upload)

	fmt.Printf("Starting http server on :8080\n")

	err := http.ListenAndServe(":8080", nil) // set listen port

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
