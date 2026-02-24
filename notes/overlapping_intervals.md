# DSA Pattern Notes

# Overlapping Intervals

### Complete Interview Preparation Guide

---

## Quick Reference Card

> **CORE INSIGHT**
> Sort intervals by start time. Two intervals [a,b] and [c,d] OVERLAP if a <= d AND c <= b.
> They DO NOT overlap when b < c (A ends before B starts).
> Compact form: max(a,c) <= min(b,d)

> **WHEN TO USE THIS PATTERN**
>
> - Problem involves intervals, ranges, or time slots
> - Keywords: merge, overlap, conflict, schedule, gaps, cover, meeting rooms
> - You need to combine, remove, or count overlapping ranges
> - Free time / available slots need to be found

---

## 1. Fundamentals

### 1.1 What Is an Interval?

An interval is a pair `[start, end]` representing a contiguous range (e.g. time, index). We always assume `start <= end`.

### 1.2 Overlap Conditions

> **Two intervals A=[a1,a2] and B=[b1,b2] OVERLAP if:** `a1 <= b2  AND  b1 <= a2`
> **They DO NOT OVERLAP if:** `a2 < b1  OR  b2 < a1` (one ends before the other starts)
>
> Inclusive: [1,3] and [3,5] share point 3 → overlap
> Exclusive: [1,3) and [3,5) → no overlap
> **ALWAYS clarify inclusive vs exclusive with interviewer!**

Visualizing all overlap cases:

```
Case 1 - No overlap (A before B):   A=[1,3]  B=[5,8]
         [---A---]            [----B----]

Case 2 - Partial overlap:           A=[1,5]  B=[3,8]
         [----A----]
                [----B----]

Case 3 - A fully contains B:        A=[1,8]  B=[3,5]
         [--------A--------]
               [--B--]

Case 4 - B fully contains A:        A=[3,5]  B=[1,8]
         [--------B--------]
               [--A--]

Case 5 - Touch at one point:        A=[1,3]  B=[3,5]
         [--A--][--B--]   <- inclusive overlap
```

---

## 2. Data Structures

### 2.1 Array / List of Intervals

Most common. Each interval is `[start, end]`. Sorting is the critical preprocessing step.

```go
// Go representation
type Interval struct{ Start, End int }

// Or use [][]int
intervals := [][]int{{1,3},{2,6},{8,10},{15,18}}

// Sort by start time - O(n log n)
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
})
```

> **SORTING — THE MOST CRITICAL STEP**
>
> - Sort by **START** time: merging, insertion, finding gaps, scheduling order
> - Sort by **END** time: greedy removal / activity selection / arrow problems
> - Tie-break: sort by end when starts are equal (problem-dependent).
> - Time cost: O(n log n) — this usually dominates overall complexity.

### 2.2 Min-Heap (Priority Queue)

Used when you need to track the EARLIEST ending interval dynamically (e.g. meeting rooms, task scheduling).

```go
// Go MinHeap for end times
import "container/heap"

type MinHeap []int
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
    old := *h; n := len(old)
    x := old[n-1]; *h = old[:n-1]; return x
}

h := &MinHeap{}
heap.Init(h)
heap.Push(h, endTime)   // O(log n)
minEnd := (*h)[0]       // peek min - O(1)
heap.Pop(h)             // remove min - O(log n)
```

> **When to use Heap vs Plain Sort**
>
> - Plain sort → merging, inserting, finding gaps (offline — process all at once)
> - Min-Heap → meeting rooms II, task scheduler (need MINIMUM end time dynamically)

### 2.3 Stack-based Merge

A stack naturally tracks the 'current merged interval' at its top. Use for merge-style problems.

```go
stack := [][]int{}
for _, interval := range sortedIntervals {
    if len(stack) == 0 || stack[len(stack)-1][1] < interval[0] {
        stack = append(stack, interval)  // no overlap, push new
    } else {
        // overlap - extend the end
        top := stack[len(stack)-1]
        if interval[1] > top[1] {
            stack[len(stack)-1][1] = interval[1]
        }
    }
}
```

---

## 3. Core Algorithms

### 3.1 Merge Overlapping Intervals [LC 56]

Foundational algorithm. Sort by start, then sweep and merge overlapping intervals.

