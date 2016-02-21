package cgsparser

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type Chapter struct {
	Number string 
	Title string
	Sections []Section
}

type Section struct {
	Number  string
	Title   string
	Body    string
	Source  string
	History string
}

func NewChapter(url string) (Chapter, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Print(err)
		return Chapter{}, err
	}

	chapter := Chapter{}

	chapter.Number = strings.TrimPrefix(doc.Find("H2.chap-no").Text(), "CHAPTER ")
	chapter.Title = doc.Find("H2.chap-name").Text()
	chapter.Sections = make([]Section, 0)

	doc.Find("P SPAN.catchln").Each(func(i int, s *goquery.Selection) {
		section := Section{}

		fullSectionSpan := strings.TrimPrefix(s.Text(), "Sec. ")
		section.Number = fullSectionSpan[0:strings.Index(fullSectionSpan, ".")]
		section.Title = fullSectionSpan[strings.Index(fullSectionSpan, " ")+1:]

		log.Printf("Found section %s", section.Number)

		section.Body = strings.TrimPrefix(s.Parent().Text(), "Sec. " + section.Number + ". " + section.Title + " ")
		section.Source = s.Parent().Next().Text()
		section.History = s.Parent().Next().Next().Text()

		chapter.Sections = append(chapter.Sections, section)
	})

	return chapter, nil
}
