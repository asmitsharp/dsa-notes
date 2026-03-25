# 🔵 Breadth First Search (BFS) — Interview-Focused DSA Notes

---

## 1. Pattern Overview

### What is BFS?

BFS is a graph/tree traversal technique that explores nodes **level by level** — visiting all neighbors at the current depth before going deeper.

Think of it like throwing a stone in water: ripples spread outward in circles — that's BFS.

### Why Is It Used? (Intuition)

- BFS **guarantees shortest path** in an **unweighted graph**.
- It explores the **closest nodes first**, then moves outward.
- It's ideal when you need the **minimum steps/hops** to reach a target.

### Real-World Analogy

> Imagine you're looking for someone in a building floor by floor.
> You check **Floor 1 completely**, then **Floor 2 completely**, and so on.
> You don't go to Floor 3 until every room on Floor 2 is checked.
> That's BFS — layer by layer.

### Time Complexity Intuition

- You visit **every node once** and process **every edge once**.
- So: **O(V + E)** where V = nodes, E = edges.
- For a grid of size M×N: **O(M × N)** — you touch each cell at most once.

---

## 2. Core Idea (MOST IMPORTANT)

### Key Insight #1: Use a Queue (FIFO)

BFS works because a **queue** naturally processes nodes in the order they were discovered — ensuring level-by-level traversal.

```
Queue: [start] → process → add neighbors → process next level
```

### Key Insight #2: Track Visited Nodes

Without a `visited` set, BFS revisits nodes and loops forever.

> ✅ Always mark a node as visited **when you add it to the queue**, NOT when you process it.

### What to Remember

```
- Queue → BFS
- Mark visited BEFORE pushing to queue
- Process layer by layer (use queue size for level separation)
- Distance/level = number of BFS rounds completed
```

---

## 3. When to Use BFS

| Signal in Question                | What It Means                                |
| --------------------------------- | -------------------------------------------- |
| "Minimum steps / hops / moves"    | BFS gives shortest path in unweighted graphs |
| "Shortest path between two nodes" | Classic BFS use case                         |
| "Level order traversal"           | BFS by definition                            |
| "Nearest / closest X"             | BFS expands outward, finds nearest first     |
| "Minimum number of operations"    | Each operation = 1 level in BFS              |
| "Find if reachable"               | BFS can detect reachability                  |
| "Spread / infection / fire"       | Multi-source BFS                             |
| "Word ladder / transformation"    | BFS on implicit graph                        |
| Grid + "minimum path"             | BFS on 2D grid                               |

---

## 4. Types / Variants of BFS

### Type 1: Standard BFS (Single Source)

Start from one node, explore all reachable nodes level by level.

```
Use when: Finding shortest path from ONE source to target.
```

### Type 2: Multi-Source BFS

Start from **multiple nodes simultaneously** (push all sources into the queue at once).

```
Use when: "Nearest X from any of these points" — e.g., nearest 0 in a grid, rotting oranges.
```

### Type 3: Level-by-Level BFS

Track levels explicitly using queue size or a sentinel.

```
Use when: "Return level order", "minimum depth", "minimum moves count".
```

### Type 4: BFS on Implicit Graph

The graph is not given explicitly — you **build neighbors on the fly** (e.g., word ladder, state machines).

```
Use when: Problem has states that transform, and you need min transformations.
```

---

## 5. Data Structures Used

| Data Structure                  | Why Used                          | Key Property          | Time Complexity      |
| ------------------------------- | --------------------------------- | --------------------- | -------------------- |
| **Queue**                       | FIFO ensures level-by-level order | First in, first out   | Enqueue/Dequeue O(1) |
| **Visited Set / Boolean Array** | Prevents revisiting nodes         | O(1) lookup           | O(1) check/insert    |
| **HashMap**                     | Store distance, parent, or state  | Fast key-value access | O(1) average         |
| **2D Array (Grid)**             | Represent the graph implicitly    | Direct index access   | O(1) access          |

### When to Choose Which:

- Use a **boolean array** for visited when nodes are integers (faster, less overhead).
- Use a **HashSet** when nodes are strings or complex states (word ladder, etc.).
- Use a **HashMap** when you need to store distance or parent info per node.

