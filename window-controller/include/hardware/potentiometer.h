#pragma once

#include "devices/analogInputDevice.h"

class Potentiometer : public AnalogInputDevice {

public:
    Potentiometer(int pin);
    float readValue();

private:
    int pin;

};
