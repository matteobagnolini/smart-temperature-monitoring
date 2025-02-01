#include "hardware/lcd.h"
#include <Arduino.h>
#include <stdarg.h>

LCD::LCD() : lcd(LCD_I2C_ADDR, LCD_COLS, LCD_ROWS) {
    lcd.init();
    lcd.backlight();
}

void LCD::display(const char *msg) {
        lcd.clear();
        lcd.print(msg);
}

void LCD::displayOnLines(int numOfLines, ...) {
    lcd.clear();
    va_list args;
    va_start(args, numOfLines);
    for (int i = 0; i < numOfLines; i++) {
        lcd.setCursor(0, i);
        lcd.print(va_arg(args, const char*));
    }
    va_end(args);
}

void LCD::clear() {
    lcd.clear();
}

void LCD::turnDisplayOn() {
    lcd.setBacklight(HIGH);
}

void LCD::turnDisplayOff() {
    lcd.setBacklight(LOW);
}
