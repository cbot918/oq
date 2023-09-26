package main

import (
	"fmt"
	"net"
)

func RunTcpServer(port string) error {
	fmt.Println("listening: ", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			lf("accept failed: ", err.Error())
			continue
		}

		go handleConn(conn)

	}

}

func handleConn(conn net.Conn) {
	lg("conn:")
	lv(conn)

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				lg("client 結束連線")
				return
			}
			lg("conn read failed")
			lg(err.Error())
			return
		}
		lg("message: ")
		lg(string(buf[:n]))

		_, err = conn.Write([]byte(httpResponse(" hihihi\n ")))
		if err != nil {
			lg("conn write failed")
			lg(err.Error())
			return
		}
		conn, ok := conn.(*net.TCPConn)
		if ok {
			err := conn.CloseWrite()
			if err != nil {
				lg("conn.CloseWrite failed: ", err.Error())
				return
			}
		}

	}

}

func httpResponse(message string) string {

	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Connection: close\r\n\r\n" +
		message

	return response

}
