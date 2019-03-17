package irail

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Connections struct {
	ConnectionList []Connection `json:"connection"`
}

type Connection struct {
	Id        string    `json:"id"`
	Departure Departure `json:"departure"`
	Arrival   Arrival   `json:"arrival"`
	Duration  int       `json:"duration"`
	Vias      struct {
		Via []Via `json:"via"`
	} `json:"vias"`
}

type Departure struct {
	Delay     int       `json:"delay"`
	Platform  string    `json:"platform"`
	Time      time.Time `json:"time"`
	Direction struct {
		Name string `json:"name"`
	} `json:"direction"`
}

type Arrival struct {
	Delay     int       `json:"delay"`
	Platform  string    `json:"platform"`
	Time      time.Time `json:"time"`
	Direction struct {
		Name string `json:"name"`
	} `json:"direction"`
}

type Via struct {
	Id string `json:"id"`
}

func (c *Connection) UnmarshalJSON(data []byte) error {
	type Alias Connection
	aux := &struct {
		Duration string `json:"duration"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	i, err := strconv.Atoi(aux.Duration)
	if err != nil {
		return err
	}
	c.Duration = i
	return nil
}

func (d *Departure) UnmarshalJSON(data []byte) error {
	type Alias Departure
	aux := &struct {
		Delay string `json:"delay"`
		Time  string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	timestamp, err := strconv.ParseInt(aux.Time, 10, 64)
	if err != nil {
		return err
	}
	d.Time = time.Unix(timestamp, 0)

	d.Delay, err = strconv.Atoi(aux.Delay)
	if err != nil {
		return err
	}
	return nil
}

func (a *Arrival) UnmarshalJSON(data []byte) error {
	type Alias Departure
	aux := &struct {
		Delay string `json:"delay"`
		Time  string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	timestamp, err := strconv.ParseInt(aux.Time, 10, 64)
	if err != nil {
		return err
	}
	a.Time = time.Unix(timestamp, 0)

	a.Delay, err = strconv.Atoi(aux.Delay)
	if err != nil {
		return err
	}
	return nil
}

// Finds the connections between two stations. The `timeSel` attribute
// can be "departure" or "arrival" and defines the results will be by
// departure or arrival `time`.
func FindConnections(fromStation string, toStation string, timeSel string, time time.Time, lang string) []Connection {
	params := map[string]string{}
	params["from"] = fromStation
	params["to"] = toStation
	params["timesel"] = timeSel
	params["date"] = toIrailDate(time)
	params["time"] = toIrailTime(time)
	params["lang"] = lang

	body := Fetch("/connections", params)
	connections := Connections{}
	err := json.Unmarshal(body, &connections)
	if err != nil {
		log.Fatal(err)
	}
	return connections.ConnectionList
}

func toIrailTime(t time.Time) string {
	h, m, _ := t.Clock()
	return fmt.Sprintf("%02d%02d", h, m)
}

func toIrailDate(t time.Time) string {
	y := t.Year() % 100 // only need the last 2 digits
	return fmt.Sprintf("%02d%02d%02d", t.Day(), t.Month(), y)
}
