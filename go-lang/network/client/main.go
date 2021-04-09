package main

import (
	"log"
	"net"
	"time"
)

func main() {
	addr := "127.0.0.1:8000"

	for i := 0; i < 10000; i++ {
		err := connect(addr)
		if err != nil {
			log.Printf("%d connect: %v", i, err)
		} else {
			log.Printf("%d connect success", i)
		}
	}
}

func connect(addr string) error {
	d := net.Dialer{
		Timeout: 1 * time.Second,
	}
	conn, err := d.Dial("tcp", addr)
	if err != nil {
		return err
	}
	_ = conn
	//conn.Close()
	return nil
}
