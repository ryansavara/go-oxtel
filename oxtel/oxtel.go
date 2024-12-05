package oxtel

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

type Oxtel struct {
	address    string
	port       uint16
	conn       net.Conn
	reader     bufio.Reader
	RxMessages chan string
	end        chan bool
}

func NewOxtel(address string, port uint16) *Oxtel {
	return &Oxtel{
		address:    address,
		port:       port,
		conn:       nil,
		RxMessages: make(chan string),
		end:        make(chan bool),
	}
}

func (o *Oxtel) Connect() error {
	address := fmt.Sprintf("%s:%d", o.address, o.port)

	c, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	o.conn = c
	o.reader = *bufio.NewReader(c)
	go o.rxLoop()
	return nil
}

func (o *Oxtel) Disconnect() error {
	if o.conn == nil {
		return nil
	}
	err := o.conn.Close()

	if err == io.EOF {
		err = nil
	}
	o.conn = nil
	return err
}

func (o *Oxtel) rxLoop() {
	for {
		line, err := o.reader.ReadString(':')
		if err != nil {
			if err == io.EOF {
				break
			}
			o.Disconnect()
			break
		}

		line = strings.TrimSpace(line)
		if len(line) > 0 {
			o.RxMessages <- line
		}
	}
}

func (o *Oxtel) sendCommand(cmd string) error {
	escapedCmd := strings.ReplaceAll(cmd, "\\", "\\5C")
	escapedCmd = strings.ReplaceAll(escapedCmd, "|", "\\7C")
	escapedCmd = strings.ReplaceAll(escapedCmd, ";", "\\3B")
	escapedCmd = strings.ReplaceAll(escapedCmd, ":", "\\3A")

	escapedCmd += ":"
	fmt.Printf("Sending command: '%s'\n", escapedCmd)
	cmdBytes := []byte(escapedCmd)

	_, err := o.conn.Write(cmdBytes)
	if err != nil {
		if err == io.EOF {
			o.Disconnect()
		}
	}
	return err
}

func (o *Oxtel) sendCommandExpectResponse(cmd string, data string) (string, error) {
	err := o.sendCommand(cmd + data)
	if err != nil && err != io.EOF {
		return "", err
	}

	timeout := time.After(5 * time.Second)
	for {
		select {
		case unsolicited := <-o.RxMessages:
			if unsolicited[:len(cmd)] == cmd {
				return unsolicited[len(cmd) : len(unsolicited)-1], nil
			}
		case <-timeout:
			return "", &TimeoutError{
				BaseError: BaseError{
					Message: "Timed out waiting for response",
				},
			}
		}
	}
}
