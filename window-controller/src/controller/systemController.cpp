#include "controller/systemController.h"

void Controller::setCurrTemp(float temp) {
    this->currTemp = temp;
}

float Controller::getCurrTemp() {
    return this->currTemp;
}

void Controller::setCurrOpening(float perc) {
    this->currOpeningPerc = perc;
}

float Controller::getCurrOpening() {
    return this->currOpeningPerc;
}

void Controller::setStateAutomatic() {
    this->isCurrStateAutomatic = true;
    this->isCurrStateManual = false;
}

void Controller::setStateManual() {
    this->isCurrStateManual = true;
    this->isCurrStateAutomatic = false;
}

bool Controller::isStateAutomatic() {
    return this->isCurrStateAutomatic;
}

bool Controller::isStateManual() {
    return this->isCurrStateManual;
}
