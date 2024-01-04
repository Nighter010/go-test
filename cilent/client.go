package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter username: ")
	userName, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Print("Enter password: ")
	passWord, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	userName = strings.TrimSpace(userName)
	passWord = strings.TrimSpace(passWord)

	data := fmt.Sprintf("%s:%s", userName, passWord)

	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(data))
	if err != nil {
		fmt.Println("Error sending data to server:", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving data from server:", err)
		return
	}

	fmt.Println("Server response:", string(buffer[:n]))
}
