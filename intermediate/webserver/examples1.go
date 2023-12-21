package webserver

import (
	"fmt"
	"net/http"
)

func helloRequest(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	fileRequest := "./" + r.URL.Path
	http.ServeFile(w, r, fileRequest)
}

func getRabbit(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Rabbit Evil")
}
