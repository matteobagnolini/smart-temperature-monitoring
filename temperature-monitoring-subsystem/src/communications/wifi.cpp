#include "communications/wifi.h"
#include "WiFi.h"

void WiFiHandler::init(const char *ssid, const char *password) {
    this->ssid = ssid;
    this->password = password;
    WiFi.mode(WIFI_STA);
    WiFi.disconnect();
    delay(100);
}

void WiFiHandler::connect() {
    WiFi.begin(ssid, password);
    while (WiFi.status() != WL_CONNECTED) {
        delay(500);
        Serial.println(".");
    }
    Serial.print("Connected to WiFi network with IP Addres: ");
    Serial.println(WiFi.localIP());
}

bool WiFiHandler::isConnectionOk() {
    return WiFi.status() == WL_CONNECTED;
}
