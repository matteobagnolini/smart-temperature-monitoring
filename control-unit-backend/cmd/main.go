package main

import (
	"control-unit-backend/pkg/models"
	"control-unit-backend/pkg/mqtt"
	"control-unit-backend/pkg/serial"
)

const BROKER_ADDR = "broker.mqtt-dashboard.com:1883"
const TOPIC = "smart-temp/esp32/temp"

const SERIAL_PORT = "/dev/cu.usbserial-14120"
const BAUD_RATE = 9600

func main() {

	models.DataSampler.StartSampling() // start sampling subroutine

	mqtt.ConnectMQTT(BROKER_ADDR, TOPIC)
	go models.StartMQTTListener()

	serial.SerialConn, _ = serial.OpenSerial(SERIAL_PORT, BAUD_RATE)

	go serial.SerialConn.Read()

	go models.StartSerialListener()

	models.Tick()

}
