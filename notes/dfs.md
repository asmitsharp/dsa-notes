# 🌊 Depth First Search (DFS) — Interview Notes

> **Mentor's Note:** DFS is one of the most important patterns in coding interviews. Once you understand the _core idea_, you can solve trees, graphs, backtracking, and path problems with the same mental model. Let's build that intuition.

---

## 1. 🧭 Pattern Overview

### What is DFS?

DFS is a traversal technique where you **go as deep as possible** into a structure (tree/graph) before backtracking. You explore one full path before trying the next.

### Why is it used?

Use DFS when you need to:

- **Explore all possibilities** (paths, combinations, subsets)
- **Find something deep** in a nested structure
- **Make a decision at each step** and backtrack if it doesn't work

### Real-World Analogy

> Imagine you're solving a **maze**. You pick one direction and walk straight until you hit a dead end. Then you backtrack to the last fork and try the next direction. You repeat this until you find the exit.

That's exactly what DFS does — commit to one path, go deep, backtrack, try the next.

### Time Complexity Intuition

| Structure    | Complexity                                    |
| ------------ | --------------------------------------------- |
| Binary Tree  | O(N) — visit every node once                  |
| Graph        | O(V + E) — every vertex + every edge          |
| Backtracking | O(N!) or O(2^N) — exploring all possibilities |

> **Intuition:** You visit each node **exactly once**, so it's always linear in the number of nodes (unless backtracking blows it up).

---

## 2. 💡 Core Idea (MOST IMPORTANT)

### Key Insight #1: Recursion = DFS

Every recursive function call is DFS under the hood. The **call stack** is your "memory" of where you've been.

```
explore(node)
  → explore(left child)
      → explore(left-left child) ← goes DEEP first
```

### Key Insight #2: The DFS Pattern = Process + Recurse + Backtrack

```
dfs(node):
  1. BASE CASE — stop condition
  2. PROCESS — do something with current node
  3. RECURSE — go deeper (left, right, neighbors)
  4. BACKTRACK — undo if needed (for backtracking problems)
```

### ⭐ What to Remember

- DFS uses the **call stack** (recursion) or an **explicit stack** (iterative)
- **Pre-order** = process BEFORE recursing (root → left → right)
- **Post-order** = process AFTER recursing (left → right → root)
- **In-order** = left → process → right (gives sorted order in BST)

---

## 3. 🎯 When to Use This Pattern

| Signal in Question               | What it Means           |
| -------------------------------- | ----------------------- |
| "Find all paths"                 | DFS + backtracking      |
| "Does a path exist?"             | DFS with early return   |
| "Count all combinations/subsets" | DFS + backtracking      |
| "Traverse a tree"                | DFS (pre/in/post order) |
| "Maximum/minimum depth"          | DFS post-order          |
| "Validate a tree property"       | DFS with return values  |
| "Find all permutations"          | DFS + backtracking      |
| "Path sum equals target"         | DFS with running sum    |
| "Nested structure"               | DFS naturally fits      |

---

## 4. 🔀 Types / Variants of DFS

### Type 1: Tree DFS (Most Common)

Traverse trees recursively. Three flavors:

- **Pre-order:** Root → Left → Right (copy tree, serialize)
- **In-order:** Left → Root → Right (BST sorted order)
- **Post-order:** Left → Right → Root (delete tree, height calculation)

### Type 2: DFS with Return Value

Each DFS call returns something useful (height, bool, count) that the parent uses.

```
// Example: return height from each subtree
height = 1 + max(dfs(left), dfs(right))
```

### Type 3: Backtracking DFS

Make a choice → go deep → **undo the choice** → try next option.
Used for: subsets, permutations, combinations, word search.

```
choose → dfs() → unchoose  ← this is the backtracking pattern
```

### Type 4: Iterative DFS (using explicit Stack)

When recursion depth is too large or recursion is not allowed.
Use a stack, push right before left (so left is processed first).

---

## 5. 🗃️ Data Structures Used

| Data Structure             | Why Used                    | Key Property         | When to Choose                    |
| -------------------------- | --------------------------- | -------------------- | --------------------------------- |
| **Call Stack** (recursion) | Natural DFS — OS manages it | LIFO, implicit       | Default choice for trees          |
| **Explicit Stack**         | Iterative DFS               | LIFO, manual control | Deep trees (avoid stack overflow) |
| **Visited Set (HashSet)**  | Avoid revisiting nodes      | O(1) lookup          | Graph DFS (cycles)                |
| **Path List/Array**        | Track current path          | Dynamic add/remove   | Backtracking problems             |
| **HashMap**                | Store computed results      | O(1) lookup          | Memoized DFS                      |

> **Key Rule:** In **tree DFS**, you usually don't need a visited set (trees have no cycles). In **graph DFS**, always use a visited set.

---

## 6. 📋 Core Templates

### Template 1: Basic Tree DFS (Recursive)

```go
// Generic tree node
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func dfs(node *TreeNode) {
    // BASE CASE: stop when node is nil
    if node == nil {
        return
    }

    // PRE-ORDER: process here for root → left → right
    fmt.Println(node.Val)

    dfs(node.Left)   // go deep left
    dfs(node.Right)  // go deep right

    // POST-ORDER: process here for left → right → root
}
```

