#pragma once

class WiFiHandler {

public:
    void init(const char *ssid, const char *password);
    void connect();
    bool isConnectionOk();

private:
    const char *ssid;
    const char *password;

};
