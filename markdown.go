package cgsparser

import (
	"bytes"
)

func (t Title) ToMarkdown(titlePrefix string) string {
	buf := new(bytes.Buffer)

	// Heading
	buf.WriteString("## ")
	if titlePrefix != "" {
		buf.WriteString(titlePrefix)
		buf.WriteString(" ")
	}
	buf.WriteString("Title ")
	buf.WriteString(t.Number)
	buf.WriteString(": ")
	buf.WriteString(t.Name)
	buf.WriteString("\n\n")

	// Brief description
	if t.Note != "" {
		buf.WriteString(t.Note)
		buf.WriteString("\n\n")
	}

	// Iterate
	for _, c := range t.Chapters {
		buf.WriteString(c.ToMarkdown())
	}

	return buf.String()
}

func (c Chapter) ToMarkdown() string {
	buf := new(bytes.Buffer)

	// Heading
	buf.WriteString("### Chapter ")
	buf.WriteString(c.Number)
	buf.WriteString(": ")
	buf.WriteString(c.Name)
	buf.WriteString("\n\n")

	return buf.String()
}

func (s Section) ToMarkdown() string {
	buf := new(bytes.Buffer)

	// Heading
	buf.WriteString("#### Section ")
	buf.WriteString(s.Number)
	buf.WriteString(": ")
	buf.WriteString(s.Name)
	buf.WriteString("\n\n")

	// Body
	buf.WriteString(s.Body)
	buf.WriteString("\n\n")

	// Source and History, if present
	if s.Source != "" {
		buf.WriteString(s.Source)
		buf.WriteString("\n\n")
	}
	if s.History != "" {
		buf.WriteString(s.History)
		buf.WriteString("\n\n")
	}

	return buf.String()
}
