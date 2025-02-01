#pragma once
#include <Arduino.h>

class UserLCD {

public:
    virtual void display(const char *msg);
    virtual void displayOnLines(int numOfLines, ...) = 0;
    virtual void clear();
    virtual void turnDisplayOn();
    virtual void turnDisplayOff();

};
