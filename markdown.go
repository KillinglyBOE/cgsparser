package cgsparser

import (
	"bytes"
	"strings"
)

func (t Title) ToMarkdown(titlePrefix string, layers int) string {
	buf := new(bytes.Buffer)

	// Heading
	buf.WriteString(layersPrefix(layers) + " ")
	if titlePrefix != "" {
		buf.WriteString(titlePrefix)
		buf.WriteString(" ")
	}
	buf.WriteString("Title ")
	buf.WriteString(t.Number)
	buf.WriteString(" - ")
	buf.WriteString(t.Name)
	buf.WriteString("\n\n")

	// Brief description
	if t.Note != "" {
		buf.WriteString(strings.Replace(t.Note, "*", "", -1))
		buf.WriteString("\n\n")
	}

	// Iterate
	for _, c := range t.Chapters {
		buf.WriteString(c.ToMarkdown(layers + 1))
	}

	return buf.String()
}

func (c Chapter) ToMarkdown(layers int) string {
	buf := new(bytes.Buffer)

	// Heading
	buf.WriteString(layersPrefix(layers) + " Chapter ")
	buf.WriteString(strings.Replace(c.Number, "*", "", -1))
	buf.WriteString(": ")
	buf.WriteString(c.Name)
	buf.WriteString("\n\n")

	// Iterate
	for _, s := range c.Sections {
		buf.WriteString(s.ToMarkdown(layers + 1))
	}

	return buf.String()
}

func (s Section) ToMarkdown(layers int) string {
	buf := new(bytes.Buffer)

	// Heading
	buf.WriteString(layersPrefix(layers) + " Section ")
	buf.WriteString(s.Number)
	buf.WriteString(": ")
	buf.WriteString(s.Name)
	buf.WriteString("\n\n")

	// Body
	buf.WriteString(strings.Replace(s.Body, "*", "\\*", -1))
	buf.WriteString("\n\n")

	// Source and History, if present
	if s.Source != "" {
		buf.WriteString("**")
		buf.WriteString(s.Source)
		buf.WriteString("**\n\n")
	}
	if s.History != "" {
		buf.WriteString("> ")
		buf.WriteString(s.History)
		buf.WriteString("\n\n")
	}
	if s.Annotation != "" {
		buf.WriteString("*")
		buf.WriteString(s.Annotation)
		buf.WriteString("*\n\n")
	}

	return buf.String()
}

func layersPrefix(layers int) string {
	x := ""
	for i := 0; i <= layers; i++ {
		x += "#"
	}
	return x
}
