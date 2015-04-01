// Tragon: Simple SMTP server.
// Receives mail and prints it on the screen.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const (
	SMTPPort = ":2525" // SMTP port.
	Timeout  = 60      // Connection timeout.

	Reply220 = "Welcome to Tragon SMTP server." // Greeting message.
	Reply250 = "Ok, I'll swallow that."         // Ok message.
	Reply354 = "Give it to me..."               // Data message.
	Reply221 = "Yum!"                           // Quit message.
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", SMTPPort)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

// Handle SMTP connection.
func handleClient(conn net.Conn) {
	defer conn.Close()

	// Initialize timeout counter.
	time.AfterFunc(Timeout*time.Second, func() { conn.Close() })

	// Mandatory greeting to start SMTP dialogue.
	if _, err := fmt.Fprintf(conn, "220 %s\n", Reply220); err != nil {
		log.Println(err)
		return
	}

	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		// Extract command keyword.
		switch strings.ToUpper(strings.Trim(line, " \t\n\r")) {
		default:
			if _, err := fmt.Fprintf(conn, "250 %s\n", Reply250); err != nil {
				log.Println(err)
				return
			}
		case "DATA":
			if _, err := fmt.Fprintf(conn, "354 %s\n", Reply354); err != nil {
				log.Println(err)
				return
			}
			handleMessage(reader, conn)
		case "QUIT":
			if _, err := fmt.Fprintf(conn, "221 %s\n", Reply221); err != nil {
				log.Println(err)
				return
			}
			conn.Close()
		}
	}
}

// Read message data.
func handleMessage(reader *bufio.Reader, conn net.Conn) {
	var line string
	for strings.Trim(line, " \t\n\r") != "." {
		fmt.Printf("%s", line)
		line, _ = reader.ReadString('\n')
	}
	if _, err := fmt.Fprintf(conn, "250 %s\n", Reply250); err != nil {
		log.Println(err)
		return
	}
}
