package line

import (
	"fmt"
	"strconv"
)

var (
	// BlackColor is a Formatter for black
	BlackColor = Color{Code: black}

	// RedColor is a Formatter for red
	RedColor = Color{Code: red}

	// GreenColor is a Formatter for green
	GreenColor = Color{Code: green}

	// YellowColor is a Formatter for yellow
	YellowColor = Color{Code: yellow}

	// BlueColor is a Formatter for blue
	BlueColor = Color{Code: blue}

	// MagentaColor is a Formatter for magenta
	MagentaColor = Color{Code: magenta}

	// CyanColor is a Formatter for cyan
	CyanColor = Color{Code: cyan}

	// WhiteColor is a Formatter for white
	WhiteColor = Color{Code: white}
)

const escape = "\x1b"
const reset = "0"

// ansi escape sequence color constants
const (
	black   = 30
	red     = 31
	green   = 32
	yellow  = 33
	blue    = 34
	magenta = 35
	cyan    = 36
	white   = 37
)

// Color formats with a given color
type Color struct {
	Code int
}

// Sprint wraps fmt.Sprint with ANSI control codes for a color
func (c Color) Sprint(a ...interface{}) string {
	str := fmt.Sprint(a...)
	return c.enable() + str + c.disable()
}
func (c Color) enable() string {
	return fmt.Sprintf("%s[%sm", escape, c.seq())
}

func (c Color) disable() string {
	return fmt.Sprintf("%s[%sm", escape, reset)
}

func (c Color) seq() string {
	return strconv.Itoa(c.Code)
}
