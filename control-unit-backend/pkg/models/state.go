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
	TOO_HIT TemperatureState = "TOO_HOT"
)

type StateMachine struct {
	SysState  SystemState
	TempState TemperatureState
	mu        sync.Mutex
}

var System = StateMachine{
	SysState:  AUTOMATIC,
	TempState: NORMAL,
}

func (s *StateMachine) GetSysState() SystemState {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.SysState
}

func (s *StateMachine) GetTempState() TemperatureState {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.TempState
}
