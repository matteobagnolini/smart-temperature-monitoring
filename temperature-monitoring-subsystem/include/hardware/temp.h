#pragma once

#include "devices/tempDetector.h"

class Temp : public tempDetector {

public:
    Temp(int pin);
    float readTemp();

private:
    int pin;

};
