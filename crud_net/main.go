package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world 2"))
	})
	http.ListenAndServe(":3000", nil)

}
