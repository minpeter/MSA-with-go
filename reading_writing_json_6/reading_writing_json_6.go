package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	// 정적 파일 핸들러
	cathandler := http.FileServer(http.Dir("./images"))
	http.Handle("/cat/", http.StripPrefix("/cat/", cathandler))

	log.Printf("Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
