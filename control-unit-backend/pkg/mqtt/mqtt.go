package mqtt

import (
	"fmt"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// Usage of MQTT.Client:
// t := client.Publish("topic", qos, retained, msg)
// https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang@v1.5.0#section-readme

func ConnectMQTT(broker string) MQTT.Client {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("control-unit-backend")
	opts.SetDefaultPublishHandler(messageHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	return client
}

func messageHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received temperature: %s\n", msg.Payload())
}
