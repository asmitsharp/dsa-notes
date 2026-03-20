# 🔄 Cyclic Sort (Index-Based)

> **DSA Pattern Notes** — Interview-Focused | FAANG Level

---

## 1. Pattern Overview

### What is this pattern?

Cyclic Sort is a technique used to sort arrays that contain numbers in a **known range** (typically `1 to N` or `0 to N-1`) by placing each number at its **correct index** directly — without using extra space.

### Why is it used? (Intuition)

If you're given an array of numbers from `1` to `N`, the number `i` **belongs at index `i-1`**.
Instead of using sorting algorithms (O(n log n)) or hash sets (O(n) space), you can exploit this structure to sort **in-place in O(n)** time.

> 💡 **Key Insight:** Every number has a "home index." If a number isn't home, swap it there.

### Real-World Analogy

Imagine 5 people standing in a line. Each person has a number on their jersey (1–5), and each numbered spot on the floor is their assigned place.

- You walk down the line.
- If the person isn't standing on their number, swap them with whoever **is** standing on their spot.
- Repeat until everyone is in the right place.

That's exactly Cyclic Sort.

### Time Complexity Intuition

- We traverse the array once: **O(n)**
- Each swap puts at least one element in its correct place
- An element is swapped at most once → **total swaps ≤ n**
- Even though there's a `while` loop inside a `for` loop, the **total iterations across all iterations is still O(n)**

---

## 2. Core Idea (MOST IMPORTANT)

### The 2 Key Insights

**Insight 1: Each number has a correct index**

```
For range [1, N]:  correct index of num = num - 1
For range [0, N]:  correct index of num = num
```

**Insight 2: If a number is NOT at its correct index → swap it there**
Keep swapping until either:

- The current number IS at its correct index, OR
- The number at the target index is a **duplicate** (don't swap, move on)

### The Core Loop (Remember This!)

```go
i := 0
for i < len(nums) {
    correctIdx := nums[i] - 1          // where this number SHOULD be
    if nums[i] != nums[correctIdx] {   // ← NOT a duplicate
        nums[i], nums[correctIdx] = nums[correctIdx], nums[i]  // swap to correct position
    } else {
        i++  // already in place or duplicate → move forward
    }
}
```

> ⚠️ **Critical:** Only increment `i` when the current element is in its correct position. Otherwise, keep swapping.

---

## 3. When to Use This Pattern

| Signal in Question                                 | What It Means                     |
| -------------------------------------------------- | --------------------------------- |
| Array of size `n` with numbers in range `[1, n]`   | Classic cyclic sort setup         |
| "Find the missing number"                          | Sort then scan for gap            |
| "Find the duplicate number"                        | Sort then scan for repeated value |
| "Find all missing/duplicate numbers"               | Generalized cyclic sort scan      |
| Array with numbers in range `[0, n]`               | Slight variant (correctIdx = num) |
| "Without extra space" or "O(1) space"              | Strong hint toward cyclic sort    |
| Positive integers, no specific sorted order needed | Cyclic sort likely applicable     |

---

## 4. Types / Variants

### Type 1: Basic Cyclic Sort — `[1, N]` range

Place every number at `index = num - 1`.
Used when: No missing/duplicate numbers, just sort in-place.

### Type 2: Find Missing Number(s)

After cyclic sort, scan: if `nums[i] != i + 1`, then `i + 1` is missing.

### Type 3: Find Duplicate Number(s)

After cyclic sort, scan: if `nums[i] != i + 1`, then `nums[i]` is the duplicate.

### Type 4: Find Missing AND Duplicate

Combine both scans — one pass after sorting is enough.

---

## 5. Data Structures Used

| Data Structure         | Used?      | Why                                                       |
| ---------------------- | ---------- | --------------------------------------------------------- |
| **Array (in-place)**   | ✅ Primary | The array itself acts as a "hash map" where index = value |
| **HashSet**            | ❌ Avoid   | Cyclic sort replaces the need — saves O(n) space          |
| **Sorting (built-in)** | ❌ Avoid   | O(n log n), cyclic sort does it in O(n)                   |

**Key Property:** The array is used as an index-to-value mapping. This is only possible when the value range is **known and bounded**.

---

## 6. Core Templates (VERY IMPORTANT)

### Template 1: Basic Cyclic Sort `[1, N]`

```go
func cyclicSort(nums []int) []int {
    i := 0
    for i < len(nums) {
        correctIdx := nums[i] - 1              // nums[i] belongs at this index
        if nums[i] != nums[correctIdx] {       // not a duplicate, so safe to swap
            nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
        } else {
            i++                                // in place or duplicate, move on
        }
    }
    return nums
}
```

---

### Template 2: Find All Missing Numbers

```go
func findMissingNumbers(nums []int) []int {
    // Step 1: Cyclic sort
    i := 0
    for i < len(nums) {
        correctIdx := nums[i] - 1
        if nums[i] != nums[correctIdx] {
            nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
        } else {
            i++
        }
    }

    // Step 2: Find mismatches
    missing := []int{}
    for i, num := range nums {
        if num != i+1 {
            missing = append(missing, i+1)   // i+1 is missing
        }
    }
    return missing
}
```

---

### Template 3: Find All Duplicates

```go
func findDuplicates(nums []int) []int {
    // Step 1: Cyclic sort
    i := 0
    for i < len(nums) {
        correctIdx := nums[i] - 1
        if nums[i] != nums[correctIdx] {
            nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
        } else {
            i++
        }
    }

    // Step 2: Find mismatches → the number there is a duplicate
    duplicates := []int{}
    for i, num := range nums {
        if num != i+1 {
            duplicates = append(duplicates, num)  // num is the duplicate
        }
    }
    return duplicates
}
```

---

## 7. Step-by-Step Dry Run

**Input:** `[3, 1, 5, 4, 2]` — range `[1, 5]`

| Step | i   | Array       | nums[i] | correctIdx | Action                          |
| ---- | --- | ----------- | ------- | ---------- | ------------------------------- |
| 1    | 0   | [3,1,5,4,2] | 3       | 2          | swap(0,2) → [5,1,3,4,2]         |
| 2    | 0   | [5,1,3,4,2] | 5       | 4          | swap(0,4) → [2,1,3,4,5]         |
| 3    | 0   | [2,1,3,4,5] | 2       | 1          | swap(0,1) → [1,2,3,4,5]         |
| 4    | 0   | [1,2,3,4,5] | 1       | 0          | nums[i]==nums[correctIdx] → i++ |
| 5    | 1   | [1,2,3,4,5] | 2       | 1          | nums[i]==nums[correctIdx] → i++ |
| ...  | ... | ...         | ...     | ...        | rest already in place           |

**Output:** `[1, 2, 3, 4, 5]` ✅

> 🔑 Notice: At step 4, `nums[0] = 1` and `nums[correctIdx=0] = 1` — same! So we move on.

---

## 8. Must-Know Problems (Deep Explanation)

### Problem 1: Missing Number — LC #268 (Easy)

**Intuition:** Array has `n` numbers in range `[0, n]`, one is missing. Place each number at its index. Then scan for `nums[i] != i`.

**Approach:**

1. Cyclic sort (handle `nums[i] == n` separately — no home in array)
2. Scan: first index where `nums[i] != i` → that `i` is missing
3. If all match, missing number is `n`

```go
func missingNumber(nums []int) int {
    i := 0
    n := len(nums)
    for i < n {
        correctIdx := nums[i]                        // range is [0, n], so correctIdx = nums[i]
        if nums[i] < n && nums[i] != nums[correctIdx] {
            nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
        } else {
            i++
        }
    }
    // Scan for the missing number
    for i := 0; i < n; i++ {
        if nums[i] != i {
            return i
        }
    }
    return n  // all [0, n-1] present, so n is missing
}
```

---

### Problem 2: Find the Duplicate Number — LC #287 (Medium)

**Intuition:** Array has `n+1` numbers in range `[1, n]`, exactly one duplicate. Cyclic sort: when you try to place a number at its correct index and find the same number already there → that's the duplicate.

**Approach:** During cyclic sort, when `nums[i] == nums[correctIdx]` AND `i != correctIdx` → duplicate found.

```go
func findDuplicate(nums []int) int {
    i := 0
    for i < len(nums) {
        correctIdx := nums[i] - 1
        if nums[i] != i+1 {                           // not in correct place
            if nums[i] != nums[correctIdx] {           // no duplicate yet, swap
                nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
            } else {
                return nums[i]                         // duplicate found!
            }
        } else {
            i++
        }
    }
    return -1
}
```

---

### Problem 3: Find All Duplicates in an Array — LC #442 (Medium)

**Intuition:** Array has `n` numbers in `[1, n]`, some appear twice. Sort cyclically. After sorting, every index where `nums[i] != i+1` means `nums[i]` appeared twice (it "stole" the spot from `i+1`).

```go
func findAllDuplicates(nums []int) []int {
    i := 0
    for i < len(nums) {
        correctIdx := nums[i] - 1
        if nums[i] != nums[correctIdx] {
            nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
        } else {
            i++
        }
    }

    result := []int{}
    for i, num := range nums {
        if num != i+1 {
            result = append(result, num)
        }
    }
    return result
}
```

---

## 9. Common Mistakes / Gotchas

| #   | Mistake                                             | Fix                                                                          |
| --- | --------------------------------------------------- | ---------------------------------------------------------------------------- |
| 1   | **Using `correctIdx = nums[i]` for `[1,N]` range**  | For `[1,N]`: `correctIdx = nums[i] - 1`. For `[0,N]`: `correctIdx = nums[i]` |
| 2   | **Infinite loop: swapping duplicates forever**      | Always check `nums[i] != nums[correctIdx]` before swapping                   |
| 3   | **Incrementing `i` after every swap**               | Only increment `i` when element is in correct place                          |
| 4   | **Not handling `nums[i] >= n` for `[0,N]` range**   | Numbers equal to `n` have no home — skip them with `i++`                     |
| 5   | **Confusing missing vs duplicate scan**             | Missing: return `i+1`. Duplicate: return `nums[i]`                           |
| 6   | **Applying to arrays with unknown/unbounded range** | Cyclic sort only works when range is known and bounded                       |
| 7   | **Forgetting the `i != correctIdx` check**          | When `i == correctIdx`, element is already home — don't swap with itself     |

---

## 10. Time & Space Complexity

| Operation                   | Time | Space | Why                               |
| --------------------------- | ---- | ----- | --------------------------------- |
| Cyclic Sort                 | O(n) | O(1)  | Each element swapped at most once |
| Find Missing                | O(n) | O(1)  | One pass after sort               |
| Find Duplicate              | O(n) | O(1)  | One pass after sort               |
| Find All Missing/Duplicates | O(n) | O(1)  | One scan after O(n) sort          |

> 💡 The `while` loop inside `for` looks like O(n²) but it's actually O(n) because each swap places one element in its final position — so total swaps ≤ n.

---

## 11. Practice Problems

### 🟢 Easy

| #   | Problem                                  | LC # | Hint                                                       |
| --- | ---------------------------------------- | ---- | ---------------------------------------------------------- |
| 1   | Missing Number                           | 268  | Range `[0,n]`: place at `nums[i]` index, scan for mismatch |
| 2   | Find All Numbers Disappeared in an Array | 448  | Cyclic sort `[1,n]`, scan where `nums[i] != i+1`           |
| 3   | Single Number                            | 136  | Not cyclic sort — XOR trick (good contrast problem)        |

---

### 🟡 Medium

| #   | Problem                         | LC # | Hint                                                          |
| --- | ------------------------------- | ---- | ------------------------------------------------------------- |
| 4   | Find the Duplicate Number       | 287  | Spot when `nums[i] == nums[correctIdx]` during sort           |
| 5   | Find All Duplicates in an Array | 442  | After sort, mismatch index reveals duplicate                  |
| 6   | Set Mismatch                    | 645  | One missing + one duplicate — solve both in one scan          |
| 7   | First Missing Positive          | 41   | Filter out-of-range, cyclic sort `[1,n]`, first gap is answer |
| 8   | Couples Holding Hands           | 765  | Cyclic groups in permutation — advanced variant               |
| 9   | Array of Doubled Pairs          | 954  | Grouping variant — understand cyclic logic                    |

---

### 🔴 Hard

| #   | Problem                            | LC # | Hint                                                 |
| --- | ---------------------------------- | ---- | ---------------------------------------------------- |
| 10  | Find the Smallest Missing Positive | 41   | Ignore negatives/out-of-range, cyclic sort remaining |
| 11  | Missing Ranges                     | 163  | Not cyclic — but builds on missing-number intuition  |
| 12  | First Missing Positive (follow-up) | 41   | Handle negatives + zeros + duplicates in one pass    |

---

## 12. Quick Revision (60 sec) ⚡

```
CYCLIC SORT CHEAT SHEET
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

KEY IDEA:
  Numbers in range [1,N] → num belongs at index (num-1)
  Place each number at its correct index by swapping

WHEN TO USE:
  ✅ Range is [1,N] or [0,N]
  ✅ "Find missing/duplicate" in O(1) space
  ✅ Array size = range size (or off by 1)

CORE TEMPLATE:
  i := 0
  for i < len(nums) {
      j := nums[i] - 1                       // correct index
      if nums[i] != nums[j] {                // not duplicate → swap
          nums[i], nums[j] = nums[j], nums[i]
      } else {
          i++                                // in place or dup → advance
      }
  }

AFTER SORT:
  nums[i] != i+1 → missing = i+1, duplicate = nums[i]

KEY RULES:
  • Only increment i when element is in correct place
  • Check nums[i] != nums[correctIdx] to avoid infinite loop on duplicates
  • For range [0,N]: correctIdx = nums[i] (not nums[i]-1)
  • For out-of-range numbers → just skip (i++)
```

---

_Happy Coding! 🚀 — Master the pattern, not the problems._
