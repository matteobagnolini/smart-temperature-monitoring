package serial

import (
	"fmt"
)

const WINDOW_PREF = "wi"
const TEMP_PREF = "te"
const STATE_PREF = "st"
const STATE_AUTO = STATE_PREF + ":aut"
const STATE_MAN = STATE_PREF + ":man"

func ParseMsg(msg string) (msgType, msgValue string) {
	if len(msg) < 4 || msg[2] != ':' {
		return "", ""
	}

	pref := msg[0:2]
	value := msg[3:]

	switch pref {
	case WINDOW_PREF:
		return "window", value
	case TEMP_PREF:
		return "temperature", value
	case STATE_PREF:
		if msg == STATE_AUTO {
			value = "automatic"
		} else if msg == STATE_MAN {
			value = "manual"
		} else {
			value = ""
		}
		return "state", value
	default:
		return "", msg
	}
}

func WindowOpeningMsg(perc int) string {
	return fmt.Sprintf("%s:%d", WINDOW_PREF, perc)
}

func TemperatureMsg(t float32) string {
	return fmt.Sprintf("%s:%.2f", TEMP_PREF, t)
}
