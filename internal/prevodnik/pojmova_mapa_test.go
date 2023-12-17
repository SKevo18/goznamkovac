package prevodnik_test

import (
	"encoding/json"
	"testing"

	"poznamkovac/internal/prevodnik"
)

var markdownText = []byte(`
# Nadpis 1

Obsah nadpisu 1

## Nadpis 2

Obsah nadpisu 2

### Nadpis 3

Obsah nadpisu 3
`)

func TestPojmovaMapa(t *testing.T) {
	t.Run("GenerovatFarbu", func(t *testing.T) {
		t.Run("RozneFarby", func(t *testing.T) {
			farba1 := prevodnik.GenerovatFarbu(1)
			farba2 := prevodnik.GenerovatFarbu(1)
			if farba1 != farba2 {
				t.Errorf("Farby sa nezhodujú: %s a %s", farba1, farba2)
			}
		})

		t.Run("RovnakeFarby", func(t *testing.T) {
			farba1 := prevodnik.GenerovatFarbu(1)
			farba2 := prevodnik.GenerovatFarbu(2)
			if farba1 == farba2 {
				t.Errorf("Farby sa zhodujú: %s a %s", farba1, farba2)
			}
		})
	})

	t.Run("NajstNadpisy", func(t *testing.T) {
		nadpisy := prevodnik.NajstNadpisy(markdownText)
		json, _ := json.Marshal(nadpisy)

		t.Logf("Nájdené nadpisy: %v", string(json))

		if len(nadpisy) != 3 {
			t.Errorf("Neočakávaný počet nájdených nadpisov: %d", len(nadpisy))
		}
	})

	t.Run("PojmovaMapa", func(t *testing.T) {
		pojmovaMapa := prevodnik.VytvoritPojmovuMapu(markdownText)
		json, _ := json.Marshal(pojmovaMapa)

		t.Logf("Vygenerovaná pojmová mapa: %v", string(json))

		if len(pojmovaMapa.Bunky) != 3 {
			t.Errorf("Neočakávaný počet vygenerovaných buniek: %d", len(pojmovaMapa.Bunky))
		}
		if len(pojmovaMapa.Spojenia) != 2 {
			t.Errorf("Neočakávaný počet vygenerovaných spojení: %d", len(pojmovaMapa.Spojenia))
		}
		if len(pojmovaMapa.Tooltipy) != 3 {
			t.Errorf("Neočakávaný počet vygenerovaných tooltipov: %d", len(pojmovaMapa.Tooltipy))
		}
	})
}
