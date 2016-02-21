package cgsparser

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	//"os"
	//"encoding/json"
)

type Title struct {
	Number   string
	Name     string
	Note     string
	Chapters []Chapter
}

func NewTitle(url string) (Title, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Print(err)
		return Title{}, err
	}

	title := Title{}

	title.Number = strings.TrimSuffix(strings.TrimPrefix(doc.Find("H1.title-no").Text(), "TITLE "), "*")
	title.Name = doc.Find("H1.title-name").Text()
	title.Note = doc.Find("P.front-note-first").Text()
	title.Chapters = make([]Chapter, 0)

	urlprefix := url[0 : strings.LastIndex(url, "/")+1]

	doc.Find("TABLE TR TD.left_40pct A").Each(func(i int, s *goquery.Selection) {
		txt, exists := s.Attr("href")
		if exists && txt != "titles.htm" {
			//log.Printf("Parsing URL : %s%s", urlprefix, txt)
			chap, err := NewChapter(urlprefix + txt)
			if err == nil {
				title.Chapters = append(title.Chapters, chap)
			}
		}
	})

	/*
		b, err := json.MarshalIndent(title, "", "\t")
		if err != nil {
			log.Print("error:", err)
		}
		os.Stdout.Write(b)
	*/

	return title, nil
}
