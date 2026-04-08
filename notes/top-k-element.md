**🏔️ Heaps --- The Data Structure Explained**

_Before the pattern. Build intuition first._

**1. What Is a Heap? (The Honest Explanation)**

A heap is a special binary tree with one rule:

→ Min-Heap: Every parent is SMALLER than its children. The top is always
the smallest.

→ Max-Heap: Every parent is LARGER than its children. The top is always
the largest.

That\'s it. One rule. Everything else follows from this.

**2. Why Should You Care?**

The problem heaps solve: \'Give me the min (or max) element FAST, even
as elements are added/removed.\'

Arrays can find min/max but it costs O(n) --- you scan everything.

A sorted array gives O(1) access but O(n) insertion.

A Heap gives you BOTH: O(1) peek at min/max + O(log n) insert/remove.

**3. Real-World Analogy**

Think of a hospital ER:

- Patients arrive at random times

- The most critical patient (highest priority) is always treated first

- When that patient is treated, the next most critical takes their
  spot

The ER queue IS a max-heap on \'criticality score\'. O(log n) to add a
patient, O(1) to see who\'s next, O(log n) to remove them after
treatment.

**4. How Does It Work Internally? (Simplified)**

A heap is stored as a simple ARRAY (not actually a tree in memory):

---

// Min-Heap stored as array:

// Index: 0 1 2 3 4 5 6

// Value: \[1, 3, 2, 7, 5, 4, 8\]

// Relationships (for node at index i):

// Left child → 2\*i + 1

// Right child → 2\*i + 2

// Parent → (i-1) / 2

---

The array \[1, 3, 2, 7, 5, 4, 8\] represents this tree:

---

1 ← root (always min)

/ \\

3 2

/ \\ / \\

7 5 4 8

---

**5. The Two Key Operations**

**Insert (Heapify Up)**

Add to end of array. Then \'bubble up\' --- swap with parent until heap
property is restored.

---

// Insert 0 into \[1, 3, 2, 7, 5, 4, 8\]:

// Step 1: append → \[1, 3, 2, 7, 5, 4, 8, 0\]

// Step 2: 0 \< parent(3)? Yes → swap → \[1, 3, 2, 0, 5, 4, 8, 7\]

// Step 3: 0 \< parent(1)? Yes → swap → \[0, 3, 2, 1, 5, 4, 8, 7\]

// Step 4: 0 is root. Done! O(log n) swaps

---

**Remove Min (Heapify Down)**

Swap root with last element. Remove last. Then \'bubble down\' --- swap
with smaller child until fixed.

---

// Remove min from \[1, 3, 2, 7, 5, 4, 8\]:

// Step 1: swap root+last → \[8, 3, 2, 7, 5, 4, 1\]

// Step 2: remove last → \[8, 3, 2, 7, 5, 4\]

// Step 3: 8 \> min(3,2)=2 → swap → \[2, 3, 8, 7, 5, 4\]

// Step 4: 8 \> min(4)=4 → swap → \[2, 3, 4, 7, 5, 8\]

// Done! O(log n)

---

**6. Heap in Go**

Go doesn\'t have a built-in heap like Python\'s heapq. You use the
\'container/heap\' package:

---

import \"container/heap\"

// Min-Heap of integers

type MinHeap \[\]int

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h\[i\] \< h\[j\] } // \<
= min-heap

func (h MinHeap) Swap(i, j int) { h\[i\], h\[j\] = h\[j\], h\[i\] }

func (h \*MinHeap) Push(x any) { \*h = append(\*h, x.(int)) }

func (h \*MinHeap) Pop() any {

old := \*h; n := len(old)

x := old\[n-1\]; \*h = old\[:n-1\]

return x

}

// Usage:

h := &MinHeap{5, 3, 1, 4, 2}

heap.Init(h) // build heap from slice: O(n)

heap.Push(h, 0) // insert: O(log n)

min := heap.Pop(h) // remove min: O(log n)

peekMin := (\*h)\[0\] // peek min: O(1)

---

For Max-Heap, just flip the Less function:

---

func (h MaxHeap) Less(i, j int) bool { return h\[i\] \> h\[j\] } // \>
= max-heap

---

**7. Heap Complexity Cheat Sheet**

---

**Operation** **Time** **Why**

---

Peek min/max O(1) Root is always at index 0

Insert O(log n) Bubble up at most tree height

Remove min/max O(log n) Bubble down at most tree height

Build from array O(n) Math trick --- heapify is O(n)
not O(n log n)

Search arbitrary O(n) Heap gives NO ordering guarantee
for non-root

