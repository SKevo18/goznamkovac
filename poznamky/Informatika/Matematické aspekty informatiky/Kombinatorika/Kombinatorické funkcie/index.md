# Kombinatorika

## Čo je to kombinatorika

**Kombinatorika** je ako veľká skladačka. Predstav si, že máš veľa rôznych dielikov a snažíš sa zistiť, koľko rôznych stavebníc môžeš z týchto dielikov poskladať.

V matematike používame kombinatoriku na počítanie, ako sa dajú rôzne veci kombinovať alebo usporiadať.

## Pravidlo súčtu a pravidlo súčinu

### Pravidlo súčtu

Pravidlo súčtu využívame, keď riešime komplexnejšiu kombinatorickú úlohu pri ktorej si ju rozdelíme na menšie časti. Výsledky potom medzi sebou spočítame.

*Príklad:*

**Koľko existuje trojciferných čísel, ktoré sa začínajú číslicou 3 alebo 7?**

Množinu všetkých hľadaných čísel si rozdelíme na dve časti:

1. Čísla, ktoré začínajú číslicou 3;
2. Čísla začínajúce číslicou 7;

Všetky čísla majú 3 cifry, to znamená že platné sú čísla od 100 po 999, t. j.: hľadáme, ako čísla dosadíme na 3 miesta:

$$
\begin{array}
    {|r|r|}
    \hline
    ? &
    ? &
    ? \\
    \hline
\end{array}
$$

- Číslicou 3 začínajú čísla 300, 301, 302, ..., 398, 399. Dokopy ich je 100.
- Číslicou 7 začínajú čísla 700, 701, 702, ..., 798, 799. Je ich tiež 100.

Pre určenie konečného výsledku použijeme kombinatorické pravidlo súčtu. To znamená, že počet všetkých hľadaných trojciferných čísel je $100 + 100 = 200$ čísel.

### Pravidlo súčinu

Pri pravidle súčinu naopak prvky medzi sebou násobíme.

*Príklad:*

**Koľkými spôsobmi si môžeme vybrať tričko a nohavice, ak máme 3 tričká a 4 páry nohavíc.**

Pre výpočet použijeme pravidlo súčinu. Môžme teda mať $3 \times 4 = 12$ rôznych kombinácií oblečenia.

## Faktoriál

**Faktoriál** je ako keby si násobil všetky čísla od 1 po číslo, ktoré máš.

Ak máš napríklad číslo 4, tak faktoriál je $1 \times 2 \times 3 \times 4$. V matematike to zapisujeme ako $4!$, a to je 24.

$$
4! = 1 \times 2 \times 3 \times 4 = 24 \\
$$

$$
\begin{array}
    {|r|r|}\hline
    \text{Výpočet} &
    \text{Výsledok} \\

    \hline 1 \times 2 & 2 \\
    \hline 2 \times 3 & 6 \\
    \hline 6 \times 4 & 24 \\
    \hline 
\end{array}
$$

Faktoriál nám pomáha pochopiť koľko rôznych spôsobov existuje na to, ako môžeme nejaké prvky usporiadať (bez toho, aby sa prvky opakovali).

Napríklad, ak máme 3 písmená (`A`, `B` a `C`) a chceme vedieť koľko spôsobov existuje na to ako ich môžeme usporiadať bez toho aby sa písmená opakovali, použijeme faktoriál. V hlave si vieme predstaviť, že písmená dosádzame do nasledovnej pomyselnej štruktúry:

$$
\begin{array}
    {|r|r|}
    \hline
    ? &
    ? &
    ? \\
    \hline
\end{array}
$$

Na prvé miesto môžeme dosadiť napríklad `A`:

$$
\begin{array}
    {|r|r|}
    \hline
    A &
    ? &
    ? \\
    \hline
\end{array}
$$

Ak sme `A` už použili a nemôže sa opakovať, tak na ďalšie miesta môžeme dosadiť už iba `B` alebo `C`. To znamená, že nám zostali 2 možnosti. Ak dosadíme jedno z písmen, tak na posledné tretie miesto môžeme dosadiť už iba písmeno, ktoré ešte nebolo použité (t. j.: na poslednom mieste máme vždy presne 1 možnosť výberu).
Z toho vyplýva, že s každou možnosťou sa nám spôsoby, akým môžeme vybrať ďalšiu možnosť, zmenší o 1. Jedná sa o princíp faktoriálu.

Odpoveďou na otázku "**Koľkými spôsobmi môžeme usporiadať 3 písmená tak, aby sa neopakovali**" je teda $3! = 1 \times 2 \times 3 = 6$. Správnosť tohto tvrdenia si môžeme overiť v tabuľke:

| A   | B   | C   |
|-----|-----|-----|
| ABC | BAC | CAB |
| ACB | BCA | CBA |

V matematike sa často výpočet faktoriálu zapisuje od najmenšieho čísla po najväčšie, takže $4! = 4 \times 3 \times 2 \times 1$. V skutočnosti nezáleží na tom, v akom poradí čísla napíšeš (pretože násobenie je komutatívne[^1]), takže $4! = 1 \times 2 \times 3 \times 4$ je to isté ako $4! = 4 \times 3 \times 2 \times 1$.

