#pragma once

#include "devices/window.h"
#include "ServoTimer2.h"

#define OPEN 180
#define CLOSE 90
#define REVERSE 0

class ServoMotor : public Window {

public:
    ServoMotor(int pin);
    void open(int percentage);  // 0% = 0 gr, 100% = 90 gr

private:
    int pin;
    ServoTimer2 motor;
    void setPosition(int angle);

};
