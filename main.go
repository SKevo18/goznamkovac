package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"poznamkovac/internal/prevodnik"
	"poznamkovac/internal/sablonovac"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Poznámkovač",
		Usage: "Aplikácia pre konvertovanie Markdown poznámok na statickú HTML stránku.",
		Action: func(c *cli.Context) error {
			wd, _ := os.Getwd()

			poznamkyCesta := c.Args().First()
			if poznamkyCesta == "" {
				poznamkyCesta = wd
			}

			// FIXME: zlá cesta k statickým pre napr.: `../poznamkovac/poznamky`
			poznamkyCesta = filepath.Clean(strings.TrimPrefix(poznamkyCesta, wd + string(os.PathSeparator)))

			vystupnaCesta := c.Args().Get(1)
			if vystupnaCesta == "" {
				vystupnaCesta = "./site"
			}

			os.RemoveAll(vystupnaCesta)
			poznamkyZoznam, chyba := prevodnik.KonvertovatVsetkyPoznamky(poznamkyCesta, vystupnaCesta)
			if chyba != nil {
				return chyba
			}

			sablonovac.KopirovatStatickeSubory(vystupnaCesta)

			chyba = prevodnik.VytvoritZoznamPoznamok(vystupnaCesta+"/index.html", poznamkyZoznam)
			if chyba != nil {
				return chyba
			}

			return nil
		},
	}

	if chyba := app.Run(os.Args); chyba != nil {
		log.Fatal(chyba)
	}
}
