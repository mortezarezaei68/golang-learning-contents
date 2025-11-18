# 🏆 Coding Challenge – Maximize Shelf Stability Score

## 📦 Problem Description

You are designing an automated product shelving system for a large Amazon warehouse.  
There are **N products**, and each product has a **weight**.

The warehouse shelves are arranged vertically:

- Shelf **level 1** is the **lowest**
- Shelf **level 2** is above it
- Shelf **level 3** is above level 2  
  …and so on

A **heavier product should ideally be placed lower** to avoid collapse and reduce stress.

---

## 🎯 Objective

You must assign every product to exactly **one shelf level**.

Each level has a **maximum total weight capacity**.

Your goal is to **maximize the total stability score**, defined as:

```
Score = Σ (product_weight × shelf_level)
```

This means:

- Heavier products should ideally go to lower levels
- Lighter products can be placed higher

---

## 📥 Input Format

```
N L
w1 w2 w3 ... wN
c1 c2 c3 ... cL
```

Where:

- `N` = number of products (1 ≤ N ≤ 2000)
- `L` = number of shelf levels (1 ≤ L ≤ 2000)
- `wi` = weight of product i (1 ≥ wi ≤ 10⁶)
- `ci` = max weight capacity of shelf level i

---

## 📤 Output Format

A **single integer**:

```
Maximum achievable stability score
```

---

## 🧾 Example

### Input
```
4 2
3 2 5 4
7 10
```

### Valid Assignment

| Product | Weight | Level |
|---------|--------|-------|
| P1      | 3      | 1 |
| P2      | 2      | 1 |
| P3      | 5      | 2 |
| P4      | 4      | 2 |

### Score Calculation

```
(3×1) + (2×1) + (5×2) + (4×2)
= 3 + 2 + 10 + 8
= 23
```

### Output
```
23
```

---

## ⛔ Constraints

- All products **must** be placed.
- Shelf weight **must not exceed** capacity.
- Products may be placed in **any order**.
- Only the **maximum score** matters.
- A heavier product should ideally be placed on a shelf with a larger capacity.

