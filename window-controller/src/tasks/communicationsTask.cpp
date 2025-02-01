#include "tasks/communicationsTask.h"
#include "kernel/msgService.h"

CommunicationsTask::CommunicationsTask(Controller *controller) {
    this->controller = controller;
}

void CommunicationsTask::init(int period) {
    Task::init(period);
}

int freeRam() {

  extern int __heap_start,*__brkval;

  int v;

  return (int)&v - (__brkval == 0  

    ? (int)&__heap_start : (int) __brkval);  

}

void display_freeram() {

    MsgService.sendMsg("- SRAM left: ");
    char *buf;
    buf = (char *)malloc(sizeof(*buf) * (10));
    sprintf(buf, "%d", freeRam());
    MsgService.sendMsg(buf);
    free(buf);
}

void CommunicationsTask::tick() {
    if (controller->isStateAutomatic()) {    // AUTOMATIC State
        if (controller->hasStateChanged()) {
            MsgService.sendMsg(STATE_AUTO);
        }
        Msg *msg = MsgService.receiveMsg();
        if (msg == NULL)
            return;
        const char *content = msg->getContent().c_str();
        if (strncmp(content, WINDOW_PREF, 2) == 0) {
            int perc = atoi(content + 3);
            controller->setCurrOpening(perc);
        }
        delete msg;

    } else if (controller->isStateManual()) {    // MANUAL State
        if (controller->hasStateChanged()) {
            MsgService.sendMsg(STATE_MAN);
        }
        Msg *msg = MsgService.receiveMsg();
        if (msg == NULL)
            return;
        const char *content = msg->getContent().c_str();
        if (strncmp(content, TEMP_PREF, 2) == 0) {
            controller->setCurrTemp(atoi(content + 3));
        }
        delete msg;
    } 
    display_freeram();
};
