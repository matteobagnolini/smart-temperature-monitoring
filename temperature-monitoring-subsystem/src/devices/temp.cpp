#include "hardware/temp.h"
#include <Arduino.h>

Temp::Temp(int pin) {
    this->pin = pin;
    pinMode(pin, INPUT);
}

float Temp::readTemp() {
    int reading = analogRead(this->pin);
    float voltage = (reading) * (3.3 / 4095.0);
    float tempC = (voltage - 0.5) * 100;
    return tempC;
}