---

### Template 2: DFS with Return Value

```go
// Returns something useful (int, bool, etc.)
func dfs(node *TreeNode) int {
    // BASE CASE
    if node == nil {
        return 0 // return neutral value
    }

    // RECURSE: get values from children
    left := dfs(node.Left)
    right := dfs(node.Right)

    // PROCESS: combine results from children
    return 1 + max(left, right) // example: height
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

---

### Template 3: Backtracking DFS

```go
func backtrack(start int, current []int, result *[][]int, nums []int) {
    // COLLECT: save current state as a result
    temp := make([]int, len(current))
    copy(temp, current)
    *result = append(*result, temp)

    for i := start; i < len(nums); i++ {
        // CHOOSE
        current = append(current, nums[i])

        // RECURSE
        backtrack(i+1, current, result, nums)

        // UNCHOOSE (backtrack)
        current = current[:len(current)-1]
    }
}
```

---

### Template 4: Iterative DFS (Explicit Stack)

```go
func dfsIterative(root *TreeNode) {
    if root == nil {
        return
    }

    stack := []*TreeNode{root}

    for len(stack) > 0 {
        // POP from top (LIFO)
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        // PROCESS
        fmt.Println(node.Val)

        // Push RIGHT first, so LEFT is processed first
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
    }
}
```

---

## 7. 🔍 Step-by-Step Example (Dry Run)

### Problem: Find max depth of a binary tree

```
Tree:
        3
       / \
      9  20
         / \
        15   7
```

**Goal:** Return 3 (deepest path: 3 → 20 → 15 or 3 → 20 → 7)

**Trace:**

```
dfs(3)
  ├── dfs(9)
  │     ├── dfs(nil) → return 0
  │     └── dfs(nil) → return 0
  │     → return 1 + max(0, 0) = 1
  │
  └── dfs(20)
        ├── dfs(15)
        │     ├── dfs(nil) → return 0
        │     └── dfs(nil) → return 0
        │     → return 1 + max(0, 0) = 1
        │
        └── dfs(7)
              ├── dfs(nil) → return 0
              └── dfs(nil) → return 0
              → return 1 + max(0, 0) = 1
        → return 1 + max(1, 1) = 2

→ return 1 + max(1, 2) = 3  ✅
```

**Key Insight:** Post-order DFS — children report their height, parent adds 1.

---

## 8. 🧩 Must-Know Problems (Deep Explanation)

---

### Problem 1: Maximum Depth of Binary Tree (LC #104) — Easy

**Intuition:** Ask each subtree: "how tall are you?" Take the taller one, add 1 for current node.

**Approach:** Post-order DFS. Each node returns its height to its parent.

```go
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)
    return 1 + max(left, right)
}

func max(a, b int) int {
    if a > b { return a }
    return b
}
```

---

### Problem 2: Path Sum (LC #112) — Easy

**Intuition:** Walk down every path. At each step, subtract node's value from target. At a leaf, check if target became 0.

**Approach:** Pre-order DFS. Pass remaining sum down the recursion.

```go
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    // Reached a leaf
    if root.Left == nil && root.Right == nil {
        return root.Val == targetSum
    }

    remaining := targetSum - root.Val
    // Check either left OR right path
    return hasPathSum(root.Left, remaining) || hasPathSum(root.Right, remaining)
}
```

**Key Line:** `root.Left == nil && root.Right == nil` → leaf node check.

---

### Problem 3: Subsets (LC #78) — Medium

**Intuition:** At each element, you have 2 choices: include it or skip it. DFS explores both choices for all elements → 2^N subsets.

**Approach:** Backtracking DFS. At each step, add current subset, then try adding more elements starting from `i+1`.

```go
func subsets(nums []int) [][]int {
    result := [][]int{}
    backtrack(0, []int{}, &result, nums)
    return result
}

