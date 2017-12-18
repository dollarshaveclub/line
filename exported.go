package line

import "os"

var (
	// Std is a os.Stdout backed Line
	Std = New(os.Stdout, "", "", nil)
)

// Prefix sets the prefix for the returned Output
func Prefix(prefix string) *Output {
	return NewOutput(Std, prefix, Std.suffix, Std.formatter)
}

// Suffix sets the suffix for the returned Output
func Suffix(suffix string) *Output {
	return NewOutput(Std, Std.prefix, suffix, Std.formatter)
}

// Format sets the Format for the returned Output
func Format(f Formatter) *Output {
	return NewOutput(Std, Std.prefix, Std.suffix, f)
}

// Print prints and returns an Output allowing for chaining
func Print(a ...interface{}) *Output {
	return NewOutput(Std, Std.prefix, Std.suffix, Std.formatter).Print(a...)
}

// Println prints the arguments with a new line
func Println(a ...interface{}) *Output {
	return NewOutput(Std, Std.prefix, Std.suffix, Std.formatter).Println(a...)
}

// Printf prints and returns an Output allowing for chaining
func Printf(format string, a ...interface{}) *Output {
	return NewOutput(Std, Std.prefix, Std.suffix, Std.formatter).Printf(format, a...)
}

// Info prints using formatting suitable for an info message
func Info(a ...interface{}) *Output {
	return Std.Black().Print(a...)
}

// Progress prints with an "-->" prefix
func Progress(a ...interface{}) *Output {
	return Std.Prefix("--> ").Print(a...)
}

// Error prints using formatting suitable for an error message
func Error(a ...interface{}) *Output {
	return Std.Red().Print(a...)
}

// Black prints black text
func Black(a ...interface{}) *Output {
	return Std.Format(BlackColor).Print(a...)
}

// Red prints red text
func Red(a ...interface{}) *Output {
	return Std.Format(RedColor).Print(a...)
}

// Green prints green text
func Green(a ...interface{}) *Output {
	return Std.Format(GreenColor).Print(a...)
}

// Yellow prints yellow text
func Yellow(a ...interface{}) *Output {
	return Std.Format(YellowColor).Print(a...)
}

// Blue prints blue text
func Blue(a ...interface{}) *Output {
	return Std.Format(BlueColor).Print(a...)
}

// Magenta prints magenta text
func Magenta(a ...interface{}) *Output {
	return Std.Format(MagentaColor).Print(a...)
}

// Cyan prints cyan text
func Cyan(a ...interface{}) *Output {
	return Std.Format(CyanColor).Print(a...)
}

// White prints white text
func White(a ...interface{}) *Output {
	return Std.Format(WhiteColor).Print(a...)
}