[^1]: Komutatívny zákon o sčítaní hovorí: Veľkosť súčtu nezávisí od poradia sčítancov. Napríklad: $2 \times 3 = 3 \times 2$.

Pre zjednodušenie zápisu môžeme násobenie jednotkou vynechať, takže $4! = 2 \times 3 \times 4$ (ak máme niečo jeden krát, tak je to to isté niečo).

## Kombinatorické funkcie

Kombinatorické funkcie nám pomáhajú počítať rôzne kombinácie vecí. Vo vzorcoch sa vyskytujú vo všeobecnosti 2 pojmy:

- $n$ - počet všetkých prvkov v množine (prvky, napríklad máme 3 písmená: `A`, `B` a `C`)
- $k$ - počet prvkov, ktoré si vyberáme (triedy, napríklad si chceme vybrať 2 písmená)

### Kombinatorické funkcie bez opakovania

#### Variácie bez opakovania

Vzorec pre variácie bez opakovania je nasledovný:

$$
V_k(n) =
\frac{n!}{(n - k)!}
$$

*Príklad:*

Ak máš písmenká `A`, `B`, `C` a chceš vedieť, koľko rôznych dvojpísmenkových "slov" môžeš vytvoriť, použiješ variácie.
Ako napríklad: `A-B`, `B-C`, ale nie `B-B`, pretože nemôžeš použiť to isté písmenko viackrát (preto sú to variácie *bez opakovania*).

#### Permutácie bez opakovania

Permutácie sú ako variácie, ale vždy vyberáme z množiny všetky prvky.

Vzorec pre permutácie bez opakovania je nasledovný:

$$
P(n) =
n!
$$

*Príklad:*

Z písmeniek `A`, `B`, `C` môžeš mať `ABC`, `ACB`, `BAC`, a tak ďalej. Ide o všetky možné poradia, pri ktorých použijeme všetky prvky, *bez opakovania* (viď. príklad pre faktoriál).

#### Kombinácie bez opakovania

Kombinácie sú ako variácie, ale tu nezáleží na poradí. Takže `AB` je to isté ako `BA`. Ak máš `A`, `B`, `C` a chceš si vybrať dve písmenká, pravdepodobne ti nebude záležať v akom poradí ich vyberieš. Ty chceš jednoducho iba dve písmenká a budeš spokojný. Tým pádom môžeš mať `AB`, `AC`, `BC`. Sú to v podstate variácie bez opakovania, mínus prípady kedy sa písmenká len prehodia. Preto je vzorec pre kombinácie bez opakovania nasledovný:

$$
C_k(n) =
{n \choose k} =
\left\{
    \begin{matrix}
        \frac{n!}{(n - k)! \times k!}\,&&\mbox{ pre }n \geq k \geq 0;
        \\
        0\,&&\mbox{inak}\qquad\qquad
    \end{matrix}
\right.
$$

### Kombinatorické funkcie s opakovaním

#### Variácie s opakovaním

Teraz môžeš použiť písmenká viackrát. Ako napríklad: `A-A`, `B-B`.

Vzorec pre variácie s opakovaním je nasledovný:

$$
V´_{(n, k)} (𝑛) =
n^k
$$

#### Permutácie s opakovaním

Keď máš napríklad dve `A` a jedno `B`, permutácie by boli `AAB`, `ABA`, `BAA`.

Vzorec pre permutácie s opakovaním je nasledovný:

$$
P´_{(n_1, n_2, \ldots n_𝑘)} (n) =
\frac{n!}{n_1! \times n_2! \ldots \times n_k!}
$$

Kde:

- $n$ je počet všetkých prvkov;
- $n_k$ je počet prvkov, ktoré sa opakujú $k$ krát;
  - $n_1$ je počet prvkov, ktoré sa opakujú 1 krát;
  - $n_2$ je počet prvkov, ktoré sa opakujú 2 krát;
  - atď...

*Príklad:*

Máme 3 žlté kocky, 1 modrú kocku a 1 červenú kocku a chceme ich poukladať do radu. Koľko spôsobov máme na to, ako ich môžeme poukladať?

*Riešenie:*

Jedná sa o permutácie s opakovaním, kde prvý prvok sa opakuje 3-krát, druhý 1-krát a tretí tiež 1-krát.

Dokopy máme 5 kociek z ktorých môžeme vyberať ($n = 3 + 1 + 1 = 5$).

$$
P´_{(3, 1, 1)} (5) =
\frac{5!}{3! \times 1! \times 1!} =
\frac{120}{6 \times 1 \times 1} =
20 \text{ možností zoradenia kociek}
$$

#### Kombinácie s opakovaním

Podobné kombináciám bez opakovania, ale teraz sa písmená môžu opakovať: `AA`, `BB`.

##### Abc

bla bla

###### Def

aaaa

## Slovné úlohy na kombinatoriku

Predstavme si úlohy, ako napríklad:

- Máš 5 druhov ovocia a chceš si vybrať 3. Koľko rôznych spôsobov je?
- Máš 4 priateľov a chceš poslať 2 listy. Koľko rôznych dvojíc priateľov môžeš vybrať?

V týchto úlohách používame kombinatorické funkcie, aby sme našli odpovede.
