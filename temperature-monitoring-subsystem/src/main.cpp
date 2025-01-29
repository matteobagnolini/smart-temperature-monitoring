#include "tasks/temperatureTask.h"
#include "communications/communications.h"
#include <WiFi.h>
#include "config.h"
#include <Arduino.h>

TemperatureTask *tempTask;
TaskHandle_t tempTaskHandle;

WiFiClient espClient;
PubSubClient client(espClient);

Communications *comms;

void tempTaskCode(void *argument);

void callback(char *topic, byte *payload, unsigned int length) {
    Serial.println(String("Message arrived on [") + topic + "] len: " + length );
}

void setup() {
    Serial.begin(115200);
    comms = new Communications(client);
    comms->init(callback);

    tempTask = new TemperatureTask(TEMP_PIN, REDLED_PIN, GREENLED_PIN);
    xTaskCreatePinnedToCore(tempTaskCode, "Temperature Task", 1000, NULL, 1, &tempTaskHandle, 1);

}

void loop() {
    comms->loop();
    Serial.println("looped");
    delay(500);
}

void tempTaskCode(void *argument) {
     //control period ecc

    for (;;) {
        tempTask->loop();
    }

}
