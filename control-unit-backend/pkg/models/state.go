package models

import (
	"sync"
)

const NORMAL_TEMP float32 = 15.0
const HOT_TEMP float32 = 17.0
const TOO_HOT_TEMP float32 = 18.0

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
	mu              sync.Mutex
}

// Singleton
var System = StateMachine{
	windPercOpening: 0,
	lastTemp:        0.0,
	sysState:        AUTOMATIC,
	tempState:       NORMAL,
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
