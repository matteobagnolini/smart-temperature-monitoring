#include "hardware/potentiometer.h"
#include <Arduino.h>

Potentiometer::Potentiometer(int pin) {
    this->pin = pin;
}

float Potentiometer::readValue() {
    int value = analogRead(pin);
    return value;
}
