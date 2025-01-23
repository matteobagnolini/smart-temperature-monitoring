#pragma once
#include <Arduino.h>

class UserLCD {

public:
    virtual void display(String msg);
    virtual void clear();
    virtual void turnDisplayOn();
    virtual void turnDisplayOff();

};
