# 🔀 K-Way Merge — Interview-Focused DSA Notes

> **Pattern:** K-Way Merge  
> **Language:** Go  
> **Level:** Beginner-friendly → Interview-ready

---

## 1. Pattern Overview

### What is this pattern?

K-Way Merge is the technique of **merging K sorted sequences** (arrays, linked lists, or streams) into a single sorted output — efficiently, without brute-forcing.

Think of it as the generalized version of the classic "merge two sorted arrays" problem, extended to **K** inputs at once.

### Why is it used? (Intuition)

When you have K sorted lists, the naive approach is to dump everything into one array and sort it — that's `O(N log N)` where N is total elements. But you're throwing away the fact that **each list is already sorted**. The K-Way Merge pattern exploits that structure to do better.

The key insight: **at any point, the globally next-smallest element must be the smallest front element across all K lists.** You only ever need to compare K candidates at a time — not all N elements.

### Real-World Analogy

Imagine you're the judge at a spelling bee with K queues of students, each queue sorted by skill level (best → worst). To create the final ranking, you don't look at everyone. You just pick the best student from the _front_ of each queue, add them to your list, and advance that queue. That's exactly K-Way Merge.

Another analogy: **merge lanes on a highway.** K sorted streams of cars merging into one lane — you always pull from whichever lane has the earliest car.

### Time Complexity Intuition

With a **min-heap** of size K:

- Every element gets pushed into the heap exactly once: `O(N log K)`
- `log K` because the heap never grows beyond K elements — one per list
- This is much better than `O(N log N)` when `K << N`

---

## 2. Core Idea (Most Important)

### Key Insight #1 — The Heap Acts as a "Smart Comparator"

Instead of comparing across all lists manually, you maintain a **min-heap of size K**. Each entry in the heap represents the _current front element_ of one of the K lists. When you pop the minimum, you immediately know where to pull the next element from (the same list).

### Key Insight #2 — What Goes Into the Heap

You don't just store values in the heap. You store **triplets**: `(value, listIndex, elementIndex)`. This is crucial — when you pop a minimum value, you need to know _which list_ it came from so you can push the next element from that list.

```
Heap entry = { value, list_index, element_index }
```

### What to Remember

> **"Push the first element of each list. Then: pop min → add to result → push next from same list."**

That single loop is the entire pattern.

---

## 3. When to Use This Pattern

| Signal in Question                             | What It Means                         |
| ---------------------------------------------- | ------------------------------------- |
| "K sorted lists/arrays"                        | Classic K-way merge                   |
| "Merge sorted linked lists"                    | K-way merge on lists                  |
| "Find the Kth smallest across K sorted arrays" | K-way merge with early stop           |
| "Smallest range covering all K lists"          | K-way merge with sliding window logic |
| "Stream of sorted data from K sources"         | K-way merge on streams                |
| "External sort" or "merge sorted files"        | K-way merge is the algorithm          |
| "Find the median from a data stream"           | Variant using two heaps               |

---

## 4. Types / Variants

### Type 1 — Merge K Sorted Lists (Full Output)

The classic form. You want the complete merged sorted sequence. Use a min-heap, process all elements, collect results.

### Type 2 — Find Kth Smallest Element

Same setup as Type 1, but you **stop as soon as you've popped K elements**. No need to collect everything — just count pops.

### Type 3 — Smallest Range Covering K Lists

More complex variant. You must always have **at least one element from every list** in your current "window." Use a min-heap to track the minimum, and a running maximum to track the range. Advance only the minimum's list.

### Type 4 — Two-Heap Split (Median Finder)

Not a merge problem in the traditional sense, but uses the same "maintain heap invariants" thinking. One max-heap for the lower half, one min-heap for the upper half.

---

## 5. Data Structures Used

### Min-Heap (Primary DS)

This is the heart of the pattern. A min-heap gives you `O(log K)` push and pop, and `O(1)` peek at the minimum — exactly what you need to always find the globally smallest front element.

**In Go**, you implement a heap using the `container/heap` interface by defining `Len`, `Less`, `Swap`, `Push`, and `Pop` on a custom type.

**Key properties:**

- Push: `O(log K)`
- Pop (extract min): `O(log K)`
- Peek min: `O(1)`
- Space: `O(K)` — only K elements live in the heap at once

### Linked List Nodes / Array Pointers

These are used alongside the heap to track _where we are_ in each of the K lists. Instead of copying data, you store pointers/indices.

### When to Choose Which DS

| Situation                         | Use                                        |
| --------------------------------- | ------------------------------------------ |
| K sorted arrays                   | Min-heap with `(value, listIdx, elemIdx)`  |
| K sorted linked lists             | Min-heap with `*ListNode` pointers         |
| Need both min and max efficiently | Two heaps (min + max)                      |
| Only 2 sorted lists               | Two-pointer merge (no heap needed, `O(N)`) |

