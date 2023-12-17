# Kombinatorika

## Kombinatorické funkcie

### Permutácie

Permutácia je usporiadanie množiny prvkov do určitého poradia. Počet všetkých permutácií množiny s `n` prvkami je `n!` (n-faktoriál).

$$
P(n) = n!
$$

### Variácie

Variácie sú usporiadané výbery prvkov z danej množiny. Počet variácií `k` prvkov z množiny `n` prvkov je daný vzorcom:

$$
V(n, k) = \frac{n!}{(n-k)!}
$$

### Kombinačné čísla

Kombinačné číslo, označované ako `C(n, k)` alebo `n` nad `k`, predstavuje počet možností, ako vybrať `k` prvkov z `n` prvkov bez ohľadu na poradie. Vypočíta sa podľa vzorca:

$$
C(n, k) = \frac{n!}{k!(n-k)!}
$$

## Grafické znázornenie kombinatorických funkcií

### Graf permutácií

```mermaid
graph LR
    A[N prvkov] -->|Permutácie| B[n! možností]
```

### Graf variácií

```pikchr
arrow right 200 from "n prvkov Variácie"
box same "V(n, k)"
arrow right 200 "k-prvkové usporiadania"
```

### Graf kombinácií

```mermaid
graph TD
    A[N prvkov] -->|Kombinácie| B[C(n, k) možností]
```
