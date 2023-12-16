package main

import (
	"log"
	"os"

	"poznamkovac/internal/prevodnik"
	"poznamkovac/internal/sablonovac"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Poznámkovač",
		Usage: "Aplikácia pre konvertovanie Markdown poznámok na statickú HTML stránku.",
		Action: func(c *cli.Context) error {
			poznamkyCesta := c.Args().First()
			if poznamkyCesta == "" {
				poznamkyCesta, _ = os.Getwd()
			}

			vystupnaCesta := c.Args().Get(1)
			if vystupnaCesta == "" {
				vystupnaCesta = "./site"
			}

			_, chyba := prevodnik.KonvertovatVsetkyPoznamky(poznamkyCesta, vystupnaCesta)
			if chyba != nil {
				return chyba
			}

			sablonovac.KopirovatStatickeSubory(vystupnaCesta)

			/* // TODO: Upraviť šablónu
			chyba = prevodnik.VytvoritZoznamPoznamok(vystupnaCesta+"/index.html", poznamkyZoznam)
			if chyba != nil {
				return chyba
			}
			*/

			return nil
		},
	}

	if chyba := app.Run(os.Args); chyba != nil {
		log.Fatal(chyba)
	}
}
