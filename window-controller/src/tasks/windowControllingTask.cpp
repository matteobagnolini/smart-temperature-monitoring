#include "tasks/windowControllingTask.h"
#include "hardware/lcd.h"
#include "controller/systemController.h"
#include <Arduino.h>

WindowControllingTask::WindowControllingTask(Controller *controller, int motorPin, int potPin, int buttonPin) {
    this->controller = controller;
    this->window = new ServoMotor(motorPin);
    this->pot = new Potentiometer(potPin);
    this->button = new Button(buttonPin);
    this->lcd = new LCD();
}

void WindowControllingTask::init(int period) {
    Task::init(period);
    this->state = AUTOMATIC;
}

void WindowControllingTask::tick() {
    if (button->isPressed())
        changeState();

    switch(state) {
        case AUTOMATIC:
            openWindowAutomatic();
        break;

        case MANUAL:
            openWindowManual();
        break;

    }
    displayMsg();
}

void WindowControllingTask::changeState() {
    if (state == AUTOMATIC)
        setState(MANUAL);
    else
        setState(AUTOMATIC);
}

void WindowControllingTask::setState(windowState s) {
    state = s;
    s == MANUAL ? controller->setStateManual() : controller->setStateAutomatic();
}

void WindowControllingTask::openWindowManual() {
    float potValue = pot->readValue();  // value between 0 to 1023
    int perc = map(potValue, 0, 1023, 0, 100);
    window->open(perc);

    controller->setCurrOpening(perc);
}

void WindowControllingTask::openWindowAutomatic() {
    int perc = controller->getCurrOpening();
    window->open(perc);
}

void WindowControllingTask::displayMsg() {
    int perc = controller->getCurrOpening();
    String windowLevelString = "Window: " + String(perc) + "%";
    String stateString;
    if (state == AUTOMATIC) {
        stateString = "State: Automatic";
        lcd->displayOnLines(
                            2,
                            windowLevelString.c_str(),
                            stateString.c_str()
        );
    } else if (state == MANUAL) {
        stateString = "State: Manual";
        String tempString = "Temperature: " + String(controller->getCurrTemp());
        lcd->displayOnLines(
                            3,
                            windowLevelString.c_str(),
                            stateString.c_str(),
                            tempString.c_str()
        );
    }
}
