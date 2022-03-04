package clockface_svg_test

import (
	"testing"
	"time"

	clockface_svg "github.com/ichang0301/learn-golang/16_maths"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := clockface_svg.Point{X: 150, Y: 150 - 90}
	got := clockface_svg.SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := clockface_svg.Point{X: 150, Y: 150 + 90}
	got := clockface_svg.SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
