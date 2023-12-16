package prevodnik

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"

	"poznamkovac/internal/sablonovac"

	"github.com/flosch/pongo2/v6"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var poznamkySablona = sablonovac.NacitatSablonu("sablony/_poznamky.html")

// zoznamSablona   = sablonovac.NacitatSablonu("sablony/_zoznam.html")

type Poznamky struct {
	nazov string
	cesta string

	prilozene_subory []string
	datum_vytvorenia string
}

// Konvertuje poznámky z Markdown súboru do HTML
func (poznamky Poznamky) MarkdownNaHTML() ([]byte, error) {
	mdPoznamky, chyba := os.ReadFile(poznamky.cesta)
	if chyba != nil {
		return nil, chyba
	}

	prevodnik := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
	)

	var buffer bytes.Buffer
	if chyba := prevodnik.Convert(mdPoznamky, &buffer); chyba != nil {
		return nil, chyba
	}

	return buffer.Bytes(), nil
}

// Konvertuje poznámky z Markdown súboru do HTML a vykreslí ich do šablóny
func (poznamky Poznamky) KonvertovatPoznamky() ([]byte, error) {
	htmlPoznamky, chyba := poznamky.MarkdownNaHTML()
	if chyba != nil {
		return nil, chyba
	}

	pojmovaMapa := VytvoritPojmovuMapu(htmlPoznamky)
	pojmovaMapaJSON, chyba := json.Marshal(pojmovaMapa)
	if chyba != nil {
		return nil, chyba
	}

	html, chyba := sablonovac.VykreslitSablonu(poznamkySablona, pongo2.Context{
		"html":         string(htmlPoznamky),
		"poznamky":     poznamky,
		"pojmova_mapa": pojmovaMapaJSON,
	})
	if chyba != nil {
		return nil, chyba
	}

	return html, nil
}

// Nájde všetky Markdown súbory pre poznámky (`poznamky.md`) v zadanom priečinku, rekurzívne
func najstMarkdownPoznamky(poznamkyCesta string) ([]Poznamky, error) {
	markdownPoznamky := make([]Poznamky, 0)

	chyba := filepath.Walk(poznamkyCesta, func(cesta string, info os.FileInfo, chyba error) error {
		if chyba != nil {
			return chyba
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Base(cesta) == "poznamky.md" {
			root := filepath.Dir(cesta)
			prilozene_subory, chyba := filepath.Glob(root + "/*")
			if chyba != nil {
				return chyba
			}

			markdownPoznamky = append(markdownPoznamky, Poznamky{
				nazov:            filepath.Base(root),
				cesta:            cesta,
				prilozene_subory: prilozene_subory[1:],
				datum_vytvorenia: info.ModTime().Format("2006-01-02 15:04:05 +0100"),
			})
		}

		return nil
	})
	if chyba != nil {
		return nil, chyba
	}

	return markdownPoznamky, nil
}

func KonvertovatVsetkyPoznamky(poznamkyCesta string, vystupnaCesta string) ([]Poznamky, error) {
	var zoznamPoznamok []Poznamky
	markdownPoznamky, chyba := najstMarkdownPoznamky(poznamkyCesta)
	if chyba != nil {
		return nil, chyba
	}

	os.MkdirAll(vystupnaCesta+"/staticke", 0o755)

	for _, poznamky := range markdownPoznamky {
		vystupnaCesta := filepath.Clean(vystupnaCesta + poznamky.cesta[len(poznamkyCesta):len(poznamky.cesta)-len(".md")] + ".html")

		html, chyba := poznamky.KonvertovatPoznamky()
		if chyba != nil {
			return nil, chyba
		}

		os.MkdirAll(filepath.Dir(vystupnaCesta), 0o755)
		if chyba := os.WriteFile(vystupnaCesta, html, 0o644); chyba != nil {
			return nil, chyba
		}
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
