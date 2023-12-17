package sablonovac_test

import (
	"testing"

	"poznamkovac/internal/sablonovac"
)

func TestSablonovac(t *testing.T) {
	t.Run("RelativnaCestaKStatickym", func(t *testing.T) {
		t.Run("PrazdnaCesta", func(t *testing.T) {
			cesta := sablonovac.RelativnaCestaKStatickym("./subor.md")
			if cesta != "staticke" {
				t.Errorf("Neočakávaná cesta: %s", cesta)
			}
		})

		t.Run("JednaHlbka", func(t *testing.T) {
			cesta := sablonovac.RelativnaCestaKStatickym("site/subor.md")
			if cesta != "../staticke" {
				t.Errorf("Neočakávaná cesta: %s", cesta)
			}
		})

		t.Run("DveHlbkyPlusAbsolutnaCesta", func(t *testing.T) {
			cesta := sablonovac.RelativnaCestaKStatickym("/site/2023/12/17/ucivo/subor.md")
			if cesta != "../../../../../staticke" {
				t.Errorf("Neočakávaná cesta: %s", cesta)
			}
		})
	})
}
