#include "hardware/servoMotor.h"
#include <Arduino.h>

ServoMotor::ServoMotor(int pin) {
    this->pin = pin;
    motor.attach(pin);
}

void ServoMotor::open(int percentage) {
    float angle = (180.0 * percentage) / 100;
    setPosition(angle);
}

void ServoMotor::setPosition(int angle) {
    float coeff = (2250.0-750.0)/180;
    motor.write(750 + angle*coeff);
}
