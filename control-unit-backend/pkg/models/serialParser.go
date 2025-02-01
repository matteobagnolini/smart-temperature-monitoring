package models

import (
	"fmt"
	"strconv"
)

const WINDOW_PREF = "wi"
const TEMP_PREF = "te"
const STATE_PREF = "st"
const STATE_AUTO = "aut"
const STATE_MAN = "man"

func ParseMsg(msg string) {
	if len(msg) < 4 || msg[2] != ':' {
		fmt.Printf("could not parse: %v", msg)
		return
	}

	pref := msg[0:2]
	value := msg[3:6]

	switch pref {
	case WINDOW_PREF:
		perc, _ := strconv.Atoi(value)
		System.SetWindPercOpening(perc)
	case STATE_PREF:
		if value == STATE_AUTO {
			System.SetSysState(AUTOMATIC)
		} else if value == STATE_MAN {
			System.SetSysState(MANUAL)
		} else {
			fmt.Printf("could not find state value: %v", value)
		}
	default:
		fmt.Printf("could not find cmd: %v", msg)
	}
}

func WindowOpeningMsg(perc int) string {
	return fmt.Sprintf("%s:%d", WINDOW_PREF, perc)
}

func TemperatureMsg(t float32) string {
	return fmt.Sprintf("%s:%.2f", TEMP_PREF, t)
}
