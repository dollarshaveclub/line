package line_test

import "github.com/dollarshaveclub/line"

func Example() {
	line.Red().Print("Hello").Green("World").Blue().Println("!!!")

	line.Prefix("--> ").Suffix(" <---").Println("Nice to meet you!")

	out := line.Prefix("--> ")
	out.Println("One").Prefix("   --> ").Print("A\nB\nC")
	out.Println("Two").Prefix("   --> ").Green().Println("D")

	line.Info("An info statement")
	line.Error("An error statement")
}
