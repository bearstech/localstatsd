package main

import (
	"fmt"
	"net"
	"time"
)

/*
Listen UDP, localhost.
*/
func listen(events chan []byte, port int) {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	var buff [1024]byte
	for {
		n, _, err := conn.ReadFromUDP(buff[:])
		if err != nil {
			panic(err)
		}
		events <- buff[:n]
	}
}

/*
Talk to the real statsd server, with TCP.
*/
func talk(events chan []byte, address string) {
	for {
		conn, err := net.DialTimeout("tcp", address, 10*time.Second)
		if err != nil {
			fmt.Println("Error: target server not available")
			time.Sleep(5 * time.Second)
			continue
		} else {
			fmt.Println("Connected")
		}
		for again := true; again; {
			event := <-events
			_, err := conn.Write(append(event, '\n'))
			if err != nil {
				fmt.Println("Error: ", err)
				events <- event
				again = false
			}
		}
	}
}

func main() {
	events := make(chan []byte, 100)
	go talk(events, "127.0.0.1:8125")
	listen(events, 8126)
}
