package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, task tracker"))
	})

	fmt.Println("Listening on 2525")
	http.ListenAndServe(":2525", nil)
}
