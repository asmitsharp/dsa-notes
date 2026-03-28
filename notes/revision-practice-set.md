# Revision Practice Set

Focused topics:

- [Prefix Sum](/Users/asmitsingh/Desktop/dsa-notes/notes/prefix-sum.md)
- [Sliding Window](/Users/asmitsingh/Desktop/dsa-notes/notes/sliding-window.md)
- [Two Pointers](/Users/asmitsingh/Desktop/dsa-notes/notes/two-pointer.md)
- [Overlapping Intervals](/Users/asmitsingh/Desktop/dsa-notes/notes/overlapping_intervals.md)

---

## How To Use This Sheet

- Do not solve these randomly. Follow the order.
- Round 1 is for pattern recognition.
- Round 2 is for core interview problems.
- Round 3 is for mixed timed revision.
- After every problem, write 4 things:
  - Trigger: why this pattern applies
  - Invariant: what stays true while solving
  - Time complexity
  - One mistake you made or almost made

Recommended cadence:

- Day 1: Prefix Sum + Sliding Window
- Day 2: Two Pointers + Overlapping Intervals
- Day 3: Mixed sets only
- Day 4+: Redo only the problems you could not explain cleanly

---

## Round 1: Pattern Recognition

These should feel almost automatic.

| Topic | Problem | Main thing to recognize |
| --- | --- | --- |
| Prefix Sum | Running Sum of 1d Array | build cumulative sum |
| Prefix Sum | Range Sum Query - Immutable | preprocess once, answer many range queries |
| Prefix Sum | Find Pivot Index | left sum vs right sum from total/prefix |
| Sliding Window | Maximum Average Subarray I | fixed-size window |
| Sliding Window | Longest Substring Without Repeating Characters | dynamic window + shrink on invalid state |
| Two Pointers | Reverse String | opposite-direction pointers |
| Two Pointers | Valid Palindrome | skip and compare from both ends |
| Two Pointers | Two Sum II - Input Array Is Sorted | sorted array + move based on sum |
| Intervals | Merge Intervals | sort by start, merge greedily |
| Intervals | Meeting Rooms | overlap detection after sorting |

Rule for Round 1:

- If you cannot identify the pattern in 30 seconds, revisit the note before coding.

---

## Round 2: Core Practice By Topic

### 1. Prefix Sum

Order matters here. Do them in sequence.

| Order | Problem | Focus |
| --- | --- | --- |
| 1 | Running Sum of 1d Array | warm-up, build prefix array |
| 2 | Range Sum Query - Immutable | sentinel prefix, range query formula |
| 3 | Find Pivot Index | left sum and right sum relation |
| 4 | Product of Array Except Self | prefix + suffix idea |
| 5 | Subarray Sum Equals K | prefix sum + hashmap frequency |
| 6 | Contiguous Array | prefix difference mapped to first index |
| 7 | Subarray Sums Divisible by K | modulo buckets with prefix sum |
| 8 | Continuous Subarray Sum | prefix modulo and earliest occurrence |

What you must be able to say out loud:

- "Prefix sum is useful when I need information about many contiguous subarrays."
- "For count problems, I usually need prefix sum + hashmap."
- "The most common bug is wrong initialization like `map[0] = 1` or bad modulo handling."

---

### 2. Sliding Window

| Order | Problem | Focus |
| --- | --- | --- |
| 1 | Maximum Average Subarray I | fixed-size window |
| 2 | Minimum Size Subarray Sum | shrink while valid |
| 3 | Longest Substring Without Repeating Characters | frequency map + dynamic window |
| 4 | Permutation in String | fixed-size frequency window |
| 5 | Find All Anagrams in a String | same pattern as permutation, but collect all starts |
| 6 | Fruits Into Baskets | at most 2 distinct elements |
| 7 | Longest Repeating Character Replacement | window validity based on max frequency |
| 8 | Max Consecutive Ones III | at most `k` bad elements |
| 9 | Minimum Window Substring | hardest classic shrinking window |

What you must be able to say out loud:

- "Sliding window only works when the problem is about a contiguous segment."
- "Fixed window means size is constant. Dynamic window means validity decides shrinking."
- "The hardest part is deciding when the window is invalid and what removing `left` changes."

---

### 3. Two Pointers

| Order | Problem | Focus |
| --- | --- | --- |
| 1 | Reverse String | basic opposite-direction movement |
| 2 | Valid Palindrome | compare while skipping |
| 3 | Two Sum II - Input Array Is Sorted | sorted array, deterministic movement |
| 4 | Remove Duplicates from Sorted Array | slow/fast write pointer |
| 5 | Move Zeroes | stable compaction using write pointer |
| 6 | Squares of a Sorted Array | compare extremes, fill from back |
| 7 | Container With Most Water | move smaller wall only |
| 8 | 3Sum | sort + outer loop + inner two pointers |
| 9 | Sort Colors | Dutch national flag style pointers |

