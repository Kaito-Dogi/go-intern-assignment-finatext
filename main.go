package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "<h1>Hello, I am Kaito-Dogi!</h1>")
}
