# 🔄 DSA Pattern: Reversal of Linked List (In-place)

> **Mentor's Note:** This is one of the most asked patterns in FAANG interviews. Master it once, and you'll handle 90% of linked list problems with confidence.

---

## 1. Pattern Overview

### What is this pattern?

Reversing a linked list means changing the direction of all pointers — instead of each node pointing **forward**, you make it point **backward** — all without using extra memory.

```
Before: 1 → 2 → 3 → 4 → 5 → nil
After:  5 → 4 → 3 → 2 → 1 → nil
```

### Why is it used? (Intuition)

Many linked list problems require you to **look back** or **compare from both ends**. Since linked lists only go forward, the only way to "go back" is to reverse the pointers. This pattern enables:

- Checking palindromes
- Comparing two halves
- Reversing specific segments (k-groups, sub-ranges)

### Real-World Analogy

Imagine a **one-way street** with cars lined up. You want to reverse the direction of traffic. You have to redirect each car one by one — you can't just flip the whole street. That's exactly what in-place reversal does: redirect each pointer, one node at a time.

### Time Complexity Intuition

You visit **every node exactly once** to flip its pointer → **O(n)**.  
You only use 3 pointers (prev, curr, next) regardless of list size → **O(1) space**.

---

## 2. Core Idea (MOST IMPORTANT)

### The 2 Key Insights

**Insight 1: You need 3 pointers**
To reverse a pointer without losing your place, you always need:

- `prev` → the node behind (where current should now point)
- `curr` → the node you're processing right now
- `next` → save the next node _before_ breaking the link

**Insight 2: The order of operations matters**

```
Save next → Flip pointer → Move prev forward → Move curr forward
```

If you flip the pointer before saving `next`, you lose the rest of the list. **Always save `next` first.**

### The Golden Lines (memorize this rhythm):

```go
next = curr.Next       // 1. Save next
curr.Next = prev       // 2. Flip the pointer
prev = curr            // 3. Move prev forward
curr = next            // 4. Move curr forward
```

---

## 3. When to Use This Pattern

| Signal in Question                   | What it Means                                |
| ------------------------------------ | -------------------------------------------- |
| "Reverse a linked list"              | Direct reversal                              |
| "Reverse between position i and j"   | Partial reversal                             |
| "Check if linked list is palindrome" | Reverse second half, compare                 |
| "Reverse in groups of k"             | K-group reversal                             |
| "Rotate list by k places"            | Involves finding new head via reversal logic |
| "Reorder list (1→n→2→n-1...)"        | Split + reverse + merge                      |

---

## 4. Types / Variants of This Pattern

### Type 1: Full Reversal

Reverse the entire list from head to tail.

```
1 → 2 → 3 → 4 → 5  →→→  5 → 4 → 3 → 2 → 1
```

### Type 2: Partial Reversal (Between positions m and n)

Only reverse a segment of the list, keeping the rest intact.

```
1 → 2 → 3 → 4 → 5  (reverse positions 2 to 4)
→→→  1 → 4 → 3 → 2 → 5
```

> Requires careful handling of the node _before_ the segment and the node _after_.

### Type 3: K-Group Reversal

Reverse every group of k nodes.

```
1 → 2 → 3 → 4 → 5  (k=2)
→→→  2 → 1 → 4 → 3 → 5
```

> Use recursion or iteration with careful group counting.

### Type 4: Half-List Reversal (Palindrome Check)

Find the middle using slow/fast pointers, reverse the second half, compare with first half.

---

## 5. Data Structures Used

| Data Structure                    | Why Used                                      | Key Properties                         |
| --------------------------------- | --------------------------------------------- | -------------------------------------- |
| **Linked List Node**              | The core structure being manipulated          | Each node has `Val` and `Next` pointer |
| **3 Pointers (prev, curr, next)** | To flip direction without losing nodes        | O(1) space, in-place                   |
| **Dummy Node**                    | Simplifies edge cases (insertion before head) | Avoids nil checks at boundaries        |
| **Slow/Fast Pointers**            | To find the middle of the list                | Used in palindrome / split problems    |

> **Why no extra array or stack?** You _could_ use a stack to reverse (push all, pop all), but that's O(n) space. In-place pointer manipulation is the expected interview approach.

---

## 6. Core Templates (VERY IMPORTANT)

### Template 1: Full List Reversal

```go
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode  // starts as nil (new tail points to nil)
    curr := head

    for curr != nil {
        next := curr.Next   // 1. Save next before breaking link
        curr.Next = prev    // 2. Flip the pointer
        prev = curr         // 3. Move prev forward
        curr = next         // 4. Move curr forward
    }

    return prev  // prev is the new head when curr == nil
}
```

