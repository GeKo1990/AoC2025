# 🎄 Day 7: Laboratories — Advent of Code 2025

After escaping the trash compactor, you end up in a **teleporter research wing** at the North Pole.  
A broken **tachyon manifold** in a leaking teleporter needs diagnostics – and of course, that means more weird grid-based physics puzzles.

Your input is a diagram of the manifold: a 2D grid where

- `S` = starting point of the tachyon beam  
- `.` = empty space  
- `^` = splitter  

A classical tachyon beam:

- always moves **downward**,
- passes freely through `.`,
- and when it hits a `^`, the **incoming beam stops** and **two new beams** continue from the **immediate left and right** of the splitter (still moving downward).

Example manifold:

```text
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
````

As the beam moves, splitters can create more beams, which may interact or overlap further down in the grid.

---

## ⭐ Part One

Simulate a **classical tachyon beam**:

1. Start from `S`, moving straight down.
2. When a beam enters a splitter `^`:

   * the beam **stops** at the splitter,
   * two new beams appear, one to the **left** and one to the **right** of the splitter, both continuing downward.
3. Beams continue until they either:

   * reach another splitter (and split again), or
   * exit the grid at the bottom.

In the example, this process ultimately produces this beam pattern:

```text
.......S.......
.......|.......
......|^|......
......|.|......
.....|^|^|.....
.....|.|.|.....
....|^|^|^|....
....|.|.|.|....
...|^|^|||^|...
...|.|.|||.|...
..|^|^|||^|^|..
..|.|.|||.|.|..
.|^|||^||.||^|.
.|.|||.||.||.|.
|^|^|^|^|^|||^|
|.|.|.|.|.|||.|
```

In this example, the beam is split a total of **21 times**.

Your task:

> **Count how many times the beam is split** (i.e. how often a beam actually hits a `^` and produces new beams).

---

## ⭐ Part Two

Now the twist: the tachyon manifold is actually **quantum**.

Instead of many classical beams, there is only **one tachyon particle**.
Each time it reaches a splitter `^`, it doesn’t choose left *or* right – it takes **both paths**, in separate timelines.

The recommended interpretation is **many-worlds quantum mechanics**:

* At every splitter:

  * In one timeline, the particle goes left.
  * In another timeline, it goes right.
* Time itself branches, creating a growing number of possible **timelines**.

Your goal is no longer, "How often was the beam split?"
Instead, you must determine:

> **How many different timelines exist after the tachyon has completed all of its possible journeys through the manifold?**

In the given example, considering all possible left/right choices at every splitter, the particle can end up on **40 different timelines**.

---

## 🔍 Summary

* **Part 1**:
  Simulate classical beams.
  Count how many times beams hit a splitter `^` and split.

* **Part 2**:
  Interpret the manifold quantum-mechanically.
  A single particle branches at every splitter.
  Count how many **distinct timelines** result from all possible paths through the manifold.

This puzzle is a nice mix of:

* grid simulation (Part 1) and
* combinatorial path counting / many-worlds interpretation (Part 2).