package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
)

// StartServer
func StartServer() {
	log.Printf("Starting %s server on %s:%s..", connType, host, port)
	listener, err := net.Listen(connType, host+":"+port)
	if err != nil {
		fmt.Println("Error starting server", err.Error())
	}

	// close server when app close
	defer listener.Close()

	// loop until close
	for {
		// accept connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accept connection", err.Error())
		}

		// logId
		logId := LogId()

		// print client connected
		log.Printf("%s, Client %s connected..", logId, conn.RemoteAddr().String())

		// handle connections concurrently
		go handleClient(logId, conn)
		// handleConnection(conn)
	}
}

func handleClient(logId string, conn net.Conn) {
	defer conn.Close()

	// get message until new line
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		conn.Write([]byte(err.Error()))
		log.Println(logId+", Error get message", err.Error())
		conn.Close()
		return
	}

	// message
	msg := string(buffer[:len(buffer)-1])

	// check '\n' character
	if len(msg) > 2 {
		lastChar := string(msg[len(msg)-2:])

		if strings.Contains(lastChar, "\\n") {
			msg = strings.ReplaceAll(msg, "\\n", "")
		}
	}

	// print message
	log.Printf("%s, Client message: %s", logId, msg)

	var message Message
	err = json.Unmarshal([]byte(msg), &message)
	if err != nil {
		conn.Write([]byte(err.Error()))
		log.Println(logId+", Error when parse message", err.Error())
		conn.Close()
		return
	}

	log.Printf("%s, Parse Message:%v", logId, message)
	sendMessage := SetClientMessage(message)
	log.Printf("%s, Send Message:%s", logId, sendMessage)

	// send mesasge
	conn.Write([]byte(sendMessage))

	// restart process
	// handleClient(conn)
}
