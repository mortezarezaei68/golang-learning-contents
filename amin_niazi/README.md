This problem, with the mandatory condition that "All products must be placed," is, no matter how I look at it, a **Mixed Integer Programming (MIP)** problem. more specifically, a **0/1 Integer Programming** problem.

MIP problems are NP-Hard. The scale of this particular problem is such that it cannot be solved on a reasonable hardware in a reasonable amount of time. Writing the MIP code from scratch is also non-trivial, and there is no pure-Go package for it available on GitHub.

If the "all products" condition is strict, the only viable option in Go seems to be the `glop` package, which is a wrapper around `lp_solve` that must be installed separately on the system.

If we could leave some items unplaced, we could model it as a **Multi-Dimensional Multi-Knapsack** problem (where different `shelf level * weight` represent different "dimensions"). However, this is also NP-Hard.

If we simplify further by assuming weight equals value (dropping the cost/dimension), we can model it as a **Multiple Knapsack** problem. For the knapsack approach, we need to define the capacity for each "knapsack" (shelf). The maximum weight per shelf wasn't explicitly given, but typically in such problems, we can assume it's the maximum items weight, so ranging same as `w_i`. We are given:
`w_i = weight of product i (1 ≤ w_i ≤ 10⁶)`

This weight range is unrealistically large. It implies a warehouse system where a one-gram item and a one-ton item could be placed on the same shelf, which is not OK I guess. The optimized versions of the algorithm that are beyond me, require at least 7.5 GB of space. Or simple versions that I can at least read along takes 4 terabytes iguess.

A practical heuristic would be to treat each shelf as a separate knapsack and fill them from the bottom up, hoping that we might place all items. The result would just be *an* answer, with no claim to optimality. However, since the problem is NP-Hard and if we find a feasible solution that places all items, no one can easily evaluate its optimality, which is actually convenient.

The new condition:
`"A heavier product should ideally be placed on a shelf with a larger capacity."`
suggests we can sort the shelves by capacity and assign items accordingly. This could be a helpful heuristic, but its importance is diminished by the condition:
`"Only the maximum score matters."`
Furthermore, it seems to conflict with the structural guideline:
`"A heavier product should ideally be placed lower to avoid collapse and reduce stress."`

Therefore, with these conflicting conditions and constraints, I cannot even formulate the MIP condition set.

Am I looking at this incorrectly, or is there a key insight I'm missing? Is there any guidance available?

### Input
```
4 2
3 2 5 4
7 10
```

### Output
```
-------------------------------
Level: 1 - filled: 7
------------ items ------------
P3 (w=5), P2 (w=2)
-------------------------------
Level: 2 - filled: 7
------------ items ------------
P4 (w=4), P1 (w=3)
------------------------------
Total Cost: 21
```