package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/foo", FooHandler)

	fmt.Printf("Starting server at port 8080\n")
	http.ListenAndServe(":8080", nil)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := `{"message": "Hello World"}`

	w.Write([]byte(response))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resposne := `{"status": "OK"}`

	w.Write([]byte(resposne))
}

func FooHandler(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Hostname string `json:"hostname"`
		Message  string `json:"message"`
	}

	hostname, _ := os.Hostname()
	message := `Hello World`

	m := Response{Hostname: hostname, Message: message}

	marshal, _ := json.Marshal(m)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(marshal))
}
