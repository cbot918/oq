package main

import (
	"fmt"
	"net/http"
)

func RunHTTPServer(port string) error {
	fmt.Println("listening: ", port)
	err := http.ListenAndServe(port, NewHTTPQ())
	if err != nil {
		return err
	}
	return nil
}

type HTTPServer struct{}

func NewHTTPQ() *HTTPServer {
	return &HTTPServer{}
}

func (q *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hihihihi")
	fmt.Println(r)
}
