#pragma once

#include "devices/light.h"

class Led : public Light {

public:
    Led(int pin);
    void switchOn();
    void switchOff();

private:
    int pin;
    bool isOn;

};
