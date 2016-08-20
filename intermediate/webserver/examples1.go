package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloRequest)
	http.HandleFunc("/", getRequest)
	http.HandleFunc("/rabbit", getRabbit)
	port := ":5000"

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func helloRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
	return
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	fileRequest := "./" + r.URL.Path
	http.ServeFile(w, r, fileRequest)
	return
}

func getRabbit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Rabbit Evil")
	return
}
