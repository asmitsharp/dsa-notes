# 🧩 DSA Pattern: Matrix Manipulation

> **Mentor's Note:** A matrix is just a 2D array. Once you stop seeing it as "grid math" and start seeing it as a **map with coordinates**, everything clicks. Almost every matrix problem is about: _navigating_, _transforming_, or _searching_ that map.

---

## 1. Pattern Overview

### What is this pattern?

Matrix Manipulation covers problems where you work on a **2D grid (rows × columns)**. Operations include traversal, rotation, spiral access, layer-by-layer processing, and in-place transformation.

### Why is it used? (Intuition)

Grids appear everywhere — game boards, images, maps, spreadsheets. The key insight is that **position = (row, col)** and every cell has up to 4 or 8 neighbors. Understanding how to move through a grid efficiently is the entire skill.

### Real-World Analogy

Think of a **chessboard**:

- Each square has coordinates (row, col)
- You can move in directions: up, down, left, right (or diagonally)
- Some operations work **layer by layer** (like peeling an onion)
- Some operations work **cell by cell** (like reading a book in a spiral)

### Time Complexity Intuition

Most matrix problems are **O(m × n)** — you visit every cell once.  
If you revisit (like BFS/DFS), it's still **O(m × n)** because each cell is visited at most once.  
Sorting anything in the matrix? That bumps you to **O(m × n × log(m × n))**.

---

## 2. Core Idea (MOST IMPORTANT)

### Key Insight #1 — Direction Vectors

Instead of writing 4 separate `if` blocks for up/down/left/right, use a **direction array**:

```
dirs = [[-1,0], [1,0], [0,-1], [0,1]]  // up, down, left, right
```

Loop over all directions → clean, scalable, no repeated code.

### Key Insight #2 — Layer-by-Layer (Onion Peeling)

Many problems (rotate, spiral, set zeroes) work on the **boundary first, then shrink inward**. Track: `top`, `bottom`, `left`, `right` boundaries and move them inward after each layer.

### Key Insight #3 — Index Relationships for Rotation

For **90° clockwise rotation** of an N×N matrix:

```
new[col][N-1-row] = old[row][col]
```

Or in-place: **Transpose first, then reverse each row.**

> 🔑 **Remember:** Transpose = swap `[i][j]` and `[j][i]`. Reverse rows = flip left-to-right.

---

## 3. When to Use This Pattern

| Signal in Question                      | What It Means                                   |
| --------------------------------------- | ----------------------------------------------- |
| "rotate the matrix"                     | Layer swap / transpose + reverse                |
| "spiral order"                          | Boundary shrinking with 4 pointers              |
| "set row and col to zero"               | In-place marking with sentinel or extra arrays  |
| "number of islands / connected regions" | DFS/BFS traversal on grid                       |
| "diagonal traversal"                    | Index math: `row + col = constant` per diagonal |
| "search in sorted matrix"               | Staircase search from top-right corner          |
| "game of life / neighbors"              | 8-directional traversal + state encoding        |
| "word search / path in grid"            | DFS + backtracking                              |
| "shortest path in grid"                 | BFS (unweighted)                                |

---

## 4. Types / Variants of This Pattern

### Type 1 — In-Place Transformation

Modify the matrix without extra space (or O(1) extra).  
**Examples:** Rotate Image, Set Matrix Zeroes, Transpose

### Type 2 — Traversal (BFS/DFS)

Visit connected cells to count, mark, or collect.  
**Examples:** Number of Islands, Flood Fill, Pacific Atlantic Water Flow

### Type 3 — Spiral / Layer Access

Access elements in a specific non-linear order using boundary pointers.  
**Examples:** Spiral Matrix I & II

### Type 4 — Search in Sorted Matrix

Exploit sorted structure to search efficiently.  
**Examples:** Search a 2D Matrix (binary search), Search a 2D Matrix II (staircase)

---

## 5. Data Structures Used

| Data Structure                   | Why Used                          | Key Property               | When to Choose                      |
| -------------------------------- | --------------------------------- | -------------------------- | ----------------------------------- |
| **2D Array (matrix itself)**     | Primary structure                 | O(1) access by index       | Always                              |
| **Visited array / boolean grid** | Track visited cells in DFS/BFS    | O(m×n) space               | When you can't mutate the grid      |
| **Queue (BFS)**                  | Level-by-level traversal          | FIFO, O(1) enqueue/dequeue | Shortest path, connected components |
| **Stack (DFS iterative)**        | Deep traversal                    | LIFO                       | Island counting, flood fill         |
| **Recursion stack (DFS)**        | Cleaner code for graph-like grids | Implicit stack             | Small grids, word search            |
| **HashMap**                      | Store diagonal index, counts      | O(1) lookup                | Diagonal grouping, unique tracking  |

---

## 6. Core Templates (VERY IMPORTANT)

### Template 1 — 4-Directional BFS Traversal

