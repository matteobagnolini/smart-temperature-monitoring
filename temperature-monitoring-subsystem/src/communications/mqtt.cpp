#include "communications/mqtt.h"
#include <WiFi.h>

MQTT::MQTT(String clientId) {
    this->clientId = clientId;
}

void MQTT::init(const char *mqtt_server, int port, const char *topic) {
    this->mqtt_server = mqtt_server;
    this->topic = topic;
    WiFiClient espClient;
    client = PubSubClient(espClient);
    client.setServer(mqtt_server, port);
    client.setCallback(callback);
}

bool MQTT::isConnected() {
    return client.connected();
}

void MQTT::reconnect() {
    while (!client.connected()) {
        Serial.print("Attempting MQTT connection...");

        if (client.connect(clientId.c_str())) {
            Serial.println("connected");
            client.subscribe(topic);
        } else {
            Serial.print("failed, rc=");
            Serial.println(client.state());
            Serial.println(" try again in 5 seconds");
            delay(5000);
        }
    }
}

void MQTT::loop() {
    client.loop();
}

static void callback(char *topic, byte *payload, unsigned int length) {
    Serial.println(String("Message arrived on [") + topic + "] len: " + length );
    // TODO:
}
