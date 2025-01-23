#include "hardware/lcd.h"
#include <Arduino.h>

LCD::LCD() : lcd(LCD_I2C_ADDR, LCD_COLS, LCD_ROWS) {
    lcd.init();
    lcd.backlight();
    currentMsg = "";
}

void LCD::display(String msg) {
    // Msg is displayed only if it's different from the current one
    if (msg != currentMsg) {
        lcd.clear();
        lcd.print(msg);
        currentMsg = msg;
    } else { }
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
