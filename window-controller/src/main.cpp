#include <Arduino.h>
#include "kernel/scheduler.h"
#include "tasks/windowControllingTask.h"

#include "config.h"

Scheduler sched;

void setup() {

    Serial.begin(9600);

    sched.init(50);

    Task *windowControllingTask = new WindowControllingTask(SERVOMOTOR_PIN, POTENTIOMETER_PIN, BUTTON_PIN);
    windowControllingTask->init(150);
    sched.addTask(windowControllingTask);
}

void loop() {
    sched.schedule();
}