---

## 6. Core Templates

### Template 1 — Merge K Sorted Arrays

```go
package main

import (
    "container/heap"
    "fmt"
)

// HeapItem stores enough info to locate the next element in its list
type HeapItem struct {
    val     int
    listIdx int // which list this came from
    elemIdx int // current position in that list
}

// MinHeap implements heap.Interface
type MinHeap []HeapItem

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(HeapItem)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}

func mergeKSortedArrays(lists [][]int) []int {
    h := &MinHeap{}
    heap.Init(h)

    // Step 1: Push the first element from each list into the heap
    for i, list := range lists {
        if len(list) > 0 {
            heap.Push(h, HeapItem{list[0], i, 0})
        }
    }

    var result []int

    // Step 2: Always pop the minimum, then push the next from the same list
    for h.Len() > 0 {
        item := heap.Pop(h).(HeapItem)
        result = append(result, item.val)

        // If the same list has more elements, push the next one
        nextIdx := item.elemIdx + 1
        if nextIdx < len(lists[item.listIdx]) {
            heap.Push(h, HeapItem{lists[item.listIdx][nextIdx], item.listIdx, nextIdx})
        }
    }

    return result
}
```

---

### Template 2 — Merge K Sorted Linked Lists

```go
// ListNode is the standard LeetCode linked list node
type ListNode struct {
    Val  int
    Next *ListNode
}

// NodeHeap is a min-heap of ListNode pointers
type NodeHeap []*ListNode

func (h NodeHeap) Len() int            { return len(h) }
func (h NodeHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *NodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *NodeHeap) Pop() interface{} {
    old := *h
    n := len(old)
    node := old[n-1]
    *h = old[:n-1]
    return node
}

func mergeKLists(lists []*ListNode) *ListNode {
    h := &NodeHeap{}
    heap.Init(h)

    // Push the head of each non-nil list
    for _, node := range lists {
        if node != nil {
            heap.Push(h, node)
        }
    }

    // Dummy head trick to simplify result list building
    dummy := &ListNode{}
    curr := dummy

    for h.Len() > 0 {
        // Pop the smallest node
        node := heap.Pop(h).(*ListNode)
        curr.Next = node
        curr = curr.Next

        // Push the next node from the same list
        if node.Next != nil {
            heap.Push(h, node.Next)
        }
    }

    return dummy.Next
}
```

---

### Template 3 — Kth Smallest in K Sorted Arrays

```go
func kthSmallest(lists [][]int, k int) int {
    h := &MinHeap{}
    heap.Init(h)

    for i, list := range lists {
        if len(list) > 0 {
            heap.Push(h, HeapItem{list[0], i, 0})
        }
    }

    count := 0
    for h.Len() > 0 {
        item := heap.Pop(h).(HeapItem)
        count++

        // Return as soon as we've popped k elements
        if count == k {
            return item.val
        }

        nextIdx := item.elemIdx + 1
        if nextIdx < len(lists[item.listIdx]) {
            heap.Push(h, HeapItem{lists[item.listIdx][nextIdx], item.listIdx, nextIdx})
        }
    }

    return -1 // k is out of range
}
```

---

## 7. Step-by-Step Dry Run

**Input:** 3 sorted arrays

```
List 0: [1, 4, 7]
List 1: [2, 5, 8]
List 2: [3, 6, 9]
```

**Step 1 — Initialize heap** with first elements:

```
Heap: [(1, L0, 0), (2, L1, 0), (3, L2, 0)]
Min = 1
```

**Step 2 — Pop 1** (from L0, idx 0). Push next from L0: value 4.

```
Result: [1]
Heap: [(2, L1, 0), (3, L2, 0), (4, L0, 1)]
```

**Step 3 — Pop 2** (from L1, idx 0). Push next from L1: value 5.

```
Result: [1, 2]
Heap: [(3, L2, 0), (4, L0, 1), (5, L1, 1)]
```

**Step 4 — Pop 3** (from L2, idx 0). Push next from L2: value 6.

```
Result: [1, 2, 3]
Heap: [(4, L0, 1), (5, L1, 1), (6, L2, 1)]
```

This pattern continues until all elements are processed.

**Final Result:** `[1, 2, 3, 4, 5, 6, 7, 8, 9]` ✅

**Key observation:** The heap never has more than K=3 elements. Each pop + push is `O(log 3)`.

---

## 8. Must-Know Problems (Deep Explanation)

### Problem 1 — Merge K Sorted Lists (LeetCode #23)

**Intuition:** Each of the K linked lists is sorted, so the next element to add to the result must be the smallest _head_ across all lists. A min-heap of size K always gives you that smallest head in `O(log K)`.

**Approach:**

