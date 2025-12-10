# 🎬 Day 9: Movie Theater

The movie theater floor is a large grid containing several **red tiles**.
Using these red tiles, the Elves want to determine the **largest possible axis-aligned rectangle** whose **opposite corners are red tiles**.

---

## ⭐ Part One

You are given a list of coordinates for all **red tiles**.
Any two red tiles may be chosen as opposite corners of a rectangle.

For any pair of red points `(x1, y1)` and `(x2, y2)`:

```
width  = |x1 - x2| + 1
height = |y1 - y2| + 1
area   = width * height
```

Your task is to find the **maximum possible area** over all red–red pairs.

---

## ⭐ Part Two

The Elves now reveal an important constraint:

> The rectangle may only include **red or green tiles**.

Where do green tiles come from?

* The input list of red tiles describes a **closed loop**.
* Each red tile is connected to the next by a **straight horizontal or vertical line** of green tiles.
* The list wraps, so the last red tile connects back to the first.
* **All tiles inside this loop are also green**.

This forms an **orthogonal polygon** composed of red and green tiles.

### Valid rectangle criteria

A rectangle is valid if:

1. Its opposite corners are red tiles (same as Part 1).
2. **Every tile inside the rectangle is red or green**, meaning:

   * the rectangle lies fully **inside or on the boundary** of the red+green polygon.

### Goal

Find the **largest rectangle** that satisfies these constraints.

---

## 🧮 Summary

* **Part 1:** largest rectangle using *any* two red tiles as corners
* **Part 2:** same, but rectangle must lie completely inside the red+green region defined by the loop

Both require checking all red–red pairs, but Part 2 introduces a geometric containment test for each rectangle.
