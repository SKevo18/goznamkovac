package prevodnik

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

// Nadpis reprezentuje jednotlivý nadpis v Markdown dokumente.
// Nejedná sa o JSON reperezentáciu, ale o pomocnú štruktúru.
type Nadpis struct {
	Uroven  int
	Titulok string
	Obsah   string
}

// Bunka reprezentuje uzol/vrchol v pojmovej mape.
type Bunka struct {
	ID     int    `json:"id"`
	Stitok string `json:"label"`
	Uroven int    `json:"level"`
	Farba  string `json:"color"`
}

// Spojenie reprezentuje hranu medzi bunkami v pojmovej mape.
type Spojenie struct {
	ZBunky  int `json:"from"`
	DoBunky int `json:"to"`
}

// Tooltip reprezentuje tooltip pre bunku.
type Tooltip struct {
	ID    int    `json:"id"`
	Obsah string `json:"content"`
}

// PojmovaMapa reprezentuje celé JSON pojmovej mapy.
type PojmovaMapa struct {
	Bunky    []Bunka    `json:"bunky"`
	Spojenia []Spojenie `json:"spojenia"`
	Tooltipy []Tooltip  `json:"tooltipy"`
}

var obshahRegex = regexp.MustCompile(`(?sm)^(#+)\s?([^#\n]+)\s*([^#]+)\s*`)

// Generuje farbu pre bunku na základe čísla. Rovnaké číslo vždy vygeneruje rovnakú farbu.
func GenerovatFarbu(cislo int) string {
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%d", cislo)))
	hashHex := hex.EncodeToString(hasher.Sum(nil))

	var svetlaFarba []string
	for i := 0; i < len(hashHex); i += 2 {
		val, _ := hex.DecodeString(hashHex[i : i+2])
		color := int(val[0])*75/100 + 75
		if color > 255 {
			color = 255
		}
		svetlaFarba = append(svetlaFarba, fmt.Sprintf("%02x", color))
	}
	return "#" + strings.Join(svetlaFarba[:3], "")
}

// Nájde nadpisy v Markdown texte.
func NajstNadpisy(markdownText []byte) []Nadpis {
	zhody := obshahRegex.FindAllSubmatch(markdownText, -1)

	var nadpisy []Nadpis
	for _, zhoda := range zhody {
		uroven := len(zhoda[1])
		titulok := string(zhoda[2])
		obsah := string(zhoda[3])
		nadpisy = append(nadpisy, Nadpis{Uroven: uroven, Titulok: titulok, Obsah: obsah})
	}

	return nadpisy
}

// Vytvorí JSON pojmovej mapy z Markdown textu.
func VytvoritPojmovuMapu(markdownText []byte) PojmovaMapa {
	nadpisy := NajstNadpisy(markdownText)

	bunky := make([]Bunka, 0)
	spojenia := make([]Spojenie, 0)
	tooltipy := make([]Tooltip, 0)

	zasobnik := make([]int, 0)
	for idBunky, nadpis := range nadpisy {
		bunka := Bunka{
			ID:     idBunky,
			Stitok: nadpis.Titulok,
			Uroven: nadpis.Uroven,
			Farba:  GenerovatFarbu(idBunky),
		}
		bunky = append(bunky, bunka)

		// Pridanie tooltipu, ak obsah nie je prázdny
		if strings.TrimSpace(nadpis.Obsah) != "" {
			tooltipObsah, chyba := MarkdownNaHTML([]byte(nadpis.Obsah))
			if chyba != nil {
				tooltipObsah = []byte(fmt.Sprintf("<p>Chyba pri konverzii obsahu tooltipu: %s</p>", chyba))
			}

			tooltip := Tooltip{
				ID:    idBunky,
				Obsah: string(tooltipObsah),
			}
			tooltipy = append(tooltipy, tooltip)
		}

		// Spracovanie spojení medzi bunkami
		for len(zasobnik) > 0 && nadpis.Uroven <= bunky[zasobnik[len(zasobnik)-1]].Uroven {
			zasobnik = zasobnik[:len(zasobnik)-1]
		}

		if len(zasobnik) > 0 {
			spojenie := Spojenie{ZBunky: zasobnik[len(zasobnik)-1], DoBunky: idBunky}
			spojenia = append(spojenia, spojenie)
		}

		zasobnik = append(zasobnik, idBunky)
	}

	return PojmovaMapa{
		Bunky:    bunky,
		Spojenia: spojenia,
		Tooltipy: tooltipy,
	}
}
