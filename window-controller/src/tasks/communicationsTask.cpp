#include "tasks/communicationsTask.h"
#include "kernel/msgService.h"

void CommunicationsTask::init(int period) {
    Task::init(period);
}

void CommunicationsTask::tick() {
    if (controller.isStateAutomatic()) {

        // Read msg and set proper states using the controller

    } else if (controller.isStateManual()) {
        Msg *msg = MsgService.receiveMsg();     // we want to know the temperature
        if (msg == NULL)
            return;
        if (msg->getContent().substring(0,3) == TEMP_PREF) {
            controller.setCurrTemp(msg->getContent().substring(4,6).toFloat());
        }
    }
};
