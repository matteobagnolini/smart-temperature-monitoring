#include "kernel/logger.h"
#include "kernel/msgService.h"

void LoggerService::log(const String &msg) {
    MsgService.sendMsg("lo:" + msg);
}
