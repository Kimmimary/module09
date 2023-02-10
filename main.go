package main

import (
	"fmt"
	"net/http"
	"time"
)

type ResponseHandler struct {
	message string
}

func (rh ResponseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rh.message)
}

func main() {
	resp := ResponseHandler{message: "hello human"}
	http.Handle("/hello", resp)
	http.Handle("/source", http.StripPrefix("/source", http.FileServer(http.Dir("/tmp"))))
	http.Handle("/longPing", http.TimeoutHandler(ResponseHandler{"Request Timeout"}, time.Second*1, "Error"))
	http.ListenAndServe(":8080", nil)
}
