# 🎯 Two Pointers — FAANG Interview Notes

> **Mentor's Note:** Two Pointers is one of the most elegant patterns in DSA. Once you "see" it, you'll spot it everywhere. The core idea is simple: instead of brute-forcing with two nested loops (O(n²)), use two variables to scan from different positions — and reduce it to O(n).

---

## 1. Pattern Overview

### What is it?

Two Pointers is a technique where you use **two index variables** to traverse a data structure (usually an array or string) — either from both ends, from the same start, or across two different arrays.

### Why is it used?

Because brute force uses two nested loops → **O(n²)**.
Two pointers let both indices cover the same data **without redundant work** → **O(n)**.

> 💡 **Intuition:** Imagine reading a book — one finger from the start, one from the end. Together they cover the whole book in one pass.

### Real-World Analogy

Imagine two people walking toward each other on a street. They'll meet somewhere in the middle. That's the **opposite-direction** variant.
Or think of two runners starting at the same spot — one fast, one slow. The fast one laps the slow one. That's the **fast/slow pointer** variant.

### Time Complexity Intuition

- Brute force = two nested loops = O(n²)
- Two pointers = both pointers travel at most n steps total = **O(n)**
- You eliminate the inner loop by making the pointer movement _intelligent_

---

## 2. Core Idea (MOST IMPORTANT)

### Key Insight #1

> **Move pointers based on a condition, not blindly.**

When you check a condition (sum too big? move right pointer left. Sum too small? move left pointer right.), you're making a _guaranteed_ progress toward the answer.

### Key Insight #2

> **The array usually needs to be sorted** (for opposite-direction variant).

Sorting gives you a guarantee: moving left pointer right _increases_ the value, moving right pointer left _decreases_ it. This is what makes the logic deterministic.

### ⭐ Remember This

```
left, right := 0, len(arr)-1

for left < right {
    if condition_met {
        // record answer
    } else if need_bigger {
        left++
    } else {
        right--
    }
}
```

---

## 3. When to Use This Pattern

