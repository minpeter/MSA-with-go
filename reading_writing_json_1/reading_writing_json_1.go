package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Starting server on port %d", port)
	// Fatal의 경우 fmt.Print 후 os.Exit(1)을 호출 한 것과 같은 기능
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello, world!"}
	data, err := json.Marshal(response)

	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}
