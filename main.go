package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8887"

func main() {
	fmt.Println("listening: ", port)
	err := http.ListenAndServe(port, New())
	if err != nil {
		log.Fatal(err)
	}
}

type Q struct{}

func New() *Q {
	return &Q{}
}

func (q *Q) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