---

## 6. Core Templates (Go)

### Template 1: Standard BFS on Graph

```go
func bfs(graph [][]int, start int) []int {
    n := len(graph)
    visited := make([]bool, n)
    queue := []int{start}
    visited[start] = true
    result := []int{}

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:] // dequeue

        result = append(result, node)

        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                visited[neighbor] = true  // mark BEFORE pushing
                queue = append(queue, neighbor)
            }
        }
    }
    return result
}
```

### Template 2: BFS with Level Tracking (Shortest Path / Min Steps)

```go
func bfsLevels(graph [][]int, start, target int) int {
    visited := make([]bool, len(graph))
    queue := []int{start}
    visited[start] = true
    level := 0

    for len(queue) > 0 {
        size := len(queue) // process current level completely
        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]

            if node == target {
                return level // found at this level = minimum steps
            }

            for _, neighbor := range graph[node] {
                if !visited[neighbor] {
                    visited[neighbor] = true
                    queue = append(queue, neighbor)
                }
            }
        }
        level++ // move to next level
    }
    return -1 // not reachable
}
```

### Template 3: BFS on 2D Grid

```go
var directions = [][2]int{{0,1},{0,-1},{1,0},{-1,0}} // right, left, down, up

func bfsGrid(grid [][]int, startR, startC int) int {
    rows, cols := len(grid), len(grid[0])
    visited := make([][]bool, rows)
    for i := range visited {
        visited[i] = make([]bool, cols)
    }

    type Point struct{ r, c int }
    queue := []Point{{startR, startC}}
    visited[startR][startC] = true
    steps := 0

    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            curr := queue[0]
            queue = queue[1:]

            for _, d := range directions {
                nr, nc := curr.r+d[0], curr.c+d[1]
                // bounds check + not visited + valid cell
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols &&
                    !visited[nr][nc] && grid[nr][nc] != 0 {
                    visited[nr][nc] = true
                    queue = append(queue, Point{nr, nc})
                }
            }
        }
        steps++
    }
    return steps
}
```

---

## 7. Step-by-Step Dry Run

### Problem: Find shortest path from node 0 to node 4

```
Graph (adjacency list):
0 → [1, 2]
1 → [3]
2 → [4]
3 → []
4 → []
```

**Step-by-step:**

```
Level 0: Queue = [0]          visited = {0}
  Process 0 → neighbors: 1, 2
  Push 1, 2

Level 1: Queue = [1, 2]       visited = {0, 1, 2}
  Process 1 → neighbors: 3
  Push 3
  Process 2 → neighbors: 4
  Push 4 ← found target! return level = 2

Answer: 2 steps (0 → 2 → 4)
```

**Key observation:** BFS found the shortest path because it explores **all paths of length 1 before length 2, all length 2 before length 3**, etc.

---

## 8. Must-Know Problems (Deep Explanation)

---

### Problem 1: Binary Tree Level Order Traversal (LC #102)

**Intuition:**
A tree is just a special graph. BFS naturally gives us level-by-level traversal because all nodes of depth D are processed before depth D+1.

**Approach:**

- Push root into queue.
- For each level, capture the current queue size — that's how many nodes are on this level.
- Process exactly that many nodes, collect values, push their children.

```go
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    result := [][]int{}
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        size := len(queue) // number of nodes at current level
        level := []int{}

        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]
            level = append(level, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        result = append(result, level)
    }
    return result
}
```

---

### Problem 2: Rotting Oranges (LC #994)

**Intuition:**
All rotten oranges rot their neighbors simultaneously — that's multi-source BFS. Start with ALL rotten oranges in the queue at once. Each BFS level = 1 minute.

**Approach:**

1. Find all rotten oranges → push ALL into queue at once.
2. Count fresh oranges.
3. BFS — each level is 1 minute. For each rotten orange, rot its fresh neighbors.
4. After BFS, if any fresh orange remains → return -1. Else return minutes.

