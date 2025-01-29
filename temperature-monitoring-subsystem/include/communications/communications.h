#pragma once

#include <PubSubClient.h>
#include "wifi.h"

class Communications {

public:
    Communications(PubSubClient client);
    void init(std::function<void (char *, uint8_t *, unsigned int)>function);
    void loop();
    bool isConnectionOk();
    void sendMessage(const char *msg);

private:
    PubSubClient client;
    WiFiHandler wifi;

};

extern Communications *comms;
