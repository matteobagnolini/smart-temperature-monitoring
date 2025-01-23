#pragma once

#include "devices/userLCD.h"
#include <LiquidCrystal_I2C.h>

#define LCD_I2C_ADDR 0x27
#define LCD_COLS 20
#define LCD_ROWS 4

class LCD : public UserLCD {

public:
    LCD();
    void display(String msg);
    void clear();
    void turnDisplayOn();
    void turnDisplayOff();

private:
    LiquidCrystal_I2C lcd;
    String currentMsg;

};

extern LCD lcd;
