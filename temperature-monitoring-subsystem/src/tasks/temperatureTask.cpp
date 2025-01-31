#include "tasks/temperatureTask.h"
#include "communications/communications.h"
#include "config.h"
#include <Arduino.h>

TemperatureTask::TemperatureTask(int tempPin, int redLedPin, int greenLedPin) {
    this->temp = new Temp(tempPin);
    this->redLed = new Led(redLedPin);
    this->greenLed = new Led(greenLedPin);

    state = PROBLEM;
    period = BASE_PERIOD;

    redLed->switchOff();
    greenLed->switchOff();
}

void TemperatureTask::loop() {
    bool connectionOk = comms->isConnectionOk();

    switch (state) {
        case PROBLEM:
            if (connectionOk) {
                redLed->switchOff();
                state = WORKING;
            }
            redLed->switchOn();
            greenLed->switchOff();
        break;

        case WORKING:
            if (!connectionOk) {
                greenLed->switchOff();
                state = PROBLEM;
            }
            greenLed->switchOn();
            redLed->switchOff();
            float currTemp = temp->readTemp();
            Serial.println(".");
            comms->sendMessage(TOPIC_TEMP, String(currTemp, 2).c_str());
        break;
    }
}

int TemperatureTask::getPeriod() {
    return period;
}

void TemperatureTask::updatePeriod(int newPeriod) {
    this->period = newPeriod;
}