```go
// O(n log n) time, O(n) space
func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    res := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        last := res[len(res)-1]
        curr := intervals[i]
        if curr[0] <= last[1] {              // overlap check
            last[1] = max(last[1], curr[1])  // extend end
        } else {
            res = append(res, curr)           // no overlap, add new
        }
    }
    return res
}

// Dry run: [[1,3],[2,6],[8,10],[15,18]]
// res=[[1,3]] -> [2,6] overlaps (2<=3) -> res=[[1,6]]
// [8,10] no overlap (8>6) -> res=[[1,6],[8,10]]
// [15,18] no overlap -> res=[[1,6],[8,10],[15,18]]
```

---

### 3.2 Insert Interval [LC 57]

Insert a new interval into a sorted non-overlapping list and merge. Three-phase approach.

```go
// O(n) time, O(n) space
func insert(intervals [][]int, newInterval []int) [][]int {
    res := [][]int{}
    i, n := 0, len(intervals)

    // Phase 1: add all intervals ending BEFORE new interval starts
    for i < n && intervals[i][1] < newInterval[0] {
        res = append(res, intervals[i])
        i++
    }

    // Phase 2: merge all OVERLAPPING intervals
    for i < n && intervals[i][0] <= newInterval[1] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    }
    res = append(res, newInterval)

    // Phase 3: add remaining intervals
    res = append(res, intervals[i:]...)
    return res
}
```

---

### 3.3 Meeting Rooms II [LC 253]

Minimum number of rooms to schedule all meetings. Two approaches.

```go
// Approach 1: Min-Heap  O(n log n) time, O(n) space
func minMeetingRooms(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    h := &MinHeap{}
    heap.Init(h)
    for _, iv := range intervals {
        if h.Len() > 0 && (*h)[0] <= iv[0] {
            heap.Pop(h)  // room freed (earliest meeting ended)
        }
        heap.Push(h, iv[1])  // allocate/reuse room
    }
    return h.Len()  // active rooms = answer
}

// Approach 2: Two sorted arrays sweep  O(n log n) time, O(n) space
func minMeetingRooms2(intervals [][]int) int {
    n := len(intervals)
    starts, ends := make([]int, n), make([]int, n)
    for i, iv := range intervals {
        starts[i], ends[i] = iv[0], iv[1]
    }
    sort.Ints(starts); sort.Ints(ends)
    rooms, e := 0, 0
    for s := 0; s < n; s++ {
        if starts[s] < ends[e] { rooms++ } else { e++ }
    }
    return rooms
}
```

---

### 3.4 Non-overlapping Intervals [LC 435]

Minimum removals to make all remaining intervals non-overlapping. Sort by END — greedy keep earliest-ending.

```go
// O(n log n) time, O(1) space
func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]  // sort by END
    })
    removed, prevEnd := 0, math.MinInt64
    for _, iv := range intervals {
        if iv[0] >= prevEnd {
            prevEnd = iv[1]   // keep this interval
        } else {
            removed++         // discard (overlaps, ends later)
        }
    }
    return removed
}
```

> **WHY SORT BY END FOR REMOVAL?**
> Activity Selection Theorem (classic greedy): Always pick the interval ending earliest.
> This maximizes the count of non-overlapping intervals we can keep.
> Sorting by start time here gives the **WRONG answer** — common interview mistake!

---

### 3.5 Interval List Intersections [LC 986]

Given two sorted lists, find all intersecting intervals. Two-pointer approach.

```go
// O(m+n) time, O(m+n) space
func intervalIntersection(A, B [][]int) [][]int {
    res := [][]int{}
    i, j := 0, 0
    for i < len(A) && j < len(B) {
        lo := max(A[i][0], B[j][0])
        hi := min(A[i][1], B[j][1])
        if lo <= hi {
            res = append(res, []int{lo, hi})  // intersection exists
        }
        // Advance pointer whose interval ends first
        if A[i][1] < B[j][1] { i++ } else { j++ }
    }
    return res
}
```

---

### 3.6 Sweep Line / Event Counting

Count maximum simultaneous events. Create +1 events at starts, -1 events at ends.

```go
// O(n log n) time, O(n) space
func maxConcurrent(intervals [][]int) int {
    events := [][2]int{}
    for _, iv := range intervals {
        events = append(events, [2]int{iv[0], 1})   // start event
        events = append(events, [2]int{iv[1], -1})  // end event
    }
    // Sort by time; at same time, ends (-1) before starts (+1)
    sort.Slice(events, func(i, j int) bool {
        if events[i][0] == events[j][0] {
            return events[i][1] < events[j][1]
        }
        return events[i][0] < events[j][0]
    })
    maxCount, count := 0, 0
    for _, e := range events {
        count += e[1]
        if count > maxCount { maxCount = count }
    }
    return maxCount
}

// Alternative: Difference Array (when range is bounded)
// diff[start]++; diff[end+1]--; then prefix sum
```

