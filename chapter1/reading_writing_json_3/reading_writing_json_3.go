package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	// 출력 필드를 "message"로 변경
	Message string `json:"message"`
	// 이 필드를 출력하지 않는다.
	Author string `json:"-"`
	// 값이 비어 있으면 출력하지 않는다.
	Date string `json:",omitempty"`
	// 출력을 문자열로 변환하고 이름을 "id"로 변경
	ID int `json:"id,string"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello, world!"}
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(&response)
}
