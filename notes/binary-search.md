# Modified Binary Search — DSA Pattern Notes

> 🎯 **Interview Focus** | Quick Revision Ready | Go Templates Included

---

## 1. Pattern Overview

### What is this pattern?

Modified Binary Search is an extension of classic binary search. Instead of searching for an **exact value** in a sorted array, you apply binary search logic to solve a broader class of problems — finding a boundary, the first/last occurrence, a rotated element, or an answer in a "search space" that may not even be a real array.

The core mechanic stays the same: **cut the search space in half every step**. What changes is _how you decide which half to keep_.

### Why is it used?

Brute force scanning an array is O(n). Any time a problem has **monotonic behavior** — meaning one side of a point satisfies a condition and the other side doesn't — you can exploit that structure to get O(log n). Modified binary search is your weapon whenever you spot that structure.

### Real-World Analogy

Imagine a physical book. You open to the middle. If the page is too early, you go right. Too late, you go left. You never re-read pages you've eliminated. Now imagine the book's pages aren't numbered sequentially — maybe they're rotated, or you're looking for the _first page_ where a certain word appears. You still bisect, but your decision rule changes slightly. That's modified binary search.

### Time Complexity Intuition

Every iteration, you **halve** your remaining work. Starting with `n` elements:

- After 1 step → n/2
- After 2 steps → n/4
- After k steps → n/2ᵏ

You stop when n/2ᵏ = 1, meaning **k = log₂(n)** steps. That's why it's O(log n). The modification doesn't change this — you still eliminate half the space each time, as long as your condition is sound.

---

## 2. Core Idea (MOST IMPORTANT)

### Key Insight #1 — Binary Search is about Eliminating Half

The real power isn't finding a midpoint. It's confidently saying: **"The answer cannot be in this half."** If you can write a condition that tells you which half to throw away, you have a binary search.

### Key Insight #2 — Think in Terms of a Predicate (True/False Boundary)

Most modified binary search problems reduce to: _"Find the first index where condition X becomes true."_

```
Index:     0  1  2  3  4  5  6
Condition: F  F  F  T  T  T  T
                   ^
              Find this boundary
```

This is the universal mental model. Train yourself to ask: **what is the predicate, and where does it flip?**

### The Essential Formula

```
mid = left + (right - left) / 2   // Always use this, never (left+right)/2 (overflow risk)
```

### What to Remember

> ✅ "Binary search is not about sorted arrays — it's about **monotonic decision functions**."
>
> ✅ "If you can write `f(mid)` that returns true/false and the results are monotonic (all F then all T), binary search works."

---

## 3. When to Use This Pattern

| Signal in Question                                           | What It Means                  |
| ------------------------------------------------------------ | ------------------------------ |
| "Find the **first/last** occurrence of X"                    | Binary search on boundary      |
| "Array is **rotated** or partially sorted"                   | Modified mid-check logic       |
| "Find **minimum/maximum** value satisfying a condition"      | Binary search on answer        |
| "**Sorted** array + search for something"                    | Classic or modified BS         |
| "Find the **peak** element"                                  | Modified BS on slope direction |
| "Answer is in a **range** [lo, hi] and is monotonic"         | Binary search on value space   |
| "Can we do X with mid as the answer?" (check feasibility)    | Binary search on answer        |
| Constraints are large (10⁸, 10⁹) and brute force is too slow | Think BS on search space       |

---

## 4. Types / Variants

### Type 1 — Classic Modified BS (Find boundary in sorted array)

Find first/last occurrence, find insert position. Array is sorted. You tweak the condition from `==` to `<=` or `>=`.

### Type 2 — Rotated / Shifted Array BS

The array was sorted but then rotated. You need to figure out which half is still sorted to decide where to search. Key trick: **one half is always sorted**.

### Type 3 — BS on Answer (Search Space BS)

The array itself isn't what you're searching — you're searching for a _value_ (like "minimum largest sum" or "minimum speed"). You define lo/hi as the value range and use a feasibility function as your predicate.

### Type 4 — Peak / Valley Finding

Find a local peak or the mountain top. You compare `mid` with its neighbors to decide which slope you're on, and move accordingly.

---

## 5. Data Structures Used

**Array** is the primary data structure. Binary search operates on index ranges, so random access in O(1) is essential — arrays give you that.

For _BS on answer_ problems, there's no explicit array at all. You create a virtual search space `[lo, hi]` and your "array" is just the range of possible answers.

| Data Structure         | Why Used                   | Key Property                     |
| ---------------------- | -------------------------- | -------------------------------- |
| Array (sorted)         | Direct index access for BS | O(1) random access               |
| Virtual range [lo, hi] | BS on answer               | No array needed, just arithmetic |
| No extra DS needed     | BS is in-place             | O(1) space                       |

> 💡 Binary search almost never needs extra memory. It's inherently O(1) space.

---

## 6. Core Templates

### Template 1 — Find Exact Value (Classic)

```go
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid          // Found it
        } else if nums[mid] < target {
            left = mid + 1      // Target is in right half
        } else {
            right = mid - 1     // Target is in left half
        }
    }
    return -1                   // Not found
}
```

