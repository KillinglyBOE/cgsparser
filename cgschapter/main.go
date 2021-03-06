package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/KillinglyBOE/cgsparser"
)

var (
	baseLevel = flag.Int("base", 0, "Base level of indent")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Printf("No chapters given.\n")
		return
	}
	for _, t := range args {
		d, _ := strconv.ParseInt(t, 10, 64)
		title, err := cgsparser.NewChapter(fmt.Sprintf("https://www.cga.ct.gov/current/pub/chap_%03d.htm", d))
		if err != nil {
			panic(err)
		}
		fmt.Printf(title.ToMarkdown(*baseLevel))
	}
}
