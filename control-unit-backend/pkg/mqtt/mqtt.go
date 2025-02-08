package mqtt

import (
	"fmt"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var TempChannel = make(chan string)

var client MQTT.Client
var pubTopic string

func ConnectMQTT(broker string, subscribeTopic string, publishTopic string) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("control-unit-backend")
	opts.SetDefaultPublishHandler(defaultMessageHandler)

	client = MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	client.Subscribe(subscribeTopic, 1, temperatureMessageHandler)
	pubTopic = publishTopic

	fmt.Println("Connected to MQTT")
}

func SendFrequencyMsg(msg string) {
	sendMsg(pubTopic, msg)
}

func sendMsg(topic string, msg string) {
	client.Publish(topic, 1, false, msg)
}

func defaultMessageHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received msg: %s\n", msg.Payload())
}

func temperatureMessageHandler(client MQTT.Client, msg MQTT.Message) {
	temp := string(msg.Payload())
	TempChannel <- temp
}
