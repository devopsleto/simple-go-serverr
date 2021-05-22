package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ErrorResponse struct {
	Code  int
	Error string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO!!!")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "root!!!")
}

func errHandler(w http.ResponseWriter, r *http.Request) {
	encoded, error := json.Marshal(ErrorResponse{
		Error: "boom",
		Code:  http.StatusInternalServerError,
	})

	if error == nil {
		http.Error(w, string(encoded), http.StatusInternalServerError)
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}

func main() {
	log.SetOutput(os.Stdout)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/errr", errHandler)
	const port = 8080
	log.Printf("Started server at %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
