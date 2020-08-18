package main

import (
	"io"
	"log"
	"net/http"
)

func handleHome(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello! This is the home page.")
}

func main() {
	http.HandleFunc("/", handleHome)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