---

### 3.7 Employee Free Time [LC 759]

Find free intervals not covered by any employee's schedule. Merge all, then find gaps.

```go
// O(n log n) time, O(n) space
// Step 1: flatten all schedules into one list
// Step 2: sort by start time
// Step 3: merge all overlapping intervals
// Step 4: find gaps between consecutive merged intervals

func findGaps(merged [][]int) [][]int {
    gaps := [][]int{}
    for i := 1; i < len(merged); i++ {
        if merged[i][0] > merged[i-1][1] {
            gaps = append(gaps, []int{merged[i-1][1], merged[i][0]})
        }
    }
    return gaps
}
```

---

### 3.8 Minimum Arrows to Burst Balloons [LC 452]

Count minimum arrows to burst all balloons (each arrow bursts all balloons it passes through).

```go
// Sort by END. Shoot at the earliest end. Burst all overlapping.
// O(n log n) time, O(1) space
func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        return points[i][1] < points[j][1]  // sort by end
    })
    arrows, arrowPos := 1, points[0][1]
    for _, p := range points[1:] {
        if p[0] > arrowPos {   // balloon starts after last arrow
            arrows++
            arrowPos = p[1]    // shoot at this balloon's end
        }
    }
    return arrows
}
```

---

## 4. Time & Space Complexity

| Algorithm / Problem              | Time           | Space     |
| -------------------------------- | -------------- | --------- |
| Merge Intervals                  | O(n log n)     | O(n)      |
| Insert Interval                  | O(n)           | O(n)      |
| Meeting Rooms I                  | O(n log n)     | O(1)      |
| Meeting Rooms II                 | O(n log n)     | O(n)      |
| Interval Intersection            | O(m + n)       | O(m + n)  |
| Non-overlapping (min remove)     | O(n log n)     | O(1)      |
| Sweep Line (max overlap)         | O(n log n)     | O(n)      |
| Employee Free Time               | O(n log n)     | O(n)      |
| Min Interval per Query [LC 1851] | O((n+q) log n) | O(n + q)  |
| Difference Array (bounded range) | O(n + maxVal)  | O(maxVal) |

---

## 5. Problem-Solving Templates

### 5.1 Decision Tree

> **STEP 1:** Does order matter? → Sort by start (or end, see below)
> **STEP 2:** Am I merging/combining? → Stack/result array approach
> **STEP 3:** Am I counting concurrent things? → Sweep line or min-heap
> **STEP 4:** Working with two sorted lists? → Two pointers
> **STEP 5:** Finding what is NOT covered? → Merge first, then find gaps
>
> **SORT BY START:** merging, insertion, finding gaps, scheduling order
> **SORT BY END:** greedy removal / activity selection / arrow problems

### 5.2 Pattern Recognition Quick Map

| Pattern Keyword             | Approach                           |
| --------------------------- | ---------------------------------- |
| "merge overlapping"         | Sort by start + stack merge        |
| "meeting rooms / min rooms" | Sort + min-heap tracking end times |
| "remove minimum intervals"  | Sort by END + greedy keep          |
| "find free time / gaps"     | Merge first, then sweep for gaps   |
| "two sorted lists overlap"  | Two pointer intersection           |
| "max overlap count"         | Sweep line events (+1/-1)          |
| "cover all points / arrows" | Sort by end, greedy shoot          |
| "dynamic insert/query"      | TreeMap / sorted set / seg tree    |
| "range coverage count"      | Difference array + prefix sum      |

---

## 6. Edge Cases & Gotchas

### 6.1 Always Check These

- **Empty input:** Return empty array. Check before sorting.
- **Single interval:** Return as-is. No merging needed.
- **All intervals identical:** Should merge to exactly one interval.
- **Completely nested:** `[1,10]` contains `[2,3],[4,5]` → all merge to `[1,10]`.
- **Touch at single point:** `[1,3]` and `[3,5]` — inclusive overlap, exclusive no overlap.
- **Negative numbers:** Common in coordinate/timestamp problems.
- **Input already sorted:** Don't assume — always sort unless stated.

### 6.2 Common Bugs

