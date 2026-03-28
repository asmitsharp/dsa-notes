# 🔁 DSA Pattern: Backtracking
### *Interview-Focused Notes | FAANG Mentor Style*

---

## 1. 🧭 Pattern Overview

### What is Backtracking?

Backtracking is a **recursive technique** where you build a solution **step by step**, and whenever you realize the current path can't lead to a valid answer, you **undo your last choice** (backtrack) and try a different one.

Think of it as **"try → fail → undo → try again"**.

### Why is it used? (Intuition)

When a problem asks you to find **all possible combinations, permutations, subsets, or arrangements**, you can't use a simple loop — there are too many nested choices. Backtracking lets you explore all of them **systematically** without missing any, and **prunes dead ends early** to avoid wasted work.

### 🌍 Real-World Analogy

Imagine navigating a **maze**. At every fork, you pick a direction. If you hit a dead end, you walk back to the last fork and try a different path. You keep doing this until you either find the exit or exhaust all paths. That's exactly backtracking.

### ⏱️ Time Complexity Intuition

Backtracking is **inherently exponential** — but that's okay, because the problem itself is exponential (you're exploring all possibilities). The key is that **pruning cuts down bad branches early**, making the actual runtime much better than the worst case.

A rough mental model:

| Problem Type | Worst-Case Complexity |
|---|---|
| Subsets of N elements | O(2ⁿ) — each element is either in or out |
| Permutations of N elements | O(N!) — N choices × (N-1) choices × ... |
| Combinations (N choose K) | O(N choose K × K) |

---

## 2. 💡 Core Idea (MOST IMPORTANT)

### Key Insight #1: The Recursion Tree

Every backtracking solution is really just a **tree traversal**. Each node in the tree is a "state" (your current partial solution). Each edge is a "choice." Leaf nodes are either valid complete solutions or dead ends.

```
                    []
              /     |      \
           [1]     [2]     [3]
          / \       |
       [1,2] [1,3] [2,3]
        |
     [1,2,3]
```

Your job is to **explore this tree**, collect valid leaves, and skip invalid branches.

### Key Insight #2: Choose → Explore → Unchoose

Every backtracking function has exactly three steps:

```
1. CHOOSE   → add an element to your current path
2. EXPLORE  → recurse deeper
3. UNCHOOSE → remove that element (backtrack)
```

This "undo" step is what makes backtracking different from normal recursion. Without it, choices from one branch contaminate other branches.

### 🔑 The Most Important Line to Remember

```go
// After recursing, UNDO your choice
path = path[:len(path)-1]  // or mark[i] = false, or any undo operation
```

This single line is the soul of backtracking. Never forget it.

---

## 3. 🚦 When to Use This Pattern

| Signal in Question | What It Means |
|---|---|
| "Find **all** combinations / subsets / permutations" | Explore entire search space |
| "Generate all valid arrangements" | Build solutions incrementally |
| "Is there a way to..." (existence check) | Early return on first valid path |
| Constraints are small (N ≤ 20) | Exponential time is acceptable |
| "Count all possible ways" | DFS + backtrack, count at leaf |
| Input involves placing/assigning items | Try each position, backtrack on conflict |
| Words like "select", "choose", "arrange" | Classic backtracking signal |

---

## 4. 🗂️ Types / Variants of This Pattern

### Type 1: Subsets / Combinations
You pick a **subset** of elements. Each element can either be included or excluded.
- The start index increases to avoid duplicates.
- Used in: Subsets, Combination Sum, Combinations.

### Type 2: Permutations
You arrange all (or some) elements in every possible order.
- A `used[]` array tracks which elements are already in the current path.
- Used in: Permutations, Permutations II.

### Type 3: Decision / Constraint Satisfaction
You place items on a board or assign values, and check rules at each step.
- If a placement violates a constraint, skip and try next.
- Used in: N-Queens, Sudoku Solver, Word Search.

