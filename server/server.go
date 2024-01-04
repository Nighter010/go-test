package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	userName := "std1"
	passWord := "p@ssw0rd"

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer) // Read() blocks until it reads some data from the network and n is the number of bytes read
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// Print the number of bytes read
	Data := string(buffer[:n])

	if Data == fmt.Sprintf("%s:%s", userName, passWord) {
		conn.Write([]byte("Hello\n"))
	} else {
		conn.Write([]byte("Invalid credentials\n"))
	}
}

// buffer for reading

func main() {
	listener, err := net.Listen("tcp", ":5000")

	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	//close the listener when the application closes
	defer listener.Close()

	fmt.Println("Server is listening on port 5000")

	//Listen for an incoming connection
	for {
		//Accept() block until a connection is made
		conn, err := listener.Accept() // Accept is three-way handshake
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue //continue to next iteration of for loop
		}

		//Handle connections is a new goroutine
		fmt.Println("New connection established")

		go handleConnection(conn)
	}
}