```go
// BUG 1: Forgot to sort first
// Always sort - never assume input is sorted

// BUG 2: Wrong sort key
sort by start -> merge    |    sort by end -> removal

// BUG 3: Shrinking the interval end
last[1] = curr[1]                // WRONG - might shrink!
last[1] = max(last[1], curr[1])  // CORRECT

// BUG 4: Wrong overlap condition
curr[0] < last[1]   // exclusive endpoints
curr[0] <= last[1]  // inclusive endpoints (usually correct)

// BUG 5: Modifying slice in place while iterating
// Always use a separate result/output slice
```

### 6.3 Follow-up Questions in Interviews

- **Stream of intervals?** → Interval tree or ordered set (balanced BST).
- **Millions of intervals?** → External sort, segment tree, offline algorithms.
- **O(n) only?** → Only achievable if already sorted; otherwise sort is bottleneck.
- **Points instead of ranges?** → Treat point x as interval `[x, x]`.
- **Max value queries?** → Offline sort queries + min-heap (LC 1851 pattern).

---

## 7. Advanced Techniques

### 7.1 Difference Array

When you need coverage count at every point, and interval values are bounded.

```go
// O(n + maxVal) time and space
func coverageCount(intervals [][]int, maxVal int) []int {
    diff := make([]int, maxVal+2)
    for _, iv := range intervals {
        diff[iv[0]]++
        diff[iv[1]+1]--
    }
    // Prefix sum: diff[i] = count of intervals covering point i
    for i := 1; i <= maxVal; i++ {
        diff[i] += diff[i-1]
    }
    return diff[:maxVal+1]
}

// Use cases: Car Pooling (LC 1094), Corporate Flight Bookings (LC 1109)
// Maximum Population Year (LC 1854)
```

### 7.2 Coordinate Compression

When values are up to 10^9 but count is small — compress coordinates to small indices.

```go
// Compress coordinates to [0, k-1] range
coords := map[int]bool{}
for _, iv := range intervals { coords[iv[0]], coords[iv[1]] = true, true }
sorted := []int{}
for k := range coords { sorted = append(sorted, k) }
sort.Ints(sorted)
rank := map[int]int{}
for i, v := range sorted { rank[v] = i }

// Now use rank[iv[0]] and rank[iv[1]] instead of actual values
// This reduces value range from [0, 10^9] to [0, 2n-1]
```

### 7.3 Offline Query Pattern [LC 1851]

Sort both intervals and queries. Use min-heap to answer each query efficiently.

```go
// Min Interval to Include Each Query - O((n+q) log n)
func minInterval(intervals [][]int, queries []int) []int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    // Sort queries but keep original indices
    qIdx := make([]int, len(queries))
    for i := range qIdx { qIdx[i] = i }
    sort.Slice(qIdx, func(i, j int) bool {
        return queries[qIdx[i]] < queries[qIdx[j]]
    })
    // Min-heap: [size, end] sorted by size
    h := &SizeHeap{}
    heap.Init(h)
    ans := make([]int, len(queries))
    j := 0
    for _, qi := range qIdx {
        q := queries[qi]
        // Add all intervals starting <= q
        for j < len(intervals) && intervals[j][0] <= q {
            size := intervals[j][1] - intervals[j][0] + 1
            heap.Push(h, [2]int{size, intervals[j][1]})
            j++
        }
        // Remove intervals that have already ended
        for h.Len() > 0 && (*h)[0][1] < q { heap.Pop(h) }
        if h.Len() > 0 { ans[qi] = (*h)[0][0] } else { ans[qi] = -1 }
    }
    return ans
}
```

---

## 8. LeetCode Problem Set (30 Problems)

