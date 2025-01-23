#pragma once

class Timer {
    
public:  
    Timer();
    void setupFreq(int freq);  
    void setupPeriod(int period);  
    void waitForNextTick();

};
