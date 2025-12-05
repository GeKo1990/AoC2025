# 🌲 Day 5: Cafeteria — Fresh or Spoiled Ingredients?

The Elves’ new inventory management system is causing chaos in the cafeteria.
Your task is to help them determine which ingredient IDs are **fresh** and which are **spoiled**.

The database consists of:

* A list of **fresh ingredient ID ranges** (inclusive)
* A **blank line**
* A list of **available ingredient IDs**

Example:

```
3-5
10-14
16-20
12-18

1
5
8
11
17
32
```

A range like `3-5` covers the IDs `3, 4, 5`.
Ranges may **overlap** — an ID is fresh if it appears in **any** range.

---

## ⭐ Part 1 — Count Fresh Available Ingredients

For each available ingredient, determine whether it falls within **at least one** fresh range.

Example evaluation:

| ID | Fresh? | Reason                    |
| -- | ------ | ------------------------- |
| 1  | No     | Not in any range          |
| 5  | Yes    | In range 3–5              |
| 8  | No     | Not in any range          |
| 11 | Yes    | In range 10–14            |
| 17 | Yes    | In ranges 16–20 and 12–18 |
| 32 | No     | Not in any range          |

In the example, **3** IDs are fresh.

---

## ⭐ Part 2 — How Many IDs Are Fresh in Total?

Here the available IDs no longer matter.

You must determine **how many unique ingredient IDs** are considered fresh across **all ranges combined**.

Example ranges:

```
3-5
10-14
16-20
12-18
```

Combined, these ranges cover:

```
3–5, 10–20
```

Which corresponds to:

```
3, 4, 5,
10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20
```

Total fresh IDs: **14**

---

## 📝 Summary

| Part  | Task                                        | Output              |
| ----- | ------------------------------------------- | ------------------- |
| **1** | Count fresh IDs among the available IDs     | **529**             |
| **2** | Count all distinct IDs covered by any range | **344260049617193** |
