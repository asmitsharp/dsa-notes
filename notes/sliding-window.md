# DSA Pattern: Sliding Window

> **Interview Readiness Level:** ⭐⭐⭐⭐
> Extremely common in FAANG interviews for **strings and subarrays**.

---

# 1. Core Idea (Most Important)

A **Sliding Window** is a **contiguous section of an array/string** that moves step-by-step.

Instead of recomputing values for every subarray (**O(n²)**), we reuse previous work and update the window in **O(1)**.

### Example

Array

```
[2,1,5,1,3,2]
```

Window size = 3

```
[2,1,5] -> sum = 8
   [1,5,1] -> sum = 7
      [5,1,3] -> sum = 9
         [1,3,2] -> sum = 6
```

Instead of recalculating each window, we do:

```
newSum = oldSum - elementLeaving + elementEntering
```

---

# 2. When To Use Sliding Window

Look for these keywords in problems:

| Signal                           | Meaning             |
| -------------------------------- | ------------------- |
| **subarray / substring**         | contiguous elements |
| **longest / smallest**           | window expansion    |
| **at most K / exactly K**        | window constraint   |
| **maximum sum of size K**        | fixed window        |
| **without repeating characters** | dynamic window      |

---

# 3. Two Types of Sliding Window

## Type 1 — Fixed Size Window

Window size **never changes**

Example problems:

- Maximum sum subarray of size K
- Average of subarray size K

### Algorithm

1. Compute first window
2. Slide window
3. Remove left element
4. Add new right element

---

### Example

**Maximum sum of subarray of size k**

```
arr = [2,1,5,1,3,2]
k = 3
```

Windows:

```
[2,1,5] sum=8
[1,5,1] sum=7
[5,1,3] sum=9
[1,3,2] sum=6
```

Answer = **9**

---

### Go Implementation

```go
func maxSumSubarray(nums []int, k int) int {

    windowSum := 0
    maxSum := 0

    for i := 0; i < k; i++ {
        windowSum += nums[i]
    }

    maxSum = windowSum

    for i := k; i < len(nums); i++ {

        windowSum += nums[i]
        windowSum -= nums[i-k]

        if windowSum > maxSum {
            maxSum = windowSum
        }
    }

    return maxSum
}
```

---

# 4. Dynamic Sliding Window

Window size **changes based on condition**

We use **two pointers**

```
left
right
```

Window grows and shrinks.

---

### Visualization

```
arr = [2,1,5,2,3]

left right

[2]
[2,1]
[2,1,5]
   shrink
[1,5]
```

---

# 5. Generic Sliding Window Template

This template solves **80% sliding window problems**.

```go
left := 0

for right := 0; right < len(nums); right++ {

    // expand window
    add nums[right]

    for window invalid {

        remove nums[left]
        left++
    }

    update answer
}
```

---

# 6. Example Problem

## Longest Substring Without Repeating Characters

LeetCode **#3**

```
s = "abcabcbb"
```

Longest substring

```
abc
```

Length = **3**

---

### Idea

Use **map to track characters in window**

If duplicate appears → shrink window.

---

### Go Implementation

```go
func lengthOfLongestSubstring(s string) int {

    charMap := map[byte]int{}
    left := 0
    maxLen := 0

    for right := 0; right < len(s); right++ {

        charMap[s[right]]++

        for charMap[s[right]] > 1 {

            charMap[s[left]]--
            left++
        }

        windowLen := right - left + 1

        if windowLen > maxLen {
            maxLen = windowLen
        }
    }

    return maxLen
}
```

---

# 7. Another Classic Problem

## Minimum Window Substring

LeetCode **#76**

Find smallest substring containing all characters.

Example

```
s = "ADOBECODEBANC"
t = "ABC"
```

Answer

```
BANC
```

---

### Core Idea

1. Expand window
2. When condition satisfied
3. Shrink window to minimize size

---

# 8. Most Common Sliding Window Problems

These **10 problems cover almost everything**.

### Easy

1️⃣ LC 643 — Maximum Average Subarray
2️⃣ LC 1343 — Subarray of Size K
3️⃣ LC 121 — Best Time to Buy Stock

---

### Medium

4️⃣ LC 3 — Longest Substring Without Repeating
5️⃣ LC 1004 — Max Consecutive Ones III
6️⃣ LC 424 — Longest Repeating Character Replacement
7️⃣ LC 209 — Minimum Size Subarray Sum

---

### Hard

8️⃣ LC 76 — Minimum Window Substring
9️⃣ LC 30 — Substring with Concatenation
10️⃣ LC 992 — Subarrays with K Distinct

---

# 9. Complexity

| Operation | Complexity              |
| --------- | ----------------------- |
| Time      | **O(n)**                |
| Space     | **O(k)** (hashmap size) |

Why?

Each element enters and leaves window **at most once**.

---

# 10. Common Interview Mistakes

### Mistake 1

Forgetting to shrink window.

```
while condition invalid
```

---

### Mistake 2

Using `if` instead of `while`

Wrong

```
if window invalid
```

Correct

```
while window invalid
```

---

### Mistake 3

Wrong window size formula

Correct

```
length = right - left + 1
```

---

# 11. Sliding Window Cheat Sheet

### Fixed Window

```
windowSum += nums[i]
windowSum -= nums[i-k]
```

---

### Dynamic Window

```
expand right
while invalid
    shrink left
```

---

# 12. 60 Second Revision

Sliding window is used when:

```
contiguous subarray or substring
```

Two types:

```
Fixed window
Dynamic window
```

Template:

```
for right
    expand window

    while invalid
        shrink window

    update answer
```

Time Complexity:

```
O(n)
```

---

# Best Way for You to Learn This Pattern

Do **only these 4 problems**:

```
LC 643
LC 3
LC 209
LC 76
```

That will cover **90% sliding window questions**.

---
