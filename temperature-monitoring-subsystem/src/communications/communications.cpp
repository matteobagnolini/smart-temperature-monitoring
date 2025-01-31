#include "communications/communications.h"
#include "WiFi.h"
#include "config.h"

void callback(char *topic, byte *payload, unsigned int length);

Communications::Communications(PubSubClient client) {
    this->client = client;
}

void Communications::init(std::function<void (char *, uint8_t *, unsigned int)> callback) {
    // WIFI connection
    wifi = WiFiHandler();
    wifi.init(SSID, PSWRD);
    wifi.connect();

    this->callback = callback;
    
    // MQTT connection
    client.setServer(MQTT_SERVER, SERVER_PORT);

    mqttReconnect();
}

void Communications::loop() {
    client.loop();      // To check for upcoming messages

    if (mqttProblem) {
        mqttReconnect();
    }
    if (wifiProblem) {
        wifiReconnect();
    }
}

bool Communications::isConnectionOk() {
    if (!client.connected()) {
        Serial.println("problem with mqtt");
        mqttProblem = true;
    }
    if (!wifi.isConnectionOk()) {
        Serial.println("problem with wifi");
        wifiProblem = true;
    }
    return !mqttProblem && !wifiProblem;
}

void Communications::sendMessage(const char *topic, const char *msg) {
    client.publish(topic, msg);
}

void Communications::mqttReconnect() {
    while (!client.connected()) {
        Serial.print("Attempting MQTT connection...");
        if (client.connect(CLIENT_ID)) {
            Serial.println("connected");
            client.subscribe(TOPIC_PERIOD);
            Serial.println("Subscribed");
        } else {
            Serial.print("failed, rc=");
            Serial.print(client.state());
            Serial.println(" try again in 5 seconds");
            delay(5000);
        }
    }
    client.setCallback(callback);
    mqttProblem = false;
}

void Communications::wifiReconnect() {
    wifi.connect();
    wifiProblem = false;
}