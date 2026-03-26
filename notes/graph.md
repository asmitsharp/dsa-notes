# 🕸️ Graphs — Complete Interview DSA Notes

> One file. Everything you need. Nothing you don't.

---

## 1. What Is a Graph?

A **graph** is a set of **nodes (vertices)** connected by **edges**.

Unlike trees, graphs can have:

- **Cycles** (loops)
- **Multiple components** (disconnected parts)
- **Directed or undirected** edges
- **Weighted** edges

```
Undirected:          Directed:
  1 — 2               1 → 2
  |   |               ↑   ↓
  3 — 4               4 ← 3
```

---

## 2. Graph Terminology

| Term                   | Meaning                                                    |
| ---------------------- | ---------------------------------------------------------- |
| **Vertex / Node**      | A point in the graph                                       |
| **Edge**               | A connection between two nodes                             |
| **Directed (Digraph)** | Edges have direction (one-way)                             |
| **Undirected**         | Edges go both ways                                         |
| **Weighted**           | Edges have a cost/weight                                   |
| **Degree**             | Number of edges connected to a node                        |
| **In-degree**          | Edges pointing INTO a node (directed)                      |
| **Out-degree**         | Edges pointing OUT of a node (directed)                    |
| **Path**               | Sequence of nodes connected by edges                       |
| **Cycle**              | A path that starts and ends at the same node               |
| **Connected Graph**    | Every node reachable from every other node                 |
| **Strongly Connected** | In a directed graph, every node reachable from every other |
| **DAG**                | Directed Acyclic Graph — directed + no cycles              |
| **Component**          | A maximal connected subgraph                               |
| **Bipartite**          | Nodes can be 2-colored with no same-color edge             |

---

## 3. Graph Representations

### Adjacency List (Most Common — Use This)

```go
// Undirected graph
graph := make([][]int, n) // n nodes
graph[u] = append(graph[u], v)
graph[v] = append(graph[v], u)

// Directed graph (only one direction)
graph[u] = append(graph[u], v)

// Weighted graph
type Edge struct { to, weight int }
graph := make([][]Edge, n)
graph[u] = append(graph[u], Edge{v, w})
```

**When to use:** Almost always. O(V + E) space.

### Adjacency Matrix

```go
matrix := make([][]int, n)
for i := range matrix {
    matrix[i] = make([]int, n)
}
matrix[u][v] = 1 // or weight
matrix[v][u] = 1 // if undirected
```

**When to use:** Dense graphs, or when you need O(1) edge lookup.

| Representation   | Space  | Add Edge | Check Edge | Iterate Neighbors |
| ---------------- | ------ | -------- | ---------- | ----------------- |
| Adjacency List   | O(V+E) | O(1)     | O(degree)  | O(degree)         |
| Adjacency Matrix | O(V²)  | O(1)     | O(1)       | O(V)              |

> 💡 For most interview problems → use **adjacency list**.

### Edge List

```go
edges := [][2]int{{0,1}, {1,2}, {2,3}}
```

**When to use:** Kruskal's algorithm (MST), or when given input as edge pairs.

---

## 4. Building Graph from Input

Most interview problems give you edges as input. Know how to build the graph:

```go
// Input: n nodes, edges [][]int where edges[i] = [u, v]
func buildGraph(n int, edges [][]int) [][]int {
    graph := make([][]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u) // remove for directed
    }
    return graph
}
```

---

## 5. DFS on Graphs

```go
// Recursive DFS
func dfs(graph [][]int, node int, visited []bool) {
    visited[node] = true
    // process node
    for _, neighbor := range graph[node] {
        if !visited[neighbor] {
            dfs(graph, neighbor, visited)
        }
    }
}

// Iterative DFS (stack)
func dfsIterative(graph [][]int, start int) {
    visited := make([]bool, len(graph))
    stack := []int{start}

    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        if visited[node] {
            continue
        }
        visited[node] = true

        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                stack = append(stack, neighbor)
            }
        }
    }
}
```

---

## 6. BFS on Graphs

```go
func bfs(graph [][]int, start int) []int {
    visited := make([]bool, len(graph))
    queue := []int{start}
    visited[start] = true
    result := []int{}

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        result = append(result, node)

        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
    return result
}
```

> **DFS vs BFS:**
>
> - BFS → shortest path (unweighted), level info
> - DFS → cycle detection, topological sort, connected components, backtracking

