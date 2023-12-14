package prevodnik

import (
    "bytes"
    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/extension"
    "github.com/yuin/goldmark/parser"
    "github.com/yuin/goldmark/renderer/html"
)

func MarkdownNaHTML(markdown_zdroj []byte) (html_vystup []byte, err error) {
    prevodnik := goldmark.New(
        goldmark.WithExtensions(extension.GFM),
        goldmark.WithParserOptions(
            parser.WithAutoHeadingID(),
        ),
        goldmark.WithRendererOptions(
            html.WithHardWraps(),
        ),
    )

    var buffer bytes.Buffer
    if chyba := prevodnik.Convert(markdown_zdroj, &buffer); chyba != nil {
        return nil, chyba
    }

    return buffer.Bytes(), nil
}
