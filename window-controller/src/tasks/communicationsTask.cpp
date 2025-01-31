#include "tasks/communicationsTask.h"
#include "kernel/msgService.h"

CommunicationsTask::CommunicationsTask(Controller *controller) {
    this->controller = controller;
}

void CommunicationsTask::init(int period) {
    Task::init(period);
}

void CommunicationsTask::tick() {
    if (controller->isStateAutomatic()) {    // AUTOMATIC State
        if (controller->hasStateChanged()) {
            MsgService.sendMsg(STATE_AUTO);
        }
        Msg *msg = MsgService.receiveMsg();
        if (msg == NULL)
            return;
        if (msg->getContent().substring(0,2) == WINDOW_PREF) {
            int perc = msg->getContent().substring(3,6).toInt();
            controller->setCurrOpening(perc);
        }

    } else if (controller->isStateManual()) {    // MANUAL State
        if (controller->hasStateChanged()) {
            MsgService.sendMsg(STATE_MAN);
        }
        Msg *msg = MsgService.receiveMsg();
        if (msg == NULL)
            return;
        if (msg->getContent().substring(0,2) == TEMP_PREF) {
            controller->setCurrTemp(msg->getContent().substring(3,6).toFloat());
        }
    }
};
