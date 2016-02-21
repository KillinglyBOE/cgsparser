package cgsparser

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	//"os"
	//"encoding/json"
)

type Chapter struct {
	Number   string
	Name     string
	Sections []Section
}

type Section struct {
	Number     string
	Name       string
	Body       string
	Source     string
	History    string
	Annotation string
}

func NewChapter(url string) (Chapter, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Print(err)
		return Chapter{}, err
	}

	chapter := Chapter{}

	chapter.Number = strings.TrimPrefix(doc.Find("H2.chap-no").Text(), "CHAPTER ")
	chapter.Name = doc.Find("H2.chap-name").Text()
	chapter.Sections = make([]Section, 0)

	doc.Find("P SPAN.catchln").Each(func(i int, s *goquery.Selection) {
		section := Section{}

		fullSectionSpan := strings.TrimPrefix(s.Text(), "Sec. ")
		section.Number = fullSectionSpan[0:strings.Index(fullSectionSpan, ".")]
		section.Name = fullSectionSpan[strings.Index(fullSectionSpan, " ")+1:]

		//log.Printf("Found section %s", section.Number)

		section.Body = strings.TrimPrefix(s.Parent().Text(), "Sec. "+section.Number+". "+section.Name+" ")
		nextone := s.Parent().Next()
		for {
			// Hack -- if there's a class, we can move on
			_, exists := nextone.Attr("class")
			if exists {
				break
			}
			section.Body += "\n\n" + nextone.Text()
			nextone = nextone.Next()
			if !nextone.Is("p") {
				break
			}
		}
		for {
			if !nextone.Is("p") {
				break
			}
			className, exists := nextone.Attr("class")
			if !exists {
				break
			}
			text := strings.Replace(nextone.Text(), "*", "\\*", -1)

			// Based on class name, build everything
			switch className {
			case "source-first":
				section.Source += text
				break
			case "source":
				section.Source += text
				break
			case "history-first":
				section.History += text
				break
			case "history":
				section.History += text
				break
			case "annotation-first":
				section.Annotation += text
				break
			case "annotation":
				section.Annotation += text
				break
			default:
				break
			}

			// Move on
			nextone = nextone.Next()
		}

		chapter.Sections = append(chapter.Sections, section)
	})

	/*
		b, err := json.MarshalIndent(chapter, "", "\t")
		if err != nil {
			log.Print("error:", err)
		}
		os.Stdout.Write(b)
	*/

	return chapter, nil
}
