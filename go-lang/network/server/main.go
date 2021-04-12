package main

import (
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Printf("net listen: %v", err)
		return
	}
	defer ln.Close()

	for {
		//conn, err := ln.Accept()
		//if err != nil {
		//	log.Printf("accept: %v", err)
		//	continue
		//}
		//conn.Close()
		time.Sleep(time.Minute)
	}
}