### Type 4: String / Path Building
You build a string character by character, checking validity at each step.
- Often involves pruning by checking open/close balance or character counts.
- Used in: Generate Parentheses, Palindrome Partitioning.

---

## 5. 🗃️ Data Structures Used

| Data Structure | Why It's Used | Key Property |
|---|---|---|
| **Slice / Array (path)** | Stores the current partial solution | Append on choose, slice on unchoose — O(1) amortized |
| **Boolean Array (used[])** | Tracks which elements are already picked (permutations) | O(1) lookup, easy to toggle |
| **Result Slice (res)** | Stores all valid solutions | Append deep copy of path when solution is found |
| **Recursion Call Stack** | Implicitly manages the DFS traversal | Each frame holds current state |
| **2D Boolean Grid** | For board problems (N-Queens, Sudoku, Word Search) | Mark/unmark cells as you recurse |

> **When to use `used[]` vs `start index`?**
> Use `start index` when order doesn't matter (subsets/combinations).
> Use `used[]` when order matters (permutations).

---

## 6. 📋 Core Templates (VERY IMPORTANT)

### Template 1: Subsets / Combinations

```go
func backtrack(nums []int, start int, path []int, res *[][]int) {
    // ✅ Every state is a valid subset — collect it
    temp := make([]int, len(path))
    copy(temp, path)
    *res = append(*res, temp)

    for i := start; i < len(nums); i++ {
        path = append(path, nums[i])       // CHOOSE
        backtrack(nums, i+1, path, res)    // EXPLORE (i+1 avoids reuse)
        path = path[:len(path)-1]          // UNCHOOSE ← THE KEY LINE
    }
}
```

### Template 2: Permutations

```go
func backtrack(nums []int, used []bool, path []int, res *[][]int) {
    // ✅ Base case: used all elements
    if len(path) == len(nums) {
        temp := make([]int, len(path))
        copy(temp, path)
        *res = append(*res, temp)
        return
    }

    for i := 0; i < len(nums); i++ {
        if used[i] {
            continue // skip already-used elements
        }
        used[i] = true                     // CHOOSE
        path = append(path, nums[i])
        backtrack(nums, used, path, res)   // EXPLORE
        path = path[:len(path)-1]          // UNCHOOSE
        used[i] = false
    }
}
```

### Template 3: Constraint Satisfaction (e.g., N-Queens style)

```go
func backtrack(board [][]byte, row int, res *[][]string) {
    // ✅ Base case: placed queens in all rows
    if row == len(board) {
        *res = append(*res, buildSolution(board))
        return
    }

    for col := 0; col < len(board[0]); col++ {
        if !isValid(board, row, col) {
            continue // PRUNE: skip invalid placements
        }
        board[row][col] = 'Q'              // CHOOSE
        backtrack(board, row+1, res)       // EXPLORE
        board[row][col] = '.'              // UNCHOOSE
    }
}
```

---

## 7. 🧪 Step-by-Step Example: Subsets of [1, 2, 3]

**Goal:** Find all subsets of `[1, 2, 3]`.

We call `backtrack(nums, start=0, path=[])`.

```
Call: start=0, path=[]         → collect []
  Pick 1 → path=[1]
    Call: start=1, path=[1]    → collect [1]
    Pick 2 → path=[1,2]
      Call: start=2, path=[1,2]  → collect [1,2]
      Pick 3 → path=[1,2,3]
        Call: start=3, path=[1,2,3] → collect [1,2,3], loop ends
      Unpick 3 → path=[1,2]
    Unpick 2 → path=[1]
    Pick 3 → path=[1,3]
      Call: start=3, path=[1,3] → collect [1,3], loop ends
    Unpick 3 → path=[1]
  Unpick 1 → path=[]
  Pick 2 → path=[2] ...
  (continues for [2], [2,3], [3])
```

**Final result:** `[[], [1], [1,2], [1,2,3], [1,3], [2], [2,3], [3]]`

**The magic:** By increasing `start` on each call, we never pick elements to the left of our current position, which naturally avoids duplicate subsets.

