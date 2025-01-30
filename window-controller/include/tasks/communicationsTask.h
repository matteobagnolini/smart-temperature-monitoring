#pragma once
#include "kernel/task.h"
#include "controller/systemController.h"

class CommunicationsTask : public Task {

public:
    CommunicationsTask(Controller *controller);
    void init(int period);
    void tick();

private:
    Controller *controller;

    void sendCurrentStates();
    void receiveUpdatedStates();

};
