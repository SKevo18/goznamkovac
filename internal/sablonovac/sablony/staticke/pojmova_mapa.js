/**
 * Vytvorí JSON pojmovej mapy z obsahu stránky.
 *
 * @param {HTMLElement} poznamkyElement Element, ktorý obsahuje poznámky.
 * @returns {object} JSON reprezentácia pojmovej mapy.
 **/
function vytvoritDataMapy(poznamkyElement) {
    const nadpisy = poznamkyElement.querySelectorAll(
        "h1:not(#obsah), h2, h3, h4, h5, h6"
    );
    let mapa = { vrcholy: [], hrany: [] };

    // Zoznam, ktorý uchováva posledný nadpis na každej úrovni
    let poslednyNadpisNaUrovni = [];

    for (let i = 0; i < nadpisy.length; i++) {
        const nadpis = nadpisy[i];
        const aktualnyLevel = ziskatLevelNadpisu(nadpis);

        const idVrchola = i + 1;
        const nazov = nadpis.innerText.slice(2);

        const idRodica = poslednyNadpisNaUrovni[aktualnyLevel - 1] || 1;
        mapa.vrcholy.push({
            id: idVrchola,
            label: nazov,
            color: generovatFarbu(idRodica * aktualnyLevel),
        });

        if (aktualnyLevel > 1) {
            mapa.hrany.push({
                from: idRodica,
                to: idVrchola,
            });
        }

        // Aktualizujte posledný nadpis na úrovni
        poslednyNadpisNaUrovni[aktualnyLevel] = idVrchola;

        // Vymažte všetky nadpisy na vyšších úrovniach
        for (
            let j = aktualnyLevel + 1;
            j < poslednyNadpisNaUrovni.length;
            j++
        ) {
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
            keyboard: {
                enabled: true,
                bindToWindow: false,
                autoFocus: false
            }
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
                roundness: 0.5,
            },
            width: 0.75,
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
                shakeTowards: "roots",
            },
        },
        physics: false,
    };

    console.log(jsonMapy);
    pojmova_mapa = new vis.Network(elementMapy, dataSiete, nastavenia);
    return pojmova_mapa;
}

/**
 * Získajte úroveň nadpisu (h1, h2, atď.).
 *
 * @param {HTMLElement} nadpis HTML element nadpisu.
 * @returns {number} Číselná reprezentácia úrovne nadpisu.
 **/
function ziskatLevelNadpisu(nadpis) {
    return parseInt(nadpis.tagName.substring(1), 10);
}

/**
 * Generuje náhodné číslo z čísla (seed). Rovnaký seed vždy vygeneruje rovnaké "náhodné" číslo v určenom rozsahu.
 *
 * https://stackoverflow.com/a/63599906
 * 
 * @param {number} seed Číslo, z ktorého sa generuje náhodné číslo.
 * @param {number[]} rozsah Rozsah, v ktorom sa náhodné číslo generuje.
 * @returns {number} Náhodné číslo.
 **/
function nahodneCislo(seed, rozsah = [0, 255]) {
    seed = String(seed)
        .split("")
        .reduce((c, n) => (n != 0 ? c * n : c * c));

    let od = rozsah[0];
    let po = rozsah[1];

    while (seed < rozsah[0] || seed > rozsah[1]) {
        if (seed > rozsah[1]) seed = Math.floor(seed / po--);
        if (seed < rozsah[0]) seed = Math.floor(seed * od++);
    }

    return seed;
}

/**
 * Generuje náhodnú farbu z čísla. Rovnaké číslo vždy vygeneruje rovnakú farbu.
 *
 * @param {number} cislo Číslo, z ktorého sa generuje farba.
 * @returns {string} Farba v RBG formáte.
 **/
function generovatFarbu(seed=1, svetla=true) {
    let od = svetla ? 125 : 0
    let po = svetla ? 255 : 125
  
    r = nahodneCislo(seed + 11, [od, po])
    g = nahodneCislo(seed + 12, [od, po])
    b = nahodneCislo(seed + 13, [od, po])

    return `rgb(${r}, ${g}, ${b})`;
}