---

## 8. 🔍 2–3 Must-Know Problems

---

### Problem 1: Subsets (LC #78) — Easy/Medium

**Intuition:** Every node in the recursion tree is a valid answer. So you collect the path at every call, not just at leaf nodes.

**Approach:** Use Template 1. Start index prevents revisiting earlier elements.

```go
func subsets(nums []int) [][]int {
    res := [][]int{}
    var backtrack func(start int, path []int)
    backtrack = func(start int, path []int) {
        // Collect current path as a valid subset
        temp := make([]int, len(path))
        copy(temp, path)
        res = append(res, temp)

        for i := start; i < len(nums); i++ {
            path = append(path, nums[i])
            backtrack(i+1, path)
            path = path[:len(path)-1]
        }
    }
    backtrack(0, []int{})
    return res
}
```

---

### Problem 2: Combination Sum (LC #39) — Medium

**Intuition:** Same as subsets, BUT you can reuse the same number. So instead of passing `i+1`, you pass `i` again. You stop when the remaining target hits 0 (success) or goes negative (prune).

**Approach:** Template 1, modified — pass `i` instead of `i+1` to allow reuse. Stop recursion early if `target < 0`.

```go
func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}
    var backtrack func(start, target int, path []int)
    backtrack = func(start, target int, path []int) {
        if target == 0 {
            // ✅ Found a valid combination
            temp := make([]int, len(path))
            copy(temp, path)
            res = append(res, temp)
            return
        }
        if target < 0 {
            return // ❌ Pruned: overshot the target
        }
        for i := start; i < len(candidates); i++ {
            path = append(path, candidates[i])
            backtrack(i, target-candidates[i], path) // i, not i+1 (allow reuse)
            path = path[:len(path)-1]
        }
    }
    backtrack(0, target, []int{})
    return res
}
```

---

### Problem 3: Generate Parentheses (LC #22) — Medium

**Intuition:** At any point, you can add `(` if you still have opens left, or `)` if there are more opens than closes so far. This constraint-based pruning is what makes it efficient.

**Approach:** Template 3 style (constraint-based). Track `open` and `close` counts. Recurse when valid.

```go
func generateParenthesis(n int) []string {
    res := []string{}
    var backtrack func(path string, open, close int)
    backtrack = func(path string, open, close int) {
        if len(path) == 2*n {
            res = append(res, path)
            return
        }
        // Can add '(' if opens used < n
        if open < n {
            backtrack(path+"(", open+1, close)
        }
        // Can add ')' only if closes < opens (ensures balance)
        if close < open {
            backtrack(path+")", open, close+1)
        }
    }
    backtrack("", 0, 0)
    return res
}
```

> **Note:** Strings are immutable in Go, so no explicit "unchoose" step is needed here — passing `path+"("` creates a new string, so the original is unchanged.

---

## 9. ⚠️ Common Mistakes / Gotchas

**1. Forgetting to deep copy the path before appending to result.**
Writing `res = append(res, path)` appends a reference to the slice, not a copy. As `path` mutates, your results get corrupted. Always `copy()` first.

**2. Forgetting the "unchoose" step.**
If you append to `path` but don't remove it after recursing, every future call inherits your choice. This is the most common backtracking bug.

