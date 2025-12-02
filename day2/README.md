# Day 2: Gift Shop 

## Part 1

You get inside and take the elevator to its only other stop: the gift shop.
"Thank you for visiting the North Pole!" gleefully exclaims a nearby sign.

A young Elf accidentally added **invalid product IDs** into the database.
Your task is to scan ID ranges and find all IDs that are made of:

> **some sequence of digits repeated *exactly twice***
>
> Examples:
>
> * `55` → `5` + `5`
> * `6464` → `64` + `64`
> * `123123` → `123` + `123`

IDs do **not** include leading zeroes (`0101` is not valid input).

### Example Input

```
11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124
```

### Example Analysis (Part 1 rules)

| Range                 | Invalid IDs |
| --------------------- | ----------- |
| 11–22                 | 11, 22      |
| 95–115                | 99          |
| 998–1012              | 1010        |
| 1188511880–1188511890 | 1188511885  |
| 222220–222224         | 222222      |
| 1698522–1698528       | *(none)*    |
| 446443–446449         | 446446      |
| 38593856–38593862     | 38593859    |
| rest                  | *(none)*    |

Sum (example): **1227775554**

--

## Part 2

The clerk takes a closer look — the Elf was *even more* chaotic.

Now, an ID is invalid if it is made of:

> **some sequence of digits repeated at least twice**
> *(2×, 3×, 4×, …)*

Examples:

* `12341234` → `1234` × 2
* `123123123` → `123` × 3
* `1111111` → `1` × 7
* `1212121212` → `12` × 5

### Updated Example (Part 2)

| Range                 | Invalid IDs |
| --------------------- | ----------- |
| 11–22                 | 11, 22      |
| 95–115                | 99, 111     |
| 998–1012              | 999, 1010   |
| 1188511880–1188511890 | 1188511885  |
| 222220–222224         | 222222      |
| 1698522–1698528       | *(none)*    |
| 446443–446449         | 446446      |
| 38593856–38593862     | 38593859    |
| 565653–565659         | 565656      |
| 824824821–824824827   | 824824824   |
| 2121212118–2121212124 | 2121212121  |

Sum (example): **4174379265**

## Code

My Go solution supports both modes:

```bash
go run main.go 1 < input.txt   # Part 1
go run main.go 2 < input.txt   # Part 2
```

---

