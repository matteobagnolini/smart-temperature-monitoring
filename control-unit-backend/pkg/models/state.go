package models

import (
	"sync"
)

const NORMAL_TEMP int = 18
const HOT_TEMP int = 21
const TOO_HOT_TEMP int = 24

type SystemState string

const (
	AUTOMATIC SystemState = "AUTOMATIC"
	MANUAL    SystemState = "MANUAL"
)

type TemperatureState string

const (
	NORMAL  TemperatureState = "NORMAL"
	HOT     TemperatureState = "HOT"
	TOO_HOT TemperatureState = "TOO_HOT"
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
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.sysState
}

func (s *StateMachine) TempState() TemperatureState {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.tempState
}

func (s *StateMachine) LastTemp() float32 {
	s.mu.Lock()
	defer s.mu.Lock()
	return s.lastTemp
}

func (s *StateMachine) WindowPercOpening() int {
	s.mu.Lock()
	defer s.mu.Lock()
	return s.windPercOpening
}
