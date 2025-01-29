#pragma once

#include <PubSubClient.h>

class MQTT {

public:
    MQTT(String clientId);
    void init(const char *mqtt_server, int port, const char *topic);
    bool isConnected();
    void reconnect();
    void loop();
    void publish(String msg);

private:
    const char *mqtt_server;
    const char *topic;
    PubSubClient client;
    String clientId;

    static void callback(char* topic, byte* payload, unsigned int length);

};
