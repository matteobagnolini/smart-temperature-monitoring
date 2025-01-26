#pragma once
#include "kernel/task.h"
#include "controller/systemController.h"

class CommunicationsTask : Task {

public:
    void init(int period);
    void tick();

private:
    Controller controller;

    void sendCurrentStates();
    void receiveUpdatedStates();

};
