#pragma once

class UserLCD {

public:
    virtual void display(const char *msg);
    virtual void clear();
    virtual void turnDisplayOn();
    virtual void turnDisplayOff();

};
