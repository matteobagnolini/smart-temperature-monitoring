#pragma once

#include "kernel/task.h"
#include "hardware/servoMotor.h"
#include "hardware/potentiometer.h"
#include "hardware/button.h"

class WindowControllingTask : public Task {

public:
    WindowControllingTask(int pin);
    void init(int period);
    void tick();

private:
    int pin;
    Window *window;
    Potentiometer* pot;
    Button* button;
    enum { AUTOMATIC, MANUAL } state;

};
