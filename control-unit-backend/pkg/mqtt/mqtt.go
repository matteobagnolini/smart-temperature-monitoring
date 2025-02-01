package mqtt

import (
	"fmt"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var TempChannel = make(chan string)

var client MQTT.Client

func ConnectMQTT(broker string, topic string) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("control-unit-backend")
	opts.SetDefaultPublishHandler(defaultMessageHandler)

	client = MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	client.Subscribe(topic, 1, temperatureMessageHandler)

	fmt.Println("Connected to MQTT")
}

func defaultMessageHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received msg: %s\n", msg.Payload())
}

func temperatureMessageHandler(client MQTT.Client, msg MQTT.Message) {
	// fmt.Printf("Received temperature: %s from Topic\n", msg.Payload())
	temp := string(msg.Payload())
	// temp, _ := strconv.ParseFloat(string(msg.Payload()), 32)
	// models.DataSampler.AddData(float32(temp), time.Now().Format(time.RFC3339)) // Add data to sampler
	TempChannel <- temp
}

func SendMsg(topic string, msg string) {
	client.Publish(topic, 1, false, msg)
}

func SendFrequencyMsg(msg string) {
	SendMsg("smart-temp/esp32/period", msg)
}
