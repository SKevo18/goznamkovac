package prevodnik

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

// Nadpis reprezentuje heading v markdown texte (pre pojmovú mapu).
type Nadpis struct {
	ID      int    `json:"id,omitempty"`
	Level   int    `json:"level"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Bunka reprezentuje uzol/vrchol v pojmovej mape.
type Bunka struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
	Level int    `json:"level"`
	Color string `json:"color"`
}

// Spojenie reprezentuje hranu medzi bunkami v pojmovej mape.
type Spojenie struct {
	From int `json:"from"`
	To   int `json:"to"`
}

// Regex pre nájdenie nadpisov v markdown texte.
var obsahRegex = regexp.MustCompile(`(#{1,6})\s+(.*?)\n(.*?)\n`)

// Generuje náhodnú farbu na základe čísla v HEX formáte. Rovnaké číslo = rovnaká farba.
func generovatFarbu(cislo int) string {
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%d", cislo)))
	hashHex := hex.EncodeToString(hasher.Sum(nil))

	var farba []string
	for i := 0; i < len(hashHex); i += 2 {
		val, _ := hex.DecodeString(hashHex[i : i+2])

		color := int(val[0])*75/100 + 75
		if color > 255 {
			color = 255 // zabránenie pretečeniu
		}

		farba = append(farba, fmt.Sprintf("%02x", color))
	}

	return "#" + strings.Join(farba[:3], "")
}

// Nájde všetky nadpisy v markdown texte a vráti ich ako pole.
func NajstNadpisy(markdownText []byte) []Nadpis {
	matches := obsahRegex.FindAllSubmatch(markdownText, -1)

	var nadpisy []Nadpis
	for _, match := range matches {
		level := len(match[1])
		title := match[2]
		content := match[3]
		nadpisy = append(nadpisy, Nadpis{
			Level:   level,
			Title:   string(title),
			Content: string(content),
		})
	}
	return nadpisy
}

func VytvoritPojmovuMapu(markdownText []byte) map[string]interface{} {
	nadpisy := NajstNadpisy(markdownText)

	bunky := make([]Bunka, 0)
	spojenia := make([]Spojenie, 0)

	zasobnik := make([]int, 0)
	idBunky := 1
	for _, nadpis := range nadpisy {
		bunky = append(bunky, Bunka{
			ID:    idBunky,
			Label: nadpis.Title,
			Level: nadpis.Level,
			Color: generovatFarbu(idBunky),
		})

		dlzkaZasobnika := len(zasobnik)
		if dlzkaZasobnika > 0 {
			posledny := zasobnik[dlzkaZasobnika-1]
			if nadpis.Level <= bunky[posledny-1].Level {
				zasobnik = zasobnik[:dlzkaZasobnika-1] // odstráni posledný prvok
			}

			spojenie := Spojenie{From: posledny, To: idBunky}
			spojenia = append(spojenia, spojenie)
		}

		zasobnik = append(zasobnik, idBunky)
		idBunky++
	}

	return map[string]interface{}{"bunky": bunky, "spojenia": spojenia}
}
