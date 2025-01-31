#include "tasks/temperatureTask.h"
#include "communications/communications.h"
#include <WiFi.h>
#include "config.h"
#include <Arduino.h>

TemperatureTask *tempTask;
TaskHandle_t tempTaskHandle;
void tempTaskCode(void *argument);

WiFiClient espClient;
PubSubClient client(espClient);
Communications *comms;
void callback(char *topic, byte *payload, unsigned int length);

void setup() {
    Serial.begin(115200);
    comms = new Communications(client);
    comms->init(callback);

    tempTask = new TemperatureTask(TEMP_PIN, REDLED_PIN, GREENLED_PIN);
    xTaskCreatePinnedToCore(tempTaskCode, "Temperature Task", 8000, NULL, 1, &tempTaskHandle, 1);
}

void loop() {
    comms->loop();
}

void tempTaskCode(void *argument) {
    TickType_t xLastWakeTime;
    const TickType_t xFrequency = tempTask->getPeriod() / portTICK_PERIOD_MS;
    for ( ;; ) {
        tempTask->loop();
        xTaskDelayUntil(&xLastWakeTime, xFrequency);
    }
}

void callback(char *topic, byte *payload, unsigned int length) {
    Serial.println(String("Message arrived on [") + topic + "] len: " + length );
}
