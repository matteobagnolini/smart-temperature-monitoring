#pragma once

#include "devices/pushDevice.h"

class Button : public PushDevice {

public:
    Button(int pin);
    bool isPressed();

private:
    int pin;

};
