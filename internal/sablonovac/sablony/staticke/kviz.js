class Kviz {
    constructor(element = document.getElementById("kviz")) {
        this.element = element;
        this.otazky = this._vsetkyOtazky();
        this.aktualnaOtazka = 0;
    }

    _vsetkyOtazky() {
        const otazky = this.element.querySelectorAll("[data-otazka]");
        if (otazky.length === 0) {
            throw new Error("Kvíz nemá žiadne elementy pre panely otázok.");
        }

        return otazky;
    }

    skontrolovatOdpovede() {
        for (const otazka of this.otazky) {
            const odpovede = otazka.querySelectorAll("[data-odpoved]");

            if (odpovede.length === 0) throw new Error("Otázka nemá žiadne elementy pre odpovede.");

            for (const odpoved of odpovede) {
                if (odpoved.dataset.spravna !== undefined) {
                    odpoved.parentElement.style.backgroundColor = odpoved.checked ? "lightgreen" : "lightcoral";
                } else if (odpoved.checked) {
                    odpoved.parentElement.style.backgroundColor = "lightcoral";
                } else {
                    odpoved.parentElement.style.backgroundColor = "lightgreen";
                }
            }
        }
    }

    istNaOtazku(otazkaCislo) {
        this.otazky[this.aktualnaOtazka].style.display = "none";
        this.aktualnaOtazka = otazkaCislo;
        this.otazky[this.aktualnaOtazka].style.display = "block";
    }

    dalsiaOtazka() {
        this.istNaOtazku(this.aktualnaOtazka + 1);
    }

    predchadzajucaOtazka() {
        this.istNaOtazku(this.aktualnaOtazka - 1);
    }
}

window.onload = () => {
    const kviz = new Kviz();

    document.getElementById("skontrolovat").addEventListener("click", () => {
        kviz.skontrolovatOdpovede();
    });
};