1. Push all K list heads into a min-heap.
2. Pop the minimum node, attach it to your result list.
3. Push that node's `.Next` (if it exists) back into the heap.
4. Repeat until heap is empty.

```go
func mergeKLists(lists []*ListNode) *ListNode {
    h := &NodeHeap{}
    heap.Init(h)
    for _, node := range lists {
        if node != nil {
            heap.Push(h, node)
        }
    }
    dummy := &ListNode{}
    curr := dummy
    for h.Len() > 0 {
        node := heap.Pop(h).(*ListNode)
        curr.Next = node
        curr = curr.Next
        if node.Next != nil {
            heap.Push(h, node.Next)
        }
    }
    return dummy.Next
}
// Time: O(N log K), Space: O(K)
```

---

### Problem 2 — Kth Smallest Element in a Sorted Matrix (LeetCode #378)

**Intuition:** Treat each _row_ of the matrix as a sorted list. You now have an N-row K-way merge problem. The Kth element popped from the min-heap is the answer.

**Approach:**

1. Push the first element of each row `(matrix[i][0], row=i, col=0)` into the heap.
2. Pop K times. Each time, push the next element in the same row.
3. The Kth pop is your answer.

```go
func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    h := &MinHeap{}
    heap.Init(h)

    // Treat each row as a sorted list; push row starts
    for i := 0; i < n; i++ {
        heap.Push(h, HeapItem{matrix[i][0], i, 0})
    }

    count := 0
    for h.Len() > 0 {
        item := heap.Pop(h).(HeapItem)
        count++
        if count == k {
            return item.val
        }
        // Push next element from the same row
        if item.elemIdx+1 < n {
            heap.Push(h, HeapItem{
                matrix[item.listIdx][item.elemIdx+1],
                item.listIdx,
                item.elemIdx + 1,
            })
        }
    }
    return -1
}
// Time: O(K log N), Space: O(N)
```

---

### Problem 3 — Find K Pairs with Smallest Sums (LeetCode #373)

**Intuition:** You have arrays `nums1` and `nums2`. The smallest sum pair starts with `(nums1[0], nums2[0])`. When you pick a pair `(i, j)`, the next candidates are `(i+1, j)` and `(i, j+1)`. This is a K-way merge of "virtual" sorted streams.

**Approach:** Initialize heap with `(nums1[i] + nums2[0], i, 0)` for all i in [0, k). For each pop, push `(i, j+1)` as the next candidate from that stream.

```go
type PairItem struct {
    sum, i, j int
}
type PairHeap []PairItem

func (h PairHeap) Len() int            { return len(h) }
func (h PairHeap) Less(a, b int) bool  { return h[a].sum < h[b].sum }
func (h PairHeap) Swap(a, b int)       { h[a], h[b] = h[b], h[a] }
func (h *PairHeap) Push(x interface{}) { *h = append(*h, x.(PairItem)) }
func (h *PairHeap) Pop() interface{} {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    if len(nums1) == 0 || len(nums2) == 0 {
        return nil
    }
    h := &PairHeap{}
    heap.Init(h)

    // Each "stream" is anchored at nums1[i], walking through nums2
    for i := 0; i < len(nums1) && i < k; i++ {
        heap.Push(h, PairItem{nums1[i] + nums2[0], i, 0})
    }

    var result [][]int
    for h.Len() > 0 && len(result) < k {
        item := heap.Pop(h).(PairItem)
        result = append(result, []int{nums1[item.i], nums2[item.j]})
        // Advance along nums2 for the same nums1[i]
        if item.j+1 < len(nums2) {
            heap.Push(h, PairItem{nums1[item.i] + nums2[item.j+1], item.i, item.j + 1})
        }
    }
    return result
}
// Time: O(K log K), Space: O(K)
```

---

## 9. Common Mistakes / Gotchas

**Mistake 1 — Storing only the value in the heap.**  
You need `(value, listIndex, elementIndex)` triplet. Without the list index and element index, you can't push the _next_ element from the correct list after a pop.

**Mistake 2 — Not handling empty lists.**  
Always check `if len(list) > 0` or `if node != nil` before pushing to the heap during initialization. A nil/empty list head will cause a panic.

**Mistake 3 — Confusing max-heap with min-heap.**  
For merging in ascending order, you always need a **min-heap**. In Go's `container/heap`, the default is a min-heap when `Less` returns `h[i] < h[j]`. Flipping this gives a max-heap — make sure you know which one you're building.

**Mistake 4 — Pushing out-of-bounds index.**  
After popping `(value, listIdx, elemIdx)`, before pushing `elemIdx+1`, always check `elemIdx+1 < len(lists[listIdx])`. Forgetting this causes an index-out-of-range panic.

