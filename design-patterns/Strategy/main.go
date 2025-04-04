package main

import (
	"fmt"
	"strings"
)

type OutputFormat int

const (
	Markdown OutputFormat = iota
	Html
)

type ListStrategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}

type MarkdownListStrategy struct{}

// * one
// * two
func (m *MarkdownListStrategy) Start(builder *strings.Builder) {
}

func (m *MarkdownListStrategy) End(builder *strings.Builder) {
}

func (m *MarkdownListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("* " + item + "\n")
}

type HtmlListStrategy struct{}

func (h *HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>")
}

func (h *HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>")
}

func (h *HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("<li>" + item + "</li>")
}

type TextProcessor struct {
	builder *strings.Builder
	list    ListStrategy
}

func NewTextProcessor(list ListStrategy) *TextProcessor {
	return &TextProcessor{&strings.Builder{}, list}
}

func (t *TextProcessor) SetOutputFormat(format OutputFormat) {
	switch format {
	case Markdown:
		t.list = &MarkdownListStrategy{}
	case Html:
		t.list = &HtmlListStrategy{}
	}
}

func (t *TextProcessor) AppendList(items []string) {
	t.list.Start(t.builder)
	for _, item := range items {
		t.list.AddListItem(t.builder, item)
	}
	t.list.End(t.builder)
}

func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

func (t *TextProcessor) String() string {
	return t.builder.String()
}

func main() {
	tp := NewTextProcessor(&MarkdownListStrategy{})
	tp.AppendList([]string{"one", "two", "three"})
	fmt.Println(tp)

	tp.Reset()
	tp.SetOutputFormat(Html)
	tp.AppendList([]string{"one", "two", "three"})
	fmt.Println(tp)

}
