class Kviz {
    constructor(element = document.getElementById("kviz")) {
        this.element = element;
        this.otazky = this._vsetkyOtazky();
        this.odpovede = {};
        this.aktualnaOtazka = 0;
    }

    _vsetkyOtazky() {
        const otazky = this.element.querySelectorAll(".otazka");
        if (otazky.length === 0) {
            throw new Error("Kvíz nemá žiadne elementy pre panely otázok.");
        }

        return otazky;
    }

    ulozitOdpovede(otazkaElement) {
        const inputy = otazkaElement.querySelectorAll("input");
        const otazkaIndex = this.otazky.indexOf(otazkaElement);

        let odpovede = [];
        for (let input of inputy) {
            odpovede.push(input.value);
        }

        this.odpovede[otazkaIndex] = odpovede;
    }

    istNaOtazku(cislo) {
        this.otazky[this.aktualnaOtazka].style.display = "none";
        this.aktualnaOtazka = cislo;
        this.otazky[this.aktualnaOtazka].style.display = "block";
    }

    dalsiaOtazka() {
        this.istNaOtazku(this.aktualnaOtazka + 1);
    }

    predchadzajucaOtazka() {
        this.istNaOtazku(this.aktualnaOtazka - 1);
    }
}
