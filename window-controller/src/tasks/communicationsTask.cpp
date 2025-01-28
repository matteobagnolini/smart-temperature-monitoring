#include "tasks/communicationsTask.h"
#include "kernel/msgService.h"

CommunicationsTask::CommunicationsTask(Controller controller) {
    this->controller = controller;
}

void CommunicationsTask::init(int period) {
    Task::init(period);
}

void CommunicationsTask::tick() {
    if (controller.isStateAutomatic()) {    // AUTOMATIC State

        // Read msg and set proper states using the controller
        Msg *msg = MsgService.receiveMsg();
        if (msg == NULL)
            return;
        if (msg->getContent().substring(0,3) == WINDOW_PREF) {  // Receiving window opening percentage
            controller.setCurrOpening(msg->getContent().substring(4,6).toFloat());
        }

    } else if (controller.isStateManual()) {    // MANUAL State
        Msg *msg = MsgService.receiveMsg();     // we want to know the temperature
        if (msg == NULL)
            return;
        if (msg->getContent().substring(0,3) == TEMP_PREF) {
            controller.setCurrTemp(msg->getContent().substring(4,6).toFloat());
        }
        // Also need to communicate to Control Unit our current state
    }
};
