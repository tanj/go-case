package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("gocase", "Convert stdin based on supplied option and output to stdout")
	upper := parser.Flag("u", "upper", &argparse.Options{Help: "Convert to all upper case"})
	lower := parser.Flag("l", "lower", &argparse.Options{Help: "Convert to all lower case"})
	title := parser.Flag("t", "title", &argparse.Options{Help: "Convert to title case"})
	source := parser.String("s", "string", &argparse.Options{Help: "Input String rather than stdin"})
	// capitalize := parser.Flag("c", "capitalize", &argparse.Options{Help: "Convert to capitalize case"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	var str string
	if *source > "" {
		str = string(*source)
	} else {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		str = string(stdin)
	}

	if *upper {
		c := cases.Upper(language.Und)
		fmt.Printf(c.String(str))
	} else if *lower {
		c := cases.Lower(language.Und)
		fmt.Printf(c.String(str))
	} else if *title {
		c := cases.Title(language.Und)
		fmt.Printf(c.String(str))
	} else {
		fmt.Printf(str)
	}
	// else if *capitalize {
	// 	fmt.Printf(strings.Title)
	// }
}
