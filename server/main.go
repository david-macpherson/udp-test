package main

import (
	"flag"
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your message "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func main() {

	port := flag.Int("port", 30000, "Port to send udp on")
	ip := flag.String("ip", "127.0.0.1", "IP to send on")

	flag.Parse()

	fmt.Printf("Sending address: %s:%v\n", *ip, *port)

	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: *port,
		IP:   net.ParseIP(*ip),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go sendResponse(ser, remoteaddr)
	}
}