**Mistake 5 — Not initializing the heap with `heap.Init`.**  
In Go, even an empty heap must be initialized with `heap.Init(h)` before use. Skipping this can cause incorrect ordering.

**Mistake 6 — Using a dummy head wrong for linked lists.**  
The dummy head trick (`dummy := &ListNode{}; curr := dummy`) simplifies list building. Return `dummy.Next` — not `dummy` — at the end.

**Mistake 7 — Trying two-pointers for K > 2.**  
Two-pointer merge only works for exactly 2 lists. For K lists, that degrades to `O(KN)` which is much worse than `O(N log K)`. Always use a heap when K > 2.

---

## 10. Time & Space Complexity

| Operation              | Time Complexity  | Why                                                             |
| ---------------------- | ---------------- | --------------------------------------------------------------- |
| Initialize heap        | `O(K log K)`     | Push K elements, each takes `O(log K)`                          |
| Process all N elements | `O(N log K)`     | Each of N elements is pushed/popped once, each op is `O(log K)` |
| **Total**              | **`O(N log K)`** | Dominant term                                                   |
| Space (heap)           | `O(K)`           | Heap holds at most one element per list                         |
| Space (result)         | `O(N)`           | For storing merged output                                       |

**Why is `O(N log K)` better than `O(N log N)`?**  
When K is small (like 10 lists of 1000 elements), `log K = log 10 ≈ 3.3` vs `log N = log 10000 ≈ 13.3`. That's a 4x improvement — and it grows with N.

---

## 11. Practice Problems

### 🟢 Easy

| #    | Problem                | Hint                                                                         |
| ---- | ---------------------- | ---------------------------------------------------------------------------- |
| 21   | Merge Two Sorted Lists | Warm-up; try two-pointer first, then generalize the idea to a heap           |
| 88   | Merge Sorted Array     | Two-pointer from the _back_ to merge in-place — foundation of merge thinking |
| 506  | Relative Ranks         | Sort with index tracking; related to "ordering across multiple sources"      |
| 1122 | Relative Sort Array    | Custom sort based on position in another array — uses sorted ordering logic  |

### 🟡 Medium

| #    | Problem                                                | Hint                                                                                      |
| ---- | ------------------------------------------------------ | ----------------------------------------------------------------------------------------- |
| 23   | Merge K Sorted Lists                                   | The core K-Way Merge template problem — must do this first                                |
| 373  | Find K Pairs with Smallest Sums                        | Treat each `(nums1[i], nums2[j])` stream as a sorted list; push `(i, j+1)` after each pop |
| 378  | Kth Smallest Element in a Sorted Matrix                | Treat each row as a sorted list; pop K times                                              |
| 632  | Smallest Range Covering K Lists                        | K-way merge + maintain a running max; advance only the minimum's list                     |
| 786  | Kth Smallest Prime Fraction                            | Min-heap on fractions; similar triplet structure to array K-way merge                     |
| 1439 | Find the Kth Smallest Sum of a Matrix With Sorted Rows | Extend K-pair logic to multiple rows — classic K-way merge generalization                 |
| 295  | Find Median from Data Stream                           | Two-heap pattern (max-heap + min-heap) to maintain median dynamically                     |

### 🔴 Hard

| #    | Problem                     | Hint                                                                                  |
| ---- | --------------------------- | ------------------------------------------------------------------------------------- |
| 4    | Median of Two Sorted Arrays | Binary search variant, but grounding in merge intuition helps                         |
| 502  | IPO                         | Two heaps — max-heap for profits, min-heap for capital; prioritize available projects |
| 1675 | Minimize Deviation in Array | Use a max-heap; repeatedly minimize the maximum while tracking the minimum            |

---

## 12. Quick Revision (60-Second Cheat Sheet)

**Core Idea:**  
"Push the first element of each list. Loop: pop min → add to result → push next from same list."

**When to Use:**

- K sorted lists/arrays → merge into one sorted output
- "Kth smallest across K sorted sequences" → stop after K pops
- Any problem where you always need the current minimum across K sorted streams

**The Non-Negotiable Rule:**  
Always store `(value, listIndex, elementIndex)` in the heap — never just the value alone.

**Template Skeleton (Go):**

```go
// 1. Define heap item with value + position info
// 2. heap.Init, push first element of each list
// 3. for h.Len() > 0:
//      item = heap.Pop(h)
//      result = append(result, item.val)
//      if item has next → heap.Push(h, next item)
```

**Complexity:**  
`O(N log K)` time, `O(K)` space — where N = total elements, K = number of lists.

**Go Heap Reminder:**  
Implement `Len`, `Less`, `Swap`, `Push`, `Pop` on your type. Min-heap: `Less(i,j) = h[i].val < h[j].val`. Always call `heap.Init` before use. Use `heap.Push` and `heap.Pop` (not direct slice ops).
