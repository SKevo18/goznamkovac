/**
 * Vytvorí JSON pojmovej mapy z obsahu stránky.
 *
 * @param {HTMLElement} poznamkyElement Element, ktorý obsahuje poznámky.
 * @returns {object} JSON reprezentácia pojmovej mapy.
 **/
function vytvoritDataMapy(poznamkyElement) {
    const nadpisy = poznamkyElement.querySelectorAll("h1:not(#obsah), h2, h3, h4, h5, h6");
    let mapa = { vrcholy: [], hrany: [] };

    // Zoznam, ktorý uchováva posledný nadpis na každej úrovni
    let poslednyNadpisNaUrovni = [];

    for (let i = 0; i < nadpisy.length; i++) {
        const nadpis = nadpisy[i];
        const aktualnyLevel = ziskatLevelNadpisu(nadpis);

        const idVrchola = i + 1;
        const nazov = nadpis.innerText.slice(2);
        mapa.vrcholy.push({ id: idVrchola, label: nazov });

        if (aktualnyLevel > 1) {
            const idRodicovskehoVrchola = poslednyNadpisNaUrovni[aktualnyLevel - 1] || 1;

            mapa.hrany.push({
                from: idRodicovskehoVrchola,
                to: idVrchola
            });
        }

        // Aktualizujte posledný nadpis na úrovni
        poslednyNadpisNaUrovni[aktualnyLevel] = idVrchola;

        // Vymažte všetky nadpisy na vyšších úrovniach
        for (let j = aktualnyLevel + 1; j < poslednyNadpisNaUrovni.length; j++) {
            delete poslednyNadpisNaUrovni[j];
        }
    }

    return mapa;
}


/**
 * Vytvorí JSON pojmovej mapy z obsahu stránky.
 * @param {HTMLElement} elementMapy Element, do ktorého sa vloží zobrazenie pojmovej mapy.
 * @param {object} jsonMapy JSON reprezentácia pojmovej mapy.
 * @returns {object} Objekt pojmovej mapy (vis-network).
 **/
function vykreslitMapu(elementMapy, jsonMapy) {
    const sirkaZobrazenia = window.innerWidth;
    const jeSirokeZobrazenie = sirkaZobrazenia > 768;

    const dataSiete = {
        nodes: new vis.DataSet(jsonMapy.vrcholy),
        edges: new vis.DataSet(jsonMapy.hrany),
    };

    const nastavenia = {
        interaction: {
            hover: true,
        },
        nodes: {
            shape: "box",
            widthConstraint: {
                maximum: 200,
            },
            margin: 10,
            labelHighlightBold: false,
        },
        edges: {
            smooth: {
                type: "dynamic",
                roundness: 1.0,
            },
            width: 0.15,
            arrows: {
                to: {
                    enabled: true,
                },
            },
        },
        layout: {
            hierarchical: {
                enabled: true,
                direction: jeSirokeZobrazenie ? "UD" : "LR",
                sortMethod: "directed",
                nodeSpacing: jeSirokeZobrazenie ? 220 : 100,
                levelSeparation: jeSirokeZobrazenie ? 100 : 250,
            },
        },
        physics: false,
    };

    console.log(jsonMapy);
    pojmova_mapa = new vis.Network(elementMapy, dataSiete, nastavenia);
    return pojmova_mapa;
}

window.addEventListener("load", function () {
    const poznamkyElement = document.getElementById("poznamky");
    const elementMapy = document.getElementById("pojmovaMapa");

    if (poznamkyElement == null || elementMapy == null) return;

    const jsonMapy = vytvoritDataMapy(poznamkyElement);
    vykreslitMapu(elementMapy, jsonMapy);
});


/**
 * Získajte úroveň nadpisu (h1, h2, atď.).
 *
 * @param {HTMLElement} nadpis HTML element nadpisu.
 * @returns {number} Číselná reprezentácia úrovne nadpisu.
 **/
function ziskatLevelNadpisu(nadpis) {
    return parseInt(nadpis.tagName.substring(1), 10);
}
