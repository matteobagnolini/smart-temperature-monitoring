#include <Arduino.h>
#include <kernel/scheduler.h>

Scheduler sched;

void setup() {
    sched.init(50);
}

void loop() {
    sched.schedule();
}