---

## 7. Cycle Detection

### Undirected Graph (DFS with parent tracking)

```go
func hasCycleUndirected(graph [][]int) bool {
    n := len(graph)
    visited := make([]bool, n)

    var dfs func(node, parent int) bool
    dfs = func(node, parent int) bool {
        visited[node] = true
        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                if dfs(neighbor, node) {
                    return true
                }
            } else if neighbor != parent {
                return true // back edge to non-parent = cycle
            }
        }
        return false
    }

    for i := 0; i < n; i++ {
        if !visited[i] {
            if dfs(i, -1) {
                return true
            }
        }
    }
    return false
}
```

### Directed Graph (DFS with 3 colors / recursion stack)

```go
// 0 = unvisited, 1 = in current path, 2 = fully processed
func hasCycleDirected(graph [][]int) bool {
    n := len(graph)
    color := make([]int, n) // 0=white, 1=gray, 2=black

    var dfs func(node int) bool
    dfs = func(node int) bool {
        color[node] = 1 // mark as in-progress
        for _, neighbor := range graph[node] {
            if color[neighbor] == 1 {
                return true // back edge = cycle
            }
            if color[neighbor] == 0 {
                if dfs(neighbor) {
                    return true
                }
            }
        }
        color[node] = 2 // fully processed
        return false
    }

    for i := 0; i < n; i++ {
        if color[i] == 0 {
            if dfs(i) {
                return true
            }
        }
    }
    return false
}
```

> ⚠️ For undirected graphs, track `parent` to avoid false cycle (A→B→A is not a cycle).
> For directed graphs, track `inStack` (gray = in current DFS path).

---

## 8. Connected Components

```go
func countComponents(n int, edges [][]int) int {
    graph := buildGraph(n, edges)
    visited := make([]bool, n)
    count := 0

    var dfs func(node int)
    dfs = func(node int) {
        visited[node] = true
        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                dfs(neighbor)
            }
        }
    }

    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(i) // explore entire component
            count++
        }
    }
    return count
}
```

---

## 9. Topological Sort (DAG Only)

> **Use when:** ordering tasks with dependencies (course schedule, build order).
> Only works on **Directed Acyclic Graphs**.

### Method 1: Kahn's Algorithm (BFS-based) ← Preferred for cycle detection too

```go
func topoSort(n int, prerequisites [][]int) []int {
    graph := make([][]int, n)
    inDegree := make([]int, n)

    for _, pre := range prerequisites {
        course, pre := pre[0], pre[1]
        graph[pre] = append(graph[pre], course)
        inDegree[course]++
    }

    queue := []int{}
    for i := 0; i < n; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i) // start with nodes that have no dependencies
        }
    }

    order := []int{}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        order = append(order, node)

        for _, neighbor := range graph[node] {
            inDegree[neighbor]--
            if inDegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }

    if len(order) != n {
        return nil // cycle detected — impossible to sort
    }
    return order
}
```

### Method 2: DFS-based Topological Sort

```go
func topoSortDFS(n int, graph [][]int) []int {
    visited := make([]bool, n)
    stack := []int{}

    var dfs func(node int)
    dfs = func(node int) {
        visited[node] = true
        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                dfs(neighbor)
            }
        }
        stack = append(stack, node) // push AFTER processing all neighbors
    }

    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(i)
        }
    }

    // reverse stack = topological order
    for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
        stack[i], stack[j] = stack[j], stack[i]
    }
    return stack
}
```

---

## 10. Shortest Path Algorithms

### BFS (Unweighted Graphs) — O(V + E)

```go
func shortestPath(graph [][]int, start, end int) int {
    visited := make([]bool, len(graph))
    queue := []int{start}
    visited[start] = true
    dist := 0

    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]
            if node == end {
                return dist
            }
            for _, neighbor := range graph[node] {
                if !visited[neighbor] {
                    visited[neighbor] = true
                    queue = append(queue, neighbor)
                }
            }
        }
        dist++
    }
    return -1
}
```

### Dijkstra's Algorithm (Weighted, Non-negative) — O((V + E) log V)

