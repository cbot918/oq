package main

import (
	"log"
)

const port = ":8887"

func main() {

	// err := RunTcpServer(port)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := RunHTTPServer(port)
	if err != nil {
		log.Fatal(err)
	}

}