---

### Template 2 — Find First True (Left Boundary / Predicate BS)

This covers ~60% of modified BS problems. Use when you need the **first index** where a condition becomes true.

```go
func findFirstTrue(nums []int, target int) int {
    left, right := 0, len(nums)   // right = len(nums), not len-1
    // Invariant: answer is in [left, right)

    for left < right {            // Note: strict less than
        mid := left + (right-left)/2

        if condition(nums, mid, target) {
            right = mid           // Mid could be the answer, keep it
        } else {
            left = mid + 1        // Mid is definitely not the answer
        }
    }

    return left                   // left == right == answer position
}
```

> 🔑 When condition is true → `right = mid` (don't exclude mid, it might be the answer)
> 🔑 When condition is false → `left = mid + 1` (mid is definitely wrong, skip it)

---

### Template 3 — Rotated Sorted Array Search

```go
func searchRotated(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        }

        // Check which half is sorted
        if nums[left] <= nums[mid] {
            // Left half is sorted
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1   // Target is in sorted left half
            } else {
                left = mid + 1    // Target is in right half
            }
        } else {
            // Right half is sorted
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1    // Target is in sorted right half
            } else {
                right = mid - 1   // Target is in left half
            }
        }
    }
    return -1
}
```

---

## 7. Step-by-Step Dry Run

**Problem:** Find the first occurrence of `7` in `[1, 3, 5, 7, 7, 7, 10]`

We use Template 2 with condition: `nums[mid] >= target`.

```
Array:  [1, 3, 5, 7, 7, 7, 10]
Index:   0  1  2  3  4  5   6
Target = 7
left = 0, right = 7 (len of array)

--- Step 1 ---
mid = 0 + (7-0)/2 = 3
nums[3] = 7 >= 7 → condition TRUE → right = 3
State: left=0, right=3

--- Step 2 ---
mid = 0 + (3-0)/2 = 1
nums[1] = 3 >= 7 → condition FALSE → left = 2
State: left=2, right=3

--- Step 3 ---
mid = 2 + (3-2)/2 = 2
nums[2] = 5 >= 7 → condition FALSE → left = 3
State: left=3, right=3

--- Loop ends (left == right) ---
Return left = 3 ✅ (first occurrence of 7 is at index 3)
```

Notice how `right = mid` (not `mid - 1`) preserved index 3 as a candidate until we confirmed it.

---

## 8. Must-Know Problems

### Problem 1 — Find Minimum in Rotated Sorted Array (LC #153)

**Intuition:** In a rotated array, the minimum is the "pivot point" — where the array resets. At any midpoint, if `nums[mid] > nums[right]`, the minimum must be in the right half (because the left half is cleanly sorted and the pivot is somewhere to the right). Otherwise, the minimum is in the left half (including mid itself).

**Approach:** Compare `nums[mid]` with `nums[right]` to decide direction.

```go
func findMin(nums []int) int {
    left, right := 0, len(nums)-1

    for left < right {
        mid := left + (right-left)/2

        if nums[mid] > nums[right] {
            // Min is definitely in right half (past mid)
            left = mid + 1
        } else {
            // Mid could be the min, or min is in left half
            right = mid
        }
    }
    return nums[left]
}
```

---

### Problem 2 — Search in Rotated Sorted Array (LC #33)

**Intuition:** Even after rotation, **one of the two halves is always sorted**. Use that sorted half to check if the target lies in it. If yes, search there. If no, search the other half.

**Approach:** Use Template 3 from above. This is the flagship rotated array problem.

```go
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        }

        if nums[left] <= nums[mid] { // Left half is sorted
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else { // Right half is sorted
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    return -1
}
```

---

### Problem 3 — Koko Eating Bananas (LC #875)

**Intuition:** This is a "binary search on answer" problem. The question is: can Koko finish all bananas at speed `k` within `h` hours? For a given speed, it's easy to check (O(n)). And crucially: if speed `k` works, any higher speed also works — it's **monotonic**. So binary search on speed.

**Approach:** `lo = 1`, `hi = max(piles)`. Find the smallest speed where the feasibility check passes.

```go
func minEatingSpeed(piles []int, h int) int {
    lo, hi := 1, 0
    for _, p := range piles {
        if p > hi {
            hi = p // hi = max pile size
        }
    }

    for lo < hi {
        mid := lo + (hi-lo)/2

        if canFinish(piles, mid, h) {
            hi = mid // mid speed works, try smaller
        } else {
            lo = mid + 1 // mid speed too slow
        }
    }
    return lo
}

func canFinish(piles []int, speed, h int) bool {
    hours := 0
    for _, p := range piles {
        hours += (p + speed - 1) / speed // ceiling division
    }
    return hours <= h
}
```

> 💡 The ceiling division trick: `ceil(a/b) = (a + b - 1) / b` — memorize this.

---

## 9. Common Mistakes / Gotchas

**1. Integer overflow on mid calculation.** Never write `mid = (left + right) / 2`. Use `mid = left + (right-left)/2` to prevent overflow when left and right are large.

**2. Off-by-one: using `right = len-1` vs `right = len`.** For "find first true" problems, initialize `right = len(nums)` (not `len-1`), because the answer might be just past the array (e.g., insert position).

**3. Infinite loop when `left < right` but `right = mid-1` isn't used.** In Template 2, when condition is false you must do `left = mid + 1`, never `left = mid`. Otherwise when `left + 1 == right`, mid equals left and you loop forever.

**4. Wrong condition for rotated array.** Always compare `nums[mid]` with `nums[right]` (not `nums[left]`) to determine which side the minimum is on. Mixing these up produces wrong results on edge cases.

**5. Forgetting that `<=` vs `<` in the loop condition matters.** `left <= right` is for exact-match BS. `left < right` is for boundary-finding BS. Mixing them causes off-by-one on the result.

**6. Not checking the feasibility function carefully.** In BS-on-answer problems, if your `canDo()` function has a bug, the binary search will silently converge to a wrong answer. Always test your predicate separately.

**7. Assuming the array must be sorted.** Modified BS works on any search space with monotonic behavior — including implicit spaces like "possible speeds" or "possible days." Don't limit your thinking to sorted arrays.

---

## 10. Time & Space Complexity

| Operation                     | Time                 | Space | Why                                   |
| ----------------------------- | -------------------- | ----- | ------------------------------------- |
| Classic binary search         | O(log n)             | O(1)  | Halving n takes log₂(n) steps         |
| Find first/last occurrence    | O(log n)             | O(1)  | Same halving, different condition     |
| Search rotated array          | O(log n)             | O(1)  | One half always eliminated            |
| BS on answer (range [lo, hi]) | O(log(hi-lo) × f(n)) | O(1)  | log steps × cost of feasibility check |
| Find peak element             | O(log n)             | O(1)  | Slope direction halves space          |

> 💡 For BS-on-answer, if the feasibility function is O(n), total time is O(n log(range)). Range is usually 10⁸ or 10⁹, and log(10⁹) ≈ 30, so it's very fast in practice.

---

## 11. Practice Problems

### Easy

| #   | Problem                      | Hint                                                      |
| --- | ---------------------------- | --------------------------------------------------------- |
| 704 | Binary Search                | Direct template application — warm up here first          |
| 35  | Search Insert Position       | Find first index where `nums[i] >= target`                |
| 374 | Guess Number Higher or Lower | Classic BS with API as comparator                         |
| 69  | Sqrt(x)                      | BS on answer: find largest `mid` where `mid*mid <= x`     |
| 278 | First Bad Version            | Find first `true` in a boolean sequence — pure Template 2 |

### Medium

| #    | Problem                              | Hint                                                |
| ---- | ------------------------------------ | --------------------------------------------------- |
| 33   | Search in Rotated Sorted Array       | One half is always sorted — use that                |
| 81   | Search in Rotated Sorted Array II    | Same as #33 but handle duplicates with `left++`     |
| 153  | Find Minimum in Rotated Sorted Array | Compare `mid` with `right`, not `left`              |
| 162  | Find Peak Element                    | If `nums[mid] < nums[mid+1]`, peak is to the right  |
| 875  | Koko Eating Bananas                  | BS on speed, feasibility = can finish in h hours    |
| 1011 | Capacity to Ship Packages            | BS on weight capacity, feasibility = days needed    |
| 74   | Search a 2D Matrix                   | Treat 2D matrix as 1D sorted array via index math   |
| 34   | Find First and Last Position         | Run Template 2 twice: once for first, once for last |
| 540  | Single Element in Sorted Array       | Pairs break at the answer; use even/odd index logic |

### Hard

| #    | Problem                     | Hint                                                         |
| ---- | --------------------------- | ------------------------------------------------------------ |
| 4    | Median of Two Sorted Arrays | BS on partition point of smaller array                       |
| 410  | Split Array Largest Sum     | BS on answer (max subarray sum), check if `k` splits suffice |
| 1231 | Divide Chocolate            | BS on minimum sweetness, check if you can get `k+1` pieces   |
| 2064 | Maximized Grid Happiness    | BS-adjacent; practice monotonic feasibility thinking         |

---

## 12. Quick Revision (60 seconds)

### Key Idea

Binary search works on any search space that is **monotonic** — not just sorted arrays. The trick is finding the predicate that splits the space into "no" and "yes".

### Universal Template (Find First True)

```go
left, right := 0, len(nums)  // or [lo, hi] for BS on answer

for left < right {
    mid := left + (right-left)/2
    if condition(mid) {
        right = mid      // keep mid as candidate
    } else {
        left = mid + 1   // mid is out
    }
}
return left  // first index where condition is true
```

### When to Use

- Sorted array + find boundary → Template 2
- Rotated array → check which half is sorted, eliminate accordingly
- "Find minimum X such that Y is possible" → BS on answer with feasibility check
- Peak/valley → compare mid with neighbor, follow the slope

### The One Rule to Remember

> If you can write a true/false condition that **never goes T→F** across your search space, you can binary search on it.

---