func backtrack(start int, current []int, result *[][]int, nums []int) {
    // Every state is a valid subset — collect it
    temp := make([]int, len(current))
    copy(temp, current)
    *result = append(*result, temp)

    for i := start; i < len(nums); i++ {
        current = append(current, nums[i])      // CHOOSE
        backtrack(i+1, current, result, nums)   // RECURSE
        current = current[:len(current)-1]      // UNCHOOSE
    }
}
```

**Key Line:** `copy(temp, current)` — always deep copy before saving, or you'll save a reference that changes later!

---

## 9. ⚠️ Common Mistakes / Gotchas

1. **Forgetting the base case (nil check)**
   - Always check `if node == nil { return }` first. Forgetting this causes nil pointer panics.

2. **Not deep copying in backtracking**
   - `*result = append(*result, current)` saves a reference! Use `copy()` to save a snapshot.

3. **Using visited set in tree DFS**
   - Trees have no cycles. You don't need a visited set. Only use it for graphs.

4. **Wrong leaf node check**
   - Leaf = `node.Left == nil && node.Right == nil`. Using just `node.Left == nil` is wrong.

5. **Pre-order vs Post-order confusion**
   - If you need **children's results first**, use post-order (process after recursing).
   - If you need to **pass info down**, use pre-order (process before recursing).

6. **Modifying slice without copy in Go**
   - In Go, slices share underlying array. Always copy before appending to result in backtracking.

7. **Stack overflow on deep trees**
   - For trees with 10,000+ depth, recursion may exceed stack limit. Use iterative DFS with explicit stack.

---

## 10. ⏱️ Time & Space Complexity

| Operation                   | Time   | Space    | Why                                             |
| --------------------------- | ------ | -------- | ----------------------------------------------- |
| Tree DFS (traversal)        | O(N)   | O(H)     | Visit each node once; H = height for call stack |
| Tree DFS (balanced)         | O(N)   | O(log N) | Height = log N for balanced tree                |
| Tree DFS (skewed)           | O(N)   | O(N)     | Height = N (like a linked list)                 |
| Backtracking (subsets)      | O(2^N) | O(N)     | 2 choices per element                           |
| Backtracking (permutations) | O(N!)  | O(N)     | N choices, then N-1, then N-2...                |
| Iterative DFS               | O(N)   | O(N)     | Stack can hold all nodes in worst case          |

> **Space Key:** Recursive DFS space = **height of tree** (call stack depth). Iterative DFS space = **max nodes in stack** at once.

---

## 11. 🏋️ Practice Problems

### 🟢 Easy

| #   | Problem                      | LC # | Hint                                                                 |
| --- | ---------------------------- | ---- | -------------------------------------------------------------------- |
| 1   | Maximum Depth of Binary Tree | 104  | Post-order: return 1 + max(left, right)                              |
| 2   | Symmetric Tree               | 101  | DFS on both sides simultaneously: compare left.left with right.right |
| 3   | Path Sum                     | 112  | Subtract node value at each step, check leaf when target == 0        |
| 4   | Invert Binary Tree           | 226  | Pre-order: swap children, then recurse                               |
| 5   | Same Tree                    | 100  | DFS both trees in sync: compare values at each node                  |

---

### 🟡 Medium

| #   | Problem                             | LC # | Hint                                                           |
| --- | ----------------------------------- | ---- | -------------------------------------------------------------- |
| 1   | Subsets                             | 78   | Backtrack: at each index, include or skip                      |
| 2   | Path Sum II                         | 113  | DFS + path tracking: add to path, remove on backtrack          |
| 3   | Binary Tree Right Side View         | 199  | DFS pre-order, track depth, save last node at each depth       |
| 4   | Lowest Common Ancestor of BST       | 235  | Use BST property: go left if both < root, right if both > root |
| 5   | Validate Binary Search Tree         | 98   | DFS with min/max bounds passed down                            |
| 6   | Permutations                        | 46   | Backtrack: pick unused element, recurse, unpick                |
| 7   | Combination Sum                     | 39   | Backtrack: include same element again (i, not i+1)             |
| 8   | Letter Combinations of Phone Number | 17   | Backtrack: at each digit, try all mapped letters               |
| 9   | Flatten Binary Tree to Linked List  | 114  | Post-order: flatten left, flatten right, relink                |
| 10  | Count Good Nodes in Binary Tree     | 1448 | DFS: track max value seen so far on path from root             |

---

### 🔴 Hard

| #   | Problem                               | LC # | Hint                                                                 |
| --- | ------------------------------------- | ---- | -------------------------------------------------------------------- |
| 1   | Binary Tree Maximum Path Sum          | 124  | Post-order: at each node, decide to extend path or start new         |
| 2   | Serialize and Deserialize Binary Tree | 297  | Pre-order DFS to serialize; use null markers for empty nodes         |
| 3   | N-Queens                              | 51   | Backtrack: place queen row by row, check column + diagonal conflicts |

---

## 12. ⚡ Quick Revision (60 sec Cheat Sheet)

### Key Idea

> Go **deep first**, then backtrack. Use recursion (call stack) naturally.

### The 4-Step DFS Template

```
1. BASE CASE  → stop condition (nil, empty, target met)
2. PROCESS    → do something with current node
3. RECURSE    → call dfs on children/neighbors
4. BACKTRACK  → undo choice (only for backtracking problems)
```

### Core Template (Go)

```go
func dfs(node *TreeNode) int {
    if node == nil { return 0 }         // base case
    left := dfs(node.Left)              // recurse
    right := dfs(node.Right)            // recurse
    return 1 + max(left, right)         // process (post-order)
}
```

### When to Use

| Clue                       | Pattern                 |
| -------------------------- | ----------------------- |
| Tree traversal             | DFS (pre/in/post order) |
| Path sum / path exists     | DFS with running value  |
| All subsets / permutations | Backtracking DFS        |
| Validate tree property     | DFS with return value   |
| Nested structure           | DFS naturally           |

### ⭐ Golden Rules

- **Trees:** No visited set needed
- **Backtracking:** Always `copy()` before saving to result
- **Post-order:** Use when children's results are needed first
- **Pre-order:** Use when passing info from parent to child
- **Leaf check:** `node.Left == nil && node.Right == nil`