---

**8. When to Reach for a Heap**

Mental trigger: \'I need the smallest/largest K things, efficiently, as
data comes in.\'

- Finding K largest / K smallest elements

- Merging K sorted arrays/lists

- Running median (two heaps trick)

- Scheduling / priority queues

- Dijkstra\'s shortest path (heap for min distances)

**🎯 Top \'K\' Elements Pattern**

_Pattern #5 --- Interview-Focused DSA Notes in Go_

**1. Pattern Overview**

**What Is This Pattern?**

The Top K Elements pattern helps you find the K largest, K smallest, K
most frequent, or K closest elements from a dataset.

The core insight: You don\'t need to sort everything. You just need to
maintain a small window of K candidates.

**Real-World Analogy**

Imagine Spotify tracking the Top 10 most-played songs globally across
millions of plays. They don\'t sort all songs every second --- they
maintain a rolling list of top 10. If a new song beats the least popular
song in the top 10, it replaces it. This is exactly how the pattern
works.

**Why Not Just Sort?**

---

**Approach** **Time** **Good When?**

---

Sort all, take K O(n log n) K is close to n

Min-Heap of size K O(n log K) K \<\< n (K is much smaller than
n)

Quickselect O(n) avg Only need Kth element, not sorted
K

---

The heap approach shines because log K is much smaller than log N when K
is small. If n=1,000,000 and K=10, you\'re doing \~10 operations per
element instead of \~20.

**2. Core Idea (Most Important)**

KEY INSIGHT #1: Use a Min-Heap of size K to find K LARGEST elements.

This seems backwards at first --- why min-heap for largest? Because:

- The min-heap\'s top is the SMALLEST of your K candidates

- When a new element arrives, compare it to the min (heap top)

- If new \> min: evict the min, insert the new element

- At the end: heap contains the K largest elements

KEY INSIGHT #2: Use a Max-Heap of size K to find K SMALLEST elements.
(Symmetric logic)

Think of the heap as a \'bouncer at a VIP club\' --- only K people
allowed. A new person only gets in if they\'re \'better\' than the worst
person currently inside.

**3. When to Use This Pattern**

---

**Signal in Question** **What It Means**

---

\'K largest / K smallest\' Direct pattern match → use heap of size K

\'K most frequent\' Frequency map + heap on frequencies

\'K closest to Heap sorted by distance
origin/point\'

\'Top K\' anything Same pattern, different comparison key

\'Kth largest/smallest\' Heap of size K, return heap top

\'Running / stream of Heap is perfect for dynamic data
data\'

---

**4. Types / Variants**

**Type 1: K Largest Elements**

Use Min-Heap of size K. At the end, heap contains the K largest.

**Type 2: K Smallest Elements**

Use Max-Heap of size K. At the end, heap contains the K smallest.

**Type 3: K Most Frequent**

Build frequency map first. Then use Min-Heap of size K sorted by
frequency. Or use bucket sort for O(n).

**Type 4: K Closest Points**

Use Max-Heap of size K sorted by distance. Same logic as K largest, just
different comparison key.

**5. Data Structures Used**

**Min-Heap (Primary DS)**

- Gives O(1) access to minimum element

- O(log k) insert and delete

- Used for finding K largest (counter-intuitive but correct)

**Max-Heap**

- Gives O(1) access to maximum element

- Used for finding K smallest

- Also used in \'K closest\' problems

**HashMap**

- Used in \'K most frequent\' to count frequencies first

- O(1) insert and lookup

- Always the first step in frequency-based problems

**Bucket Sort Array (Alternative)**

- For K most frequent: buckets\[frequency\] = list of elements

- O(n) time, avoids heap entirely

- Use when you need O(n) and frequency is bounded by array length

**6. Core Templates**

**Template 1: K Largest Elements (Min-Heap)**

---

import \"container/heap\"

type MinHeap \[\]int

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h\[i\] \< h\[j\] }

func (h MinHeap) Swap(i, j int) { h\[i\], h\[j\] = h\[j\], h\[i\] }

func (h \*MinHeap) Push(x any) { \*h = append(\*h, x.(int)) }

func (h \*MinHeap) Pop() any {

old := \*h; n := len(old)

x := old\[n-1\]; \*h = old\[:n-1\]

return x

}

func kLargest(nums \[\]int, k int) \[\]int {

h := &MinHeap{}

heap.Init(h)

for \_, num := range nums {

heap.Push(h, num)

if h.Len() \> k {

heap.Pop(h) // remove smallest --- only keep top K

}

}

return \[\]int(\*h) // heap contains K largest

}

