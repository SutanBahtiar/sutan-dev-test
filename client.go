package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func StartClient() {
	// Connecting to server
	conn, err := net.Dial(connType, sockAdrr)
	if err != nil {
		log.Println("Error connecting to server", err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		log.Print("Text to send: ")

		// input message
		input, _ := reader.ReadString('\n')

		// print send message
		log.Println("Send message:", input)

		// send message
		conn.Write([]byte(input))

		// get response message
		message, err := bufio.NewReader(conn).ReadString('\n')

		// print response message
		log.Println("Response message:", message)

		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}

	}

}
