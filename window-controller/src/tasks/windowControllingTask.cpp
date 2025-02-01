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
    const  char *lcdMsg = prepareDisplayMsg();

    switch(state) {
        case AUTOMATIC:
            lcd->display(lcdMsg);
            openWindowAutomatic();
        break;

        case MANUAL:
            lcd->display(lcdMsg);
            openWindowManual();
        break;

    }
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
    float perc = controller->getCurrOpening();
    window->open(perc);
}

const char *WindowControllingTask::prepareDisplayMsg() {
    float perc = controller->getCurrOpening();
    const char *msg = "Window Level";
    // TODO:
    
    return msg;
}
