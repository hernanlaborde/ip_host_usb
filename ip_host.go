package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// Get preferred outbound ip of this machine
// Recovered from: https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {

	ip := GetOutboundIP()
	fmt.Println("IP: ", ip)

	// Recovered from: https://www.cloudhadoop.com/2018/12/golang-examples-hostname-and-ip-address.html
	hostname, error := os.Hostname()
	if error != nil {
		panic(error)
	}
	fmt.Println("hostname: ", hostname)

}
