#pragma once

#include "tasks/task.h"
#include "hardware/temp.h"
#include "hardware/led.h"

#define BASE_PERIOD 250


class TemperatureTask : public Task {

public:
    TemperatureTask(int tempPin, int redLedPin, int greenLedPin);
    void loop();
    void updatePeriod(int newPeriod);

private:
    int period;
    Temp *temp;
    Led *redLed;
    Led *greenLed;
    enum { PROBLEM, WORKING } state;

};
