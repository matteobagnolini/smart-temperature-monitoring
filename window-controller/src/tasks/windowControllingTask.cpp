#include "tasks/windowControllingTask.h"
#include "hardware/lcd.h"
#include <Arduino.h>

WindowControllingTask::WindowControllingTask(int motorPin, int potPin, int buttonPin) {
    this->window = new ServoMotor(motorPin);
    this->pot = new Potentiometer(potPin);
    this->button = new Button(buttonPin);
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
            // lcd.display("MODE: Automatic\nTEMP:0.0");
            // set window based on received data from serial line
        break;

        case MANUAL:
            // lcd.display("MODE: Manual\nTEMP:0.0");
            float potValue = pot->readValue();  // value between 0 to 1023
            Serial.print("Potvalue: ");
            Serial.println(potValue);
            int perc = map(potValue, 0, 1023, 0, 100);
            Serial.print("Perc: ");
            Serial.println(perc);
            window->open(perc);
        break;

    }
}

void WindowControllingTask::changeState() {
    if (state == AUTOMATIC)
        state = MANUAL;
    else
        state = AUTOMATIC;
}