### Template 2: Partial Reversal (between left and right positions)

```go
func reverseBetween(head *ListNode, left, right int) *ListNode {
    dummy := &ListNode{Next: head}  // dummy prevents nil edge cases
    prev := dummy

    // Step 1: Walk to the node just before 'left'
    for i := 1; i < left; i++ {
        prev = prev.Next
    }

    curr := prev.Next    // first node of the sublist to reverse
    var tail *ListNode   // will be tracked if needed

    // Step 2: Reverse (right - left) times
    for i := 0; i < right-left; i++ {
        next := curr.Next
        curr.Next = next.Next   // disconnect next from its place
        next.Next = prev.Next   // plug next at the front of reversed segment
        prev.Next = next        // update prev's next to new front
    }

    return dummy.Next
}
```

### Template 3: K-Group Reversal

```go
func reverseKGroup(head *ListNode, k int) *ListNode {
    // Check if k nodes exist
    node := head
    for i := 0; i < k; i++ {
        if node == nil {
            return head  // fewer than k nodes left, don't reverse
        }
        node = node.Next
    }

    // Reverse k nodes
    var prev *ListNode
    curr := head
    for i := 0; i < k; i++ {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    // head is now the tail of this group — connect to next group
    head.Next = reverseKGroup(curr, k)

    return prev  // prev is the new head of this group
}
```

---

## 7. Step-by-Step Dry Run

**Problem:** Reverse `1 → 2 → 3 → 4 → 5`

```
Initial:  prev=nil, curr=1

Step 1:
  next = 2
  1.Next = nil  (flip)
  prev = 1, curr = 2
  List so far: nil ← 1    2 → 3 → 4 → 5

Step 2:
  next = 3
  2.Next = 1  (flip)
  prev = 2, curr = 3
  List so far: nil ← 1 ← 2    3 → 4 → 5

Step 3:
  next = 4
  3.Next = 2  (flip)
  prev = 3, curr = 4
  List so far: nil ← 1 ← 2 ← 3    4 → 5

Step 4:
  next = 5
  4.Next = 3  (flip)
  prev = 4, curr = 5

Step 5:
  next = nil
  5.Next = 4  (flip)
  prev = 5, curr = nil

curr == nil → STOP
Return prev = 5

Result: 5 → 4 → 3 → 2 → 1 → nil ✅
```

---

## 8. Must-Know Problems (Deep Explanation)

### Problem 1: Reverse Linked List — LeetCode #206 (Easy)

**Intuition:** Classic pointer flip. Base case for all reversal problems.

**Approach:** Use 3 pointers (prev, curr, next). Iterate through the list once, flipping each pointer.

```go
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}
```

---

### Problem 2: Reverse Linked List II — LeetCode #92 (Medium)

**Intuition:** You don't reverse the whole list. Only reverse a window between two positions. The key challenge: re-attach the reversed segment to the rest.

**Approach:**

1. Walk to node just _before_ position `left` using dummy + prev pointer
2. Use the "insert at front" trick to reverse the segment in-place
3. No need to track the tail separately

```go
func reverseBetween(head *ListNode, left, right int) *ListNode {
    dummy := &ListNode{Next: head}
    prev := dummy

    for i := 1; i < left; i++ {
        prev = prev.Next
    }

    curr := prev.Next
    for i := 0; i < right-left; i++ {
        next := curr.Next
        curr.Next = next.Next
        next.Next = prev.Next
        prev.Next = next
    }

    return dummy.Next
}
```

---

### Problem 3: Reverse Nodes in K-Group — LeetCode #25 (Hard)

**Intuition:** Reverse in chunks. If the remaining nodes are fewer than k, leave them as-is. Recursion makes it clean.

**Approach:**

1. Check if k nodes exist ahead — if not, return head as-is
2. Reverse exactly k nodes using the standard 3-pointer technique
3. Recursively call for the remaining list and attach it to the current tail (old head)

```go
func reverseKGroup(head *ListNode, k int) *ListNode {
    node := head
    for i := 0; i < k; i++ {
        if node == nil {
            return head
        }
        node = node.Next
    }

    var prev *ListNode
    curr := head
    for i := 0; i < k; i++ {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    head.Next = reverseKGroup(curr, k)
    return prev
}
```

---

## 9. Common Mistakes / Gotchas

1. **Forgetting to save `next` before flipping**

   > `curr.Next = prev` destroys the path forward. Always do `next := curr.Next` first.

