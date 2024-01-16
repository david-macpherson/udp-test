package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {

	port := flag.Int("port", 30000, "Port to send udp on")
	ip := flag.String("ip", "127.0.0.1", "IP to send on")
	timeoutFlag := flag.Int("timeout", 2, "Timeout to dial the udp server")

	flag.Parse()

	fmt.Printf("Sending address: %s:%v\n", *ip, *port)

	p := make([]byte, 2048)
	conn, err := net.DialTimeout("udp", fmt.Sprintf("%s:%v", *ip, *port), time.Duration(*timeoutFlag)*time.Second)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
