package prevodnik

import (
	"fmt"

	"poznamkovac/internal/sablonovac"

	"github.com/flosch/pongo2/v6"
	"github.com/goccy/go-yaml"
)

var kvizSablona = sablonovac.NacitatSablonu("sablony/_kviz.html")

// Kviz je štruktúra reprezentujúca kvíz
type Kviz struct {
	Nazov  string
	Otazky []Otazka
}

// Otazka je štruktúra reprezentujúca otázku v kvíze
type Otazka struct {
	Otazka   string
	Odpovede []Odpoved
}

func (otazka Otazka) Html() string {
	html, err := MarkdownNaHtml([]byte(otazka.Otazka))
	if err != nil {
		return err.Error()
	}

	return string(html)
}

// Odpoved je štruktúra reprezentujúca odpoveď na otázku
type Odpoved struct {
	Typ      string
	Spravna  string
	Atributy map[string]interface{}
}

func (odpoved Odpoved) Html() string {
	atributy := ""
	for atribut, hodnota := range odpoved.Atributy {
		atributy += fmt.Sprintf(` %s="%s"`, atribut, hodnota)
	}

	return fmt.Sprintf(`<input type="%s" data-odpoved data-spravna="%s" %s />`, odpoved.Typ, odpoved.Spravna, atributy)
}

func (kviz *Kviz) Vykreslit() ([]byte, error) {
	html, err := sablonovac.VykreslitSablonu(kvizSablona, pongo2.Context{
		"kviz": kviz,
	})

	if err != nil {
		return nil, err
	}

	return html, nil
}

func nacitatKviz(yamlData []byte) (*Kviz, error) {
	var kviz Kviz
	if chyba := yaml.Unmarshal(yamlData, &kviz); chyba != nil {
		return nil, chyba
	}

	return &kviz, nil
}