```go
func orangesRotting(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    queue := [][]int{}
    fresh := 0

    // Push all rotten oranges at start (multi-source BFS)
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 2 {
                queue = append(queue, []int{r, c})
            } else if grid[r][c] == 1 {
                fresh++
            }
        }
    }

    if fresh == 0 {
        return 0
    }

    dirs := [][2]int{{0,1},{0,-1},{1,0},{-1,0}}
    minutes := 0

    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            curr := queue[0]
            queue = queue[1:]
            for _, d := range dirs {
                nr, nc := curr[0]+d[0], curr[1]+d[1]
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
                    grid[nr][nc] = 2 // rot it
                    fresh--
                    queue = append(queue, []int{nr, nc})
                }
            }
        }
        minutes++
    }

    if fresh > 0 {
        return -1
    }
    return minutes
}
```

---

### Problem 3: Word Ladder (LC #127)

**Intuition:**
Treat each word as a node. Two words are connected if they differ by exactly 1 character. We need the **minimum transformations** = shortest path in this implicit graph. BFS on strings!

**Approach:**

1. Put `beginWord` in queue.
2. At each step, try changing every character (a–z) of the current word.
3. If new word is in `wordList` and not visited → add to queue.
4. Return level when `endWord` is reached.

```go
func ladderLength(beginWord string, endWord string, wordList []string) int {
    wordSet := map[string]bool{}
    for _, w := range wordList {
        wordSet[w] = true
    }
    if !wordSet[endWord] {
        return 0
    }

    queue := []string{beginWord}
    visited := map[string]bool{beginWord: true}
    level := 1

    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            word := queue[0]
            queue = queue[1:]

            // try every 1-character transformation
            wordBytes := []byte(word)
            for j := 0; j < len(wordBytes); j++ {
                orig := wordBytes[j]
                for c := byte('a'); c <= byte('z'); c++ {
                    wordBytes[j] = c
                    newWord := string(wordBytes)
                    if newWord == endWord {
                        return level + 1
                    }
                    if wordSet[newWord] && !visited[newWord] {
                        visited[newWord] = true
                        queue = append(queue, newWord)
                    }
                }
                wordBytes[j] = orig // restore
            }
        }
        level++
    }
    return 0
}
```

---

## 9. Common Mistakes / Gotchas

| #   | Mistake                                                | What Goes Wrong                                               | Fix                                                                                                 |
| --- | ------------------------------------------------------ | ------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| 1   | **Marking visited AFTER dequeue (not before enqueue)** | Same node added to queue multiple times → TLE or wrong answer | Mark `visited[node] = true` BEFORE or AT THE TIME of pushing to queue                               |
| 2   | **Not tracking levels when counting steps**            | Returns wrong minimum distance                                | Use `size := len(queue)` pattern to separate levels                                                 |
| 3   | **Forgetting bounds check on grid**                    | Index out of range panic                                      | Always check `nr >= 0 && nr < rows && nc >= 0 && nc < cols`                                         |
| 4   | **Using DFS when problem asks for shortest path**      | DFS finds A path, not the SHORTEST path                       | BFS guarantees shortest path in unweighted graphs                                                   |
| 5   | **Not seeding multi-source BFS correctly**             | Only processes one source → wrong answer                      | Add ALL sources to queue before starting BFS loop                                                   |
| 6   | **Modifying `queue` while iterating**                  | Logic errors, skipping elements                               | Use index-based loop with `size` captured at start of each level                                    |
| 7   | **Off-by-one in level counting**                       | Returns level+1 or level-1                                    | Decide: increment level AFTER processing a full level, or initialize to 1 if start counts as step 1 |

---

## 10. Time & Space Complexity

| Scenario                                 | Time          | Space    | Why                                                          |
| ---------------------------------------- | ------------- | -------- | ------------------------------------------------------------ |
| BFS on Graph (V nodes, E edges)          | O(V + E)      | O(V)     | Each node and edge visited once; queue holds at most V nodes |
| BFS on Grid (M × N)                      | O(M × N)      | O(M × N) | Each cell visited once; visited array is M×N                 |
| BFS with level tracking                  | O(V + E)      | O(V)     | Same as standard BFS, just grouped by level                  |
| Multi-source BFS                         | O(V + E)      | O(V)     | All sources pushed initially; still visits each node once    |
| Word Ladder (L = word length, N = words) | O(N × L × 26) | O(N)     | For each word, try L positions × 26 characters               |

