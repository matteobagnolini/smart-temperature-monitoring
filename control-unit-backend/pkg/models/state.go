package models

import (
	"sync"
)

const NORMAL_TEMP float32 = 18.0
const HOT_TEMP float32 = 20.5
const TOO_HOT_TEMP float32 = 21.5

const TOO_HOT_MAX_TIME_S int = 10

type SystemState string

const (
	AUTOMATIC        SystemState = "AUTOMATIC"
	MANUAL           SystemState = "MANUAL"
	DASHBOARD_MANUAL SystemState = "DASHBOARD_MANUAL"
)

type TemperatureState string

const (
	NORMAL  TemperatureState = "NORMAL"
	HOT     TemperatureState = "HOT"
	TOO_HOT TemperatureState = "TOO_HOT"
	ALARM   TemperatureState = "ALARM"
)

type StateMachine struct {
	windPercOpening int
	lastTemp        float32
	sysState        SystemState
	tempState       TemperatureState
	alarmOk         bool
	mu              sync.Mutex
}

var System = StateMachine{
	windPercOpening: 0,
	lastTemp:        0.0,
	sysState:        AUTOMATIC,
	tempState:       NORMAL,
	alarmOk:         true,
}

/* ====== Setters & Getters ======*/
func (s *StateMachine) SetWindPercOpening(perc int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.windPercOpening = perc
}

func (s *StateMachine) SetLastTemp(t float32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastTemp = t
}

func (s *StateMachine) SetSysState(state SystemState) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sysState = state
}

func (s *StateMachine) SetTempState(state TemperatureState) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tempState = state
}

func (s *StateMachine) SysState() SystemState {
	return s.sysState
}

func (s *StateMachine) TempState() TemperatureState {
	return s.tempState
}

func (s *StateMachine) LastTemp() float32 {
	return s.lastTemp
}

func (s *StateMachine) WindowPercOpening() int {
	return s.windPercOpening
}

func (s *StateMachine) ResolveAlarm() {
	if s.tempState == TemperatureState(ALARM) && !s.alarmOk {
		s.alarmOk = true
	}
}
