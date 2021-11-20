package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//handler
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	//servermux
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", helloWorld)

	//webserver
	log.Fatal(http.ListenAndServe(":8080", router))
}
