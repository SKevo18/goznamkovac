package prevodnik

import (
	"bytes"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func MarkdownNaHTML(markdown_zdroj []byte) ([]byte, map[string]interface{}, error) {
	prevodnik := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			meta.Meta,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
	)

	var buffer bytes.Buffer
    context := parser.NewContext()
	if chyba := prevodnik.Convert(markdown_zdroj, &buffer, parser.WithContext(context)); chyba != nil {
		return nil, nil, chyba
	}
    metaData := meta.Get(context)

	return buffer.Bytes(), metaData, nil
}
