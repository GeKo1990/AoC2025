# 🎄 Advent of Code 2025 — Day 4

## **Printing Department**

You ride the escalator down into the **Printing Department**, where large rolls of paper (`@`) are stacked everywhere. The Elves are preparing decorations, but you need their forklifts to break through a wall so you can continue deeper underground.

To help them free up time, you must optimize how forklifts access paper rolls.

Your puzzle input is a **grid** showing the location of each roll (`@`) and empty space (`.`).
Example:

```
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
```

---

## ⭐ Part 1 — Available Paper Rolls

A forklift can access a roll of paper **only if fewer than four adjacent positions** (out of the 8 surrounding cells) contain other rolls.

For the example above, 13 rolls meet this condition. These have been marked with `x`:

```
..xx.xx@x.
x@@.@.@.@@
@@@@@.x.@@
@.@@@@..@.
x@.@@@@.@x
.@@@@@@@.@
.@.@.@.@@@
x.@@@.@@@@
.@@@@@@@@.
x.x.@@@.x.
```

Your task:

> **Count the number of rolls (`@`) that have fewer than four `@` neighbors.**

---

## ⭐ Part 2 — Removing Rolls Iteratively

Once a roll can be accessed, a forklift can **remove** it.
After removing rolls, the grid changes — meaning **more** rolls may become accessible.

Repeat this process:

1. Find all accessible rolls (those with `< 4` `@` neighbors).
2. Mark/remove them.
3. Update the grid.
4. Continue until no more rolls can be accessed.

In the example, this happens in several waves:

* First removal: 13 rolls
* Second removal: 12 rolls
* Third removal: 7 rolls
* …
* Final removal: 1 roll

A total of **43** rolls are removed in the example.

Your task:

> **Simulate this recursive removal process. How many total rolls are removed?**

---

### 📝 Notes

* Adjacency uses all **8 neighboring cells** (N, NE, E, SE, S, SW, W, NW).
* Removal happens in **waves**, not one-by-one.
* The final count includes **every roll removed across all waves**.
