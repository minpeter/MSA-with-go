package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

type helloWorldRequest struct {
	Name string `json:"name"`

	// Method는 Http Method를 저장하는 필드이다.
	Method string

	// Header는 서버가 수신한 요청 헤더 필드를 가지고 있다.
	Header http.Header

	// Body는 요청의 본문이다.
	Body io.ReadCloser
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var request helloWorldRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello, " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(&response)
}
