# 🎄 Day 6: Trash Compactor — Advent of Code 2025

After helping the Elves in the kitchen, an over-enthusiastic garbage-chute re-enactment goes wrong and you fall straight into a **trash compactor**!
While you're stuck inside, a family of cephalopods offers help—but their youngest needs assistance with her **math homework** first.

The worksheet is presented in a long horizontal strip where **each problem is arranged vertically** and **problems are separated by a full column of spaces**. At the bottom of each problem is the operator (`+` or `*`) that applies to all numbers above it.

Example worksheet:

```
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
```

You can ignore left/right alignment inside each problem.
From this worksheet, the four problems are:

* `123 * 45 * 6 = 33210`
* `328 + 64 + 98 = 490`
* `51 * 387 * 215 = 4243455`
* `64 + 23 + 314 = 401`

Adding all answers together:

```
33210 + 490 + 4243455 + 401 = 4,277,556
```

Your task:
**Unroll the full worksheet, parse all problems, evaluate them, and compute the grand total.**

---

## ⭐ Part One

Solve all vertically-stacked problems and compute the sum of their results.

> **Example result:** `4,277,556`

Your puzzle answer: **7644505810277**

---

## ⭐ Part Two

The cephalopods reveal that their math is actually written **right-to-left in columns**.
Numbers are now formed by reading each **digit column from right to left**, and within each digit column **top to bottom**. Problems are still separated by columns of only spaces, and the last row still contains the operator.

Using the same example:

```
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
```

Reading right-to-left:

* Rightmost problem → digits: `4`, `431`, `623`
  → `4 + 431 + 623 = 1058`

* Second from right → digits: `175`, `581`, `32`
  → `175 * 581 * 32 = 3,253,600`

* Third → digits: `8`, `248`, `369`
  → `8 + 248 + 369 = 625`

* Leftmost → digits: `356`, `24`, `1`
  → `356 * 24 * 1 = 8,544`

Grand total:

```
1058 + 3,253,600 + 625 + 8,544 = 3,263,827
```

Your task:
**Re-interpret the worksheet with cephalopod rules and compute the new grand total.**