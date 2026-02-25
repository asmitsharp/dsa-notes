# DSA Pattern: Prefix Sum / Prefix Product

> **Interview Readiness Level:** ⭐⭐⭐⭐⭐ — High frequency, appears in FAANG interviews regularly across Easy→Hard spectrum.

---

## SECTION 1 — Quick Reference Card

### Core Insight

```
prefix[i] = prefix[i-1] + arr[i]        ← cumulative sum up to index i
rangeSum(l, r) = prefix[r] - prefix[l-1] ← O(1) range query after O(n) build
```

- Store cumulative aggregates so any subarray aggregate is a **single subtraction** away.
- Build once in **O(n)**, query infinitely in **O(1)** per query — the fundamental tradeoff.
- Extends beyond sums: products, XOR, GCD, max — anything with a compatible **inverse operation**.
- The "extra cell" trick (`prefix[0] = 0`, array is 1-indexed in prefix) eliminates all boundary checks.

---

### When to Use This Pattern

> **Trigger Keywords:** subarray sum, range query, contiguous subarray, cumulative, between indices, multiple queries on same array, count subarrays with property, equilibrium index, pivot index

| Signal in Problem                    | Technique                       |
| ------------------------------------ | ------------------------------- |
| "sum of elements from index l to r"  | Basic prefix sum                |
| "count subarrays with sum = k"       | Prefix sum + HashMap            |
| "subarray with sum divisible by k"   | Prefix sum + modular arithmetic |
| "product of subarray excluding self" | Prefix + suffix product         |
| "2D grid range sum"                  | 2D prefix sum                   |
| "multiple range update queries"      | Difference array                |
| "running XOR / running AND"          | Prefix XOR / Prefix AND         |

---

## SECTION 2 — Fundamentals

### What Is This Pattern?

A **Prefix Sum** array `prefix[]` is a derived array where each element stores the cumulative sum of the original array from index `0` to that index. The key property is that any contiguous subarray sum can be computed in **O(1)** by subtracting two prefix values, after an **O(n)** precomputation.

**From first principles:**

Given `arr = [a₀, a₁, a₂, a₃, a₄]`

```
prefix[0] = 0            (sentinel / guard value)
prefix[1] = a₀
prefix[2] = a₀ + a₁
prefix[3] = a₀ + a₁ + a₂
prefix[4] = a₀ + a₁ + a₂ + a₃
prefix[5] = a₀ + a₁ + a₂ + a₃ + a₄
```

**Range query formula (0-indexed arr, 1-indexed prefix):**

```
sum(l, r) = prefix[r+1] - prefix[l]
```

**Range query formula (both 0-indexed):**

```
sum(l, r) = prefix[r] - (l > 0 ? prefix[l-1] : 0)
```

---

### ASCII Visualization

```
Original Array (0-indexed):
Index:  0    1    2    3    4
        [3]  [1]  [4]  [1]  [5]

Prefix Array (1-indexed, prefix[0]=0 sentinel):
Index:  0    1    2    3    4    5
        [0]  [3]  [4]  [8]  [9]  [14]
         ^                         ^
      sentinel               total sum

Query: sum(1, 3) = arr[1]+arr[2]+arr[3] = 1+4+1 = 6
Via prefix: prefix[4] - prefix[1] = 9 - 3 = 6  ✓

Query: sum(0, 4) = prefix[5] - prefix[0] = 14 - 0 = 14  ✓
```

---

### Properties, Rules, and Constraints

1. **Monotonicity:** If all elements ≥ 0, prefix array is non-decreasing.
2. **Inverse operation required:** For prefix XOR, the inverse is XOR itself. For prefix product, inverse is division (watch for zeros). For prefix sum, inverse is subtraction.
3. **Off-by-one is the #1 bug** — the sentinel at index 0 is your best friend.
4. **Integer overflow:** With large arrays of large values, use `int64` in Go.
5. **Prefix array has length n+1** (one more than original array).

---

### Important Variants

| Variant                | Description                              | Use Case                                 |
| ---------------------- | ---------------------------------------- | ---------------------------------------- |
| **1D Prefix Sum**      | Standard cumulative sum                  | Range sum queries                        |
| **1D Prefix Product**  | Cumulative product                       | Product of subarray                      |
| **Prefix XOR**         | Cumulative XOR                           | XOR of subarray range                    |
| **2D Prefix Sum**      | 2D grid cumulative sum                   | Rectangle sum queries                    |
| **Prefix + HashMap**   | Store prefix sums in map for O(1) lookup | Count subarrays with target sum          |
| **Difference Array**   | Inverse of prefix sum                    | Range update queries                     |
| **Suffix Sum/Product** | Right-to-left cumulative                 | Product except self, right contributions |

---

## SECTION 3 — Data Structures

### 3.1 — Slice (Go `[]int` / `[]int64`)

**What it is:** The primary data structure — a dynamic array holding prefix values.

**Why used here:** O(1) random access by index is essential for both building (sequential write) and querying (random read).

**Key operations:**
| Operation | Time | Space |
|---|---|---|
| Build prefix array | O(n) | O(n) |
| Point query prefix[i] | O(1) | O(1) |
| Range query sum(l, r) | O(1) | O(1) |

```go
// Building a prefix sum slice
func buildPrefix(arr []int) []int {
    n := len(arr)
    prefix := make([]int, n+1) // n+1 size, prefix[0] = 0 sentinel
    for i := 0; i < n; i++ {
        prefix[i+1] = prefix[i] + arr[i]
    }
    return prefix
}

// Range query
func rangeSum(prefix []int, l, r int) int {
    // l and r are 0-indexed positions in original array
    return prefix[r+1] - prefix[l]
}
```

---

### 3.2 — HashMap (`map[int]int`)

**What it is:** A hash map storing `prefixSum → frequency` or `prefixSum → first_index`.

**Why used here:** When we need to find how many previous prefix sums satisfy `prefix[j] - prefix[i] = target`, we need O(1) lookup of `prefix[j] - target`.

**When to choose map over slice:**

