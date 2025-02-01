package models

import (
	"control-unit-backend/pkg/mqtt"
	"control-unit-backend/pkg/serial"
	"fmt"
	"strconv"
	"time"
)

const WINDOW_CLOSE = 0
const WINDOW_OPEN = 100

const NORMAL_PERIOD = 1000
const HOT_PERIOD = 800
const TOO_HOT_PERIOD = 500

func StartSerialListener() {
	go func() {
		for msg := range serial.SerialChannel {
			ParseMsg(msg)
		}
	}()
}

func StartMQTTListener() {
	go func() {
		for msg := range mqtt.TempChannel {
			processMQTTMessage(msg)
		}
	}()
}

func processMQTTMessage(msg string) {
	temp, _ := strconv.ParseFloat(string(msg), 32)
	DataSampler.AddData(float32(temp), time.Now().Format(time.RFC3339)) // Add data to sampler
}

func Tick() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for range ticker.C {
		System.SetLastTemp(DataSampler.GetLastData().Temp)
		switch System.sysState {
		case AUTOMATIC:
			handleAutomatic()
		case MANUAL:
			handleManual()
		}
	}
}

func handleManual() {
	serial.SerialConn.Write(TemperatureMsg(System.lastTemp))
}

var tooHotStartTime time.Time

func handleAutomatic() {
	switch System.tempState {
	case NORMAL:
		serial.SerialConn.Write(WindowOpeningMsg(WINDOW_CLOSE))

		if System.lastTemp > HOT_TEMP && System.lastTemp < TOO_HOT_TEMP {
			System.SetTempState(HOT)
			mqtt.SendFrequencyMsg(strconv.Itoa(HOT_PERIOD))
			fmt.Println("HOT")
		}
	case HOT:
		serial.SerialConn.Write(WindowOpeningMsg(computeOpeningWindow(System.lastTemp)))
		if System.lastTemp > TOO_HOT_TEMP {
			fmt.Println("TOO_HOT")
			System.SetTempState(TOO_HOT)
			mqtt.SendFrequencyMsg(strconv.Itoa(TOO_HOT_PERIOD))
			tooHotStartTime = time.Now()
		}
		if System.lastTemp < HOT_TEMP {
			fmt.Println("NORMAL")
			System.SetTempState(NORMAL)
			mqtt.SendFrequencyMsg(strconv.Itoa(NORMAL_PERIOD))
			serial.SerialConn.Write(WindowOpeningMsg(WINDOW_CLOSE))
		}
	case TOO_HOT:
		if System.lastTemp < TOO_HOT_TEMP {
			fmt.Println("HOT")
			System.SetTempState(HOT)
			mqtt.SendFrequencyMsg(strconv.Itoa(HOT_PERIOD))
		}
		if time.Since(tooHotStartTime) > time.Duration(TOO_HOT_MAX_TIME_S)*time.Second {
			System.SetTempState(ALARM)
		}
	case ALARM:
		fmt.Println("ALARM")
	}
}

func computeOpeningWindow(temp float32) int {
	// val := int((99*(temp-TOO_HOT_TEMP))/(HOT_TEMP-TOO_HOT_TEMP) + 1)
	return 50
}