```go
import "container/heap"

type Item struct { node, dist int }
type MinHeap []Item

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() interface{}   { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

func dijkstra(n int, graph [][]Item, start int) []int {
    dist := make([]int, n)
    for i := range dist { dist[i] = 1<<31 - 1 }
    dist[start] = 0

    pq := &MinHeap{{start, 0}}
    heap.Init(pq)

    for pq.Len() > 0 {
        curr := heap.Pop(pq).(Item)
        node, d := curr.node, curr.dist

        if d > dist[node] { continue } // outdated entry

        for _, neighbor := range graph[node] {
            newDist := dist[node] + neighbor.dist
            if newDist < dist[neighbor.node] {
                dist[neighbor.node] = newDist
                heap.Push(pq, Item{neighbor.node, newDist})
            }
        }
    }
    return dist
}
```

### Bellman-Ford (Negative Weights, Detect Negative Cycles) — O(V × E)

```go
func bellmanFord(n int, edges [][3]int, start int) []int {
    dist := make([]int, n)
    for i := range dist { dist[i] = 1<<31 - 1 }
    dist[start] = 0

    // Relax all edges V-1 times
    for i := 0; i < n-1; i++ {
        for _, e := range edges {
            u, v, w := e[0], e[1], e[2]
            if dist[u] != 1<<31-1 && dist[u]+w < dist[v] {
                dist[v] = dist[u] + w
            }
        }
    }

    // Check for negative cycles (Vth relaxation still improves)
    for _, e := range edges {
        u, v, w := e[0], e[1], e[2]
        if dist[u]+w < dist[v] {
            return nil // negative cycle exists
        }
    }
    return dist
}
```

### Floyd-Warshall (All Pairs Shortest Path) — O(V³)

```go
func floydWarshall(n int, dist [][]int) [][]int {
    // Initialize: dist[i][j] = weight if edge exists, else infinity, dist[i][i] = 0
    for k := 0; k < n; k++ {
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                if dist[i][k]+dist[k][j] < dist[i][j] {
                    dist[i][j] = dist[i][k] + dist[k][j]
                }
            }
        }
    }
    return dist
}
```

### Algorithm Cheat Sheet

| Scenario                   | Algorithm         | Complexity     |
| -------------------------- | ----------------- | -------------- |
| Unweighted graph, min hops | BFS               | O(V + E)       |
| Weighted, non-negative     | Dijkstra          | O((V+E) log V) |
| Weighted, negative weights | Bellman-Ford      | O(V × E)       |
| All pairs shortest path    | Floyd-Warshall    | O(V³)          |
| DAG shortest path          | Topo Sort + Relax | O(V + E)       |

---

## 11. Union-Find (Disjoint Set Union — DSU)

> Use for: grouping, connected components, detecting cycles in undirected graphs, Kruskal's MST.

```go
type DSU struct {
    parent []int
    rank   []int
}

func NewDSU(n int) *DSU {
    parent := make([]int, n)
    rank := make([]int, n)
    for i := range parent { parent[i] = i }
    return &DSU{parent, rank}
}

func (d *DSU) Find(x int) int {
    if d.parent[x] != x {
        d.parent[x] = d.Find(d.parent[x]) // path compression
    }
    return d.parent[x]
}

func (d *DSU) Union(x, y int) bool {
    px, py := d.Find(x), d.Find(y)
    if px == py { return false } // already connected — cycle!
    // union by rank
    if d.rank[px] < d.rank[py] {
        d.parent[px] = py
    } else if d.rank[px] > d.rank[py] {
        d.parent[py] = px
    } else {
        d.parent[py] = px
        d.rank[px]++
    }
    return true
}
```

**Complexity:** Find and Union are nearly O(1) — amortized O(α(N)) with path compression + rank.

**When to use DSU vs DFS for components:**

- DSU when you get edges one at a time (dynamic connectivity)
- DFS when full graph is given upfront

---

## 12. Minimum Spanning Tree (MST)

> MST connects all nodes with minimum total edge weight. Only for undirected weighted graphs.

### Kruskal's Algorithm (Sort edges + DSU)

```go
func kruskal(n int, edges [][3]int) int { // edges: [u, v, weight]
    // sort edges by weight
    sort.Slice(edges, func(i, j int) bool {
        return edges[i][2] < edges[j][2]
    })

    dsu := NewDSU(n)
    totalCost := 0

    for _, e := range edges {
        u, v, w := e[0], e[1], e[2]
        if dsu.Union(u, v) { // if not already connected
            totalCost += w
        }
    }
    return totalCost
}
```

