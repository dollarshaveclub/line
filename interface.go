package line

// compile time check of Interface conformance
var (
	_ Interface = &Line{}
	_ Interface = &Output{}
)

// Interface is Line's interface
type Interface interface {

	// Prefix sets the prefix for the returned Output
	Prefix(prefix string) *Output

	// Suffix sets the suffix for the returned Output
	Suffix(suffix string) *Output

	// Format sets the Format for the returned Output
	Format(f Formatter) *Output

	// Print prints and returns an Output allowing for chaining
	Print(a ...interface{}) *Output

	// Println prints the arguments with a new line
	Println(a ...interface{}) *Output

	// Printf prints and returns an Output allowing for chaining
	Printf(format string, a ...interface{}) *Output

	// Info prints using formatting suitable for an info message
	Info(a ...interface{}) *Output

	// Progress prints with an "-->" prefix
	Progress(a ...interface{}) *Output

	// Error prints using formatting suitable for an error message
	Error(a ...interface{}) *Output

	// Black prints black text
	Black(a ...interface{}) *Output

	// Red prints red text
	Red(a ...interface{}) *Output

	// Green prints green text
	Green(a ...interface{}) *Output

	// Yellow prints yellow text
	Yellow(a ...interface{}) *Output

	// Blue prints blue text
	Blue(a ...interface{}) *Output

	// Magenta prints magenta text
	Magenta(a ...interface{}) *Output

	// Cyan prints cyan text
	Cyan(a ...interface{}) *Output

	// White prints white text
	White(a ...interface{}) *Output
}

// Formatter formats arguments to be suitable for printing
type Formatter interface {
	Sprint(a ...interface{}) string
}
