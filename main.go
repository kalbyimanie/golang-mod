package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

func handleRequest(port string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/home", homePage)
	router.HandleFunc("/feed", feedPage)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	port := os.Getenv("PORT_NUMBER")
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("Server starting on port %s....", port)
	fmt.Println()
	handleRequest(port)
}