### Prim's Algorithm (Greedy + MinHeap, starts from a node)

```go
func prim(n int, graph [][]Item) int { // Item = {neighbor, weight}
    visited := make([]bool, n)
    pq := &MinHeap{{0, 0}} // start from node 0
    heap.Init(pq)
    total := 0

    for pq.Len() > 0 {
        curr := heap.Pop(pq).(Item)
        node, cost := curr.node, curr.dist
        if visited[node] { continue }
        visited[node] = true
        total += cost
        for _, neighbor := range graph[node] {
            if !visited[neighbor.node] {
                heap.Push(pq, neighbor)
            }
        }
    }
    return total
}
```

| Algorithm | Best For                       | Complexity |
| --------- | ------------------------------ | ---------- |
| Kruskal's | Sparse graphs, easy to code    | O(E log E) |
| Prim's    | Dense graphs, adjacency matrix | O(E log V) |

---

## 13. Bipartite Graph Check

> A graph is bipartite if nodes can be 2-colored such that no two adjacent nodes share a color.
> Use BFS or DFS with coloring.

```go
func isBipartite(graph [][]int) bool {
    n := len(graph)
    color := make([]int, n)
    for i := range color { color[i] = -1 }

    for start := 0; start < n; start++ {
        if color[start] != -1 { continue }

        queue := []int{start}
        color[start] = 0

        for len(queue) > 0 {
            node := queue[0]
            queue = queue[1:]
            for _, neighbor := range graph[node] {
                if color[neighbor] == -1 {
                    color[neighbor] = 1 - color[node] // alternate color
                    queue = append(queue, neighbor)
                } else if color[neighbor] == color[node] {
                    return false // same color = not bipartite
                }
            }
        }
    }
    return true
}
```

---

## 14. Graph on 2D Grid

Most grid problems are implicit graphs. Cells are nodes; adjacent cells (4 or 8 directions) are edges.

```go
var dirs = [][2]int{{0,1},{0,-1},{1,0},{-1,0}} // 4-directional
// For 8-directional add: {1,1},{1,-1},{-1,1},{-1,-1}

func validCell(r, c, rows, cols int) bool {
    return r >= 0 && r < rows && c >= 0 && c < cols
}

// BFS on grid template
func bfsGrid(grid [][]byte, sr, sc int) {
    rows, cols := len(grid), len(grid[0])
    visited := make([][]bool, rows)
    for i := range visited { visited[i] = make([]bool, cols) }

    type P struct{ r, c int }
    queue := []P{{sr, sc}}
    visited[sr][sc] = true

    for len(queue) > 0 {
        curr := queue[0]; queue = queue[1:]
        for _, d := range dirs {
            nr, nc := curr.r+d[0], curr.c+d[1]
            if validCell(nr, nc, rows, cols) && !visited[nr][nc] && grid[nr][nc] == '1' {
                visited[nr][nc] = true
                queue = append(queue, P{nr, nc})
            }
        }
    }
}
```

---

## 15. Strongly Connected Components (SCC) — Kosaraju's

> SCC: A maximal set of nodes where every node is reachable from every other.

```go
// Step 1: DFS on original graph, push to stack by finish time
// Step 2: Transpose the graph
// Step 3: DFS on transposed graph in reverse finish order
// Each DFS tree in step 3 = one SCC

func kosaraju(n int, graph [][]int) int {
    visited := make([]bool, n)
    stack := []int{}

    var dfs1 func(node int)
    dfs1 = func(node int) {
        visited[node] = true
        for _, neighbor := range graph[node] {
            if !visited[neighbor] { dfs1(neighbor) }
        }
        stack = append(stack, node)
    }

    for i := 0; i < n; i++ {
        if !visited[i] { dfs1(i) }
    }

    // Build transposed graph
    transposed := make([][]int, n)
    for u := 0; u < n; u++ {
        for _, v := range graph[u] {
            transposed[v] = append(transposed[v], u)
        }
    }

    visited = make([]bool, n)
    sccCount := 0

    var dfs2 func(node int)
    dfs2 = func(node int) {
        visited[node] = true
        for _, neighbor := range transposed[node] {
            if !visited[neighbor] { dfs2(neighbor) }
        }
    }

    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if !visited[node] {
            dfs2(node)
            sccCount++
        }
    }
    return sccCount
}
```

