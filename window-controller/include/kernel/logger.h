#pragma once

#include <Arduino.h>

class LoggerService {

public:
    void log(const String &msg);

};

extern LoggerService Logger;

