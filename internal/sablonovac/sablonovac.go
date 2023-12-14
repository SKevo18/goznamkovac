package sablonovac

import (
	"bytes"
	"embed"
	"html/template"
	"log"
)

var (
	//go:embed sablony
	sablonyFS  embed.FS
	sablonovac = template.Must(template.ParseFS(sablonyFS, "sablony/*.html"))
)

type Data map[string]any

// Načíta šablónu z daného súboru a vráti ju ako pointer na šablónu.
func NacitatSablonu(cesta string) *template.Template {
	sablona, chyba := sablonovac.ParseFiles(cesta)
	if chyba != nil {
		log.Fatalf("Nepodarilo sa načítať šablónu %s: %s", cesta, chyba)
	}

	return sablona
}

// Vykreslí danú šablónu s danými dátami a vráti výstup v bytoch.
// V prípade chyby pri vykresľovaní šablóny sa program ukončí.
func VykreslitSablonu(sablona *template.Template, data Data) []byte {
	vystup := new(bytes.Buffer)

	chyba := sablona.ExecuteTemplate(vystup, sablona.Name(), data)
	if chyba != nil {
		log.Fatalf("Nepodarilo sa vykresliť šablónu: %s", chyba)
	}

	return vystup.Bytes()
}
