package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

func main() {

	port := flag.Int("port", 30000, "Port to send udp on")
	ip := flag.String("ip", "127.0.0.1", "IP to send on")

	flag.Parse()

	p := make([]byte, 2048)
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%v", *ip, *port))
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
