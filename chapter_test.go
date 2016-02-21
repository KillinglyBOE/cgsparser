package cgsparser

import (
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
	if ch.Title != "CONNECTICUT HUMANITIES GRANTS" {
		t.Fatalf("CONNECTICUT HUMANITIES GRANTS != %s", ch.Title)
	}
	if len(ch.Sections) != 2 {
		t.Fatalf("Should have found 2 sections, found %d", len(ch.Sections))
	}
}
