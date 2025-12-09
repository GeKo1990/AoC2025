# 🎄 Day 8: Playground — Advent of Code 2025

After repairing the teleporter, you rematerialize inside a vast underground **playground filled with hanging electrical junction boxes**.
Elves are planning to connect these boxes with strings of lights. Most boxes do not provide power — but when **two boxes are connected**, electricity can flow between them, making them part of the same **circuit**.

Your input lists the **3D coordinates (X,Y,Z)** of every junction box.

Example:

```
162,817,812
57,618,57
906,360,560
...
```

Each line describes the location of one junction box in 3D space.

---

## ⭐ Problem Interpretation

* Each **junction box** = a **node** in 3D space
* Any connection between two boxes = an **edge**
* The length of the edge is the **Euclidean distance** between the two coordinates
* A **circuit** = a **connected component** of the graph

The Elves want to connect boxes starting from the **shortest straight-line distances**, gradually forming larger and larger circuits.

This is effectively the same as **sorting all possible edges by distance and processing them in ascending order** — exactly what **Kruskal’s algorithm** does when building a Minimum Spanning Tree (MST).
But the puzzle stops the algorithm early or uses the final edge of the process for Part 2.

---

# ⭐ Part One

> Consider the **1000 shortest connections** between boxes.
> Each connection is a pair of boxes sorted by distance.
> You attempt to connect them in order — regardless of whether they merge circuits.

This means:

* For each of the **first 1000 shortest pairs**, attempt to union the two boxes.
* Some pairs do nothing (if they’re already in the same circuit).
* After considering all 1000 closest pairs, count the sizes of all resulting circuits.
* Take the **three largest** circuits and multiply their sizes.

Your program must output:

```
Part 1: largest component sizes A, B, C -> product = X
```

---

# ⭐ Part Two

> Continue connecting pairs (still in increasing distance order)
> until **all junction boxes are part of one single circuit**.

During this process, record the **last successful connection** — i.e., the final edge whose union merges the last two remaining components.

Let that final connection be between boxes:

```
(X1, Y1, Z1) and (X2, Y2, Z2)
```

The puzzle asks for:

> **Multiply the X coordinates of these two boxes:**
> `X1 * X2`

Your program outputs it as:

```
Part 2: last connection between (X1, Y1, Z1) and (X2, Y2, Z2)
Part 2: X product = P
```

---

## 🧠 Algorithms Used

### Kruskal’s Algorithm (conceptual)

We:

1. Compute all pairwise distances
2. Sort all edges by distance
3. Process edges in increasing order
4. Use **Union-Find** to merge circuits
5. Track component sizes

More here:
[https://en.wikipedia.org/wiki/Kruskal%27s_algorithm](https://en.wikipedia.org/wiki/Kruskal%27s_algorithm)

### Union-Find (Disjoint Set Union)

Efficiently tracks which boxes belong to which circuit.

Operations:

* `find(x)` → returns the representative of x’s circuit
* `union(a, b)` → merges two circuits if they aren’t already connected

With path compression + union by size, these operations are nearly O(1).

More here:
[https://en.wikipedia.org/wiki/Disjoint-set_data_structure](https://en.wikipedia.org/wiki/Disjoint-set_data_structure)

---

## 🔧 Notes About Implementation

* Distances use **squared distance**, not full Euclidean distance — avoids `sqrt`, preserves ordering.
* Edges are sorted once, used for both Part 1 and continuation in Part 2.
* Part 1 and Part 2 share the same Union-Find state; Part 2 continues where Part 1 stopped.
* Last merging edge gives the answer for Part 2.