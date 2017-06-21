package main

import (
	"fmt"
	"io"
	"net/http"
)

func respond(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func StartServer() {
	fmt.Println("Server Started")
	http.HandleFunc("/", respond)
	http.ListenAndServe(":3000", nil)
}

func main() {
	fmt.Println("Starting Server")
	StartServer()
}
