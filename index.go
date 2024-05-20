package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Timer struct {
	name     string
	duration int
}

func parseDurationText(text string) int {
	s := strings.ReplaceAll(text, " ", "")

	reg := regexp.MustCompile("[^0-9]+")

	durNoNum := reg.ReplaceAllString(text, "")

	dur, err := strconv.Atoi(durNoNum)

	if err != nil {
		return 0
	}

	// check if string contains minutes
	if strings.Contains(s, "m") {
		return dur * 60
	} else if strings.Contains(s, "h") {
		return dur * 60 * 60
	} else if strings.Contains(s, "s") {
		return dur
	}

	return 0
}

func newTimer(name string, duration string) *Timer {
	durationArr := strings.Split(duration, " ")

	fmt.Println(durationArr)

	completeDur := 0

	for i := range len(durationArr) {
		durText := durationArr[i]
		fmt.Printf("Parsing %s\n", durText)
		dur := parseDurationText(durText)

		completeDur += dur
	}

	t1 := Timer{name: name, duration: completeDur}

	return &t1
}

func poormodorTicker(done *chan bool, ticker *time.Ticker, startTime *time.Time) {
	for {
		select {
		case <-*done:
			return
		case t := <-ticker.C:
			diff := startTime.Sub(t)
			fmt.Printf("\r %s", diff.String())
		}
	}
}

func getInitialTime() {

}

func main() {
	var (
		name     string
		duration string
	)

	flag.StringVar(&name, "n", "New timer", "Your timer name")
	flag.StringVar(&duration, "d", "30m", "Your timer duration")

	flag.Parse()

	startTime := time.Now()

	t1 := newTimer(name, duration)

	ticker := time.NewTicker(500 * time.Millisecond)

	fmt.Printf("Timer %s will run for %ds\n", t1.name, t1.duration)

	done := make(chan bool)

	go poormodorTicker(&done, ticker, &startTime)

	time.Sleep(time.Duration(t1.duration) * time.Second)
	ticker.Stop()
	done <- true

	fmt.Printf("Timer %s has stopped\n", t1.name)
}
