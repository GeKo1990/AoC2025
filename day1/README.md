# 🎄 Advent of Code 2025 — Day 1

## **Secret Entrance**

### Overview

The Elves discovered that the safe protecting the secret entrance to the North Pole has a dial numbered **0–99**. Your puzzle input is a list of rotation commands:

* `L<number>` — rotate left (towards lower numbers)
* `R<number>` — rotate right (towards higher numbers)

The dial starts at **50**.
Moving past 0 wraps to 99, and moving past 99 wraps to 0.

---

## Part 1 — Final-Position Zero Counts

Each rotation moves the dial by a given number of clicks.
A *click* happens at each integer the dial passes.

For part 1, we only count how many times the dial **ends a rotation** pointing at `0`.

### Example

```
L68 → 82
L30 → 52
R48 → 0  ← counts as 1
...
```

In the sample, the dial ends at `0` **three times**.

##  Part 2 — Method `0x434C49434B`

A newly discovered protocol changes the rules:

> Count *every* time the dial points at `0` during rotation —
> not just at the end.

This includes all intermediate clicks.

Example:

* `R48` from 52 → hits zero once (ending at 0)
* `R60` from 95 → hits zero once during rotation
* `L82` from 14 → hits zero once during rotation

In the sample, this method increases the count from **3 → 6**.

### Important Detail

A rotation like `R1000` from 50 causes the dial to pass zero **10 times**, since the dial loops every 100 clicks.

---

## 🛠️ Notes

* The dial math uses modular arithmetic over range `0..99`.
* Part 2 required efficiently counting zero hits without simulating each click.
* The chosen solution computes the first hit and how many full 100-step loops fit.
