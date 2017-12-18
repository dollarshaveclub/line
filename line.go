package line

import (
	"fmt"

	"io"
	"sync"
)

// Line is an easy way to stylize terminal output
type Line struct {
	prefix    string
	suffix    string
	last      rune
	formatter Formatter
	w         io.Writer
	mu        sync.Mutex
}

// New creates a new instance of Line
func New(out io.Writer, prefix string, suffix string, formatter Formatter) *Line {
	return &Line{w: out, prefix: prefix, suffix: suffix, formatter: formatter, last: '\n'} // last was a newline so that the prefix will print on the first line
}

// Prefix sets the prefix for the returned Output
func (l *Line) Prefix(prefix string) *Output {
	return NewOutput(l, prefix, l.suffix, l.formatter)
}

// Suffix sets the suffix for the returned Output
func (l *Line) Suffix(suffix string) *Output {
	return NewOutput(l, l.prefix, suffix, l.formatter)
}

// Format sets the Format for the returned Output
func (l *Line) Format(f Formatter) *Output {
	return NewOutput(l, l.prefix, l.suffix, f)
}

// Print prints and returns an Output allowing for chaining
func (l *Line) Print(a ...interface{}) *Output {
	return NewOutput(l, l.prefix, l.suffix, l.formatter).Print(a...)
}

// Println prints the arguments with a new line
func (l *Line) Println(a ...interface{}) *Output {
	return NewOutput(l, l.prefix, l.suffix, l.formatter).Println(a...)
}

// Printf prints and returns an Output allowing for chaining
func (l *Line) Printf(format string, a ...interface{}) *Output {
	return NewOutput(l, l.prefix, l.suffix, l.formatter).Printf(format, a...)
}

// Info prints using formatting suitable for an info message
func (l *Line) Info(a ...interface{}) *Output {
	return l.Yellow().Print(a...)
}

// Progress prints with an "-->" prefix
func (l *Line) Progress(a ...interface{}) *Output {
	return l.Prefix("--> ").Print(a...)
}

// Error prints using formatting suitable for an error message
func (l *Line) Error(a ...interface{}) *Output {
	return l.Red().Print(a...)
}

// Black prints black text
func (l *Line) Black(a ...interface{}) *Output {
	return l.Format(BlackColor).Print(a...)
}

// Red prints red text
func (l *Line) Red(a ...interface{}) *Output {
	return l.Format(RedColor).Print(a...)
}

// Green prints green text
func (l *Line) Green(a ...interface{}) *Output {
	return l.Format(GreenColor).Print(a...)
}

// Yellow prints yellow text
func (l *Line) Yellow(a ...interface{}) *Output {
	return l.Format(YellowColor).Print(a...)
}

// Blue prints blue text
func (l *Line) Blue(a ...interface{}) *Output {
	return l.Format(BlueColor).Print(a...)
}

// Magenta prints magenta text
func (l *Line) Magenta(a ...interface{}) *Output {
	return l.Format(MagentaColor).Print(a...)
}

// Cyan prints cyan text
func (l *Line) Cyan(a ...interface{}) *Output {
	return l.Format(CyanColor).Print(a...)
}

// White prints white text
func (l *Line) White(a ...interface{}) *Output {
	return l.Format(WhiteColor).Print(a...)
}

func (l *Line) doFprint(formatter Formatter, out string) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	print := out
	if formatter != nil {
		print = formatter.Sprint(print)
	}

	c, err := fmt.Fprint(l.w, print)
	if len(out) > 0 {
		l.last = rune(out[len(out)-1])
	}

	return c, err
}

func (l *Line) getLast() rune {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.last
}
