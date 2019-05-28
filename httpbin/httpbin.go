package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
	"strings"
)

type HeaderBody struct {
	Args 	map[string]string 	`json:"args"`
	Headers map[string]string 	`json:"headers"`
	Origin 	string 				`json:"origin"`
	Url 	string 				`json:"url"`
}

var Port = "4000"

func handler(w http.ResponseWriter, r *http.Request) {
	headers := make(map[string]string)
	args := make(map[string]string)

	// headersBody
	origin := strings.Split(r.RemoteAddr, ":")[0]
	headersBody := HeaderBody{Origin: origin, Url: "http://" + r.Host + r.URL.String()}

	// headers
	Host := strings.ReplaceAll(r.Host, ":" + Port, "")
	headers["Host"] = Host

	for k, v := range r.Header {
		headers[k] = v[0]
	}

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
	log.Fatal(http.ListenAndServe("localhost:" + Port, nil))
}
