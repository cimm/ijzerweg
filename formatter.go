package main

import (
	"fmt"
	"github.com/cimm/ijzerweg/irail"
)

const (
	shortDate = "Mon Jan 2 15:04"
)

func departureFormat(c irail.Connection) string {
	t := c.Departure.Time
	return fmt.Sprintf("%v \033[37m%v\033[0m", t.Format("Mon Jan 2"), t.Format("15:04"))
}

func delayFormat(c irail.Connection) string {
	if c.Arrival.Delay == 0 {
		return ""
	}

	return fmt.Sprintf("\033[31m+%v'\033[0m", c.Arrival.Delay/60)
}

func transferFormat(c irail.Connection) string {
	switch len(c.Vias.Via) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%v transfer", len(c.Vias.Via))
	default:
		return fmt.Sprintf("%v transfers", len(c.Vias.Via))
	}
}
