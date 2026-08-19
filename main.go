package main

import (
	"bufio"
	"fmt"
	"io"
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

		reader := bufio.NewReader(os.Stdin)
		go receiveMessages(conn)
		for {
			// fmt.Println("Enter the Message")

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
	reader := bufio.NewReader(conn)
	for {

		// s := make([]byte, 1024)
		// n, err := conn.Read(s)

		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
			} else {
				fmt.Println("Server failed to read:", err)
			}

			return
		}
		fmt.Println("Received: ", msg)

		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Server failed to send Reply: ", err)
			return
		}
	}
}

func receiveMessages(conn net.Conn) {

	reader := bufio.NewReader(conn)
	for {

		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Server disconnected")
			} else {
				fmt.Println("Client failed to read:", err)
			}

			return
		}

		fmt.Printf("Server sent : %s\n", msg)

	}
}
