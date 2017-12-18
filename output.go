package line

import (
	"fmt"
	"strings"
)

// Output represents an output
type Output struct {
	l         *Line
	prefix    string
	suffix    string
	formatter Formatter
	err       error
	n         int
}

// Err is the error that occurred in the Print call that produced this output
func (o *Output) Err() error {
	return o.err
}

// N number of bytes written on the Print call that produced this output
func (o *Output) N() int {
	return o.n
}

// NewOutput creates a new output
func NewOutput(l *Line, prefix, suffix string, formatter Formatter) *Output {
	return &Output{l: l, prefix: prefix, suffix: suffix, formatter: formatter}
}

// Prefix sets the prefix for the returned Output
func (o *Output) Prefix(prefix string) *Output {
	return NewOutput(o.l, prefix, o.suffix, o.formatter)
}

// Suffix sets the suffix for the returned Output
func (o *Output) Suffix(suffix string) *Output {
	return NewOutput(o.l, o.prefix, suffix, o.formatter)
}

// Format sets the Format for the returned Output
func (o *Output) Format(f Formatter) *Output {
	return NewOutput(o.l, o.prefix, o.suffix, f)
}

// Print prints and returns an Output allowing for chaining
func (o *Output) Print(a ...interface{}) *Output {
	n, err := o.doPrint(a...)
	output := &Output{l: o.l, prefix: o.prefix, suffix: o.suffix, formatter: o.formatter, err: err, n: n}
	return output
}

// Println prints the arguments with a new line
func (o *Output) Println(a ...interface{}) *Output {
	a = append(a, "\n")
	n, err := o.doPrint(a...)
	output := &Output{l: o.l, prefix: o.prefix, suffix: o.suffix, formatter: o.formatter, err: err, n: n}
	return output
}

// Printf prints and returns an Output allowing for chaining
func (o *Output) Printf(format string, a ...interface{}) *Output {
	n, err := o.doPrint(fmt.Sprintf(format, a...))
	output := &Output{l: o.l, prefix: o.prefix, suffix: o.suffix, formatter: o.formatter, err: err, n: n}
	return output
}

// Info prints using formatting suitable for an info message
func (o *Output) Info(a ...interface{}) *Output {
	return o.Yellow().Print(a...)
}

// Progress prints with an "-->" prefix
func (o *Output) Progress(a ...interface{}) *Output {
	return o.Prefix("--> ").Print(a...)
}

// Error prints using formatting suitable for an error message
func (o *Output) Error(a ...interface{}) *Output {
	return o.Red().Print(a...)
}

// Black prints black text
func (o *Output) Black(a ...interface{}) *Output {
	return o.Format(BlackColor).Print(a...)
}

// Red prints red text
func (o *Output) Red(a ...interface{}) *Output {
	return o.Format(RedColor).Print(a...)
}

// Green prints green text
func (o *Output) Green(a ...interface{}) *Output {
	return o.Format(GreenColor).Print(a...)
}

// Yellow prints yellow text
func (o *Output) Yellow(a ...interface{}) *Output {
	return o.Format(YellowColor).Print(a...)
}

// Blue prints blue text
func (o *Output) Blue(a ...interface{}) *Output {
	return o.Format(BlueColor).Print(a...)
}

// Magenta prints magenta text
func (o *Output) Magenta(a ...interface{}) *Output {
	return o.Format(MagentaColor).Print(a...)
}

// Cyan prints cyan text
func (o *Output) Cyan(a ...interface{}) *Output {
	return o.Format(CyanColor).Print(a...)
}

// White prints white text
func (o *Output) White(a ...interface{}) *Output {
	return o.Format(WhiteColor).Print(a...)
}
func (o *Output) doPrint(a ...interface{}) (int, error) {
	out := fmt.Sprint(a...)
	if len(out) == 0 {
		return 0, nil
	}

	skipFirst := o.l.getLast() != '\n'
	skipLast := out[len(out)-1] == '\n'
	out = o.insertFixes(skipFirst, skipLast, out)

	return o.l.doFprint(o.formatter, out)
}

func (o *Output) insertFixes(skipFirstPrefix bool, skipLastPrefix bool, output string) string {
	lines := strings.Split(output, "\n")
	if len(lines) == 1 { // no newlines
		if skipFirstPrefix {
			return output
		}

		return fmt.Sprintf("%v%v", o.prefix, output)
	}

	fixes := []string{}
	for y, line := range lines {
		out := ""
		if (y == 0 && !skipFirstPrefix) || (y == len(lines)-1 && !skipLastPrefix) || (y > 0 && y < len(lines)-1) {
			out = o.prefix
		}

		out = fmt.Sprintf("%s%s", out, line)

		if y != len(lines)-1 {
			out = fmt.Sprintf("%s%s", out, o.suffix)
		}

		fixes = append(fixes, out)
	}

	return strings.Join(fixes, "\n")
}
