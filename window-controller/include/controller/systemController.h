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
    int getCurrOpening(); // TODO: change this name in smtg like getOpeningFromControlUnit()
    void setStateManual();
    void setStateAutomatic();
    bool isStateManual();
    bool isStateAutomatic();

private:
    float currTemp;
    int currOpeningPerc;
    bool isCurrStateManual;
    bool isCurrStateAutomatic;

};
