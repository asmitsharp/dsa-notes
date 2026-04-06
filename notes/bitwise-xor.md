# 🧠 DSA Pattern Notes: Bitwise XOR

> **Mentor's note:** XOR is one of those rare patterns where a single operator unlocks elegant solutions to problems that otherwise feel impossible. Learn the 3–4 key properties and you'll spot XOR problems instantly.

---

## 1. Pattern Overview

### What is this pattern?

Bitwise XOR (`^`) is a binary operation that compares two bits and returns `1` if they're **different**, and `0` if they're the **same**.

```
0 ^ 0 = 0
1 ^ 1 = 0
0 ^ 1 = 1
1 ^ 0 = 1
```

In interviews, XOR is used to solve problems involving **pairs**, **missing numbers**, **duplicates**, and **unique elements** — often in O(n) time and O(1) space.

### Why it's used (intuition)

XOR has a magical property: **it cancels itself out**. If you XOR a number with itself, you get 0. If you XOR a number with 0, you get the number back. This means when you XOR a whole array, **duplicate values vanish**, and only the "leftover" values survive. This is incredibly powerful.

### Real-world analogy

Think of XOR like a **toggle switch**. Every time you flip a switch (XOR with 1), it changes state. If you flip it twice, you're back to where you started — the two flips **cancel out**. XOR works exactly like this with bits. Pair something with itself → it disappears. Pair something with nothing (0) → it stays.

### Time complexity intuition

Most XOR-based problems run in **O(n)** time because you just scan the array once, XOR-ing as you go. There's no sorting, no nested loops, no extra hash map lookups. That single pass is why XOR solutions are so clean and fast.

---

## 2. Core Idea (MOST IMPORTANT)

The entire pattern rests on **four properties** of XOR. Internalize these and everything else follows naturally.

```
a ^ a = 0          → A number XOR'd with itself is always 0
a ^ 0 = a          → A number XOR'd with 0 is unchanged
a ^ b = b ^ a      → Order doesn't matter (commutative)
(a ^ b) ^ c = a ^ (b ^ c)  → Grouping doesn't matter (associative)
```

### The key insight

Because XOR is commutative and associative, you can XOR numbers **in any order**. All pairs cancel to 0, and 0 XOR'd with anything leaves that thing untouched. So if you XOR every element in an array, only the **unpaired** elements survive.

> 💡 **Remember this:** `[2, 3, 2, 4, 3]` → XOR all → `2^3^2^4^3` → `(2^2)^(3^3)^4` → `0^0^4` → **4** is the answer. The unique element survives.

### Bonus: XOR for swapping

```go
a = a ^ b
b = a ^ b   // b becomes original a
a = a ^ b   // a becomes original b
```

No temp variable needed. Useful to know, occasionally asked.

---

## 3. When to Use This Pattern

| Signal in Question                           | What It Means                       |
| -------------------------------------------- | ----------------------------------- |
| "Find the single/unique number"              | Pairs cancel → XOR all elements     |
| "Find the missing number in range [1..n]"    | XOR array with XOR of [1..n]        |
| "Two numbers appear once, rest appear twice" | XOR all, then split using a set bit |
| "Find the number that appears odd times"     | XOR all → odd-frequency survivor    |
| "Swap without extra space"                   | XOR swap trick                      |
| "Complement / flip bits"                     | XOR with all-1s mask                |
| O(1) space required + O(n) time              | Strong hint to think XOR            |

---

## 4. Types / Variants of This Pattern

### Type 1: Single Unique Element (pairs cancel)

XOR all elements. Everything paired cancels to 0. The unpaired element remains.

**Example:** `[4, 1, 2, 1, 2]` → result is `4`

### Type 2: Missing Number in Range

XOR the array with XOR of `[0..n]` (or `[1..n]`). Every present number cancels, leaving the missing one.

**Example:** `[3, 0, 1]` in range 0–3 → XOR all with 0^1^2^3 → missing `2` survives.

### Type 3: Two Unique Elements (two singles, rest are pairs)

XOR all elements to get `x ^ y`. Find any set bit in the result (use `result & -result` to isolate the rightmost set bit). Use that bit to split all numbers into two groups. XOR within each group gives you `x` and `y` separately.

**Example:** `[1, 2, 1, 3, 2, 5]` → unique are `3` and `5`

### Type 4: XOR in a Range / Prefix XOR

XOR from `[l..r]` can be computed using prefix XOR: `xor(l, r) = prefix[r] ^ prefix[l-1]`. Useful for range queries. There's also a pattern for `XOR(0..n)` that cycles every 4 values: `n, 1, n+1, 0` depending on `n % 4`.

---

## 5. Data Structures Used

XOR problems are special because they often need **no extra data structures at all** — just a single integer variable acting as an accumulator. This is why they're loved in interviews: O(1) space solutions.

| Data Structure               | When Used            | Why                                                                |
| ---------------------------- | -------------------- | ------------------------------------------------------------------ |
| Single integer (accumulator) | Almost always        | XOR is applied in-place; no storage needed                         |
| Array (input only)           | Always               | We just scan it once                                               |
| Bitmask / integer            | Type 3 (two uniques) | Store the combined XOR `x^y`, then isolate a differentiating bit   |
| Prefix XOR array             | Range XOR queries    | Precompute `prefix[i] = a[0]^a[1]^...^a[i]` for O(1) range queries |

