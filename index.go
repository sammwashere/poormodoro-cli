package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func main() {
	var (
		name     string
		duration string
	)

	flag.StringVar(&name, "n", "New timer", "Your timer name")
	flag.StringVar(&duration, "d", "30m", "Your timer duration")

	flag.Parse()

	fmt.Println("duration", duration)

	t1 := newTimer(name, duration)

	fmt.Printf("My new timer is called %s and it will run for %ds\n", t1.name, t1.duration)
}
