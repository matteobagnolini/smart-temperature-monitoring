#pragma once

#include <Arduino.h>
#include "msgPrefixes.h"

/* ======= Message class ======= */
class Msg {
  String content;

public:
  Msg(String content){
    this->content = content;
  }
  
  String getContent() const {
    return content;
  }
};

class Pattern {
public:
  virtual boolean match(const Msg& m) = 0;  
};

/* ======= Messages Handler ======= */
class MsgServiceClass {
    
public: 
  
  Msg* currentMsg;
  bool msgAvailable;

  void init();  

  bool isMsgAvailable();
  Msg* receiveMsg();

  bool isMsgAvailable(Pattern& pattern);

  /* note: message deallocation is responsibility of the client */
  Msg* receiveMsg(Pattern& pattern);
  
  void sendMsg(const char *msg);

};

extern MsgServiceClass MsgService;
