package main

import (
	"control-unit-backend/pkg/db"
	"control-unit-backend/pkg/http"
	"control-unit-backend/pkg/models"
	"control-unit-backend/pkg/mqtt"
	"control-unit-backend/pkg/serial"
)

const SERVER_IP_ADDR = ""
const WEB_PORT = "3333"

const BROKER_ADDR = "broker.mqtt-dashboard.com:1883"
const SUB_TOPIC = "smart-temp/esp32/temp"
const PUB_TOPIC = "smart-temp/esp32/period"

const SERIAL_PORT = "/dev/cu.usbserial-14120"
const BAUD_RATE = 9600

func main() {

	db.InitDb()

	models.DataSampler.StartSampling()
	http.StartHttpServer("", "3333")

	mqtt.ConnectMQTT(BROKER_ADDR, SUB_TOPIC, PUB_TOPIC)
	go models.StartMQTTListener()

	serial.StartSerial(SERIAL_PORT, BAUD_RATE)

	go serial.SerialConn.Read()

	go models.StartSerialListener()

	models.Tick()

}
