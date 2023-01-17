package client

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type Telnet struct {
	timeout time.Duration
	host    string
	port    string
}

func NewTelnet(timeout time.Duration, host, port string) *Telnet {
	return &Telnet{
		timeout: timeout,
		host:    host,
		port:    port,
	}
}

func (t *Telnet) Start() {
	conn, err := net.DialTimeout("tcp", addressBuilder(t.host, t.port), t.timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		cmd, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Fprintln(os.Stdout, "Closing connection...")
			}
			fmt.Fprintln(os.Stderr, err)
			return
		}
		_, err = fmt.Fprint(conn, cmd)
		if err != nil {
			fmt.Fprintln(conn, "Error while writing message...")
		}
		serverReader := bufio.NewReader(conn)
		text, err := serverReader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Fprintln(os.Stdout, "Server closed connection...")
			}
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Fprint(os.Stdout, text)
	}

}

func addressBuilder(host, port string) string {
	return host + port
}
