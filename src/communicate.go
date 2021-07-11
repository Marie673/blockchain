package main

import (
	"net"
	"time"
)

func OpenServer() {
	port := ":100000"
	addr, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			HandleClient(connection)
		}()
	}
}

func HandleClient(connection net.Conn) {
	err := connection.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return
	}

}
