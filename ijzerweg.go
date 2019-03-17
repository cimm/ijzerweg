package main

import (
	"flag"
	"fmt"
	"github.com/cimm/ijzerweg/irail"
	"log"
	"time"
)

const (
	maxInputLen = 50
)

func main() {
	fromPtr := flag.String("from", "", "Departure station name")
	langPtr := flag.String("lang", "nl", "Preferred languaged nl, fr, or en")
	timePtr := flag.String("time", "", "Departure date RFC3339 formatted")
	toPtr := flag.String("to", "", "Destination station name")
	flag.Parse()

	if len(*fromPtr) > maxInputLen || len(*toPtr) > maxInputLen {
		log.Fatal("err: departure or arrival station name too long")
	}

	time, err := time.Parse(time.RFC3339, *timePtr)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	connections := irail.FindConnections(*fromPtr, *toPtr, "departure", time, *langPtr)
	for _, c := range connections {
		fmt.Printf("%v\t%v %v\t%v\n", c.Id, departureFormat(c), delayFormat(c), c.Departure.Direction.Name)
		fmt.Printf("\t%v %v\t%v\n", c.Arrival.Time.Format(shortDate), delayFormat(c), c.Departure.Platform)
		fmt.Printf("\t%v'\t%v\n", c.Duration/60, transferFormat(c))
	}
}
