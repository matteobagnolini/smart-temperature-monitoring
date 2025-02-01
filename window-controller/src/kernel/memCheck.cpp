#include "kernel/memCheck.h"
#include "kernel/msgService.h"

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
