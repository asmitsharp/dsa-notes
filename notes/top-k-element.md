# 🏔️ Heaps — The Data Structure That Always Knows the Best

> **Mentor Note:** Before diving into Top K Elements, you MUST understand Heaps. Most people skip this and then struggle forever. This guide will make heaps feel obvious.

---

## 1. What Problem Does a Heap Solve?

Imagine you're running a hospital emergency room. Patients keep arriving, and you always want to treat the **most critical patient first** — not the one who arrived first, not the last one, but always the most critical.

Now, you could sort all patients every time a new one arrives. But that's wasteful — O(n log n) every time.

What if there was a structure that could **always give you the most critical patient in O(1)**, and inserting a new patient only costs **O(log n)**?

**That's a Heap.**

A heap is a special tree that always keeps the most important element at the top — ready to be grabbed instantly.

---

## 2. The Core Idea — What IS a Heap?

A heap is a **complete binary tree** with one special rule:

> **Every parent is always greater than (or less than) its children.**

That's it. That one rule is the entire secret.

**Max-Heap:** Parent is always LARGER than children → Top is the MAXIMUM element.

**Min-Heap:** Parent is always SMALLER than children → Top is the MINIMUM element.

```
Max-Heap Example:
        10
       /  \
      7    8
     / \  /
    3   4 5

Min-Heap Example:
        1
       / \
      3    2
     / \  /
    7   4 5
```

**Key insight:** You do NOT need the whole tree to be sorted. You only guarantee the parent-child relationship. This is why heaps are fast — maintaining full sort order is expensive, but maintaining the parent > child rule is cheap.

---

## 3. How a Heap Is Actually Stored

Here's the beautiful trick: **we never actually build a tree with pointers**. We store the heap in a plain array!

```
Array: [10, 7, 8, 3, 4, 5]
Index:   0  1  2  3  4  5

For any element at index i:
  - Left child  → index 2*i + 1
  - Right child → index 2*i + 2
  - Parent      → index (i-1) / 2
```

This means:

- `10` is at index 0 → its children are at index 1 (`7`) and index 2 (`8`) ✓
- `7` is at index 1 → its children are at index 3 (`3`) and index 4 (`4`) ✓

No pointers. No memory overhead. Just array math. This is why heaps are incredibly memory efficient.

---

## 4. The Two Core Operations

### 4.1 Insert (Push) — O(log n)

1. Add the new element at the **end of the array** (bottom of the tree).
2. **"Bubble up"** (also called sift-up): Compare with parent. If it violates the heap rule, swap. Repeat until the rule is satisfied.

```
Insert 9 into Max-Heap [10, 7, 8, 3, 4, 5]:

Step 1: Add at end → [10, 7, 8, 3, 4, 5, 9]
Step 2: 9 > parent(8) → swap → [10, 7, 9, 3, 4, 5, 8]
Step 3: 9 < parent(10) → stop!

Final: [10, 7, 9, 3, 4, 5, 8]  ✓
```

### 4.2 Extract Top (Pop) — O(log n)

1. The top element (index 0) is your answer — grab it.
2. Move the **last element to the top**.
3. **"Bubble down"** (also called sift-down): Compare with children. Swap with the largest/smallest child if it violates the heap rule. Repeat until rule is satisfied.

```
Extract max from [10, 7, 9, 3, 4, 5, 8]:

Step 1: Grab 10 (answer)
Step 2: Move last (8) to top → [8, 7, 9, 3, 4, 5]
Step 3: 8 < child(9) → swap → [9, 7, 8, 3, 4, 5]
Step 4: 9 > both children → stop!

Extracted: 10 ✓
```

### 4.3 Peek Top — O(1)

Just look at index 0. Done. This is the magic — the best element is always sitting right there.

---

## 5. Complexity Table

| Operation             | Time     | Why                                 |
| --------------------- | -------- | ----------------------------------- |
| Peek (get min/max)    | O(1)     | Always at index 0                   |
| Insert                | O(log n) | Bubble up at most tree height       |
| Extract min/max       | O(log n) | Bubble down at most tree height     |
| Build heap from array | O(n)     | Heapify algorithm (not O(n log n)!) |
| Search for element    | O(n)     | Heap isn't sorted — must scan all   |

> **Why O(n) to build a heap?** Intuitively, most elements are near the leaves and barely need to bubble down. The math works out to O(n). This is a well-known result.

---

## 6. Heaps in Go — Using `container/heap`

Go doesn't have a built-in heap like Python's `heapq`, but it provides the `container/heap` package which is clean once you understand the pattern.

### 6.1 Min-Heap in Go

```go
package main

import (
    "container/heap"
    "fmt"
)

// MinHeap is just a slice of ints
type MinHeap []int

// These 5 methods are REQUIRED by container/heap interface
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-Heap: smaller = higher priority
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]       // grab last element
    *h = old[:n-1]      // shrink slice
    return x
}

func main() {
    h := &MinHeap{5, 2, 8, 1, 9}
    heap.Init(h)            // Build heap from existing slice — O(n)

    heap.Push(h, 3)         // Insert 3 — O(log n)
    fmt.Println((*h)[0])    // Peek min — O(1) → prints 1

    min := heap.Pop(h)      // Extract min — O(log n)
    fmt.Println(min)        // prints 1
}
```