```go
func bfsGrid(grid [][]int, startR, startC int) {
    rows, cols := len(grid), len(grid[0])
    dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

    queue := [][2]int{{startR, startC}}
    grid[startR][startC] = 0 // mark visited (mutate in-place)

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        r, c := curr[0], curr[1]

        for _, d := range dirs {
            nr, nc := r+d[0], c+d[1]
            // bounds check + condition check
            if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
                grid[nr][nc] = 0 // mark visited
                queue = append(queue, [2]int{nr, nc})
            }
        }
    }
}
```

### Template 2 — Spiral Order Traversal

```go
func spiralOrder(matrix [][]int) []int {
    res := []int{}
    top, bottom := 0, len(matrix)-1
    left, right := 0, len(matrix[0])-1

    for top <= bottom && left <= right {
        // → move right along top row
        for c := left; c <= right; c++ {
            res = append(res, matrix[top][c])
        }
        top++

        // ↓ move down along right col
        for r := top; r <= bottom; r++ {
            res = append(res, matrix[r][right])
        }
        right--

        // ← move left along bottom row (if still valid)
        if top <= bottom {
            for c := right; c >= left; c-- {
                res = append(res, matrix[bottom][c])
            }
            bottom--
        }

        // ↑ move up along left col (if still valid)
        if left <= right {
            for r := bottom; r >= top; r-- {
                res = append(res, matrix[r][left])
            }
            left++
        }
    }
    return res
}
```

### Template 3 — In-Place 90° Clockwise Rotation

```go
func rotate(matrix [][]int) {
    n := len(matrix)

    // Step 1: Transpose (swap [i][j] and [j][i])
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }

    // Step 2: Reverse each row
    for i := 0; i < n; i++ {
        left, right := 0, n-1
        for left < right {
            matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
            left++
            right--
        }
    }
}
```

---

## 7. Step-by-Step Example (Dry Run)

**Problem:** Spiral Order of a 3×3 matrix

```
Matrix:
1  2  3
4  5  6
7  8  9
```

**Pointers:** `top=0, bottom=2, left=0, right=2`

| Step | Action                  | Elements Added | Pointer Update |
| ---- | ----------------------- | -------------- | -------------- |
| 1    | → Right along top row   | 1, 2, 3        | top = 1        |
| 2    | ↓ Down along right col  | 6, 9           | right = 1      |
| 3    | ← Left along bottom row | 8, 7           | bottom = 1     |
| 4    | ↑ Up along left col     | 4              | left = 1       |
| 5    | → Right along new top   | 5              | top = 2        |
| Done | top > bottom            | —              | Loop ends      |

**Result:** `[1, 2, 3, 6, 9, 8, 7, 4, 5]` ✅

---

## 8. Must-Know Problems (Deep Explanation)

---

### Problem 1 — Number of Islands (LC #200) 🟡 Medium

**Intuition:**  
Each `'1'` that hasn't been visited is the start of a new island. Flood fill from it (DFS/BFS), marking all connected land as visited. Count how many times you start a new flood fill.

**Approach:**

1. Loop every cell
2. If cell is `'1'`, increment count, then DFS to sink the entire island (mark as `'0'`)

```go
func numIslands(grid [][]byte) int {
    count := 0
    for r := 0; r < len(grid); r++ {
        for c := 0; c < len(grid[0]); c++ {
            if grid[r][c] == '1' {
                count++
                dfs(grid, r, c) // sink the island
            }
        }
    }
    return count
}

func dfs(grid [][]byte, r, c int) {
    if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != '1' {
        return
    }
    grid[r][c] = '0' // mark visited
    dfs(grid, r+1, c)
    dfs(grid, r-1, c)
    dfs(grid, r, c+1)
    dfs(grid, r, c-1)
}
```

---

### Problem 2 — Rotate Image (LC #48) 🟡 Medium

**Intuition:**  
Don't try to directly compute the new position — that's confusing. Instead: **transpose + reverse rows** achieves the same effect cleanly.

**Why does this work?**

- Transpose flips along the diagonal
- Reversing each row mirrors left-to-right
- Combined = 90° clockwise rotation

_(See Template 3 above for the code)_

---

### Problem 3 — Search a 2D Matrix II (LC #240) 🟡 Medium

**Intuition:**  
Start from the **top-right corner**. It has a special property:

- Go **left** → values decrease
- Go **down** → values increase

This lets you eliminate one row or column per step → **O(m + n)** search.

```go
func searchMatrix(matrix [][]int, target int) bool {
    r, c := 0, len(matrix[0])-1 // top-right corner

    for r < len(matrix) && c >= 0 {
        if matrix[r][c] == target {
            return true
        } else if matrix[r][c] > target {
            c-- // too big, go left
        } else {
            r++ // too small, go down
        }
    }
    return false
}
```

---

## 9. Common Mistakes / Gotchas

1. **Forgetting bounds check before accessing neighbors**  
   Always check `nr >= 0 && nr < rows && nc >= 0 && nc < cols` _before_ accessing `grid[nr][nc]`.

2. **Mutating input when you shouldn't**  
   In problems like "Game of Life", changing cells in-place affects neighbors. Use state encoding (e.g., `2` = was alive, now dead) or a copy.

3. **Off-by-one in spiral boundaries**  
   Always check `if top <= bottom` before the bottom-row traversal and `if left <= right` before the left-col traversal.

