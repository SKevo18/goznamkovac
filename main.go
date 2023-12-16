package main

import (
	"log"
	"os"
	"path/filepath"

	"goznamkovac/internal/prevodnik"
	"goznamkovac/internal/sablonovac"

	"github.com/flosch/pongo2/v6"
)

const (
	rootCesta     = "poznamky"
	sablonaCesta  = "sablony"
	vystupnaCesta = "site"
)

var poznamkySablona = sablonovac.NacitatSablonu(sablonaCesta + "/_poznamky.html")

func konvertovatPoznamky(markdown_cesta string, html_vystup string) (chyba error) {
	mdPoznamky, chyba := os.ReadFile(markdown_cesta)
	if chyba != nil {
		return chyba
	}

	htmlPoznamky, metaData, chyba := prevodnik.MarkdownNaHTML(mdPoznamky)
	if chyba != nil {
		return chyba
	}

	html, chyba := sablonovac.VykreslitSablonu(poznamkySablona, pongo2.Context{
		"poznamky":     string(htmlPoznamky),
		"meta":         metaData,
		"pojmova_mapa": nil,
	})
	if chyba != nil {
		return chyba
	}

	chyba = os.WriteFile(html_vystup, html, 0o644)
	if chyba != nil {
		return chyba
	}

	return nil
}

func najstMarkdownPoznamky() []string {
	markdownPoznamky, _ := filepath.Glob(rootCesta + "/*.md")
	if markdownPoznamky == nil {
		log.Fatalf("Neboli nájdené žiadne markdown súbory v `%s`.", rootCesta)
	}

	return markdownPoznamky
}

func main() {
	markdownPoznamky := najstMarkdownPoznamky()
	os.MkdirAll(vystupnaCesta + "/staticke", 0o755)

	for _, markdown_poznamky := range markdownPoznamky {
		html_vystup := vystupnaCesta + "/" + markdown_poznamky[len(rootCesta)+1:len(markdown_poznamky)-3] + ".html"

		chyba := konvertovatPoznamky(markdown_poznamky, html_vystup)
		if chyba != nil {
			log.Fatalf("Nepodarilo sa konvertovať poznámky pre `%s`: %s", markdown_poznamky, chyba)
		}
	}
}
