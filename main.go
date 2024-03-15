package main

import "net/http"

func main() {
	// serve a basic http server that returns hello world
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8080", nil)
}