4. **Confusing row vs column indices**  
   `matrix[row][col]` — row goes with `i` (vertical), col goes with `j` (horizontal). Mix these up and everything breaks.

5. **Transpose loop starting from wrong index**  
   In transpose, the inner loop must start from `j = i+1`, not `j = 0`. Starting from 0 undoes your own swaps.

6. **Not marking visited cells → infinite DFS loop**  
   Always mark a cell as visited _before_ pushing to queue or recursing, not after.

7. **Using wrong direction count (4 vs 8)**  
   Most grid problems use 4 directions (no diagonal). Only use 8 directions when the problem explicitly involves diagonals (e.g., surrounded regions, queen attacks).

---

## 10. Time & Space Complexity

| Operation                      | Time        | Space  | Why                                     |
| ------------------------------ | ----------- | ------ | --------------------------------------- |
| Simple traversal               | O(m×n)      | O(1)   | Visit each cell once                    |
| BFS/DFS on grid                | O(m×n)      | O(m×n) | Queue/stack + visited array             |
| In-place rotation              | O(n²)       | O(1)   | Every cell touched once, no extra array |
| Spiral traversal               | O(m×n)      | O(1)   | Each cell added to result once          |
| Staircase search               | O(m+n)      | O(1)   | Eliminate 1 row or col per step         |
| Binary search in sorted matrix | O(log(m×n)) | O(1)   | Treat as 1D array                       |

---

## 11. Practice Problems

### 🟢 Easy

| #   | Problem                                   | LeetCode | Hint                                                        |
| --- | ----------------------------------------- | -------- | ----------------------------------------------------------- |
| 1   | Flood Fill                                | #733     | DFS from source, replace color                              |
| 2   | Transpose Matrix                          | #867     | Swap `[i][j]` with `[j][i]`                                 |
| 3   | Toeplitz Matrix                           | #766     | Every diagonal has same value; check `[i][j] == [i-1][j-1]` |
| 4   | Reshape the Matrix                        | #566     | Use `(r*cols + c)` to map 1D index back to 2D               |
| 5   | Count Negative Numbers in a Sorted Matrix | #1351    | Start from bottom-left, use staircase                       |

---

### 🟡 Medium

| #   | Problem                     | LeetCode | Hint                                                             |
| --- | --------------------------- | -------- | ---------------------------------------------------------------- |
| 1   | Number of Islands           | #200     | DFS/BFS flood fill, count trigger points                         |
| 2   | Rotate Image                | #48      | Transpose then reverse each row                                  |
| 3   | Spiral Matrix               | #54      | 4 boundary pointers, shrink inward                               |
| 4   | Spiral Matrix II            | #59      | Same as #54 but fill values 1..n²                                |
| 5   | Set Matrix Zeroes           | #73      | Use first row/col as markers to avoid extra space                |
| 6   | Search a 2D Matrix II       | #240     | Start top-right; go left if big, down if small                   |
| 7   | Pacific Atlantic Water Flow | #417     | BFS from both oceans, find intersection                          |
| 8   | Rotting Oranges             | #994     | Multi-source BFS from all rotten oranges                         |
| 9   | Word Search                 | #79      | DFS + backtracking, mark visited per path                        |
| 10  | Game of Life                | #289     | Encode states (2=dead←alive, 3=alive←dead) to transform in-place |

---

### 🔴 Hard

| #   | Problem                        | LeetCode | Hint                                                |
| --- | ------------------------------ | -------- | --------------------------------------------------- |
| 1   | Maximal Rectangle              | #85      | Build histogram per row, use stack (LC #84) per row |
| 2   | Shortest Path in Binary Matrix | #1091    | BFS with 8 directions, level = distance             |
| 3   | N-Queens                       | #51      | Backtracking on grid; track cols, diagonals         |
| 4   | Word Search II                 | #212     | DFS + Trie to prune search                          |
| 5   | Dungeon Game                   | #174     | DP from bottom-right to top-left                    |

---

## 12. Quick Revision (60 sec) ⚡

```
KEY IDEA:
  Matrix = 2D map. Navigate with direction arrays.

DIRECTION ARRAY (use this always):
  dirs := [][2]int{{-1,0},{1,0},{0,-1},{0,1}}

BOUNDS CHECK (use this always):
  nr >= 0 && nr < rows && nc >= 0 && nc < cols

ROTATION TRICK:
  90° clockwise = Transpose + Reverse each row

SPIRAL TRICK:
  top, bottom, left, right pointers → shrink after each side

SEARCH TRICK (sorted matrix):
  Start top-right → go left (too big) or down (too small)

BFS = shortest path / level-order
DFS = flood fill / connected components / backtracking

WHEN TO USE:
  Grid traversal    → BFS or DFS
  Transform matrix  → In-place with index math
  Spiral/layers     → Boundary pointers
  Search in grid    → Staircase or binary search
```

---

> 💡 **Final Mentor Tip:** Don't try to memorize formulas. Draw the matrix on paper, trace your pointer movements with your finger, and the pattern will become second nature. Every matrix problem is just a variation of "move through a grid cleverly."
