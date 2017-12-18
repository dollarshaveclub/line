package line

import (
	"bytes"
	"fmt"
	"testing"
)

func newTestLine() (*Line, *bytes.Buffer) {
	return newTestLineFixes("", "", nil)
}

func newTestLineFixes(prefix, suffix string, formatter Formatter) (*Line, *bytes.Buffer) {
	buf := new(bytes.Buffer)
	return New(buf, prefix, suffix, formatter), buf
}

func newTestOutput() (*Output, *bytes.Buffer) {
	return newTestOutputFixes("", "", nil)
}

func newTestOutputFixes(prefix, suffix string, formatter Formatter) (*Output, *bytes.Buffer) {
	l, b := newTestLine()
	return NewOutput(l, prefix, suffix, formatter), b
}

func TestPrint(t *testing.T) {
	l, b := newTestLine()

	// Line Test
	l.Print("PRINT")
	expected := "PRINT"
	received := b.String()

	if expected != received {
		t.Errorf("Expected %q but received %q", expected, received)
	}

	output, b := newTestOutput()

	// Output Test
	output.Print("PRINT")
	expected = "PRINT"
	received = b.String()

	if expected != received {
		t.Errorf("Expected %q but received %q", expected, received)
	}
}

func TestPrintln(t *testing.T) {
	l, b := newTestLine()

	// Line Test
	l.Println("PRINT")
	expected := "PRINT\n"
	received := b.String()

	if expected != received {
		t.Errorf("Expected %q but received %q", expected, received)
	}

	output, b := newTestOutput()

	// Output Test
	output.Println("PRINT")
	expected = "PRINT\n"
	received = b.String()

	if expected != received {
		t.Errorf("Expected %q but received %q", expected, received)
	}
}

func TestPrintf(t *testing.T) {
	l, b := newTestLine()

	// Line Test
	l.Printf("%v %v", "test", "test")
	expected := "test test"
	received := b.String()

	if expected != received {
		t.Errorf("Expected %q but received %q", expected, received)
	}

	output, b := newTestOutput()

	// Output Test
	output.Printf("%v %v", "test", "test")
	expected = "test test"
	received = b.String()

	if expected != received {
		t.Errorf("Expected %q but received %q", expected, received)
	}
}

func TestColors(t *testing.T) {
	l, bl := newTestLine()
	o, bo := newTestOutput()

	l.Black("b")
	l.Red("r")
	l.Green("g")
	l.Yellow("y")
	l.Blue("bl")
	l.Magenta("m")
	l.Cyan("c")
	l.White("w")

	o.Black("b")
	o.Red("r")
	o.Green("g")
	o.Yellow("y")
	o.Blue("bl")
	o.Magenta("m")
	o.Cyan("c")
	o.White("w")

	expected := "\x1b[30mb\x1b[0m\x1b[31mr\x1b[0m\x1b[32mg\x1b[0m\x1b[33my\x1b[0m\x1b[34mbl\x1b[0m\x1b[35mm\x1b[0m\x1b[36mc\x1b[0m\x1b[37mw\x1b[0m"
	receivedLine := bl.String()
	receivedOutput := bo.String()

	if expected != receivedLine {
		t.Errorf("Expected %q but received %q", expected, receivedLine)
	}

	if expected != receivedOutput {
		t.Errorf("Expected %q but received %q", expected, receivedOutput)
	}
}

func TestPrefixSuffix(t *testing.T) {
	l, bl := newTestLineFixes("--> ", " <--", nil)
	o, bo := newTestOutputFixes("--> ", " <--", nil)

	l.Println("TEST")
	o.Println("TEST")

	expected := "--> TEST <--\n"
	receivedLine := bl.String()
	receivedOutput := bo.String()

	if expected != receivedLine {
		t.Errorf("Expected %q but received %q", expected, receivedLine)
	}

	if expected != receivedOutput {
		t.Errorf("Expected %q but received %q", expected, receivedOutput)
	}
}

func TestChainedPrefixSuffix(t *testing.T) {
	l, bl := newTestLine()
	o, bo := newTestOutput()

	l.Prefix("--> ").Println("TEST")
	l.Suffix(" <--").Println("TEST")

	o.Prefix("--> ").Println("TEST")
	o.Suffix(" <--").Println("TEST")

	expected := "--> TEST\nTEST <--\n"
	receivedLine := bl.String()
	receivedOutput := bo.String()

	if expected != receivedLine {
		t.Errorf("Expected %q but received %q", expected, receivedLine)
	}

	if expected != receivedOutput {
		t.Errorf("Expected %q but received %q", expected, receivedOutput)
	}
}

func TestSemantic(t *testing.T) {
	l, bl := newTestLine()
	o, bo := newTestOutput()

	l.Info("info\n")
	l.Error("error\n")
	l.Progress("progress\n")

	o.Info("info\n")
	o.Error("error\n")
	o.Progress("progress\n")

	expected := "\x1b[33minfo\n\x1b[0m\x1b[31merror\n\x1b[0m--> progress\n"
	receivedLine := bl.String()
	receivedOutput := bo.String()

	if expected != receivedLine {
		t.Errorf("Expected %q but received %q", expected, receivedLine)
	}

	if expected != receivedOutput {
		t.Errorf("Expected %q but received %q", expected, receivedOutput)
	}
}

type errWriter struct {
	err error
	n   int
}

func (e errWriter) Write(p []byte) (n int, err error) {
	return e.n, e.err
}

func TestErrN(t *testing.T) {
	expectedErr := fmt.Errorf("test error")
	expectedN := 10
	writer := errWriter{err: expectedErr, n: expectedN}
	l := New(writer, "", "", nil)

	o := l.Print("fail")

	receivedErr := o.Err()
	receivedN := o.N()

	if expectedErr != receivedErr {
		t.Errorf("Expected %q but received %q", expectedErr, receivedErr)
	}

	if expectedN != receivedN {
		t.Errorf("Expected %q but received %q", expectedN, receivedN)
	}
}
