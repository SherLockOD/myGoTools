package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type HeaderBody struct {
	Args    map[string]string `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}

var Port = "4000"

func handler(w http.ResponseWriter, r *http.Request) {
	headers := make(map[string]string)
	args := make(map[string]string)
	headersBody := HeaderBody{Origin: "", Url: "http://" + r.Host + r.URL.String()}
	for k, v := range r.Header {
		headers[k] = v[0]
	}
	Host := strings.ReplaceAll(r.Host, ":"+Port, "")
	headers["Host"] = Host
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		args[k] = v[0]
	}
	headersBody.Args = args
	headersBody.Headers = headers
	b, err := json.Marshal(headersBody)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Fprint(w, string(b))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:"+Port, nil))
}
