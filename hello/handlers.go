package main

import "net/http"

func helloWorldHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!\n"))
}