The beauty of XOR: **you need almost nothing**. One variable, one pass.

---

## 6. Core Templates

### Template 1: Find Single Unique Element

```go
func singleNumber(nums []int) int {
    result := 0
    for _, n := range nums {
        result ^= n  // pairs cancel out, unique survives
    }
    return result
}
```

### Template 2: Find Missing Number in [0..n]

```go
func missingNumber(nums []int) int {
    n := len(nums)
    xor := 0
    for i := 0; i <= n; i++ {
        xor ^= i  // XOR with full range [0..n]
    }
    for _, n := range nums {
        xor ^= n  // cancel out present numbers
    }
    return xor  // only missing number remains
}
```

### Template 3: Find Two Unique Numbers (rest appear twice)

```go
func singleNumberIII(nums []int) []int {
    // Step 1: XOR all to get x^y
    xorAll := 0
    for _, n := range nums {
        xorAll ^= n
    }

    // Step 2: Find a bit where x and y differ
    // rightmost set bit isolates one differing bit
    diffBit := xorAll & (-xorAll)

    // Step 3: Split into two groups by that bit
    x, y := 0, 0
    for _, n := range nums {
        if n & diffBit == 0 {
            x ^= n  // group where bit is 0
        } else {
            y ^= n  // group where bit is 1
        }
    }
    return []int{x, y}
}
```

### Template 4: Prefix XOR for Range Queries

```go
func buildPrefixXOR(nums []int) []int {
    prefix := make([]int, len(nums)+1)
    for i, n := range nums {
        prefix[i+1] = prefix[i] ^ n
    }
    return prefix
}

// XOR of nums[l..r] (0-indexed)
func rangeXOR(prefix []int, l, r int) int {
    return prefix[r+1] ^ prefix[l]
}
```

---

## 7. Step-by-Step Example (Dry Run)

**Problem:** Find the single number in `[2, 3, 4, 3, 2]`

We XOR every element one by one:

```
result = 0

Step 1: result = 0 ^ 2  = 2        (binary: 010)
Step 2: result = 2 ^ 3  = 1        (010 ^ 011 = 001)
Step 3: result = 1 ^ 4  = 5        (001 ^ 100 = 101)
Step 4: result = 5 ^ 3  = 6        (101 ^ 011 = 110)
Step 5: result = 6 ^ 2  = 4        (110 ^ 010 = 100)

Final: result = 4 ✅
```

**Why did 4 survive?** Because `2^2 = 0` and `3^3 = 0`, they silently cancelled each other out regardless of position. XOR is commutative so it doesn't matter they weren't adjacent. Only `4` had no partner.

---

## 8. Must-Know Problems (Deep Explanation)

### Problem 1: Single Number (LC #136) — Easy

**Intuition:** Every number appears twice except one. XOR all elements. Pairs cancel to 0. Unique number remains.

**Approach:** One pass, one XOR accumulator.

```go
func singleNumber(nums []int) int {
    result := 0
    for _, n := range nums {
        result ^= n
    }
    return result
}
```

**Time:** O(n) | **Space:** O(1)

---

### Problem 2: Missing Number (LC #268) — Easy

**Intuition:** XOR the array with the "complete" range `[0..n]`. Every present number cancels. The missing one has no partner, so it survives.

**Approach:** XOR indices `0..n` along with array values in one loop.

```go
func missingNumber(nums []int) int {
    result := len(nums)  // start with n since loop goes 0..n-1
    for i, n := range nums {
        result ^= i ^ n  // XOR index and value together
    }
    return result
}
```

**Dry run for `[3,0,1]`:**

```
result = 3
i=0: result ^= 0^3 = 3^0^3 = 0
i=1: result ^= 1^0 = 0^1^0 = 1
i=2: result ^= 2^1 = 1^2^1 = 2  ✅
```

**Time:** O(n) | **Space:** O(1)

---

### Problem 3: Single Number III (LC #260) — Medium

**Intuition:** Two unique numbers exist. XOR all to get `x^y`. Since x ≠ y, at least one bit differs. Use that bit to split the array into two groups — x lands in one, y in the other. XOR within each group gives the answer.

**Key trick:** `xorAll & (-xorAll)` isolates the rightmost set bit (the bit where x and y differ).

```go
func singleNumber(nums []int) []int {
    xorAll := 0
    for _, n := range nums {
        xorAll ^= n
    }

    // isolate rightmost set bit
    diffBit := xorAll & (-xorAll)

    x, y := 0, 0
    for _, n := range nums {
        if n & diffBit == 0 {
            x ^= n
        } else {
            y ^= n
        }
    }
    return []int{x, y}
}
```

**Time:** O(n) | **Space:** O(1)

---

## 9. Common Mistakes / Gotchas

**1. Forgetting `a^a = 0` needs the element to appear EXACTLY twice.**
If a number appears 4 times it also cancels. But 3 times? It survives (odd count). Make sure you understand the problem says "exactly twice."

