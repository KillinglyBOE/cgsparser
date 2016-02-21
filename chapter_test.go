package cgsparser

import (
	//"fmt"
	"testing"
)

func TestChapterParse(t *testing.T) {
	ch, err := NewChapter("https://www.cga.ct.gov/current/pub/chap_181a.htm")
	if err != nil {
		t.Fatalf("%v", err)
	}
	if ch.Number != "181a" {
		t.Fatalf("181a != %s", ch.Number)
	}
	if ch.Name != "CONNECTICUT HUMANITIES GRANTS" {
		t.Fatalf("CONNECTICUT HUMANITIES GRANTS != %s", ch.Name)
	}
	if len(ch.Sections) != 2 {
		t.Fatalf("Should have found 2 sections, found %d", len(ch.Sections))
	}
}

/*
func TestChapterFullParse(t *testing.T) {
	ch, err := NewChapter("https://www.cga.ct.gov/current/pub/chap_164.htm")
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Println(ch.ToMarkdown())
}
*/
