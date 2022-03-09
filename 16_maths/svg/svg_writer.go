package svg

import (
	"fmt"
	"io"
	"time"

	clockface_svg "github.com/ichang0301/learn-golang/16_maths"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCentreX     = 150
	clockCentreY     = 150
	svgStart         = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
    width="100%"
    height="100%"
    viewBox="0 0 300 300"
    version="2.0">`
	bezel  = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
	svgEnd = `</svg>`
)

//SVGWriter writes an SVG representation of an analogue clock, showing the time t, to the writer w
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

// secondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func secondHand(w io.Writer, t time.Time) {
	p := makeHand(clockface_svg.SecondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="%d" y1="%d" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, clockCentreX, clockCentreY, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(clockface_svg.MinuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, `<line x1="%d" y1="%d" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, clockCentreX, clockCentreY, p.X, p.Y)
}

func hourHand(w io.Writer, t time.Time) {
	p := makeHand(clockface_svg.HourHandPoint(t), hourHandLength)
	fmt.Fprintf(w, `<line x1="%d" y1="%d" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, clockCentreX, clockCentreY, p.X, p.Y)
}

func makeHand(p clockface_svg.Point, length float64) clockface_svg.Point {
	p = clockface_svg.Point{X: p.X * length, Y: p.Y * length}                // scale
	p = clockface_svg.Point{X: p.X, Y: -p.Y}                                 // flip
	return clockface_svg.Point{X: p.X + clockCentreX, Y: p.Y + clockCentreY} // translate
}
