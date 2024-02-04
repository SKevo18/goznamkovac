class Kviz {
    constructor(element = document.getElementById("kviz")) {
        this.element = element;
        this.otazky = this._vsetkyOtazky();
        this.aktualnaOtazka = 0;

        this.postup = document.getElementById("postup");
        this.postup.max = this.otazky.length;
        this.postupText = document.getElementById("postupText");

        this.tlcSkontrolovatOdpovede = document.getElementById("skontrolovatOdpovede");
        this.tlcDalsiaOtazka = document.getElementById("dalsiaOtazka");
        this.tlcPredchadzajucaOtazka = document.getElementById("predchadzajucaOtazka");

        this._nastavitUdalosti();
        this.istNaOtazku(this.aktualnaOtazka);
    }

    _nastavitUdalosti() {
        this.tlcSkontrolovatOdpovede.addEventListener("click", () => {
            this.skontrolovatOdpovede();
        });

        this.tlcPredchadzajucaOtazka.addEventListener("click", () => {
            this.predchadzajucaOtazka();
        });

        this.tlcDalsiaOtazka.addEventListener("click", () => {
            this.dalsiaOtazka();
        });
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

    istNaOtazku(indexOtazky) {
        this.otazky[this.aktualnaOtazka].style.display = "none";

        this.aktualnaOtazka = indexOtazky;
        this.otazky[this.aktualnaOtazka].style.display = "block";

        this.postupText.innerText = `${this.aktualnaOtazka + 1} / ${this.otazky.length}`;
        this.postup.value = this.aktualnaOtazka + 1;

        if (this.aktualnaOtazka === 0) {
            this.tlcPredchadzajucaOtazka.style.display = "none";
        } else {
            this.tlcPredchadzajucaOtazka.style.display = "inline";
        }

        if (this.aktualnaOtazka === this.otazky.length - 1) {
            this.tlcDalsiaOtazka.style.display = "none";
        } else {
            this.tlcDalsiaOtazka.style.display = "inline";
        }
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
};
