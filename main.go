package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("run the program properly u idiot")
		return
	}
	mode := os.Args[1]
	if mode == "server" {
		listener, err := net.Listen("tcp", ":4040")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		defer listener.Close()

		for {

			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			go handleConnection(conn)

		}

	} else if mode == "client" {

		conn, err := net.Dial("tcp", "localhost:4040")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		defer conn.Close()

		// var str string
		// fmt.Scanln(&str)
		reader := bufio.NewReader(os.Stdin)
		go recieveMessages(conn)
		for {
			fmt.Println("Enter the Message")

			str, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			_, err = conn.Write([]byte(str))
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			fmt.Println("Message sent")

		}
	}

}

func handleConnection(conn net.Conn) {

	defer conn.Close()
	for {

		s := make([]byte, 1024)
		n, err := conn.Read(s)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Received: ", string(s[:n]))

		// reader := bufio.NewReader(os.Stdin)
		// str, err := reader.ReadString('\n')

		// if err != nil {
		// 	fmt.Println("Error :", err)
		// 	return
		// }

		_, err = conn.Write([]byte(s[:n]))
		if err != nil {
			fmt.Println("Server failed to send Reply: ", err)
			return
		}
	}
}

func receiveMessages(conn net.Conn) {
	for {

		s := make([]byte, 1024)
		n, err := conn.Read(s)
		if err != nil {
			fmt.Println("Client failed to read:", err)
			return
		}

		fmt.Printf("Server sent : %s\n", string(s[:n]))

	}
}
