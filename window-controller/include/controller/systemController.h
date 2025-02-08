#pragma once

/**
 * This class is the mediator between the window controlling task and the
 * communication task. It also represents a temporary snapshot of the current
 * state (temperature, window opening, tasks states) of the subsystem.
 */
class Controller {

public:
    void init();
    void setCurrTemp(float temp);
    float getCurrTemp();
    void setCurrOpening(int perc);
    int getCurrOpening();
    void setStateManual();
    void setStateAutomatic();
    bool isStateManual();
    bool isStateAutomatic();
    bool hasStateChanged(); // To check if the current state has changed in last update

private:
    float currTemp;
    int currOpeningPerc;
    bool isCurrStateManual;
    bool isCurrStateAutomatic;
    bool stateHasChanged;

};