### 6.2 Max-Heap in Go

The ONLY change from MinHeap: **flip the Less function**.

```go
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // FLIPPED → Max-Heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}
```

### 6.3 Heap of Pairs (most useful in interviews!)

When you need to store (value, index) or (frequency, element) pairs:

```go
// Pair represents (priority, value)
type Pair struct {
    priority int
    value    int
}

type PairMinHeap []Pair

func (h PairMinHeap) Len() int           { return len(h) }
func (h PairMinHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
func (h PairMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PairMinHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *PairMinHeap) Pop() interface{} {
    old := *h
    x := old[len(old)-1]
    *h = old[:len(old)-1]
    return x
}
```

---

## 7. When to Use Min-Heap vs Max-Heap

This trips up a LOT of people. Here's the intuition:

| You want               | Use                    | Why                                                         |
| ---------------------- | ---------------------- | ----------------------------------------------------------- |
| K largest elements     | **Min-Heap of size K** | Keep K largest; min-heap lets you easily evict the smallest |
| K smallest elements    | **Max-Heap of size K** | Keep K smallest; max-heap lets you easily evict the largest |
| Always get the minimum | **Min-Heap**           | Minimum is always at the top                                |
| Always get the maximum | **Max-Heap**           | Maximum is always at the top                                |
| Merge K sorted lists   | **Min-Heap**           | Always pick the smallest current element                    |
| Running median         | **Both**               | Max-heap for lower half, min-heap for upper half            |

> **Memory trick for K Largest:** "I want large, so I use a min-heap." This sounds backward but makes sense: you maintain K large elements and throw away anything smaller than your current minimum. The min-heap top tells you what to evict.

---

## 8. Real Problems Heaps Solve

### Problem 1: Always get the next best element

- Finding the shortest path (Dijkstra's algorithm)
- Task scheduling by priority
- Finding K closest points

### Problem 2: Maintain a running rank

- "Is this new salary above the median?" — requires knowing the median at all times
- Sliding window maximum/minimum

### Problem 3: Merging sorted sequences

- Merge K sorted arrays — always pick the smallest unprocessed element across all arrays

### Problem 4: Top K anything

- Top K frequent words
- K closest numbers
- K largest elements

---

## 9. Common Mistakes with Heaps

**Mistake 1: Confusing which heap to use for Top K.** Remember: Top K Largest → Min-Heap of size K.

**Mistake 2: Forgetting to call `heap.Init()`** when building from an existing slice. Without it, the heap property is not established.

**Mistake 3: Reading the top without popping.** To peek without removing, access `(*h)[0]`. Calling `heap.Pop()` removes the element.

**Mistake 4: Modifying the slice directly** instead of using `heap.Push` and `heap.Pop`. If you modify the underlying slice, the heap property breaks.

**Mistake 5: Using heap.Pop return value without type assertion.** `heap.Pop` returns `interface{}`, so always cast: `val := heap.Pop(h).(int)`.

---

## 10. Quick Mental Model

Think of a heap as a **VIP queue**:

- **Min-Heap** = queue where the person with the smallest number goes first (like a waiting room sorted by ticket number)
- **Max-Heap** = queue where the most important person goes first (like a hospital by severity)
- **Insert** = new person joins and finds their correct relative position
- **Pop** = the front person leaves, and the queue reorganizes
- **Peek** = glance at who's next without removing them

The heap doesn't care about the full order of everyone in the queue — it only guarantees who's **next**.

---
## 11. Practice Problems

**🟢 Easy**

| # | Problem | Hint |
| --- | --- | --- |
| 703 | Kth Largest in a Stream | Min-heap of size K; each new element: push + maybe pop |
| 1337 | K Weakest Rows in Matrix | Heap on `(soldiers_count, row_index)`; keep K smallest |
| 1046 | Last Stone Weight | Max-heap; simulate: pop two, push difference if any |
| 2231 | Largest Number After Digit Swaps by Parity | Min-heap on digit value; straightforward heap exercise |

**🟡 Medium**

| # | Problem | Hint |
| --- | --- | --- |
| 215 | Kth Largest Element in Array | Min-heap size K; root = answer |
| 347 | Top K Frequent Elements | Freq map first, then min-heap on frequency |
| 973 | K Closest Points to Origin | Max-heap on `x² + y²`; evict farthest |
| 658 | Find K Closest Elements | Min-heap on `|x - target|`; collect K elements |
| 1985 | Find the Kth Largest Integer in Array | Strings as numbers; min-heap with string comparison |
| 373 | Find K Pairs with Smallest Sums | Min-heap on sum; classic K-way merge variant |
| 767 | Reorganize String | Max-heap on char frequency; greedily pick most frequent |
| 1167 | Minimum Cost to Connect Sticks | Min-heap; always merge two smallest (classic greedy) |

**🔴 Hard**

| # | Problem | Hint |
| --- | --- | --- |
| 295 | Find Median from Data Stream | Two heaps: max-heap (lower half) + min-heap (upper half) |
| 358 | Rearrange String K Distance Apart | Max-heap + cooldown queue of size K |
| 1439 | Find the Kth Smallest Sum of a Matrix | Min-heap on `(sum, i, j)`; expand neighbors |

_Now that you understand heaps, the Top K Elements pattern will feel completely natural._