---

**Template 2: K Most Frequent (Frequency Map + Heap)**

---

type FreqItem struct{ val, freq int }

type FreqHeap \[\]FreqItem

func (h FreqHeap) Len() int { return len(h) }

func (h FreqHeap) Less(i, j int) bool { return h\[i\].freq \<
h\[j\].freq } // min-heap on freq

func (h FreqHeap) Swap(i, j int) { h\[i\], h\[j\] = h\[j\], h\[i\] }

func (h \*FreqHeap) Push(x any) { \*h = append(\*h, x.(FreqItem)) }

func (h \*FreqHeap) Pop() any {

old := \*h; n := len(old)

x := old\[n-1\]; \*h = old\[:n-1\]

return x

}

func topKFrequent(nums \[\]int, k int) \[\]int {

// Step 1: count frequencies

freq := map\[int\]int{}

for \_, n := range nums { freq\[n\]++ }

// Step 2: maintain min-heap of size K on frequency

h := &FreqHeap{}

heap.Init(h)

for val, f := range freq {

heap.Push(h, FreqItem{val, f})

if h.Len() \> k { heap.Pop(h) }

}

// Step 3: extract results

res := make(\[\]int, k)

for i := k - 1; i \>= 0; i\-- {

res\[i\] = heap.Pop(h).(FreqItem).val

}

return res

}

---

**Template 3: Kth Largest Element**

---

func findKthLargest(nums \[\]int, k int) int {

h := &MinHeap{}

heap.Init(h)

for \_, num := range nums {

heap.Push(h, num)

if h.Len() \> k { heap.Pop(h) }

}

return (\*h)\[0\] // root = Kth largest

}

---

**7. Step-by-Step Dry Run**

Problem: Find 3 largest from \[3, 1, 5, 12, 2, 11\]. K = 3

We maintain a Min-Heap of size K=3. The top is always the smallest of
our candidates.

---

Process 3 → heap: \[3\] size=1 ≤ 3, no pop

Process 1 → heap: \[1, 3\] size=2 ≤ 3, no pop

Process 5 → heap: \[1, 3, 5\] size=3 ≤ 3, no pop

Process 12 → heap: \[1, 3, 5, 12\] size=4 \> 3!

→ pop min(1) → heap: \[3, 5, 12\]

Process 2 → heap: \[2, 3, 5, 12\]\... 2 \< min(3)? No wait\...

→ push 2 → heap: \[2, 3, 12, 5\] size=4 \> 3

→ pop min(2) → heap: \[3, 5, 12\] ✓ 2 didn\'t make it

Process 11 → push → \[3, 5, 12, 11\] size=4 \> 3

→ pop min(3) → heap: \[5, 11, 12\]

Result: \[5, 11, 12\] ✓ The 3 largest elements!

---

**8. Must-Know Problems**

**Problem 1: Kth Largest Element in an Array (LC 215) --- Medium**

Intuition: Maintain a min-heap of size K. After processing all elements,
the root IS the Kth largest because K-1 elements are larger than it.

---

func findKthLargest(nums \[\]int, k int) int {

h := &MinHeap{}

heap.Init(h)

for \_, num := range nums {

heap.Push(h, num)

if h.Len() \> k {

heap.Pop(h)

}

}

return (\*h)\[0\] // Kth largest is at root

}

// Time: O(n log k) Space: O(k)

---

**Problem 2: Top K Frequent Elements (LC 347) --- Medium**

Intuition: First count frequencies (O(n) hashmap), then find K largest
by frequency (heap on freq). Classic two-step approach.

---

func topKFrequent(nums \[\]int, k int) \[\]int {

freq := map\[int\]int{}

for \_, n := range nums { freq\[n\]++ }

h := &FreqHeap{} // min-heap on frequency

heap.Init(h)

for val, f := range freq {

heap.Push(h, FreqItem{val, f})

if h.Len() \> k { heap.Pop(h) }

}

res := make(\[\]int, k)

for i := k - 1; i \>= 0; i\-- {

res\[i\] = heap.Pop(h).(FreqItem).val

}

return res

}

// Time: O(n log k) Space: O(n) for map

---

**Problem 3: K Closest Points to Origin (LC 973) --- Medium**

