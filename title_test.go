package cgsparser

import (
	"testing"
)

func TestTitleParse(t *testing.T) {
	title, err := NewTitle("https://www.cga.ct.gov/current/pub/title_10.htm")
	if err != nil {
		t.Fatalf("%v", err)
	}
	if title.Number != "10" {
		t.Fatalf("10 != %s", title.Number)
	}
}
