package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {

	portFlag := flag.Int("port", 30000, "Port to send udp on")
	ipFlag := flag.String("ip", "127.0.0.1", "IP to send on")
	timeoutFlag := flag.Int("timeout", 2, "Timeout to dial the udp server")

	flag.Parse()

	ip := *ipFlag
	port := *portFlag
	timeout := time.Duration(*timeoutFlag) * time.Second

	fmt.Printf("IP:     %s\n", ip)
	fmt.Printf("Port:   %v\n", port)
	fmt.Printf("Timeout %v\n", timeout)

	p := make([]byte, 2048)
	conn, err := net.DialTimeout("udp", fmt.Sprintf("%s:%v", ip, port), time.Duration(*timeoutFlag)*time.Second)
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
