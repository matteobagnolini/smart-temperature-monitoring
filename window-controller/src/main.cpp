#include <Arduino.h>
#include "kernel/scheduler.h"
#include "tasks/windowControllingTask.h"
#include "tasks/communicationsTask.h"
#include "controller/systemController.h"

#include "config.h"

Scheduler sched;
Controller contr;

void setup() {

    Serial.begin(9600);

    sched.init(50);

    contr.init();

    Task *windowControllingTask =
        new WindowControllingTask(&contr, SERVOMOTOR_PIN, POTENTIOMETER_PIN, BUTTON_PIN);
    windowControllingTask->init(150);
    sched.addTask(windowControllingTask);

    Task *commTask = new CommunicationsTask(&contr);
    commTask->init(200);
    sched.addTask(commTask);
}

void loop() {
    sched.schedule();
}
