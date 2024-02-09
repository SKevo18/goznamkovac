package prevodnik

import (
	"poznamkovac/internal/sablonovac"

	"github.com/flosch/pongo2/v6"
	"github.com/goccy/go-yaml"
)

var kvizSablona = sablonovac.NacitatSablonu("sablony/_kviz.html")

// Kviz je štruktúra reprezentujúca kvíz
type Kviz struct {
	Nazov  string
	Otazky []struct {
		Otazka   string
		Odpovede []struct {
			Typ     string
			Spravna string
			// Atributy map[string]interface{}
		}
	}
}

func nacitatKviz(yamlData []byte) (*Kviz, error) {
	var kviz Kviz
	if chyba := yaml.Unmarshal(yamlData, &kviz); chyba != nil {
		return nil, chyba
	}

	for i := range kviz.Otazky {
		html, err := MarkdownNaHTML([]byte(kviz.Otazky[i].Otazka))
		if err != nil {
			return nil, err
		}

		kviz.Otazky[i].Otazka = string(html)
	}

	return &kviz, nil
}

func vykreslitKviz(kviz *Kviz) ([]byte, error) {
	html, err := sablonovac.VykreslitSablonu(kvizSablona, pongo2.Context{
		"kviz": kviz,
	})
	if err != nil {
		return nil, err
	}

	return html, nil
}
