package prevodnik

import (
	"bytes"
	"os"
	"path/filepath"

	"poznamkovac/internal/sablonovac"

	"github.com/flosch/pongo2/v6"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"

	pikchr "github.com/jchenry/goldmark-pikchr"
	mathjax "github.com/litao91/goldmark-mathjax"
	"go.abhg.dev/goldmark/anchor"
	"go.abhg.dev/goldmark/mermaid"
	"go.abhg.dev/goldmark/toc"
)

var (
	poznamkySablona = sablonovac.NacitatSablonu("sablony/_poznamky.html")
	zoznamSablona   = sablonovac.NacitatSablonu("sablony/_zoznam.html")
)

// Konvertuje poznámky z Markdown súboru do HTML
func MarkdownNaHtml(markdownPoznamky []byte) ([]byte, error) {
	prevodnik := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Footnote,
			extension.DefinitionList,
			extension.Typographer,

			mathjax.MathJax,
			&anchor.Extender{
				Position: anchor.Before,
			},
			&toc.Extender{
				Title:   "Obsah",
				TitleID: "obsah",
				ListID:  "obsah-list",
			},
			&mermaid.Extender{},
			&pikchr.Extender{DarkMode: true},
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	var buffer bytes.Buffer
	if chyba := prevodnik.Convert(markdownPoznamky, &buffer); chyba != nil {
		return nil, chyba
	}

	return buffer.Bytes(), nil
}

type Poznamky struct {
	Nazov         string
	MarkdownCesta string

	PrilozeneSubory []string
	ObsahujeKviz    bool
	DatumUpravy     string
}

// Vráti výstupnú cestu poznámok (t. j.: cesta mínus koreňový priečinok - priečinok poznámok sa vo výstupe nenachádza)
func (poznamky Poznamky) VystupnaCesta() string {
	return poznamky.MarkdownCesta[len(filepath.Base(poznamky.MarkdownCesta))+1:len(poznamky.MarkdownCesta)-len(".md")] + ".html"
}

// Konvertuje poznámky z Markdown súboru do HTML a vykreslí ich do šablóny
func (poznamky Poznamky) KonvertovatPoznamky() ([]byte, error) {
	markdownPoznamky, chyba := os.ReadFile(poznamky.MarkdownCesta)
	if chyba != nil {
		return nil, chyba
	}

	htmlPoznamky, chyba := MarkdownNaHtml(markdownPoznamky)
	if chyba != nil {
		return nil, chyba
	}

	html, chyba := sablonovac.VykreslitSablonu(poznamkySablona, pongo2.Context{
		"html":                string(htmlPoznamky),
		"poznamky":            poznamky,
		"zobrazitPojmovuMapu": true,
		"staticke":            sablonovac.RelativnaCestaKStatickym(poznamky.VystupnaCesta()),
	})
	if chyba != nil {
		return nil, chyba
	}

	return html, nil
}

func najstPrilozeneSubory(rootCesta string) ([]string, error) {
	prilozene_subory, chyba := filepath.Glob(rootCesta + "/*")
	if chyba != nil {
		return nil, chyba
	}

	for i, subor := range prilozene_subory {
		preskocitSubory := []string{"index.md", "kviz.yaml"}

		for _, preskocit := range preskocitSubory {
			if filepath.Base(subor) == preskocit {
				prilozene_subory = append(prilozene_subory[:i], prilozene_subory[i+1:]...)
				break
			}
		}
	}

	return prilozene_subory, nil
}

// Nájde všetky Markdown súbory pre poznámky (`index.md`) v zadanom priečinku, rekurzívne
// (každý priečinok s poznámkami musí mať súbor `index.md` aby boli poznámky platné).
func najstMarkdownPoznamky(poznamkyCesta string) ([]Poznamky, error) {
	markdownPoznamky := make([]Poznamky, 0)

	chyba := filepath.Walk(poznamkyCesta, func(cesta string, info os.FileInfo, chyba error) error {
		if chyba != nil {
			return chyba
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Base(cesta) == "index.md" {
			root := filepath.Dir(cesta)
			prilozene_subory, chyba := najstPrilozeneSubory(root)
			if chyba != nil {
				return chyba
			}

			obsahujeKviz := false
			if _, chyba := os.Stat(root + "/kviz.yaml"); chyba == nil {
				obsahujeKviz = true
			}

			markdownPoznamky = append(markdownPoznamky, Poznamky{
				Nazov:           filepath.Base(root),
				MarkdownCesta:   cesta,
				PrilozeneSubory: prilozene_subory,
				ObsahujeKviz:    obsahujeKviz,
				DatumUpravy:     info.ModTime().Format("2006-01-02 15:04:05"),
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
	markdownPoznamky, chyba := najstMarkdownPoznamky(poznamkyCesta)
	if chyba != nil {
		return nil, chyba
	}

	os.MkdirAll(vystupnaCesta+"/staticke", 0o755)

	for _, poznamky := range markdownPoznamky {
		html, chyba := poznamky.KonvertovatPoznamky()
		if chyba != nil {
			return nil, chyba
		}

		htmlCesta := vystupnaCesta + "/" + poznamky.VystupnaCesta()
		os.MkdirAll(filepath.Dir(htmlCesta), 0o755)

		if chyba := os.WriteFile(htmlCesta, html, 0o644); chyba != nil {
			return nil, chyba
		}

		// Kvíz
		if poznamky.ObsahujeKviz {
			kvizYaml, err := os.ReadFile(filepath.Dir(poznamky.MarkdownCesta) + "/kviz.yaml")
			if err != nil {
				return nil, err
			}

			kviz, chyba := nacitatKviz(kvizYaml)
			if chyba != nil {
				return nil, chyba
			}

			html, chyba := kviz.Vykreslit()
			if chyba != nil {
				return nil, chyba
			}

			kvizCesta := vystupnaCesta + "/" + filepath.Dir(poznamky.VystupnaCesta()) + "/kviz.html"
			if chyba := os.WriteFile(kvizCesta, html, 0o644); chyba != nil {
				return nil, chyba
			}
		}
	}

	return markdownPoznamky, nil
}

func VytvoritZoznamPoznamok(cestaZoznamu string, zoznamPoznamok []Poznamky) error {
	html, chyba := sablonovac.VykreslitSablonu(zoznamSablona, pongo2.Context{
		"zoznam_poznamok": zoznamPoznamok,
		"staticke":        "staticke",
	})
	if chyba != nil {
		return chyba
	}

	if chyba := os.WriteFile(cestaZoznamu, html, 0o644); chyba != nil {
		return chyba
	}

	return nil
}
