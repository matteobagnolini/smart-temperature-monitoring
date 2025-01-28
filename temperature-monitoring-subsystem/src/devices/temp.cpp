#include "hardware/temp.h"
#include <Arduino.h>

Temp::Temp(int pin) {
    this->pin = pin;
    pinMode(pin, INPUT);
}

float Temp::readTemp() {
    int reading = analogRead(this->pin);
    Serial.print("Voltage: ");
    float voltage = (reading) / 1023.0;
    Serial.println(voltage);
    float tempC = (voltage - 0.5) * 100;
    return tempC;
}