2. **Returning `curr` instead of `prev`**

   > When the loop ends, `curr == nil`. The new head is `prev`. Classic mistake under pressure.

3. **Not using a dummy node for partial reversal**

   > If `left == 1`, you're reversing from the head. Without a dummy, you'll have nil pointer issues. Always use `dummy.Next = head`.

4. **Off-by-one in walking to position `left`**

   > You want to stop at the node _before_ `left`, not at `left`. Loop condition: `i < left`, not `i <= left`.

5. **Infinite loop in K-group when not checking remaining nodes**

   > Always check that k nodes exist before reversing. If you skip this, you'll reverse partial groups incorrectly.

6. **Losing connection between reversed segment and the rest of the list**

   > In partial reversal, the old `curr` (tail of reversed segment) must connect to `node` after position `right`. Template 2 handles this automatically.

7. **Mutating input head without returning new head**
   > After full reversal, the original head is now the tail. Always return `prev` (the new head).

---

## 10. Time & Space Complexity

| Operation                       | Time | Space  | Why                                   |
| ------------------------------- | ---- | ------ | ------------------------------------- |
| Full Reversal                   | O(n) | O(1)   | Visit each node once, 3 pointers only |
| Partial Reversal (m to n)       | O(n) | O(1)   | Walk to position + reverse segment    |
| K-Group Reversal (iterative)    | O(n) | O(1)   | Each node processed exactly once      |
| K-Group Reversal (recursive)    | O(n) | O(n/k) | Recursion stack depth = n/k frames    |
| Palindrome Check (reverse half) | O(n) | O(1)   | Find mid + reverse + compare          |

---

## 11. Practice Problems

### 🟢 Easy

| #   | Problem                   | LeetCode | Hint                                                  |
| --- | ------------------------- | -------- | ----------------------------------------------------- |
| 1   | Reverse Linked List       | #206     | Classic 3-pointer template                            |
| 2   | Palindrome Linked List    | #234     | Find mid with slow/fast, reverse second half, compare |
| 3   | Middle of the Linked List | #876     | Slow/fast pointer — prerequisite for many problems    |

---

### 🟡 Medium

| #   | Problem                             | LeetCode | Hint                                                  |
| --- | ----------------------------------- | -------- | ----------------------------------------------------- |
| 1   | Reverse Linked List II              | #92      | Use dummy node + in-place insert-at-front trick       |
| 2   | Reorder List                        | #143     | Find mid → reverse second half → merge alternating    |
| 3   | Rotate List                         | #61      | Find new tail at (n - k % n), break and reconnect     |
| 4   | Swap Nodes in Pairs                 | #24      | K-group reversal with k=2                             |
| 5   | Remove Nth Node From End            | #19      | Two pointers n apart, then delete                     |
| 6   | Odd Even Linked List                | #328     | Re-link odd/even indices — uses pointer rearrangement |
| 7   | Reverse Nodes in Even Length Groups | #2074    | Find group length, conditionally reverse              |

---

### 🔴 Hard

| #   | Problem                                               | LeetCode | Hint                                                    |
| --- | ----------------------------------------------------- | -------- | ------------------------------------------------------- |
| 1   | Reverse Nodes in K-Group                              | #25      | Check k nodes exist, reverse chunk, recurse             |
| 2   | Reverse Linked List in Groups of K (Even-Length Only) | #2074    | Conditional K-group reversal with length check          |
| 3   | Sort List                                             | #148     | Merge sort on linked list — uses reversal + merge logic |

---

## 12. ⚡ Quick Revision (60 seconds)

### Key Idea

> Flip each node's pointer one at a time using 3 pointers: `prev`, `curr`, `next`. Always save `next` before flipping. Return `prev` as the new head.

### Core Template

```go
var prev *ListNode
curr := head
for curr != nil {
    next := curr.Next   // save
    curr.Next = prev    // flip
    prev = curr         // advance prev
    curr = next         // advance curr
}
return prev
```

### When to Use

| Keyword                | Pattern                                  |
| ---------------------- | ---------------------------------------- |
| "Reverse list"         | Full reversal template                   |
| "Reverse from i to j"  | Partial reversal + dummy node            |
| "K-group"              | K-group template + check k nodes first   |
| "Palindrome"           | Reverse second half, compare             |
| "Reorder / interleave" | Split at mid, reverse second half, merge |

### 3 Things to Never Forget

1. ✅ Save `next` **before** flipping the pointer
2. ✅ Return `prev`, not `curr`
3. ✅ Use a **dummy node** for partial reversal to avoid nil head issues