---

## 16. Key Patterns & When to Use What

| Problem Signal                           | Use This                                             |
| ---------------------------------------- | ---------------------------------------------------- |
| "Shortest path, unweighted"              | BFS                                                  |
| "Shortest path, weighted (non-negative)" | Dijkstra                                             |
| "Shortest path, negative weights"        | Bellman-Ford                                         |
| "All pairs shortest path"                | Floyd-Warshall                                       |
| "Is there a path / reachable?"           | DFS or BFS                                           |
| "Number of connected components"         | DFS loop + count, or DSU                             |
| "Is there a cycle?"                      | DFS (directed: color; undirected: parent)            |
| "Task ordering / dependencies"           | Topological Sort (Kahn's BFS)                        |
| "Can all tasks be completed?"            | Topo Sort — check if order length == n               |
| "Minimum cost to connect all nodes"      | MST — Kruskal's or Prim's                            |
| "Group nodes dynamically"                | DSU / Union-Find                                     |
| "2-color graph check"                    | Bipartite BFS/DFS                                    |
| "Islands / regions in grid"              | DFS/BFS from each unvisited cell                     |
| "Minimum steps in grid"                  | BFS on grid                                          |
| "State-space minimum moves"              | BFS on implicit graph (encode state as string/tuple) |

---

## 17. Common Mistakes / Gotchas

| #   | Mistake                                                | Fix                                                                |
| --- | ------------------------------------------------------ | ------------------------------------------------------------------ |
| 1   | **Not handling disconnected graph**                    | Always loop over ALL nodes for DFS/BFS, not just node 0            |
| 2   | **Marking visited after dequeue (not before enqueue)** | Mark when adding to queue — prevents duplicate entries             |
| 3   | **Using DFS for shortest path in unweighted graph**    | DFS gives A path, not the SHORTEST — always use BFS                |
| 4   | **Cycle detection: undirected vs directed confusion**  | Undirected needs parent tracking; directed needs in-stack tracking |
| 5   | **Topo sort on a graph with cycles**                   | Kahn's output length < n → cycle exists, impossible                |
| 6   | **Dijkstra with negative weights**                     | Use Bellman-Ford instead — Dijkstra fails with negatives           |
| 7   | **Building graph with 0-indexed vs 1-indexed nodes**   | Allocate `n+1` nodes if input is 1-indexed                         |
| 8   | **Forgetting bounds check on grid**                    | Always `r >= 0 && r < rows && c >= 0 && c < cols`                  |
| 9   | **Mutating graph while traversing**                    | Work on a copy or use visited array properly                       |
| 10  | **Not handling self-loops or multiple edges**          | Add `if u == v { continue }` where needed                          |

---

## 18. Complexity Reference

| Algorithm                   | Time           | Space |
| --------------------------- | -------------- | ----- |
| DFS / BFS                   | O(V + E)       | O(V)  |
| Topological Sort (Kahn's)   | O(V + E)       | O(V)  |
| Dijkstra (MinHeap)          | O((V+E) log V) | O(V)  |
| Bellman-Ford                | O(V × E)       | O(V)  |
| Floyd-Warshall              | O(V³)          | O(V²) |
| Kruskal's MST               | O(E log E)     | O(V)  |
| Prim's MST                  | O(E log V)     | O(V)  |
| Union-Find (w/ compression) | O(α(N)) per op | O(N)  |
| Kosaraju SCC                | O(V + E)       | O(V)  |

---

## 19. Practice Problems

### 🟢 Easy

| #   | Problem                      | LC # | Hint                              |
| --- | ---------------------------- | ---- | --------------------------------- |
| 1   | Find if Path Exists in Graph | 1971 | Simple BFS/DFS or DSU             |
| 2   | Find the Town Judge          | 997  | In-degree == n-1, out-degree == 0 |
| 3   | Flood Fill                   | 733  | BFS/DFS from starting cell        |
| 4   | Number of Islands            | 200  | DFS from each unvisited '1'       |
| 5   | Ransom Note (graph concept)  | 383  | Frequency count                   |

### 🟡 Medium

| #   | Problem                        | LC # | Hint                              |
| --- | ------------------------------ | ---- | --------------------------------- |
| 1   | Course Schedule                | 207  | Topo sort (Kahn's) — detect cycle |
| 2   | Course Schedule II             | 210  | Topo sort — return order          |
| 3   | Number of Connected Components | 323  | DFS loop count or DSU             |
| 4   | Clone Graph                    | 133  | BFS + HashMap old→new             |
| 5   | Pacific Atlantic Water Flow    | 417  | Reverse BFS from both oceans      |
| 6   | Rotting Oranges                | 994  | Multi-source BFS                  |
| 7   | 01 Matrix                      | 542  | Multi-source BFS from all 0s      |
| 8   | Graph Valid Tree               | 261  | n-1 edges + no cycle (DSU)        |
| 9   | Is Graph Bipartite?            | 785  | BFS 2-coloring                    |
| 10  | Network Delay Time             | 743  | Dijkstra, return max dist         |
| 11  | Redundant Connection           | 684  | DSU — find edge creating cycle    |
| 12  | Word Ladder                    | 127  | BFS on implicit graph             |
| 13  | Walls and Gates                | 286  | Multi-source BFS from gates       |
| 14  | Max Area of Island             | 695  | DFS/BFS — track area size         |
| 15  | Surrounded Regions             | 130  | BFS from border 'O's, mark safe   |

### 🔴 Hard

| #   | Problem                           | LC # | Hint                                     |
| --- | --------------------------------- | ---- | ---------------------------------------- |
| 1   | Word Ladder II                    | 126  | BFS levels + DFS path reconstruction     |
| 2   | Alien Dictionary                  | 269  | Topo sort on character DAG               |
| 3   | Swim in Rising Water              | 778  | Dijkstra / binary search + BFS           |
| 4   | Minimum Cost to Reach Destination | 787  | Bellman-Ford with k stops                |
| 5   | Critical Connections in Network   | 1192 | Tarjan's bridge finding (DFS + disc/low) |

---

## 20. Tarjan's Bridge Finding (Bonus — for Hard Problems)

```go
// A bridge is an edge whose removal disconnects the graph
var disc, low []int
var timer int
var bridges [][]int

func findBridges(n int, graph [][]int) [][]int {
    disc = make([]int, n)
    low = make([]int, n)
    for i := range disc { disc[i] = -1 }
    timer = 0
    bridges = [][]int{}

    var dfs func(node, parent int)
    dfs = func(node, parent int) {
        disc[node] = timer
        low[node] = timer
        timer++

        for _, neighbor := range graph[node] {
            if disc[neighbor] == -1 {
                dfs(neighbor, node)
                low[node] = min(low[node], low[neighbor])
                if low[neighbor] > disc[node] {
                    bridges = append(bridges, []int{node, neighbor}) // bridge!
                }
            } else if neighbor != parent {
                low[node] = min(low[node], disc[neighbor])
            }
        }
    }

    for i := 0; i < n; i++ {
        if disc[i] == -1 { dfs(i, -1) }
    }
    return bridges
}
```

---

## 21. Quick Revision (60 sec Cheat Sheet)

```
GRAPH REPRESENTATIONS:
  Adjacency List → use by default (O(V+E) space)
  Adjacency Matrix → dense graphs or O(1) edge lookup

DFS vs BFS:
  DFS  → cycle detection, topo sort, components, paths
  BFS  → shortest path (unweighted), level-based problems

CYCLE DETECTION:
  Undirected → DFS + track parent (back edge to non-parent = cycle)
  Directed   → DFS + 3-color (gray = in stack, back to gray = cycle)

TOPOLOGICAL SORT:
  Only for DAGs. Use Kahn's (BFS + in-degree).
  If output length < n → cycle exists.

SHORTEST PATH:
  Unweighted → BFS
  Weighted, positive → Dijkstra (MinHeap)
  Weighted, negative → Bellman-Ford
  All pairs → Floyd-Warshall

UNION-FIND:
  Group nodes, detect cycles, count components.
  Find (path compression) + Union (by rank) ≈ O(1).

MULTI-SOURCE BFS:
  Push ALL sources at once → spread simultaneously.
  e.g., Rotting Oranges, 01 Matrix, Walls & Gates.

KEY FORMULA:
  Always loop ALL nodes (graph may be disconnected!)
  Mark visited AT enqueue time, not dequeue time.

MST:
  Kruskal = sort edges + DSU (better for sparse)
  Prim    = MinHeap from one node (better for dense)
```

---

_Graph Notes | Interview-Ready | All-in-One_
