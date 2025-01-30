package serial

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

type SerialConnection struct {
	port serial.Port
}

func openSerial(portName string, baudRate int) (*SerialConnection, error) {
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
		fmt.Printf("%v", string(buf[:n]))
	}
}

func (s *SerialConnection) Write(msg string) {
	n, err := s.port.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)
}

// TODO: need to parse messages from/to serial line
