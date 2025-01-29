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
    
    // MQTT connection
    client.setServer(MQTT_SERVER, SERVER_PORT);

    while (!client.connected()) {
        Serial.print("Attempting MQTT connection...");
        if (client.connect(CLIENT_ID)) {
            Serial.println("connected");
            client.subscribe(TOPIC);
            Serial.println("Subscribed");
        } else {
            Serial.print("failed, rc=");
            Serial.print(client.state());
            Serial.println(" try again in 5 seconds");
            delay(5000);
        }
    }
    client.setCallback(callback);
}

void Communications::loop() {
    client.loop();      // To check for upcoming messages
}

bool Communications::isConnectionOk() {
    return client.connected() && wifi.isConnectionOk();
}

void Communications::sendMessage(const char *msg) {
    client.publish(TOPIC, msg);
}
