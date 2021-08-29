package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Welcome to the HomePage!")
}

func feedPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the FeedPage!")
	fmt.Println("Welcome to the FeedPage!")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/home", homePage)
	router.HandleFunc("/feed", feedPage)
	log.Fatal(http.ListenAndServe(":80", router))
}

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println("Server starting on port 80....")
	handleRequest()
}
