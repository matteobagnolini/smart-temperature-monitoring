#include <Arduino.h>
#include "kernel/scheduler.h"
#include "tasks/windowControllingTask.h"
#include "controller/systemController.h"

#include "config.h"

Scheduler sched;
Controller contr;

void setup() {

    Serial.begin(9600);

    sched.init(50);

    Task *windowControllingTask =
        new WindowControllingTask(contr, SERVOMOTOR_PIN, POTENTIOMETER_PIN, BUTTON_PIN);
    windowControllingTask->init(150);
    sched.addTask(windowControllingTask);
}

void loop() {
    sched.schedule();
}
