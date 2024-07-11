package main

import "net/http"

func main() {
	r := http.NewServeMux()
	r.Handle("/", http.FileServer(http.Dir(".build/")))
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