**2. Confusing XOR with OR or AND.**
XOR (`^`) is different from OR (`|`) and AND (`&`). In Go, `^` is XOR for binary operands. Don't mix them up under interview pressure.

**3. Missing number: forgetting to include `n` itself.**
In LC #268, the range is `[0..n]`. Many people XOR only `[0..n-1]` and miss the last value. Start `result = n` or loop `for i := 0; i <= n; i++`.

**4. Two unique numbers: not isolating the differing bit correctly.**
The expression `xorAll & (-xorAll)` isolates the **rightmost** set bit. Some people try to find any set bit manually, which works but is messier. Learn this idiom — it's clean and standard.

**5. Assuming XOR works for "find duplicates" when count > 2.**
XOR only naturally handles elements that appear in even/odd counts. If an element appears 3 times in a problem asking about triples, a different approach (like modular counters or bit counting) is needed.

**6. Off-by-one in range XOR.**
When computing `XOR(0..n)`, the cycle pattern is `[n, 1, n+1, 0]` based on `n % 4`. Applying the wrong formula because you shifted the range by 1 is a common error.

**7. Using the wrong bit isolation for Type 3 problems.**
Some candidates use `xorAll & 1` to just check the last bit. This works only if x and y differ in the last bit, which isn't guaranteed. Always use `xorAll & (-xorAll)` to correctly isolate the rightmost set bit.

---

## 10. Time & Space Complexity

| Operation                     | Time | Space | Why                                |
| ----------------------------- | ---- | ----- | ---------------------------------- |
| Single unique element         | O(n) | O(1)  | One pass, one variable             |
| Missing number                | O(n) | O(1)  | One pass, XOR both range and array |
| Two unique elements           | O(n) | O(1)  | Two passes (or one), two variables |
| Prefix XOR build              | O(n) | O(n)  | Store prefix array                 |
| Range XOR query (with prefix) | O(1) | O(1)  | Just two lookups and one XOR       |
| XOR of [0..n] (formula)       | O(1) | O(1)  | Uses the 4-cycle pattern           |

---

## 11. Practice Problems

### 🟢 Easy

| #   | Problem                   | LC # | Hint                                            |
| --- | ------------------------- | ---- | ----------------------------------------------- |
| 1   | Single Number             | 136  | XOR everything — pairs cancel                   |
| 2   | Missing Number            | 268  | XOR array with full range [0..n]                |
| 3   | Number Complement         | 476  | XOR with a mask of all 1s of same bit length    |
| 4   | Find the Difference       | 389  | XOR both strings — the extra character survives |
| 5   | XOR Operation in an Array | 1486 | Build XOR array then XOR a subrange             |

### 🟡 Medium

| #   | Problem                                              | LC # | Hint                                                                     |
| --- | ---------------------------------------------------- | ---- | ------------------------------------------------------------------------ |
| 6   | Single Number III                                    | 260  | XOR all → isolate differing bit → split into two groups                  |
| 7   | Total Hamming Distance                               | 477  | Count 1s at each bit position across all numbers                         |
| 8   | Decode XOR'd Array                                   | 1720 | Recover original array using XOR reverse: `a[i] = a[i-1] ^ encoded[i-1]` |
| 9   | Maximum XOR of Two Numbers                           | 421  | Use bit-by-bit greedy with a prefix trie                                 |
| 10  | Minimum XOR Sum of Two Arrays                        | 1879 | DP with bitmask — XOR cost minimization                                  |
| 11  | Find XOR Sum of All Pairs Bitwise AND                | 1835 | Distribute XOR over AND using algebra                                    |
| 12  | Count Triplets That Can Form Two Arrays of Equal XOR | 1442 | If `XOR(i..k) = 0`, any split works — count pairs using prefix XOR       |

### 🔴 Hard

| #   | Problem                                                   | LC # | Hint                                                    |
| --- | --------------------------------------------------------- | ---- | ------------------------------------------------------- |
| 13  | Maximum XOR With an Element From Array                    | 1707 | Offline queries + sort + trie — process by element size |
| 14  | Minimum Number of Operations to Make Array XOR Equal to K | 2997 | Count differing bits between current XOR and k          |
| 15  | Xor Queries of a Subarray                                 | 1310 | Classic prefix XOR — warm-up for harder range problems  |

---

## 12. Quick Revision (60 sec)

### Key Idea

XOR cancels pairs. Unique/missing elements survive. Use it when you need O(1) space with pairs.

### The 4 Properties to Memorize

```
a ^ a = 0      → self-cancels
a ^ 0 = a      → identity
commutative    → order doesn't matter
associative    → grouping doesn't matter
```

### Core Templates

```go
// Find single unique (pairs cancel)
result := 0
for _, n := range nums { result ^= n }

// Find missing in [0..n]
result := len(nums)
for i, n := range nums { result ^= i ^ n }

// Isolate rightmost set bit
diffBit := xorAll & (-xorAll)
```

### When to Use

- "Find the unique/single number" → XOR all
- "Find missing number in range" → XOR with full range
- "Two unique numbers" → XOR all, isolate set bit, split groups
- O(1) space required → think XOR first
