package sablonovac

import (
	"embed"
	"log"

	"github.com/flosch/pongo2/v6"
)

var (
	//go:embed sablony/*.html
	sablonyFS embed.FS

	// Globálny šablónovač, ktorý sa používa na načítanie šablón.
	sablonovac = pongo2.NewSet("sablony", pongo2.NewFSLoader(sablonyFS))
)

// Načíta šablónu z daného súboru a vráti ju ako pointer na šablónu.
func NacitatSablonu(cesta string) *pongo2.Template {
	sablona, chyba := sablonovac.FromFile(cesta)
	if chyba != nil {
		log.Fatalf("Nepodarilo sa načítať šablónu %s: %s", cesta, chyba)
	}

	return sablona
}

// Vykreslí danú šablónu s danými dátami a vráti výstup v bytoch, alebo chybu.
func VykreslitSablonu(sablona *pongo2.Template, data pongo2.Context) ([]byte, error) {
	vystup, chyba := sablona.ExecuteBytes(data)
	if chyba != nil {
		return nil, chyba
	}

	return vystup, nil
}
