package main

import (
	"control-unit-backend/pkg/models"
	"control-unit-backend/pkg/mqtt"
	"control-unit-backend/pkg/serial"
	"log"
	"strconv"
	"time"
)

func main() {

	sampler := models.Sampler{}
	sampler.StartSampling() // start sampling subroutine

	client := mqtt.ConnectMQTT("broker.mqtt-dashboard.com:1883")
	defer client.Disconnect(250)

	serialConn, err := serial.OpenSerial("/dev/cu.usbserial-14120", 9600)
	if err != nil {
		log.Fatal(err)
	}

	go serialConn.Read()

	perc := 0
	for {
		serialConn.Write("wi:" + strconv.Itoa(perc) + "\n")
		perc = (perc + 50) % 150
		time.Sleep(2 * time.Second)
	}

}
