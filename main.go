package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Response struct {
	Message string
}

func mainHandler(w http.ResponseWriter, req *http.Request) {
	responseObject := Response{"Hello world !"}
	jsonResponse, _ := json.Marshal(responseObject)
    fmt.Fprintf(w, string(jsonResponse))
}

func sayHello(w http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["name"]
	if !ok || keys[0] == "" {
		responseObject := Response{"Error : cannot find name"}
		jsonResponse, _ := json.Marshal(responseObject)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, string(jsonResponse))
	} else {
		responseObject := Response{"Hello " + keys[0]}
		jsonResponse, _ := json.Marshal(responseObject)
		fmt.Fprintf(w, string(jsonResponse))
	}
}

var port string = "8000"

func main() {
	fmt.Println("Server running on port " + port)
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/hello", sayHello)
	http.ListenAndServe(":" + port, nil)
}
