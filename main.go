package main

import (
	"fmt"
	"io"
	"net/http"
)

func respond(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func startServer() {
	fmt.Println("Server Started")
	http.HandleFunc("/", respond)
	http.ListenAndServe(":3000", nil)
}

func StartMyApp() {
	go startServer()
}

func main() {
	fmt.Println("Starting Server")
	startServer()
}
