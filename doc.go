/*

Package line is an easy to use package for stylizing terminal output. The API focuses on usability via chaining and, consequently, is quite flexible.

Use the color methods to print in a wide selection of colors:

	line.Red().Print("Hello").Green("World").Blue().Println("!!!")

Output with a prefix and a suffix:

	line.Prefix("--> ").Suffix(" <---").Println("Nice to meet you!")

Prefix all output:

	out := line.Prefix("--> ")
	out.Println("One").Println("Two").Println("Three")
*/
package line