Stretch problem:

- Trapping Rain Water

What you must be able to say out loud:

- "Two pointers is not one pattern, it has variants: opposite-direction, slow/fast, and multi-array."
- "Pointer movement must be justified. If I cannot explain why a pointer moves, I do not understand the solution."
- "Sorting often makes the pointer movement valid."

---

### 4. Overlapping Intervals

| Order | Problem | Focus |
| --- | --- | --- |
| 1 | Merge Intervals | sort by start, merge into result |
| 2 | Insert Interval | three phases: before, merge, after |
| 3 | Interval List Intersections | overlap generation between two sorted lists |
| 4 | Meeting Rooms | detect any overlap |
| 5 | Meeting Rooms II | min-heap of ending times |
| 6 | Non-overlapping Intervals | greedy by end time |
| 7 | Minimum Number of Arrows to Burst Balloons | greedy by end time again |
| 8 | Employee Free Time | merge schedule first, then find gaps |

What you must be able to say out loud:

- "For interval problems, sorting is usually the real first step."
- "Sort by start for merging. Sort by end for greedy keep/remove decisions."
- "I must clarify whether touching boundaries count as overlap."

---

## Round 3: Mixed Timed Revision Sets

### Set A

Target: 75 to 90 minutes

| Order | Problem | Topic |
| --- | --- | --- |
| 1 | Find Pivot Index | Prefix Sum |
| 2 | Minimum Size Subarray Sum | Sliding Window |
| 3 | Two Sum II - Input Array Is Sorted | Two Pointers |
| 4 | Merge Intervals | Intervals |
| 5 | Subarray Sum Equals K | Prefix Sum |
| 6 | Longest Substring Without Repeating Characters | Sliding Window |

### Set B

Target: 90 to 120 minutes

| Order | Problem | Topic |
| --- | --- | --- |
| 1 | Product of Array Except Self | Prefix Sum |
| 2 | Fruits Into Baskets | Sliding Window |
| 3 | Squares of a Sorted Array | Two Pointers |
| 4 | Meeting Rooms II | Intervals |
| 5 | Subarray Sums Divisible by K | Prefix Sum |
| 6 | Longest Repeating Character Replacement | Sliding Window |

### Set C

Target: 120 minutes

| Order | Problem | Topic |
| --- | --- | --- |
| 1 | Continuous Subarray Sum | Prefix Sum |
| 2 | Minimum Window Substring | Sliding Window |
| 3 | 3Sum | Two Pointers |
| 4 | Non-overlapping Intervals | Intervals |
| 5 | Interval List Intersections | Intervals |
| 6 | Max Consecutive Ones III | Sliding Window |

---

## Final Confidence Test

If you can solve at least 4 out of 5 cleanly, your revision is working.

| Problem | Topic | What it tests |
| --- | --- | --- |
| Subarray Sum Equals K | Prefix Sum | hashmap initialization and counting logic |
| Minimum Window Substring | Sliding Window | hardest common dynamic window |
| 3Sum | Two Pointers | sorting, duplicate handling, inner pointer logic |
| Meeting Rooms II | Intervals | heap + overlap scheduling |
| Non-overlapping Intervals | Intervals | greedy proof and sort-by-end intuition |

---

## Self-Review Checklist

After each session, mark the ones that were weak:

- [ ] I confused sliding window with two pointers without knowing the exact invariant.
- [ ] I missed that the problem required contiguity.
- [ ] I wrote prefix sum correctly but initialized the hashmap incorrectly.
- [ ] I forgot the extra `0` prefix / sentinel trick.
- [ ] I moved a pointer without being able to justify why.
- [ ] I forgot to sort intervals before processing.
- [ ] I used sort-by-start when the greedy solution actually needed sort-by-end.
- [ ] I had an off-by-one bug in window length or range sum.
- [ ] I handled duplicates badly in `3Sum`.
- [ ] I understood the code after seeing it, but could not derive it myself.

---

## Minimum Goal For This Revision Cycle

Before moving on, you should be able to do these from memory:

- Prefix Sum: `Find Pivot Index`, `Subarray Sum Equals K`
- Sliding Window: `Longest Substring Without Repeating Characters`, `Minimum Size Subarray Sum`
- Two Pointers: `Two Sum II`, `Squares of a Sorted Array`
- Intervals: `Merge Intervals`, `Non-overlapping Intervals`

If these 8 feel clean, your base is strong enough to build on.