**3. Using `i+1` when reuse is allowed (or `i` when it's not).**
In Combination Sum, reuse is allowed → pass `i`. In Subsets/Combinations, no reuse → pass `i+1`. Confusing these produces wrong answers silently.

**4. Not pruning early.**
If `target < 0`, return immediately. If `row == n`, collect and return. Pruning early is what makes backtracking feasible — without it, TLE is likely.

**5. Incorrect duplicate handling in arrays with repeated elements.**
For problems like Subsets II or Permutations II, you must **sort first** and skip `nums[i] == nums[i-1]` when `i > start` (or `!used[i-1]`). Missing this produces duplicate results.

**6. Modifying the loop variable inside recursion.**
The `for i := start; i < len(nums); i++` loop uses `i` as the boundary for the next call. Mutating `nums` or `i` inside the loop causes subtle bugs.

**7. Confusing `used[]` with `visited[]` in grid problems.**
In Word Search, you mark a cell as visited during DFS, but must unmark it when backtracking. Forgetting the unmark makes cells permanently unavailable.

---

## 10. ⏱️ Time & Space Complexity

| Operation | Time Complexity | Why |
|---|---|---|
| Generate all subsets | O(N × 2ᴺ) | 2ᴺ subsets, each takes O(N) to copy |
| Generate all permutations | O(N × N!) | N! permutations, each takes O(N) to copy |
| Combination Sum | O(N^(T/M)) | T = target, M = min candidate; tree depth is T/M |
| Generate Parentheses | O(4ᴺ / √N) | Catalan number — known closed form |
| N-Queens | O(N!) | At each row, valid column choices decrease |
| **Space (call stack)** | O(N) or O(depth) | Recursion depth = path length |

---

## 11. 🏋️ Practice Problems

### 🟢 Easy

| # | Problem | LeetCode | Hint |
|---|---|---|---|
| 1 | Subsets | #78 | Collect path at every node, use start index |
| 2 | Letter Case Permutation | #784 | At each char, branch into lowercase and uppercase |
| 3 | Binary Watch | #401 | Treat hour/minute bits as a combination problem |

---

### 🟡 Medium

| # | Problem | LeetCode | Hint |
|---|---|---|---|
| 1 | Subsets II (with duplicates) | #90 | Sort + skip `nums[i]==nums[i-1]` when `i > start` |
| 2 | Combination Sum | #39 | Pass `i` (not `i+1`) to allow element reuse |
| 3 | Combination Sum II | #40 | Sort + skip same element at same depth level |
| 4 | Permutations | #46 | Use `used[]` array, pick all unused elements each time |
| 5 | Permutations II (with duplicates) | #47 | Sort + skip if `used[i-1]==false` and `nums[i]==nums[i-1]` |
| 6 | Generate Parentheses | #22 | Add `(` if open < n, add `)` only if close < open |
| 7 | Palindrome Partitioning | #131 | At each index, try every valid palindrome prefix, recurse on rest |
| 8 | Combinations | #77 | Fix start index, stop when path length equals K |
| 9 | Target Sum | #494 | At each element, try adding or subtracting it |
| 10 | Restore IP Addresses | #93 | Try segments of length 1, 2, 3; validate each segment |

---

### 🔴 Hard

| # | Problem | LeetCode | Hint |
|---|---|---|---|
| 1 | N-Queens | #51 | Track cols, diagonals; place one queen per row |
| 2 | Sudoku Solver | #37 | For each empty cell, try 1–9; skip if row/col/box conflict |
| 3 | Expression Add Operators | #282 | Track running value AND last operand (for multiplication undo) |

---

## 12. ⚡ Quick Revision (60 sec Cheat Sheet)

**Key Idea:** Try a choice → recurse deeper → undo the choice. Repeat for all options.

**The Golden Rule:**
```go
path = append(path, choice)     // CHOOSE
backtrack(...)                  // EXPLORE
path = path[:len(path)-1]       // UNCHOOSE ← never forget this
```

**When to Use:**
- "Find ALL combinations / permutations / subsets"
- "Generate all valid strings / arrangements"
- "Place items on a board with constraints"
- N is small (≤ 20)

**Key Decisions:**
```
Need reuse?        → pass i (not i+1)
Need no reuse?     → pass i+1
Order matters?     → use used[] (permutations)
Order doesn't?     → use start index (combinations)
Has duplicates?    → sort first, skip same adjacent at same depth
```

**Always deep copy before adding to result:**
```go
temp := make([]int, len(path))
copy(temp, path)
res = append(res, temp)
```

---

*Made with ❤️ for interview prep. Master the template, understand the tree, and 80% of backtracking problems become straightforward.*