| Signal in Question                    | What It Means                           |
| ------------------------------------- | --------------------------------------- |
| "Find pair with sum = target"         | Classic opposite-direction two pointers |
| "Remove duplicates / in-place"        | Same-direction (slow/fast) two pointers |
| "Sorted array" mentioned              | Strong hint for two pointers            |
| "Subarray / substring with condition" | Sliding window (two pointers variant)   |
| "Palindrome check"                    | Opposite direction from both ends       |
| "Merge two sorted arrays"             | Two pointers on different arrays        |
| "Cycle detection in linked list"      | Fast/slow pointer (Floyd's algorithm)   |
| "Minimize/maximize window size"       | Sliding window (two pointer style)      |

---

## 4. Types / Variants

### Type 1: Opposite Direction (Left ↔ Right)

- Both pointers start at **opposite ends**
- Move toward each other
- **Use when:** Array is sorted, finding pairs, reversing, palindrome check

```
[1, 2, 3, 4, 5]
 L           R   → move inward based on condition
```

### Type 2: Same Direction (Slow → Fast)

- Both pointers start at the **same end**
- One moves fast, one moves slow
- **Use when:** Remove duplicates, partition, find middle, detect cycle

```
[1, 1, 2, 3, 3]
 S  F            → fast explores, slow writes valid answers
```

### Type 3: Two Arrays (Array1 & Array2)

- One pointer per array, both move forward
- **Use when:** Merge sorted arrays, intersection of arrays

```
arr1: [1, 3, 5]    i →
arr2: [2, 3, 6]    j →
```

### Type 4: Fast/Slow (Floyd's — Linked List)

- Fast moves 2 steps, slow moves 1
- **Use when:** Detect cycle, find middle of linked list

---

## 5. Data Structures Used

| Data Structure          | Why Used                                | Key Property            | Time Complexity |
| ----------------------- | --------------------------------------- | ----------------------- | --------------- |
| **Array (sorted)**      | Primary — enables directional guarantee | Random access O(1)      | Access: O(1)    |
| **String**              | Treated like array of chars             | Immutable in most langs | Access: O(1)    |
| **Linked List**         | Fast/slow pointer problems              | Sequential access only  | Access: O(n)    |
| **HashMap** (sometimes) | When sorting not allowed                | O(1) lookup             | O(1) avg        |

> 💡 **Key rule:** Arrays must usually be **sorted** for two pointers to work correctly. If not sorted, either sort first or use a HashMap instead.

---

## 6. Core Templates (VERY IMPORTANT)

### Template 1: Opposite Direction — Find Pair with Target Sum

```go
func twoSumSorted(arr []int, target int) (int, int) {
    left, right := 0, len(arr)-1

    for left < right {
        sum := arr[left] + arr[right]

        if sum == target {
            return left, right   // found the pair
        } else if sum < target {
            left++               // need bigger sum → move left forward
        } else {
            right--              // need smaller sum → move right back
        }
    }

    return -1, -1 // no pair found
}
```

---

### Template 2: Same Direction — Remove Duplicates In-Place

```go
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    slow := 0 // slow = last position of valid (non-duplicate) element

    for fast := 1; fast < len(nums); fast++ {
        if nums[fast] != nums[slow] { // found a new unique element
            slow++
            nums[slow] = nums[fast]  // write it to the next valid position
        }
        // if duplicate: fast keeps moving, slow stays
    }

    return slow + 1 // length of deduplicated array
}
```

---

### Template 3: Two Arrays — Merge Sorted Arrays

```go
func mergeSorted(a, b []int) []int {
    result := []int{}
    i, j := 0, 0

    for i < len(a) && j < len(b) {
        if a[i] <= b[j] {
            result = append(result, a[i])
            i++
        } else {
            result = append(result, b[j])
            j++
        }
    }

    // Append remaining elements (one of the arrays may have leftovers)
    result = append(result, a[i:]...)
    result = append(result, b[j:]...)

    return result
}
```

---

## 7. Step-by-Step Dry Run

**Problem:** Two Sum II — find indices of two numbers that add up to 17 in `[2, 7, 11, 15]`

```
Array:  [2,  7,  11,  15]
Index:   0   1    2    3

Step 1: left=0, right=3 → arr[0]+arr[3] = 2+15 = 17 ✅ → return [0,3]
```

Simple case. Let's try target = 18:

```
Array:  [2,  7,  11,  15]

Step 1: left=0, right=3 → 2+15=17 < 18 → need bigger → left++
Step 2: left=1, right=3 → 7+15=22 > 18 → need smaller → right--
Step 3: left=1, right=2 → 7+11=18 == 18 ✅ → return [1,2]
```

> 🔑 Notice: We never revisit any pair. Each step eliminates one element from consideration. This is why it's O(n).

---

## 8. Must-Know Problems (Deep Explanation)

---

### Problem 1: Two Sum II — Input Array Is Sorted (LC #167)

**Intuition:**
Array is sorted → two pointers from both ends. If sum is too small, increase left. If too big, decrease right.

**Approach:**

1. Start `left=0`, `right=n-1`
2. Check `arr[left] + arr[right]` vs target
3. Adjust pointer based on comparison

```go
func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1

    for left < right {
        sum := numbers[left] + numbers[right]
        if sum == target {
            return []int{left + 1, right + 1} // 1-indexed answer
        } else if sum < target {
            left++
        } else {
            right--
        }
    }

    return []int{}
}
```

**Time:** O(n) | **Space:** O(1)

---

### Problem 2: Remove Duplicates from Sorted Array (LC #26)

**Intuition:**
Use slow pointer to track where the next unique element should go. Fast pointer scans ahead. When fast finds a new value, copy it to slow's next position.

**Approach:**

1. `slow=0` is the "write head"
2. `fast` scans from 1 to end
3. If `nums[fast] != nums[slow]` → unique found → `slow++`, write

```go
func removeDuplicates(nums []int) int {
    slow := 0

    for fast := 1; fast < len(nums); fast++ {
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
    }

    return slow + 1
}
```

**Time:** O(n) | **Space:** O(1)

---

### Problem 3: Container With Most Water (LC #11)

**Intuition:**
Area = min(height[left], height[right]) × (right - left).
To maximize area: always move the pointer with the **shorter** height inward. Moving the taller one can only decrease or keep width the same without any gain guarantee.

**Approach:**

1. Start from both ends
2. Calculate area
3. Move the shorter height pointer inward
4. Track max area

```go
func maxArea(height []int) int {
    left, right := 0, len(height)-1
    maxWater := 0

    for left < right {
        // width = right - left, height = min of both sides
        h := min(height[left], height[right])
        area := h * (right - left)

        if area > maxWater {
            maxWater = area
        }

        // Move the shorter side — moving taller side can't help
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxWater
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

**Time:** O(n) | **Space:** O(1)

---

## 9. Common Mistakes / Gotchas

| #   | Mistake                                                                   | Fix                                                                                              |
| --- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------ |
| 1   | **Using `left <= right` instead of `left < right`**                       | For pair-finding, use `left < right` — when they meet, no valid pair remains                     |
| 2   | **Forgetting to sort the array first**                                    | Opposite-direction two pointers only works on sorted arrays                                      |
| 3   | **Moving both pointers when answer is found**                             | Only move one pointer (or return immediately); moving both can skip valid answers                |
| 4   | **Off-by-one in slow/fast template**                                      | `slow` starts at 0 (first element already valid). `fast` starts at 1                             |
| 5   | **Infinite loop**                                                         | Always ensure at least one pointer moves in every iteration                                      |
| 6   | **Using two pointers when array isn't sorted and sorting is not allowed** | Use HashMap instead (e.g., original Two Sum LC #1)                                               |
| 7   | **Confusing two pointers with sliding window**                            | Sliding window has a dynamic window size condition; two pointers is simpler directional movement |

---

## 10. Time & Space Complexity

| Operation                      | Time Complexity | Space Complexity    |
| ------------------------------ | --------------- | ------------------- |
| Opposite direction (find pair) | O(n)            | O(1)                |
| Same direction (remove dups)   | O(n)            | O(1)                |
| Two arrays merge               | O(n + m)        | O(n + m) for result |
| With sorting needed            | O(n log n)      | O(1)                |
| Cycle detection (linked list)  | O(n)            | O(1)                |

**WHY O(n)?**
Both pointers together travel at most `n` steps total (each pointer only moves forward, never backward). So the inner work is linear regardless of how many iterations the loop runs.

---

## 11. Practice Problems

### 🟢 Easy

| #   | Problem                             | LC # | Hint                                                  |
| --- | ----------------------------------- | ---- | ----------------------------------------------------- |
| 1   | Two Sum II                          | 167  | Sort + opposite pointers; move based on sum vs target |
| 2   | Valid Palindrome                    | 125  | Skip non-alphanumeric; compare chars from both ends   |
| 3   | Remove Duplicates from Sorted Array | 26   | Slow writes, fast explores                            |
| 4   | Move Zeroes                         | 283  | Slow = next non-zero position; fast scans ahead       |
| 5   | Merge Sorted Array                  | 88   | Two pointers from the END to avoid overwriting        |

### 🟡 Medium

| #   | Problem                          | LC # | Hint                                                                |
| --- | -------------------------------- | ---- | ------------------------------------------------------------------- |
| 1   | Container With Most Water        | 11   | Move shorter-height pointer inward each time                        |
| 2   | 3Sum                             | 15   | Fix one element, two-pointer on the rest; skip duplicates carefully |
| 3   | 3Sum Closest                     | 16   | Same as 3Sum but track closest sum                                  |
| 4   | Sort Colors (Dutch Flag)         | 75   | Three pointers: low, mid, high                                      |
| 5   | Remove Nth Node from End of List | 19   | Fast pointer leads by n steps, then both move together              |
| 6   | Squares of a Sorted Array        | 977  | Largest squares come from both ends; fill result from right         |
| 7   | Partition Labels                 | 763  | Track last occurrence; two pointers define partition range          |
| 8   | Boats to Save People             | 881  | Sort + greedily pair heaviest with lightest                         |

### 🔴 Hard

| #   | Problem                  | LC # | Hint                                                                    |
| --- | ------------------------ | ---- | ----------------------------------------------------------------------- |
| 1   | Trapping Rain Water      | 42   | Track max from left and right; water at i = min(maxL, maxR) - height[i] |
| 2   | 4Sum                     | 18   | Fix two, two-pointer on rest; careful duplicate skipping at all levels  |
| 3   | Minimum Window Substring | 76   | Sliding window (two pointer variant); expand right, shrink left         |

---

## 12. Quick Revision (60 sec Cheat Sheet)

```
🔑 KEY IDEA
─────────────────────────────────────────
Two pointers = eliminate the inner loop
Move pointers based on condition, not blindly
Sorted array → opposite direction
Duplicates/partition → same direction (slow/fast)

📐 TEMPLATE (Opposite Direction)
─────────────────────────────────────────
left, right := 0, len(arr)-1
for left < right {
    if condition_met  → record + move both (or return)
    if need_bigger    → left++
    else              → right--
}

📐 TEMPLATE (Same Direction)
─────────────────────────────────────────
slow := 0
for fast := 1; fast < len; fast++ {
    if valid_condition(nums[fast]) {
        slow++
        nums[slow] = nums[fast]
    }
}

🚦 WHEN TO USE
─────────────────────────────────────────
✅ Sorted array + find pair/triplet
✅ In-place modification (remove dups, partition)
✅ Palindrome check
✅ Merge two sorted arrays
✅ Cycle detection (fast/slow)
❌ Unsorted + can't sort → use HashMap

⚠️  TOP GOTCHAS
─────────────────────────────────────────
• Use left < right (not <=) for pair finding
• Always sort first (opposite direction)
• Skip duplicates in 3Sum/4Sum
• Ensure pointer always moves → no infinite loop
```

---

_Happy Coding! Remember: the goal is to understand the intuition, not memorize the code. Once you see WHY two pointers work, you can reconstruct any template from scratch._ 🚀
