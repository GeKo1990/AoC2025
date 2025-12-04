# 🎄 Advent of Code 2025 — Day 3

## **Lobby**

You descend a short staircase and enter a surprisingly vast lobby, quickly passing through the security checkpoint. When you reach the main elevators, you notice that each one has a red light above it — they’re all offline.

An Elf nearby apologizes while tinkering with a control panel:

> “Some kind of electrical surge fried them. I’ll try to get them back online soon.”

You explain that you need to get deeper underground.

> “Well, you *could* take the escalator down to the printing department…
> if it weren’t also offline.”

Luckily, the escalator isn’t broken; it just needs power.
Nearby are emergency batteries, each labeled with a joltage rating **1–9**.
Your puzzle input lists these batteries. For example:

```
987654321111111
811111111111119
234234234234278
818181911112111
```

Each **line** is a battery **bank**.
From each bank, you must turn on **exactly two batteries**, keeping them in order.
The joltage produced is the **two-digit number** made from those digits:

* Bank `12345`, choosing batteries 2 and 4 → joltage **24**
* You **cannot** rearrange digits.

### 🧩 Part One

Find the **largest** possible two-digit joltage for each bank.

Example (with input above):

* `987654321111111` → **98**
* `811111111111119` → **89**
* `234234234234278` → **78**
* `818181911112111` → **92**

Total output joltage for the example:

```
98 + 89 + 78 + 92 = 357
```

---

## ⭐ Part Two

The Elf presses the “**joltage limit safety override**” button again and again until the escalator allows higher power.

Now, for each bank, you must turn on **exactly twelve batteries** (still in order).
The produced joltage is the 12-digit number formed by those batteries.

Using the same example input:

* `987654321111111` → **987654321111**
* `811111111111119` → **811111111119**
* `234234234234278` → **434234234278**
* `818181911112111` → **888911112111**

Total example joltage:

```
987654321111
+ 811111111119
+ 434234234278
+ 888911112111
= 3121910778619
```

Your task: compute the **total output joltage** using this new rule.
