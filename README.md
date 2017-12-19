# line   &nbsp;<img src="doc/images/line-v.png" width="25" height="25"/>

line is an easy to use package for stylizing terminal output. line focuses on usability via chaining and, consequently, is quite flexible. line also boasts compatibility with the popular [Color](github.com/faith/color) package.

## Install

```bash
go get github.com/dollarshaveclub/line
```

## Usage

### Simple

```go
package main

import "github.com/dollarshaveclub/line"

func main() {
    line.Red().Print("Hello ").Green("World").Blue().Println("!!!")
}
```

### Prefix / Suffix

```go
package main

import "github.com/dollarshaveclub/line"

func main() {
	line.Prefix("--> ").Suffix(" <---").Println("Nice to meet you!").Println("And you too!")
}
```

### Complex

```go
package main

import "github.com/dollarshaveclub/line"
import "github.com/fatih/color"

func main() {
	output := line.White()
	output.Println("Welcome! Here is a list:")

	li := output.Prefix("--> ").Red()
	li.Println("one").Println("two").Println("sub")

	subli := li.Prefix("  --> ").Green()
	subli.Println("a").Println("b")

	output.Println().Println()

	boldgreen := color.New(color.Bold, color.FgMagenta)
	output.Format(boldgreen).Println("Have a nice day!")
}
```
