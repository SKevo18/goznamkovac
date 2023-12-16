package prevodnik

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"poznamkovac/internal/sablonovac"

	"github.com/flosch/pongo2/v6"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var poznamkySablona = sablonovac.NacitatSablonu("sablony/_poznamky.html") // zoznamSablona   = sablonovac.NacitatSablonu("sablony/_zoznam.html")

func markdownNaHTML(markdown_zdroj []byte) ([]byte, map[string]interface{}, error) {
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

func konvertovatPoznamky(markdown_cesta string) ([]byte, map[string]interface{}, error) {
	mdPoznamky, chyba := os.ReadFile(markdown_cesta)
	if chyba != nil {
		return nil, nil, chyba
	}

	htmlPoznamky, metaData, chyba := markdownNaHTML(mdPoznamky)
	if chyba != nil {
		return nil, nil, chyba
	}

	html, chyba := sablonovac.VykreslitSablonu(poznamkySablona, pongo2.Context{
		"poznamky":     string(htmlPoznamky),
		"meta":         metaData,
		"pojmova_mapa": nil,
	})
	if chyba != nil {
		return nil, nil, chyba
	}

	return html, metaData, nil
}

type Poznamka struct {
	nazov string
	cesta string

	datum_vytvorenia string
	autor            string
}

func najstMarkdownPoznamky(poznamkyCesta string) ([]string, error) {
	markdownPoznamky, _ := filepath.Glob(poznamkyCesta + "/*.md")
	if markdownPoznamky == nil {
		return nil, fmt.Errorf("neboli nájdené žiadne markdown súbory v `%s`", poznamkyCesta)
	}

	return markdownPoznamky, nil
}

func KonvertovatVsetkyPoznamky(poznamkyCesta string, vystupnaCesta string) ([]Poznamka, error) {
	var zoznamPoznamok []Poznamka
	markdownPoznamky, chyba := najstMarkdownPoznamky(poznamkyCesta)
	if chyba != nil {
		return nil, chyba
	}

	os.MkdirAll(vystupnaCesta+"/staticke", 0o755)

	for _, markdown_poznamky := range markdownPoznamky {
		html_cesta := filepath.Clean(vystupnaCesta + markdown_poznamky[len(poznamkyCesta):len(markdown_poznamky)-len(".md")] + ".html")

		html, metaData, chyba := konvertovatPoznamky(markdown_poznamky)
		if chyba != nil {
			return nil, chyba
		}

		os.MkdirAll(filepath.Dir(html_cesta), 0o755)
		if chyba := os.WriteFile(html_cesta, html, 0o644); chyba != nil {
			return nil, chyba
		}

		var nazov string
		if metaData["Nazov"] == nil {
			nazov = filepath.Base(markdown_poznamky)
		} else {
			nazov = metaData["Nazov"].(string)
		}

		var datum_vytvorenia string
		if metaData["Datum"] == nil {
			stat, _ := os.Stat(markdown_poznamky)
			datum_vytvorenia = stat.ModTime().Format("2006-01-02")
		} else {
			datum_vytvorenia = metaData["Datum"].(string)
		}

		var autor string
		if metaData["Autor"] == nil {
			autor = "Neznámy"
		} else {
			autor = metaData["Autor"].(string)
		}

		zoznamPoznamok = append(zoznamPoznamok, Poznamka{
			nazov:            nazov,
			cesta:            html_cesta,
			datum_vytvorenia: datum_vytvorenia,
			autor:            autor,
		})
	}

	return zoznamPoznamok, nil
}

/*func VytvoritZoznamPoznamok(cestaZoznamu string, zoznamPoznamok []Poznamka) error {
	html, chyba := sablonovac.VykreslitSablonu(zoznamSablona, pongo2.Context{
		"poznamky": zoznamPoznamok,
	})
	if chyba != nil {
		return chyba
	}

	if chyba := os.WriteFile(cestaZoznamu, html, 0o644); chyba != nil {
		return chyba
	}

	return nil
}
*/
