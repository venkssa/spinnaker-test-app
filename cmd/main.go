package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port = flag.Uint("port", 8080, "server port.")
)

func main() {
	flag.Parse()
	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()
		name := query.Get("name")
		if name == "" {
			name = "Anonymous"
		}
		fmt.Fprintln(rw, "Hello ", name)
	})
	log.Println("Launching at ", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		panic(err)
	}
}
