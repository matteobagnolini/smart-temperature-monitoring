#include "controller/systemController.h"

void Controller::init() {
    setStateAutomatic();
}

void Controller::setCurrTemp(float temp) {
    this->currTemp = temp;
}

float Controller::getCurrTemp() {
    return this->currTemp;
}

void Controller::setCurrOpening(int perc) {
    this->currOpeningPerc = perc;
}

int Controller::getCurrOpening() {
    return this->currOpeningPerc;
}

void Controller::setStateAutomatic() {
    this->isCurrStateAutomatic = true;
    this->isCurrStateManual = false;
    this->stateHasChanged = true;
}

void Controller::setStateManual() {
    this->isCurrStateManual = true;
    this->isCurrStateAutomatic = false;
    this->stateHasChanged = true;
}

bool Controller::isStateAutomatic() {
    return this->isCurrStateAutomatic;
}

bool Controller::isStateManual() {
    return this->isCurrStateManual;
}

bool Controller::hasStateChanged() {
    if (stateHasChanged) {
        stateHasChanged = false;
        return true;
    }
    return stateHasChanged;
}