| #   | Problem                                    | LC # | Difficulty | Key Technique                |
| --- | ------------------------------------------ | ---- | ---------- | ---------------------------- |
| 1   | Merge Intervals                            | 56   | Medium     | Sort by start + stack merge  |
| 2   | Insert Interval                            | 57   | Medium     | 3-phase insert               |
| 3   | Meeting Rooms                              | 252  | Easy       | Sort + check adjacent pair   |
| 4   | Meeting Rooms II                           | 253  | Medium     | Min-heap / two-array sweep   |
| 5   | Non-overlapping Intervals                  | 435  | Medium     | Greedy, sort by end time     |
| 6   | Interval List Intersections                | 986  | Medium     | Two pointer on sorted lists  |
| 7   | Employee Free Time                         | 759  | Hard       | Merge all, then find gaps    |
| 8   | My Calendar I                              | 729  | Medium     | Overlap check with BST       |
| 9   | My Calendar II                             | 731  | Medium     | Track double bookings        |
| 10  | My Calendar III                            | 732  | Hard       | Sweep line / segment tree    |
| 11  | Min Arrows to Burst Balloons               | 452  | Medium     | Sort by end, greedy shoot    |
| 12  | Partition Labels                           | 763  | Medium     | Implicit intervals + greedy  |
| 13  | Teemo Attacking                            | 495  | Easy       | Overlapping poison durations |
| 14  | Car Pooling                                | 1094 | Medium     | Difference array on stops    |
| 15  | Corporate Flight Bookings                  | 1109 | Medium     | Difference array on seats    |
| 16  | Remove Covered Intervals                   | 1288 | Medium     | Sort, check containment      |
| 17  | Count Days Without Meetings                | 3169 | Medium     | Sort + gap sweep             |
| 18  | Count Integers in Intervals                | 2276 | Hard       | Dynamic sorted merge set     |
| 19  | Divide Intervals Into Min Groups           | 2406 | Medium     | Sweep line = meeting rooms   |
| 20  | Min Interval to Include Each Query         | 1851 | Hard       | Offline queries + min-heap   |
| 21  | Data Stream as Disjoint Intervals          | 352  | Hard       | Online sorted list + merge   |
| 22  | Range Module                               | 715  | Hard       | Interval tree / sorted map   |
| 23  | Falling Squares                            | 699  | Hard       | Coord compress + seg tree    |
| 24  | Maximum Population Year                    | 1854 | Easy       | Difference array on years    |
| 25  | Video Stitching                            | 1024 | Medium     | Greedy interval cover        |
| 26  | Jump Game II                               | 45   | Medium     | Implicit greedy intervals    |
| 27  | Maximum Sum of 3 Non-Overlapping Subarrays | 689  | Hard       | DP + interval windows        |
| 28  | Determine if String Halves are Alike       | 1704 | Easy       | Counting, range awareness    |
| 29  | Check if Array Is Sorted and Rotated       | 1752 | Easy       | Finding rotation point       |
| 30  | Summary Ranges                             | 228  | Easy       | Interval building from array |

### 8.1 Recommended Study Order

- **Week 1 — Foundation:** #56 (Merge), #57 (Insert), #252 (Meeting I), #253 (Meeting II), #452 (Arrows), #763 (Partition)
- **Week 2 — Intermediate:** #435 (Non-overlap), #986 (Intersection), #759 (Free Time), #1288 (Remove Covered), #1094 (Car Pooling)
- **Week 3 — Hard:** #732 (Calendar III), #1851 (Min Interval Query), #352 (Data Stream), #715 (Range Module)

---

## 9. Interview Strategy

### 9.1 Clarification Questions

```
1. "Are endpoints inclusive or exclusive?"
2. "Can intervals have equal start and end? (point intervals like [3,3])"
3. "Is the input already sorted? Can I assume it?"
4. "Can start > end? Should I normalize?"
5. "What are the constraints on values?"
   (for choosing diff array vs segment tree approach)
```

### 9.2 Interview Communication Template

```
Step 1 - APPROACH:  "I'll sort by start time - O(n log n)."
                    "Then one linear pass tracking the last merged interval."
                    "Total: O(n log n) time, O(n) space."

Step 2 - TRACE:     Always trace a small example before coding.
                    [[1,3],[2,6],[8,10]] -> [[1,6],[8,10]]

Step 3 - CODE:      Write clean code with helper max() min() functions.

Step 4 - TEST:      Check: empty input, single interval,
                    all overlap, none overlap, nested intervals.
```

---

## 10. Quick Revision Summary

> **CORE OVERLAP CHECK:** `max(a1,b1) <= min(a2,b2)`
> **DEFAULT SORT:** by start time ascending
> **MERGE:** if `curr.start <= last.end` then `last.end = max(last.end, curr.end)`
> **MEETING ROOMS II:** min-heap of end times, heap size = answer
> **REMOVAL / ACTIVITY SELECTION:** sort by END, greedy keep earliest-ending
> **INTERSECTION:** two pointers, advance pointer whose interval ends first
> **SWEEP LINE:** +1 at start, -1 at end, prefix max = max concurrent
> **GAPS:** merge all intervals first, then sweep between merged pairs
> **DIFFERENCE ARRAY:** `diff[start]++`, `diff[end+1]--`, then prefix sum
> **ARROWS / COVERAGE:** sort by end, shoot greedily at earliest end
> **ALWAYS CLARIFY:** inclusive or exclusive endpoints?