> Use a **map** when prefix values can be large/negative (can't use as array index), or when you need frequency counts of prefix sums. Use a **slice** when values are bounded and non-negative (e.g., 2D grid problems).

| Operation                  | Time     | Space |
| -------------------------- | -------- | ----- |
| Insert/lookup prefix sum   | O(1) avg | O(n)  |
| Count subarrays with sum=k | O(n)     | O(n)  |

```go
// Pattern: prefix sum + hashmap for count of subarrays with sum = k
func subarraySum(nums []int, k int) int {
    count := 0
    prefixSum := 0
    freq := map[int]int{0: 1} // CRITICAL: empty prefix (sum=0) seen once

    for _, num := range nums {
        prefixSum += num
        // If (prefixSum - k) exists in map, those many subarrays end here
        count += freq[prefixSum-k]
        freq[prefixSum]++
    }
    return count
}
```

---

### 3.3 — 2D Slice (`[][]int`)

**What it is:** A 2D array of size `(m+1) × (n+1)` for 2D prefix sums.

**Why used here:** Enables O(1) rectangle sum queries on a 2D grid after O(m×n) build.

**Inclusion-Exclusion Formula (the key insight):**

```
rect(r1,c1,r2,c2) = P[r2+1][c2+1] - P[r1][c2+1] - P[r2+1][c1] + P[r1][c1]
```

```
ASCII: Visualizing 2D prefix inclusion-exclusion

     c1      c2
r1 [ . . . | . . ]
   [ . . . | . . ]
   [--------+----]
r2 [ . . . | X  ]   ← we want sum of X region

P[r2+1][c2+1]           = whole top-left rectangle (A+B+C+X)
- P[r1][c2+1]           = remove top strip (A+B)
- P[r2+1][c1]           = remove left strip (A+C)
+ P[r1][c1]             = add back double-subtracted corner (A)
= X  ✓
```

---

### 3.4 — Difference Array

**What it is:** An array `diff[]` where range updates become O(1) point updates, and the final array is recovered via prefix sum.

**When to use:** Multiple range update queries `add(l, r, val)` followed by reading final array. The inverse of prefix sum.

```go
// diff[l] += val, diff[r+1] -= val → recover by prefix sum
func buildDiffArray(n int) []int {
    return make([]int, n+1) // n+1 to handle diff[r+1] when r = n-1
}

func rangeUpdate(diff []int, l, r, val int) {
    diff[l] += val
    if r+1 < len(diff) {
        diff[r+1] -= val
    }
}

func recoverArray(diff []int) []int {
    result := make([]int, len(diff))
    result[0] = diff[0]
    for i := 1; i < len(diff); i++ {
        result[i] = result[i-1] + diff[i]
    }
    return result
}
```

---

## SECTION 4 — Core Algorithms

### Algorithm 1: Basic Prefix Sum + Range Query

**LeetCode:** #303 — Range Sum Query - Immutable

**Steps:**

1. Build prefix array of size n+1 with sentinel `prefix[0] = 0`
2. Fill: `prefix[i+1] = prefix[i] + arr[i]`
3. Answer `sum(l, r) = prefix[r+1] - prefix[l]`

```go
// LC 303 - Range Sum Query Immutable
type NumArray struct {
    prefix []int
}

func Constructor(nums []int) NumArray {
    prefix := make([]int, len(nums)+1)
    for i, v := range nums {
        prefix[i+1] = prefix[i] + v
    }
    return NumArray{prefix: prefix}
}

func (na *NumArray) SumRange(left, right int) int {
    return na.prefix[right+1] - na.prefix[left]
}

// Dry run: nums = [−2, 0, 3, −5, 2, −1]
// prefix = [0, −2, −2, 1, −4, −2, −3]
// SumRange(0,2) = prefix[3] - prefix[0] = 1 - 0 = 1 ✓ (−2+0+3=1)
// SumRange(2,5) = prefix[6] - prefix[2] = −3 − (−2) = −1 ✓ (3−5+2−1=−1)
```

---

### Algorithm 2: Count Subarrays with Sum = K

**LeetCode:** #560 — Subarray Sum Equals K

**Key insight:** We need count of pairs `(i, j)` where `prefix[j] - prefix[i] = k`, i.e., `prefix[i] = prefix[j] - k`. As we scan left to right, look up `currentPrefix - k` in our map.

```go
// LC 560 - Subarray Sum Equals K
func subarraySum(nums []int, k int) int {
    count := 0
    prefixSum := 0
    freq := map[int]int{0: 1} // empty subarray has sum 0

    for _, num := range nums {
        prefixSum += num
        // How many previous indices had prefixSum - k?
        count += freq[prefixSum-k]
        freq[prefixSum]++
    }
    return count
}

// Dry run: nums = [1, 1, 1], k = 2
// Step 0: prefixSum=0, freq={0:1}
// Step 1: num=1, prefixSum=1, lookup freq[1-2]=freq[-1]=0, count=0, freq={0:1,1:1}
// Step 2: num=1, prefixSum=2, lookup freq[2-2]=freq[0]=1, count=1, freq={0:1,1:1,2:1}
// Step 3: num=1, prefixSum=3, lookup freq[3-2]=freq[1]=1, count=2, freq={...,3:1}
// Result: 2 ✓ ([1,1] at idx 0-1 and [1,1] at idx 1-2)
```

---

### Algorithm 3: Product of Array Except Self

**LeetCode:** #238 — Product of Array Except Self

**Key insight:** Use prefix product from the left AND suffix product from the right. `result[i] = leftProduct[i] * rightProduct[i]`.

```go
// LC 238 - Product of Array Except Self (O(1) extra space)
func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)

    // Pass 1: fill result with prefix products (product of all elements LEFT of i)
    result[0] = 1
    for i := 1; i < n; i++ {
        result[i] = result[i-1] * nums[i-1]
    }

    // Pass 2: multiply by suffix products (product of all elements RIGHT of i)
    suffix := 1
    for i := n - 1; i >= 0; i-- {
        result[i] *= suffix
        suffix *= nums[i]
    }

    return result
}

// Dry run: nums = [1, 2, 3, 4]
// After Pass 1 (prefix):  result = [1, 1, 2, 6]
//   result[0]=1, result[1]=1*1=1, result[2]=1*2=2, result[3]=2*3=6
// Pass 2 (suffix), suffix starts=1:
//   i=3: result[3]=6*1=6,  suffix=1*4=4
//   i=2: result[2]=2*4=8,  suffix=4*3=12
//   i=1: result[1]=1*12=12, suffix=12*2=24
//   i=0: result[0]=1*24=24, suffix=24*1=24
// result = [24, 12, 8, 6]  ✓
```

---

### Algorithm 4: 2D Prefix Sum (Matrix Region Sum)

**LeetCode:** #304 — Range Sum Query 2D - Immutable

```go
// LC 304 - Range Sum Query 2D Immutable
type NumMatrix struct {
    prefix [][]int
}

func Constructor2D(matrix [][]int) NumMatrix {
    m, n := len(matrix), len(matrix[0])
    // (m+1) x (n+1) with zero-row and zero-column padding
    prefix := make([][]int, m+1)
    for i := range prefix {
        prefix[i] = make([]int, n+1)
    }

    for r := 1; r <= m; r++ {
        for c := 1; c <= n; c++ {
            prefix[r][c] = matrix[r-1][c-1] +
                prefix[r-1][c] +    // top
                prefix[r][c-1] -    // left
                prefix[r-1][c-1]    // top-left (added twice, remove once)
        }
    }
    return NumMatrix{prefix: prefix}
}

func (nm *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
    // Convert to 1-indexed prefix coordinates
    return nm.prefix[row2+1][col2+1] -
        nm.prefix[row1][col2+1] -
        nm.prefix[row2+1][col1] +
        nm.prefix[row1][col1]
}

// Dry run: matrix = [[3,0,1,4],[5,6,3,2],[1,2,0,1],[4,1,0,1]]
// prefix[2][3] = 3+0+1+5+6+3 = 18  (top-left 2x3 region sum)
// SumRegion(2,1,4,3) uses inclusion-exclusion
```

---

### Algorithm 5: Subarray Sum Divisible by K

**LeetCode:** #974 — Subarray Sums Divisible by K

**Key insight:** `prefix[j] - prefix[i]` divisible by K iff `prefix[j] % K == prefix[i] % K`. Handle negative mods carefully.

```go
// LC 974 - Subarray Sums Divisible by K
func subarraysDivByK(nums []int, k int) int {
    count := 0
    prefixSum := 0
    // freq stores count of each remainder seen
    freq := make([]int, k)
    freq[0] = 1 // empty prefix has remainder 0

    for _, num := range nums {
        prefixSum += num
        // In Go, % can be negative for negative numbers
        mod := ((prefixSum % k) + k) % k // always positive remainder
        count += freq[mod]
        freq[mod]++
    }
    return count
}

// Dry run: nums = [4,5,0,-2,-3,1], k = 5
// prefixSums:   4, 9, 9, 7, 4, 5
// remainders:   4, 4, 4, 2, 4, 0
// freq[0]=1 initially
// i=0: rem=4, count+=freq[4]=0, freq[4]=1
// i=1: rem=4, count+=freq[4]=1→count=1, freq[4]=2
// i=2: rem=4, count+=freq[4]=2→count=3, freq[4]=3
// i=3: rem=2, count+=freq[2]=0, freq[2]=1
// i=4: rem=4, count+=freq[4]=3→count=6, freq[4]=4
// i=5: rem=0, count+=freq[0]=1→count=7, freq[0]=2
// Result: 7 ✓
```

---

### Algorithm 6: Pivot Index / Equilibrium Index

**LeetCode:** #724 — Find Pivot Index

**Key insight:** Pivot at `i` means `leftSum == rightSum`. Since `total = leftSum + arr[i] + rightSum`, we need `leftSum == (total - arr[i] - leftSum)` → `2*leftSum + arr[i] == total`.

```go
// LC 724 - Find Pivot Index
func pivotIndex(nums []int) int {
    total := 0
    for _, v := range nums {
        total += v
    }

    leftSum := 0
    for i, v := range nums {
        // rightSum = total - leftSum - nums[i]
        if leftSum == total-leftSum-v {
            return i
        }
        leftSum += v
    }
    return -1
}

// Dry run: nums = [1, 7, 3, 6, 5, 6], total = 28
// i=0: v=1, leftSum=0, right=28-0-1=27, 0≠27
// i=1: v=7, leftSum=1, right=28-1-7=20, 1≠20
// i=2: v=3, leftSum=8, right=28-8-3=17, 8≠17
// i=3: v=6, leftSum=11, right=28-11-6=11, 11==11 → return 3 ✓
```

---

### Algorithm 7: Difference Array for Range Updates

**LeetCode:** #1094 — Car Pooling, #1109 — Corporate Flight Bookings

```go
// LC 1109 - Corporate Flight Bookings
// bookings[i] = [first, last, seats]: add seats to flights [first, last]
func corpFlightBookings(bookings [][]int, n int) []int {
    diff := make([]int, n+1) // 1-indexed flights, diff[n] as guard

    for _, b := range bookings {
        first, last, seats := b[0]-1, b[1]-1, b[2] // convert to 0-indexed
        diff[first] += seats
        if last+1 <= n-1 {
            diff[last+1] -= seats
        }
    }

    // Recover array via prefix sum
    result := make([]int, n)
    result[0] = diff[0]
    for i := 1; i < n; i++ {
        result[i] = result[i-1] + diff[i]
    }
    return result
}

// Dry run: bookings = [[1,2,10],[2,3,20],[2,5,25]], n=5
// diff (0-indexed): after all updates:
//   booking [0,1,10]: diff[0]+=10, diff[2]-=10
//   booking [1,2,20]: diff[1]+=20, diff[3]-=20
//   booking [1,4,25]: diff[1]+=25, diff[5]-=25 (guard)
// diff = [10, 45, -10, -20, 0, -25]
// prefix: [10, 55, 45, 25, 25]  ✓
```

---

### Algorithm 8: Prefix XOR for Subarray XOR Queries

**LeetCode:** #1310 — XOR Queries of a Subarray

**Key insight:** XOR is its own inverse. `XOR(l, r) = prefixXOR[r] ^ prefixXOR[l-1]` (using 0-indexed with sentinel).

```go
// LC 1310 - XOR Queries of a Subarray
func xorQueries(arr []int, queries [][]int) []int {
    n := len(arr)
    prefix := make([]int, n+1) // prefix[0] = 0 sentinel
    for i, v := range arr {
        prefix[i+1] = prefix[i] ^ v
    }

    result := make([]int, len(queries))
    for i, q := range queries {
        l, r := q[0], q[1]
        result[i] = prefix[r+1] ^ prefix[l] // XOR(l, r)
    }
    return result
}

// Dry run: arr = [1, 3, 4, 8]
// prefix = [0, 1, 2, 6, 14]
//   prefix[1] = 0^1 = 1
//   prefix[2] = 1^3 = 2
//   prefix[3] = 2^4 = 6
//   prefix[4] = 6^8 = 14
// query [0,1]: prefix[2]^prefix[0] = 2^0 = 2 = 1^3 ✓
// query [1,3]: prefix[4]^prefix[1] = 14^1 = 15 = 3^4^8 ✓
```

---

### Algorithm 9: Longest Subarray with Sum = K (Prefix + HashMap with First Index)

**LeetCode:** #325 — Maximum Size Subarray Sum Equals k

**Key insight:** For maximum length, store the **first occurrence** index of each prefix sum (not frequency). Then `maxLen = max(maxLen, i - firstIndex[prefix-k])`.

```go
// LC 325 - Maximum Size Subarray Sum Equals k
func maxSubArrayLen(nums []int, k int) int {
    // firstIndex[sum] = earliest index where this prefixSum was seen
    firstIndex := map[int]int{0: -1} // prefix sum 0 at virtual index -1
    prefixSum := 0
    maxLen := 0

    for i, num := range nums {
        prefixSum += num
        if idx, ok := firstIndex[prefixSum-k]; ok {
            length := i - idx
            if length > maxLen {
                maxLen = length
            }
        }
        // Only store FIRST occurrence (don't overwrite)
        if _, exists := firstIndex[prefixSum]; !exists {
            firstIndex[prefixSum] = i
        }
    }
    return maxLen
}

// Dry run: nums = [1, -1, 5, -2, 3], k = 3
// i=0: ps=1,  lookup ps-k=-2 → not found. Store first[1]=0
// i=1: ps=0,  lookup ps-k=-3 → not found. Store first[0]=1? NO, first[0]=-1 already set
// i=2: ps=5,  lookup ps-k=2  → not found. Store first[5]=2
// i=3: ps=3,  lookup ps-k=0  → found at idx=-1, len=3-(-1)=4 → maxLen=4
// i=4: ps=6,  lookup ps-k=3  → found at idx=3, len=4-3=1
// Result: 4 ✓ (subarray [1,-1,5,-2])
```

---

### Algorithm 10: Contiguous Array (Binary → Prefix)

**LeetCode:** #525 — Contiguous Array (Equal 0s and 1s)

**Key insight:** Replace 0s with -1s. Now equal 0s and 1s means subarray sum = 0. Find longest subarray with sum = 0 using prefix + hashmap.

```go
// LC 525 - Contiguous Array
func findMaxLength(nums []int) int {
    firstIndex := map[int]int{0: -1}
    prefixSum := 0
    maxLen := 0

    for i, num := range nums {
        if num == 0 {
            prefixSum-- // treat 0 as -1
        } else {
            prefixSum++
        }

        if idx, ok := firstIndex[prefixSum]; ok {
            if i-idx > maxLen {
                maxLen = i - idx
            }
        } else {
            firstIndex[prefixSum] = i
        }
    }
    return maxLen
}

// Dry run: nums = [0, 1, 0, 0, 1, 1, 0]
// Treat as:      [-1, 1,-1,-1, 1, 1,-1]
// prefixSums: -1, 0, -1, -2, -1, 0, -1
// i=0: ps=-1, not in map, store first[-1]=0
// i=1: ps=0,  found at -1, len=1-(-1)=2, maxLen=2
// i=2: ps=-1, found at 0,  len=2-0=2
// i=3: ps=-2, not in map, store
// i=4: ps=-1, found at 0,  len=4-0=4, maxLen=4
// i=5: ps=0,  found at -1, len=5-(-1)=6, maxLen=6
// i=6: ps=-1, found at 0,  len=6-0=6
// Result: 6 ✓
```

---

## SECTION 5 — Time & Space Complexity

| Algorithm / Problem                   | Time Complexity | Space Complexity | Notes                     |
| ------------------------------------- | --------------- | ---------------- | ------------------------- |
| Build 1D Prefix Array                 | O(n)            | O(n)             | One-time cost             |
| Range Sum Query (after build)         | O(1)            | O(1)             | Per query                 |
| Build 2D Prefix Array                 | O(m × n)        | O(m × n)         | Grid size                 |
| 2D Range Query (after build)          | O(1)            | O(1)             | Per query                 |
| Count Subarrays with Sum=K (LC 560)   | O(n)            | O(n)             | HashMap for freq          |
| Max Length Subarray Sum=K (LC 325)    | O(n)            | O(n)             | HashMap for first index   |
| Subarray Sum Divisible by K (LC 974)  | O(n)            | O(k)             | Remainder array of size k |
| Product of Array Except Self (LC 238) | O(n)            | O(1)\*           | \*Excluding output array  |
| Contiguous Array (LC 525)             | O(n)            | O(n)             | HashMap                   |
| Difference Array Build                | O(n)            | O(n)             |                           |
| Difference Array Range Update         | O(1)            | O(1)             | Per update                |
| Difference Array Recover              | O(n)            | O(1)             | In-place prefix sum       |
| Prefix XOR Queries                    | O(n + q)        | O(n)             | q = num queries           |
| Pivot Index (LC 724)                  | O(n)            | O(1)             | Single pass after total   |

> **Space Optimization Note:** The prefix array can be eliminated if queries are batched and sorted. For Product Except Self, the output array doubles as the prefix array, achieving O(1) auxiliary space. For the HashMap approach, space is unavoidable since we need lookups of arbitrary past prefix values.

---

## SECTION 6 — Problem-Solving Templates

### Universal Decision Tree

```
Is the problem about SUBARRAYS / RANGES?
│
├─ YES: Does it involve UPDATES (add/remove from range)?
│       ├─ YES → Difference Array (then prefix to recover)
│       └─ NO  → Prefix Sum / Prefix Product
│
├─ Are there MULTIPLE QUERIES on FIXED array?
│       ├─ YES, 1D → Build prefix once, answer in O(1)
│       └─ YES, 2D → Build 2D prefix once, answer in O(1)
│
├─ COUNT subarrays with PROPERTY (sum=k, divisible by k)?
│       └─ Prefix Sum + HashMap
│               ├─ count  → store FREQUENCY in map
│               └─ length → store FIRST INDEX in map
│
├─ Product / XOR / GCD instead of sum?
│       ├─ XOR → Prefix XOR (XOR is self-inverse)
│       ├─ Product → Prefix × Suffix (avoid division when zeros possible)
│       └─ GCD → Not easily invertible, different approach
│
└─ 2D grid, find rectangle with max/target sum?
        └─ 2D Prefix + (sliding window or binary search for max)
```

---

### Pattern Recognition Quick-Map

| Keyword / Signal                           | Technique                            |
| ------------------------------------------ | ------------------------------------ |
| "sum from index l to r"                    | Prefix sum + range query             |
| "multiple queries, immutable array"        | Build prefix once                    |
| "count subarrays with sum = k"             | Prefix + freq map                    |
| "count subarrays divisible by k"           | Prefix + mod + remainder array       |
| "longest subarray with sum = k"            | Prefix + first-index map             |
| "equal 0s and 1s"                          | Convert 0→-1, longest sum=0 subarray |
| "product except self"                      | Prefix × Suffix products             |
| "range update, then read"                  | Difference array                     |
| "XOR of range"                             | Prefix XOR                           |
| "equilibrium / pivot index"                | Total sum - left sum                 |
| "rectangle sum in grid"                    | 2D prefix sum                        |
| "k consecutive ones"                       | Sliding window (not prefix)          |
| "minimum subarray sum" (negative elements) | Prefix + deque or Kadane's           |

---

### Universal Code Template

```go
// ==== TEMPLATE: Prefix Sum + HashMap ====
// Covers 80% of "count/find subarray with property" problems

func prefixSumTemplate(nums []int, target int) int {
    result := 0
    prefixSum := 0

    // KEY: initialize with the "empty prefix" case
    // For COUNT: freq map with freq[0] = 1
    // For LENGTH: index map with firstIndex[0] = -1
    freq := map[int]int{0: 1}

    for _, num := range nums {
        prefixSum += num // or: ^= for XOR, modify for other operations

        // QUERY: what are we looking for?
        needed := prefixSum - target // for sum = target
        // needed = prefixSum % k     // for divisible by k (then lookup freq[needed])
        result += freq[needed]

        // UPDATE: record current prefix sum
        freq[prefixSum]++
        // For FIRST INDEX: if _, ok := firstIndex[prefixSum]; !ok { firstIndex[prefixSum] = i }
    }
    return result
}

// ==== TEMPLATE: Basic Range Query ====
func buildAndQuery(arr []int, queries [][2]int) []int {
    n := len(arr)
    prefix := make([]int, n+1)
    for i, v := range arr {
        prefix[i+1] = prefix[i] + v
    }

    results := make([]int, len(queries))
    for i, q := range queries {
        l, r := q[0], q[1]
        results[i] = prefix[r+1] - prefix[l]
    }
    return results
}

// ==== TEMPLATE: Difference Array ====
func differenceArrayTemplate(n int, updates [][3]int) []int {
    diff := make([]int, n+1)
    for _, u := range updates {
        l, r, val := u[0], u[1], u[2]
        diff[l] += val
        if r+1 < len(diff) {
            diff[r+1] -= val
        }
    }
    // Recover via prefix sum
    for i := 1; i < n; i++ {
        diff[i] += diff[i-1]
    }
    return diff[:n]
}
```

---

## SECTION 7 — Edge Cases & Gotchas

### 7 Critical Edge Cases

1. **Single element array:** `sum(0, 0) = prefix[1] - prefix[0] = arr[0]`. Works correctly with 1-indexed prefix. Always test this.

2. **All negative numbers:** Prefix sum is decreasing. Max subarray sum algorithms (Kadane's) are different. Prefix sums still work for range queries — they're not restricted to positive values.

3. **Zeros in product array:** Division-based approaches to prefix product FAIL. Always use the "pass from left, then pass from right" pattern. Zeros require special tracking (`countZero`, `nonZeroProduct`).

4. **Overflow:** `n = 10^5`, values `up to 10^9` → prefix sums can reach `10^14`. Always use `int64` in Go for such constraints.

5. **Negative modulo in Go:** `(-3) % 5 = -3` in Go (not 2). Fix: `((prefixSum % k) + k) % k`. Forgetting this will fail all negative-number divisibility problems.

6. **Empty subarray / empty range:** The sentinel `prefix[0] = 0` handles the case where the subarray starts at index 0. Without it, you'd need special-case handling for `l == 0`. Always add the sentinel.

7. **First-index vs frequency in HashMap:** For **count** problems, store frequency (`map[sum]count`). For **max length** problems, store **first index** (`map[sum]index`) and **never overwrite** existing entries. Mixing these up is a very common bug.

---

### Top 5 Common Bugs: WRONG vs CORRECT

```go
// BUG 1: Missing sentinel in HashMap — misses subarrays starting at index 0
// WRONG:
freq := map[int]int{}  // empty map

// CORRECT:
freq := map[int]int{0: 1}  // sum=0 seen once (empty prefix)

// ─────────────────────────────────────────────

// BUG 2: Off-by-one in range query (both 0-indexed arr and prefix)
// WRONG:
sum := prefix[r] - prefix[l]  // misses arr[r], double-counts boundary

// CORRECT (with n+1 sized prefix, prefix[0]=0 sentinel):
sum := prefix[r+1] - prefix[l]  // arr[l..r] inclusive

// ─────────────────────────────────────────────

// BUG 3: Negative modulo
// WRONG (fails for negative prefixSum):
mod := prefixSum % k

// CORRECT:
mod := ((prefixSum % k) + k) % k

// ─────────────────────────────────────────────

// BUG 4: Overwriting first index in "max length" problems
// WRONG:
firstIndex[prefixSum] = i  // overwrites earlier, shorter result

// CORRECT:
if _, exists := firstIndex[prefixSum]; !exists {
    firstIndex[prefixSum] = i  // only store FIRST occurrence
}

// ─────────────────────────────────────────────

// BUG 5: Using division for prefix product (fails with zeros)
// WRONG:
func productRange(prefix []int, l, r int) int {
    return prefix[r] / prefix[l-1]  // division by zero if prefix[l-1]=0!
}

// CORRECT: Use separate left-product and right-product passes
// (see Algorithm 3 — Product Except Self)
```

---

### Interviewer Follow-Up Questions (with Brief Answers)

| Follow-Up                                                | Brief Answer                                                                                                      |
| -------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| "Can you do it in O(1) space?"                           | For range queries: no (need prefix). For product-except-self: yes (use output array + running variable).          |
| "What if the array is updated frequently?"               | Use a Fenwick Tree (BIT) or Segment Tree — O(log n) update + O(log n) query.                                      |
| "Can you handle 2D range sum queries?"                   | Yes, 2D prefix sum: O(mn) build, O(1) per query using inclusion-exclusion.                                        |
| "What if you need both range updates AND range queries?" | Use Segment Tree with lazy propagation or Fenwick Tree with range updates.                                        |
| "What if values can be very large?"                      | Use `int64`. Also, for modular problems, keep sums modulo the given modulus.                                      |
| "Can prefix sum find the maximum subarray sum?"          | Only with all-positive arrays (then max subarray is whole array). With negatives, use Kadane's or prefix + deque. |
| "How would you parallelize this?"                        | Parallel prefix sum (scan algorithm) — prefix computation can be parallelized in O(log n) depth.                  |

---

## SECTION 8 — Advanced Techniques

### Advanced Technique 1: Prefix Sum + Binary Search (Minimum Size Subarray Sum)

**LeetCode:** #209 — Minimum Size Subarray Sum (with all positive numbers)

Since all values are positive, the prefix array is strictly increasing. For each right boundary `r`, use binary search to find smallest `l` such that `prefix[r] - prefix[l-1] >= target`.

```go
// LC 209 - Minimum Size Subarray Sum (Binary Search approach, O(n log n))
// Note: Sliding window is O(n) and preferred, but this shows prefix+BS combo
func minSubArrayLen(target int, nums []int) int {
    n := len(nums)
    prefix := make([]int, n+1)
    for i, v := range nums {
        prefix[i+1] = prefix[i] + v
    }

    minLen := n + 1

    for r := 1; r <= n; r++ {
        // Find smallest l such that prefix[r] - prefix[l] >= target
        // i.e., prefix[l] <= prefix[r] - target
        needed := prefix[r] - target
        // Binary search for rightmost index l with prefix[l] <= needed
        lo, hi := 0, r
        for lo < hi {
            mid := (lo + hi + 1) / 2
            if prefix[mid] <= needed {
                lo = mid
            } else {
                hi = mid - 1
            }
        }
        // lo is the largest l where prefix[l] <= needed
        // subarray is [lo+1 .. r] in original (0-indexed: lo..r-1)
        if prefix[lo] <= needed {
            length := r - lo
            if length < minLen {
                minLen = length
            }
        }
    }

    if minLen == n+1 {
        return 0
    }
    return minLen
}
```

---

### Advanced Technique 2: 2D Prefix Sum + Sliding Window (Max Sum Rectangle ≤ K)

**LeetCode:** #363 — Max Sum of Rectangle No Larger Than K

**Idea:** Fix left and right column boundaries. Compute row-wise prefix sums to get 1D "column sum" array, then find max subarray sum ≤ k using a sorted set + binary search.

```go
import "sort"

// LC 363 - Max Sum of Rectangle No Larger Than K, O(m² × n log n)
func maxSumSubmatrix(matrix [][]int, k int) int {
    m, n := len(matrix), len(matrix[0])
    result := -1 << 62 // negative infinity

    // Fix left column
    for left := 0; left < n; left++ {
        rowSum := make([]int, m) // cumulative sum for each row between left..right

        // Expand right column
        for right := left; right < n; right++ {
            for r := 0; r < m; r++ {
                rowSum[r] += matrix[r][right]
            }

            // Now find max subarray sum <= k in rowSum[]
            // Using sorted prefix sums + binary search
            sortedPrefix := []int{0}
            prefixSum := 0

            for _, v := range rowSum {
                prefixSum += v
                // Find smallest prefix in sortedPrefix >= prefixSum - k
                target := prefixSum - k
                idx := sort.SearchInts(sortedPrefix, target)
                if idx < len(sortedPrefix) {
                    candidate := prefixSum - sortedPrefix[idx]
                    if candidate > result {
                        result = candidate
                    }
                }
                // Insert prefixSum into sortedPrefix (maintain sorted order)
                ins := sort.SearchInts(sortedPrefix, prefixSum)
                sortedPrefix = append(sortedPrefix, 0)
                copy(sortedPrefix[ins+1:], sortedPrefix[ins:])
                sortedPrefix[ins] = prefixSum
            }
        }
    }
    return result
}
```

---

### Advanced Technique 3: Prefix Sum + DP (Combination)

**LeetCode:** #1074 — Number of Submatrices That Sum to Target

**Idea:** Use 2D prefix sums + the "count subarrays with sum=k" technique (prefix + hashmap) applied row-by-row.

```go
// LC 1074 - Number of Submatrices That Sum to Target, O(m² × n)
func numSubmatrixSumTarget(matrix [][]int, target int) int {
    m, n := len(matrix), len(matrix[0])

    // Build row-wise prefix sums
    for r := 0; r < m; r++ {
        for c := 1; c < n; c++ {
            matrix[r][c] += matrix[r][c-1]
        }
    }

    result := 0

    // Fix left and right column, reduce to 1D subarray sum = target
    for left := 0; left < n; left++ {
        for right := left; right < n; right++ {
            freq := map[int]int{0: 1}
            prefixSum := 0

            for r := 0; r < m; r++ {
                // Column sum from left to right in row r
                colSum := matrix[r][right]
                if left > 0 {
                    colSum -= matrix[r][left-1]
                }
                prefixSum += colSum
                result += freq[prefixSum-target]
                freq[prefixSum]++
            }
        }
    }
    return result
}
```

---

### Advanced Technique 4: Fenwick Tree (BIT) — Dynamic Prefix Sums

When the array is **mutable** (point updates happen), prefix sum recomputation is O(n). A **Fenwick Tree** gives O(log n) update AND O(log n) prefix query.

```go
// Fenwick Tree (Binary Indexed Tree) for dynamic prefix sums
type BIT struct {
    tree []int
    n    int
}

func NewBIT(n int) *BIT {
    return &BIT{tree: make([]int, n+1), n: n}
}

// Update: add delta to index i (1-indexed)
func (b *BIT) Update(i, delta int) {
    for ; i <= b.n; i += i & (-i) { // i += lowest set bit
        b.tree[i] += delta
    }
}

// Query: prefix sum [1..i]
func (b *BIT) Query(i int) int {
    sum := 0
    for ; i > 0; i -= i & (-i) { // i -= lowest set bit
        sum += b.tree[i]
    }
    return sum
}

// Range sum [l..r] (1-indexed)
func (b *BIT) RangeQuery(l, r int) int {
    return b.Query(r) - b.Query(l-1)
}

// Example: LC 307 - Range Sum Query Mutable
// Update: O(log n), Query: O(log n) — vs O(n) rebuild for plain prefix array
```

---

## SECTION 9 — LeetCode Problem Set (30 Problems)

| #   | Problem                                      | LC # | Difficulty | Key Technique                                      |
| --- | -------------------------------------------- | ---- | ---------- | -------------------------------------------------- |
| 1   | Running Sum of 1D Array                      | 1480 | Easy       | Basic prefix sum                                   |
| 2   | Find Pivot Index                             | 724  | Easy       | Total sum - left sum                               |
| 3   | Find the Middle Index in Array               | 1991 | Easy       | Same as pivot index                                |
| 4   | Range Sum Query - Immutable                  | 303  | Easy       | Prefix sum, OOP                                    |
| 5   | Left and Right Sum Differences               | 2574 | Easy       | Prefix + suffix sums                               |
| 6   | Subarray Sum Equals K                        | 560  | Medium     | Prefix + freq map                                  |
| 7   | Contiguous Array                             | 525  | Medium     | 0→-1 transform, prefix + map                       |
| 8   | Product of Array Except Self                 | 238  | Medium     | Prefix × suffix product                            |
| 9   | Range Sum Query 2D - Immutable               | 304  | Medium     | 2D prefix sum                                      |
| 10  | Subarray Sums Divisible by K                 | 974  | Medium     | Prefix + mod + remainder array                     |
| 11  | Maximum Size Subarray Sum Equals k           | 325  | Medium     | Prefix + first-index map                           |
| 12  | Minimum Size Subarray Sum                    | 209  | Medium     | Prefix + binary search / sliding window            |
| 13  | Corporate Flight Bookings                    | 1109 | Medium     | Difference array                                   |
| 14  | Car Pooling                                  | 1094 | Medium     | Difference array                                   |
| 15  | XOR Queries of a Subarray                    | 1310 | Medium     | Prefix XOR                                         |
| 16  | Find the Longest Substring Containing Vowels | 1371 | Medium     | Prefix XOR bitmask + map                           |
| 17  | Matrix Block Sum                             | 1314 | Medium     | 2D prefix sum                                      |
| 18  | Sum of Subarray Minimums                     | 907  | Medium     | Prefix sum + monotonic stack                       |
| 19  | Number of Ways to Split Array                | 2270 | Medium     | Prefix sum comparison                              |
| 20  | Count Number of Nice Subarrays               | 1248 | Medium     | Prefix sum (odd count) + map                       |
| 21  | Continuous Subarray Sum                      | 523  | Medium     | Prefix + mod, divisibility check                   |
| 22  | Make Sum Divisible by P                      | 1590 | Medium     | Prefix mod + hashmap                               |
| 23  | Random Pick with Weight                      | 528  | Medium     | Prefix sum + binary search                         |
| 24  | Path Sum III                                 | 437  | Medium     | Prefix sum on tree (DFS)                           |
| 25  | Number of Submatrices That Sum to Target     | 1074 | Hard       | 2D prefix + count subarrays                        |
| 26  | Max Sum of Rectangle No Larger Than K        | 363  | Hard       | 2D prefix + sorted set                             |
| 27  | Count of Range Sum                           | 327  | Hard       | Prefix sum + merge sort                            |
| 28  | Subarrays with K Different Integers          | 992  | Hard       | Prefix count (exactly k = at most k - at most k-1) |
| 29  | Shortest Subarray with Sum at Least K        | 862  | Hard       | Prefix sum + monotonic deque                       |
| 30  | Maximum Subarray Sum with One Deletion       | 1186 | Hard       | Prefix sum + DP                                    |

---

### Recommended Study Order — 3-Week Breakdown

**Week 1: Foundation (Problems #1–10)**
Build intuition for prefix sum construction and basic range queries. Understand the sentinel. Learn the prefix + hashmap pattern deeply via LC 560 and LC 525.

| Day     | Problems                                                      |
| ------- | ------------------------------------------------------------- |
| Day 1-2 | #1480, #724, #1991, #303 — Basic prefix sum, sentinel concept |
| Day 3-4 | #560, #525 — Prefix + HashMap (the core advanced pattern)     |
| Day 5-6 | #238, #304, #2574 — Prefix/Suffix product, 2D prefix          |
| Day 7   | Revisit and solve without hints                               |

**Week 2: Intermediate Patterns (Problems #11–22)**
Difference arrays, XOR, modular arithmetic, 2D applications, tree prefix sums.

| Day       | Problems                                                          |
| --------- | ----------------------------------------------------------------- |
| Day 8-9   | #974, #325, #209 — Mod arithmetic, first-index map, binary search |
| Day 10-11 | #1109, #1094 — Difference arrays                                  |
| Day 12-13 | #1310, #1371 — Prefix XOR, bitmask prefix                         |
| Day 14    | #1314, #907, #437 — 2D prefix, stack combo, tree prefix           |

**Week 3: Advanced & Hard (Problems #23–30)**
Complex combinations, offline queries, hard constraints.

| Day       | Problems                                                  |
| --------- | --------------------------------------------------------- |
| Day 15-16 | #2270, #1248, #523, #528 — Variety of medium applications |
| Day 17-18 | #1590, #1074 — Hard mediums and 2D hard                   |
| Day 19-20 | #363, #327 — Hard problems (sorted set, merge sort)       |
| Day 21    | #992, #862, #1186 — Hardest problems, review all patterns |

---

## SECTION 10 — Interview Strategy

### Clarification Questions to Ask

```
1. "Are the array values guaranteed to be positive, or can there be negatives/zeros?"
   → Affects whether prefix is monotone; zeros break product division approach.

2. "Is the array mutable? Will there be update operations between queries?"
   → If yes → Fenwick Tree or Segment Tree instead of static prefix.

3. "What is the range of values and array size?"
   → Check for overflow. n=10^5, values=10^9 → need int64.

4. "Are query indices guaranteed to be valid (l ≤ r, within bounds)?"
   → Determines how much bounds-checking to add.

5. "For 2D problems: is the matrix guaranteed to be non-empty?"
   → Edge case: empty matrix.

6. "Should I optimize for time or space? Is O(n) extra space acceptable?"
   → Determines whether you pursue the O(1) space variants.

7. "Are there multiple test cases sharing the same array, or is it single-use?"
   → Justifies O(n) build time (amortized over many queries).
```

---

### Step-by-Step Communication Template

```
STEP 1 — APPROACH (30 seconds, before writing any code)
"I see this is asking for [range sum / count subarrays / etc.].
 My approach: precompute a prefix sum array in O(n), then each query
 is O(1) subtraction. Total: O(n + q) time, O(n) space."

STEP 2 — TRACE (1-2 minutes on example)
"Let me trace through the example:
 arr = [1, 2, 3], prefix = [0, 1, 3, 6]
 query(0, 2) = prefix[3] - prefix[0] = 6 - 0 = 6 ✓"

STEP 3 — CODE (write with commentary)
"I'll start with the build phase — notice I'm using size n+1 with
 prefix[0] = 0 as a sentinel, so I don't need special-case logic
 for queries starting at index 0..."

STEP 4 — TEST (edge cases aloud)
"Let me check: single element → prefix[1]-prefix[0] = arr[0] ✓
 First element query (l=0): prefix[r+1]-prefix[0] = prefix[r+1] ✓
 Negative values: prefix can decrease, but subtraction still correct ✓
 Overflow: values up to 10^9, n up to 10^5 → max sum 10^14, need int64"
```

---

### Pattern-Specific Interviewer Red Flags

> **Watch for these in your own code during interview:**

1. **Missing sentinel** → `freq := map[int]int{}` instead of `{0: 1}` — will silently miss subarrays starting at index 0.
2. **Wrong prefix array size** → `make([]int, n)` instead of `make([]int, n+1)` — index out of bounds at runtime.
3. **Overwriting first-index map** → `firstIndex[sum] = i` inside loop without existence check — gives wrong max length.
4. **Negative mod** → `prefixSum % k` without `+k)%k` correction — fails all negative test cases.
5. **Division by zero in product** → Attempting `totalProduct / arr[i]` when `arr[i] == 0`.
6. **Premature optimization** → Trying to avoid the prefix array entirely and botching the logic.

---

## SECTION 11 — Quick Revision Summary

> ### 🚀 Prefix Sum — 60-Second Cheat Sheet
>
> - **Build:** `prefix[i+1] = prefix[i] + arr[i]`, size `n+1`, `prefix[0] = 0` sentinel
> - **Range query:** `sum(l, r) = prefix[r+1] - prefix[l]` (0-indexed arr, 1-indexed prefix)
> - **2D range:** `P[r2+1][c2+1] - P[r1][c2+1] - P[r2+1][c1] + P[r1][c1]`
> - **Count subarrays sum=k:** Prefix + HashMap with `freq[0] = 1`, add `freq[prefix - k]` before updating
> - **Max length subarray sum=k:** Prefix + first-index map, **never overwrite** existing entries
> - **Divisible by k:** Store `((prefix % k) + k) % k` remainders, same count pattern
> - **Product except self:** Left-product pass → then right-product pass in-place (avoids division)
> - **Prefix XOR:** `XOR(l,r) = prefix[r+1] ^ prefix[l]`, XOR is self-inverse
> - **Difference array:** `diff[l]+=val, diff[r+1]-=val` → recover via prefix sum → O(1) range update
> - **Mutable array?** → Switch to Fenwick Tree: O(log n) update + O(log n) query
> - **Negative mod in Go:** Always `((x % k) + k) % k` to ensure non-negative remainder
> - **Overflow guard:** Use `int64` when `n × maxVal > 2^31`
