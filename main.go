package main

import (
	"log"
	"os"
	"path/filepath"

	"goznamkovac/internal/prevodnik"
	"goznamkovac/internal/sablonovac"
)

const (
	rootCesta     = "poznamky"
	sablonaCesta  = "sablony"
	vystupnaCesta = "vystup"
)

var poznamkySablona = sablonovac.NacitatSablonu(sablonaCesta + "/_poznamky.html")

func konvertovatPoznamky(markdown_cesta string, html_vystup string) (chyba error) {
	mdPoznamky, chyba := os.ReadFile(markdown_cesta)
	if chyba != nil {
		return chyba
	}

	htmlPoznamky, chyba := prevodnik.MarkdownNaHTML(mdPoznamky)
	if chyba != nil {
		return chyba
	}

	html := sablonovac.VykreslitSablonu(poznamkySablona, sablonovac.Data{
		"poznamky":     htmlPoznamky,
		"pojmova_mapa": nil,
	})

	chyba = os.WriteFile(html_vystup, html, 0o644)
	if chyba != nil {
		return chyba
	}

	return nil
}

func main() {
	markdownPoznamky, _ := filepath.Glob(rootCesta + "/**/*.md")
	if markdownPoznamky == nil {
		log.Fatalf("Neboli nájdené žiadne markdown súbory v %s", rootCesta)
	}

	for _, markdown_poznamky := range markdownPoznamky {
		html_vystup := vystupnaCesta + "/" + markdown_poznamky[len(rootCesta)+1:len(markdown_poznamky)-3] + ".html"

		chyba := konvertovatPoznamky(markdown_poznamky, html_vystup)
		if chyba != nil {
			log.Fatalf("Nepodarilo sa konvertovať poznámky pre `%s`: %s", markdown_poznamky, chyba)
		}
	}
}
