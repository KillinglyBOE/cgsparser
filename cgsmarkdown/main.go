package main

import (
	"flag"
	"fmt"

	"github.com/KillinglyBOE/cgsparser"
)

var (
	baseLevel = flag.Int("base", 0, "Base level of indent")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Printf("No titles given.\n")
		return
	}
	for _, t := range args {
		title, err := cgsparser.NewTitle(fmt.Sprintf("https://www.cga.ct.gov/current/pub/title_%s.htm", t))
		if err != nil {
			panic(err)
		}
		fmt.Printf(title.ToMarkdown("Appendix: CGS", *baseLevel))
	}
}
