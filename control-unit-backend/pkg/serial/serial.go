package serial

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

type SerialConnection struct {
	port serial.Port
}

var SerialConn *SerialConnection

var SerialChannel = make(chan string)

func OpenSerial(portName string, baudRate int) (*SerialConnection, error) {
	mode := &serial.Mode{
		BaudRate: baudRate,
	}

	port, err := serial.Open(portName, mode)

	if err != nil {
		return nil, err
	}

	return &SerialConnection{port: port}, nil
}

func (s *SerialConnection) Read() {
	buf := make([]byte, 100)
	for {
		n, err := s.port.Read(buf)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			fmt.Println("\nEOF") // read until no bytes available
			break
		}
		SerialChannel <- string(buf[:n])
	}
}

func (s *SerialConnection) Write(msg string) {
	_, err := s.port.Write([]byte(msg + "\n"))
	if err != nil {
		log.Fatal(err)
	}
}
