#pragma once

#include "kernel/task.h"
#include "hardware/servoMotor.h"
#include "hardware/potentiometer.h"
#include "hardware/button.h"
#include "hardware/lcd.h"
#include "controller/systemController.h"
#include <Arduino.h>

class WindowControllingTask : public Task {

public:
    WindowControllingTask(Controller *controller, int motorPin, int potPin, int buttonPin);
    void init(int period);
    void tick();

private:
    Controller *controller;
    Window *window;
    Potentiometer* pot;
    Button* button;
    UserLCD *lcd;
    enum windowState { AUTOMATIC, MANUAL } state;

    void changeState();
    void setState(windowState s);

    void openWindowManual();
    void openWindowAutomatic();
    const char *prepareDisplayMsg();

};