Intuition: Distance = x² + y² (don\'t need sqrt). Use Max-Heap of size K
on distance --- evict the farthest when size exceeds K.

---

type Point struct{ x, y, dist int }

type MaxDistHeap \[\]Point

func (h MaxDistHeap) Len() int { return len(h) }

func (h MaxDistHeap) Less(i, j int) bool { return h\[i\].dist \>
h\[j\].dist } // MAX-heap

func (h MaxDistHeap) Swap(i, j int) { h\[i\], h\[j\] = h\[j\], h\[i\] }

func (h \*MaxDistHeap) Push(x any) { \*h = append(\*h, x.(Point)) }

func (h \*MaxDistHeap) Pop() any {

old := \*h; n := len(old)

x := old\[n-1\]; \*h = old\[:n-1\]

return x

}

func kClosest(points \[\]\[\]int, k int) \[\]\[\]int {

h := &MaxDistHeap{}

heap.Init(h)

for \_, p := range points {

d := p\[0\]\*p\[0\] + p\[1\]\*p\[1\]

heap.Push(h, Point{p\[0\], p\[1\], d})

if h.Len() \> k { heap.Pop(h) } // remove farthest

}

res := make(\[\]\[\]int, k)

for i := range res {

pt := heap.Pop(h).(Point)

res\[i\] = \[\]int{pt.x, pt.y}

}

return res

}

// Time: O(n log k) Space: O(k)

---

**9. Common Mistakes / Gotchas**

- ❌ Using Max-Heap for K largest (should be Min-Heap). Remember:
  Min-Heap keeps K largest because small elements get evicted.

- ❌ Forgetting heap.Init(h) before use. Always initialize before
  push/pop.

- ❌ Using sqrt for distance --- never needed! x²+y² comparison gives
  same order.

- ❌ Off-by-one: popping when h.Len() \>= k instead of \> k. Pop when
  size EXCEEDS k.

- ❌ Returning heap top for \'K largest\' problems when you want all K
  elements --- collect the whole heap.

- ❌ Forgetting to build a frequency map before \'K most frequent\'
  problems. Always map first.

- ❌ Type assertion panic: heap.Pop(h).(int) --- make sure your
  Push/Pop types match.

**10. Time & Space Complexity**

---

**Operation** **Complexity** **Why**

---

Build heap from O(n) heap.Init is O(n) --- math
scratch magic

Insert element O(log k) Bubble up in heap of size k

Remove min/max O(log k) Bubble down in heap of size k

Full n-element scan O(n log k) n elements × O(log k) per
insert

Space O(k) Heap only holds k elements at
once

Frequency map step O(n) One pass to count

---

**11. Practice Problems**

**🟢 Easy**

---

**\#** **Problem** **Hint**

---

703 Kth Largest in a Stream Min-heap of size K; each new element:
push + maybe pop

1337 K Weakest Rows in Matrix Heap on (soldiers_count, row_index);
keep K smallest

1046 Last Stone Weight Max-heap; simulate: pop two, push
difference if any

2231 Largest Number After Digit Min-heap on digit value;
Swaps by Parity straightforward heap exercise

---

**🟡 Medium**

---

**\#** **Problem** **Hint**

---

215 Kth Largest Element in Min-heap size K; root = answer
Array

347 Top K Frequent Elements Freq map first, then min-heap on
frequency

973 K Closest Points to Origin Max-heap on x²+y²; evict farthest

658 Find K Closest Elements Min-heap on \|x - target\|; collect K
elements

1985 Find the Kth Largest Strings as numbers; min-heap with
Integer in Array string comparison

373 Find K Pairs with Smallest Min-heap on sum; classic K-way merge
Sums variant

767 Reorganize String Max-heap on char frequency; greedily
pick most frequent

1167 Minimum Cost to Connect Min-heap; always merge two smallest
Sticks (classic greedy)

---

**🔴 Hard**

---

**\#** **Problem** **Hint**

---

295 Find Median from Data Two heaps: max-heap (lower half) +
Stream min-heap (upper half)

358 Rearrange String K Max-heap + cooldown queue of size K
Distance Apart

1439 Find the Kth Smallest Sum Min-heap on (sum, i, j); expand
of a Matrix neighbors

---

**12. Quick Revision (60 sec)**

---

KEY IDEA:

K Largest → Min-Heap of size K (evict min when size \> K)

K Smallest → Max-Heap of size K (evict max when size \> K)

K Frequent → freq map + min-heap on frequency

K Closest → max-heap on distance (evict farthest)

TEMPLATE (K Largest):

for each num in nums:

heap.Push(h, num)

if h.Len() \> k: heap.Pop(h)

return heap contents

WHEN TO USE:

Signal words: \'top K\', \'K largest/smallest\', \'K most frequent\'

Stream of data where you can\'t sort everything

Need better than O(n log n) when K \<\< n

COMPLEXITY: O(n log k) time, O(k) space

---
