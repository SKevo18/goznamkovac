package sablonovac

import (
	"embed"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2/v6"
	"github.com/otiai10/copy"
)

var (
	//go:embed all:sablony
	sablonyFS embed.FS

	// Globálny šablónovač, ktorý sa používa na načítanie šablón.
	sablonovac = pongo2.NewSet("sablony", pongo2.NewFSLoader(sablonyFS))
)

const PathSeparatorStr = string(os.PathSeparator)

// Vráti cestu ako zoznam adresárov, napr.: `sablony/_poznamky.html` -> `["sablony", "_poznamky.html"]`
func CestaAkoZoznam(cesta string) []string {
	return strings.Split(cesta, PathSeparatorStr)
}

// Vráti cestu k statickému priečinku (v koreňovom priečinku) z relatívnej cesty.
// Napr.: ak je cesta `/site/2023/12/17/ucivo/subor.md`, vráti `../../../../../staticke`.
func RelativnaCestaKStatickym(odRelativnejCesty string) string {
	cesta := ""

	dir := strings.TrimPrefix(filepath.Dir(odRelativnejCesty), PathSeparatorStr)
	if dir == "." {
		return "staticke"
	}

	for i := 0; i < len(CestaAkoZoznam(dir)); i++ {
		cesta += "../"
	}

	return cesta + "staticke"
}

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

// Skopíruje statické súbory do daného priečinku.
func KopirovatStatickeSubory(cesta string) error {
	chyba := copy.Copy("sablony/staticke", cesta+"/staticke", copy.Options{FS: sablonyFS})
	if chyba != nil {
		return chyba
	}

	return nil
}
