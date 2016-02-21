package main

import (
	"fmt"
	"github.com/KillinglyBOE/cgsparser"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf("No titles given.\n")
		return
	}
	for _, t := range args {
		title, err := cgsparser.NewTitle(fmt.Sprintf("https://www.cga.ct.gov/current/pub/title_%s.htm", t))
		if err != nil {
			panic(err)
		}
		fmt.Printf(title.ToMarkdown("Appendix: Connecticut General Statutes"))
	}
}
