#include "tasks/temperatureTask.h"
#include "communications/communications.h"
#include <Arduino.h>

TemperatureTask::TemperatureTask(int tempPin, int redLedPin, int greenLedPin) {
    this->temp = new Temp(tempPin);
    this->redLed = new Led(redLedPin);
    this->greenLed = new Led(greenLedPin);

    state = PROBLEM;
    period = BASE_PERIOD;
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
        break;

        case WORKING:
            if (!connectionOk) {
                greenLed->switchOff();
                state = PROBLEM;
            }
            greenLed->switchOn();
            float currTemp = temp->readTemp();
            // send currTemp;
        break;
    }
}