---

## 11. Practice Problems

### 🟢 Easy

| #   | Problem                           | LeetCode # | Hint                                                                  |
| --- | --------------------------------- | ---------- | --------------------------------------------------------------------- |
| 1   | Binary Tree Level Order Traversal | 102        | Use queue; capture `size` at each level                               |
| 2   | Maximum Depth of Binary Tree      | 104        | BFS levels = depth; count levels                                      |
| 3   | Same Tree                         | 100        | BFS both trees simultaneously, compare node values                    |
| 4   | Symmetric Tree                    | 101        | BFS with pairs: push left.left + right.right, left.right + right.left |
| 5   | Flood Fill                        | 733        | BFS from starting cell, change color of connected cells               |

---

### 🟡 Medium

| #   | Problem                     | LeetCode # | Hint                                                    |
| --- | --------------------------- | ---------- | ------------------------------------------------------- |
| 1   | Rotting Oranges             | 994        | Multi-source BFS; all rotten oranges start together     |
| 2   | Number of Islands           | 200        | BFS from each unvisited '1', mark entire island visited |
| 3   | 01 Matrix                   | 542        | Multi-source BFS from all 0s simultaneously             |
| 4   | Course Schedule             | 207        | BFS on DAG (Topological sort via Kahn's algorithm)      |
| 5   | Clone Graph                 | 133        | BFS + HashMap to map original node → cloned node        |
| 6   | Walls and Gates             | 286        | Multi-source BFS from all gates (0s), fill distances    |
| 7   | Pacific Atlantic Water Flow | 417        | Reverse BFS from both oceans; find intersection         |
| 8   | Minimum Knight Moves        | 1197       | BFS on implicit graph; each move = 1 level              |
| 9   | Snakes and Ladders          | 909        | BFS on board; level = number of dice rolls              |
| 10  | Word Ladder                 | 127        | BFS on implicit graph of word transformations           |

---

### 🔴 Hard

| #   | Problem                                      | LeetCode # | Hint                                                                |
| --- | -------------------------------------------- | ---------- | ------------------------------------------------------------------- |
| 1   | Word Ladder II                               | 126        | BFS to find level + DFS/backtrack to reconstruct all shortest paths |
| 2   | Sliding Puzzle                               | 773        | BFS on board state (encode as string); find min moves               |
| 3   | Jump Game IV                                 | 1345       | BFS on array indices; group same values for O(n) jumps              |
| 4   | Cut Off Trees for Golf Event                 | 675        | BFS repeated from each tree in sorted order                         |
| 5   | Minimum Moves to Reach Target with Rotations | 1210       | BFS on state = (row, col, orientation of snake)                     |

---

## 12. Quick Revision (60 sec Cheat Sheet)

```
KEY IDEA:
  BFS = Level-by-level traversal using a Queue (FIFO)
  → Guarantees shortest path in unweighted graphs
  → Mark visited BEFORE pushing to queue

WHEN TO USE:
  ✅ Shortest path / minimum steps
  ✅ Level order traversal
  ✅ Nearest node / spread problems
  ✅ Word transformations (implicit graph)
  ❌ NOT for weighted shortest path (use Dijkstra)
  ❌ NOT for detecting cycles efficiently (use DFS)

CORE TEMPLATE:
  queue := []Node{start}
  visited[start] = true

  for len(queue) > 0 {
      size := len(queue)          // current level size
      for i := 0; i < size; i++ {
          node := queue[0]; queue = queue[1:]
          // process node
          for each neighbor {
              if !visited[neighbor] {
                  visited[neighbor] = true
                  queue = append(queue, neighbor)
              }
          }
      }
      level++                     // done with this level
  }

MULTI-SOURCE BFS:
  Push ALL sources into queue BEFORE the loop starts.

GRID BFS:
  dirs := [][2]int{{0,1},{0,-1},{1,0},{-1,0}}
  Check: bounds + not visited + valid cell

COMPLEXITY:
  Time: O(V + E) or O(M×N) for grids
  Space: O(V) for queue + visited
```

---

_Notes by FAANG DSA Mentor | BFS Pattern | Interview-Ready_
