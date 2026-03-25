# 🌳 Trees — Complete Interview DSA Notes

> One file. Everything you need. Nothing you don't.

---

## 1. What Is a Tree?

A **tree** is a connected, acyclic graph with a hierarchy. It has:

- A **root** (top node, no parent)
- **Parent / Child** relationships
- **Leaves** (nodes with no children)

```
        1          ← root
       / \
      2   3        ← internal nodes
     / \
    4   5          ← leaves
```

**Key vocabulary:**
| Term | Meaning |
|---|---|
| Height of node | Longest path from node to a leaf |
| Depth of node | Distance from root to node |
| Height of tree | Height of root |
| Diameter | Longest path between any two nodes (may not pass root) |
| Balanced tree | Height difference of left/right subtrees ≤ 1 for every node |
| Complete tree | All levels full except last, filled left to right |
| Full tree | Every node has 0 or 2 children |
| Perfect tree | All internal nodes have 2 children, all leaves at same level |

---

## 2. Tree Node Structure (Go)

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
```

---

## 3. Core Traversals (MUST KNOW)

### The Big 4 Traversals

```
Tree:      1
          / \
         2   3
        / \
       4   5
```

| Traversal       | Order               | Result        | Memory Hook                 |
| --------------- | ------------------- | ------------- | --------------------------- |
| **Inorder**     | Left → Root → Right | 4, 2, 5, 1, 3 | "In order" = sorted for BST |
| **Preorder**    | Root → Left → Right | 1, 2, 4, 5, 3 | Root is **pre**-first       |
| **Postorder**   | Left → Right → Root | 4, 5, 2, 3, 1 | Root is **post**-last       |
| **Level Order** | Level by level      | 1, 2, 3, 4, 5 | BFS                         |

---

### Inorder (Left → Root → Right)

```go
func inorder(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    result := []int{}
    result = append(result, inorder(root.Left)...)
    result = append(result, root.Val)
    result = append(result, inorder(root.Right)...)
    return result
}
```

**Iterative Inorder** (stack-based — common interview ask):

```go
func inorderIterative(root *TreeNode) []int {
    result := []int{}
    stack := []*TreeNode{}
    curr := root

    for curr != nil || len(stack) > 0 {
        // go as left as possible
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        // backtrack
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, curr.Val)
        curr = curr.Right
    }
    return result
}
```

---

### Preorder (Root → Left → Right)

```go
func preorder(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    result := []int{root.Val}
    result = append(result, preorder(root.Left)...)
    result = append(result, preorder(root.Right)...)
    return result
}
```

---

### Postorder (Left → Right → Root)

```go
func postorder(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    result := []int{}
    result = append(result, postorder(root.Left)...)
    result = append(result, postorder(root.Right)...)
    result = append(result, root.Val)
    return result
}
```

---

### Level Order (BFS)

```go
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    result := [][]int{}
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        size := len(queue)
        level := []int{}
        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]
            level = append(level, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        result = append(result, level)
    }
    return result
}
```

---

## 4. The DFS Recursive Template (Golden Template)

> 90% of tree problems use this shape. Master this.

```go
func solve(root *TreeNode) ReturnType {
    // BASE CASE
    if root == nil {
        return baseValue
    }

    // RECURSE on children
    left := solve(root.Left)
    right := solve(root.Right)

    // COMBINE / compute at current node
    return combine(left, right, root.Val)
}
```

**Examples of what "combine" means:**

- Max depth → `1 + max(left, right)`
- Path sum → `root.Val + left + right`
- Is valid BST → `left && right && inRange(root.Val)`
- Diameter → `max(globalMax, left+right)` + return `1 + max(left, right)`

---

## 5. Binary Search Tree (BST)

### BST Property

```
For every node:
  ALL values in LEFT subtree < node.Val
  ALL values in RIGHT subtree > node.Val
```

```
      5
     / \
    3   7
   / \ / \
  2  4 6  8
```

### Key BST Facts

- **Inorder traversal of BST = sorted array** ← very frequently tested
- Search, insert, delete: **O(H)** where H = height
  - Balanced BST: H = O(log N)
  - Skewed BST (worst case): H = O(N)

### BST Search

```go
func search(root *TreeNode, target int) *TreeNode {
    if root == nil || root.Val == target {
        return root
    }
    if target < root.Val {
        return search(root.Left, target)
    }
    return search(root.Right, target)
}
```

### BST Insert

```go
func insert(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    if val < root.Val {
        root.Left = insert(root.Left, val)
    } else {
        root.Right = insert(root.Right, val)
    }
    return root
}
```

### BST Delete

```go
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else {
        // node found — 3 cases:
        if root.Left == nil {
            return root.Right  // case 1: no left child
        }
        if root.Right == nil {
            return root.Left   // case 2: no right child
        }
        // case 3: two children → replace with inorder successor (min of right)
        minNode := findMin(root.Right)
        root.Val = minNode.Val
        root.Right = deleteNode(root.Right, minNode.Val)
    }
    return root
}

func findMin(node *TreeNode) *TreeNode {
    for node.Left != nil {
        node = node.Left
    }
    return node
}
```

### Validate BST

```go
// Pass min/max bounds down — not just checking parent
func isValidBST(root *TreeNode) bool {
    return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int) bool {
    if node == nil {
        return true
    }
    if node.Val <= min || node.Val >= max {
        return false
    }
    return validate(node.Left, min, node.Val) &&
           validate(node.Right, node.Val, max)
}
```

> ⚠️ Common mistake: comparing only with parent. Must track full range.

---

## 6. Tree Height, Depth, Diameter

### Max Depth / Height

```go
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
```

### Min Depth

```go
// NOT the same as maxDepth with min! Must handle one-sided nodes.
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil {
        return 1 + minDepth(root.Right)
    }
    if root.Right == nil {
        return 1 + minDepth(root.Left)
    }
    return 1 + min(minDepth(root.Left), minDepth(root.Right))
}
```

### Diameter of Binary Tree

```go
// Diameter = longest path between any two nodes
// At each node: left_height + right_height is the path through this node
var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
    diameter = 0
    height(root)
    return diameter
}

func height(node *TreeNode) int {
    if node == nil {
        return 0
    }
    left := height(node.Left)
    right := height(node.Right)
    diameter = max(diameter, left+right) // update global max
    return 1 + max(left, right)          // return height to parent
}
```

> 💡 Key insight: The answer might not pass through root. Use a global variable and update at every node.

---

## 7. Path Problems (Very Common in Interviews)

### Path Sum (Root to Leaf)

```go
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil {
        return root.Val == targetSum // at leaf
    }
    return hasPathSum(root.Left, targetSum-root.Val) ||
           hasPathSum(root.Right, targetSum-root.Val)
}
```

### Max Path Sum (Any Node to Any Node)

```go
// Classic hard problem — path can go up then down through any node
var maxSum int

func maxPathSum(root *TreeNode) int {
    maxSum = math.MinInt32
    gainFrom(root)
    return maxSum
}

func gainFrom(node *TreeNode) int {
    if node == nil {
        return 0
    }
    // only take positive contributions
    leftGain := max(0, gainFrom(node.Left))
    rightGain := max(0, gainFrom(node.Right))

    // update answer: path through this node
    maxSum = max(maxSum, node.Val+leftGain+rightGain)

    // return max single-branch gain to parent
    return node.Val + max(leftGain, rightGain)
}
```

---

## 8. Lowest Common Ancestor (LCA)

```go
// Works for both general tree and BST
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {
        return root // found one of them (or end of tree)
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    if left != nil && right != nil {
        return root // p and q are in different subtrees → root is LCA
    }
    if left != nil {
        return left  // both in left
    }
    return right     // both in right
}
```

**For BST specifically (faster):**

```go
func lcaBST(root, p, q *TreeNode) *TreeNode {
    if p.Val < root.Val && q.Val < root.Val {
        return lcaBST(root.Left, p, q)
    }
    if p.Val > root.Val && q.Val > root.Val {
        return lcaBST(root.Right, p, q)
    }
    return root // split point = LCA
}
```

---

## 9. Tree Construction

### Build Tree from Preorder + Inorder

```go
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{Val: preorder[0]}
    mid := 0
    for i, v := range inorder {
        if v == preorder[0] {
            mid = i
            break
        }
    }
    root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
    root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
    return root
}
```

> 💡 Key insight: preorder[0] is always the root. Find it in inorder to split left/right.

### Sorted Array → Balanced BST

```go
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    mid := len(nums) / 2
    root := &TreeNode{Val: nums[mid]}
    root.Left = sortedArrayToBST(nums[:mid])
    root.Right = sortedArrayToBST(nums[mid+1:])
    return root
}
```

---

## 10. Symmetric, Same Tree, Invert

### Check Symmetric Tree

```go
func isSymmetric(root *TreeNode) bool {
    return isMirror(root.Left, root.Right)
}
func isMirror(l, r *TreeNode) bool {
    if l == nil && r == nil { return true }
    if l == nil || r == nil { return false }
    return l.Val == r.Val &&
           isMirror(l.Left, r.Right) &&
           isMirror(l.Right, r.Left)
}
```

### Same Tree

```go
func isSameTree(p, q *TreeNode) bool {
    if p == nil && q == nil { return true }
    if p == nil || q == nil { return false }
    return p.Val == q.Val &&
           isSameTree(p.Left, q.Left) &&
           isSameTree(p.Right, q.Right)
}
```

### Invert Binary Tree

```go
func invertTree(root *TreeNode) *TreeNode {
    if root == nil { return nil }
    root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
    return root
}
```

---

## 11. Serialization & Deserialization

```go
// Serialize using preorder with "null" markers
func serialize(root *TreeNode) string {
    if root == nil {
        return "null,"
    }
    return strconv.Itoa(root.Val) + "," +
           serialize(root.Left) +
           serialize(root.Right)
}

func deserialize(data string) *TreeNode {
    parts := strings.Split(data, ",")
    idx := 0
    var build func() *TreeNode
    build = func() *TreeNode {
        if parts[idx] == "null" {
            idx++
            return nil
        }
        val, _ := strconv.Atoi(parts[idx])
        idx++
        node := &TreeNode{Val: val}
        node.Left = build()
        node.Right = build()
        return node
    }
    return build()
}
```

---

## 12. When to Use Which Tree Approach

| Signal                                 | Use This                                |
| -------------------------------------- | --------------------------------------- |
| "Sorted output / kth smallest"         | Inorder DFS on BST                      |
| "Level by level / min depth / nearest" | BFS (level order)                       |
| "Max depth / diameter / path sum"      | DFS with return value                   |
| "Any-to-any path"                      | DFS + global variable                   |
| "Common ancestor"                      | LCA pattern                             |
| "Validate structure"                   | DFS with bounds (min/max)               |
| "Build tree from traversals"           | Preorder + Inorder construction         |
| "Subtree problems"                     | Recursion returning state from children |

---

## 13. Common Mistakes / Gotchas

| #   | Mistake                                                 | Fix                                                            |
| --- | ------------------------------------------------------- | -------------------------------------------------------------- |
| 1   | Checking `left < root` only for BST validity            | Track full min/max range at every node                         |
| 2   | `minDepth` = just `min(left, right)` recursively        | Handle one-child nodes — don't count nil as a valid leaf       |
| 3   | Diameter only through root                              | Update global at every node, not just root                     |
| 4   | Missing null checks before accessing `.Left` / `.Right` | Always check `if root == nil` at the start                     |
| 5   | Forgetting BST inorder = sorted                         | Leverage this — don't sort again                               |
| 6   | Max path sum forgetting to take `max(0, childGain)`     | Negative subtrees should contribute 0, not drag sum down       |
| 7   | Confusing height and depth                              | Height = bottom-up (from leaves); Depth = top-down (from root) |

---

## 14. Complexity Reference

| Operation         | BST (Balanced) | BST (Skewed) | General Tree      |
| ----------------- | -------------- | ------------ | ----------------- |
| Search            | O(log N)       | O(N)         | O(N)              |
| Insert            | O(log N)       | O(N)         | O(1) with pointer |
| Delete            | O(log N)       | O(N)         | O(N)              |
| Traversal         | O(N)           | O(N)         | O(N)              |
| Space (recursion) | O(log N)       | O(N)         | O(H)              |

---

## 15. Practice Problems

### 🟢 Easy

| #   | Problem                      | LC # | Hint                              |
| --- | ---------------------------- | ---- | --------------------------------- |
| 1   | Maximum Depth of Binary Tree | 104  | `1 + max(left, right)`            |
| 2   | Invert Binary Tree           | 226  | Swap left and right at every node |
| 3   | Symmetric Tree               | 101  | Mirror check: l.left vs r.right   |
| 4   | Same Tree                    | 100  | Compare val + recurse both sides  |
| 5   | Path Sum                     | 112  | Subtract val, check at leaf       |
| 6   | Merge Two Binary Trees       | 617  | Add vals, recurse both            |

### 🟡 Medium

| #   | Problem                            | LC # | Hint                                |
| --- | ---------------------------------- | ---- | ----------------------------------- |
| 1   | Binary Tree Level Order Traversal  | 102  | BFS with size capture               |
| 2   | Validate Binary Search Tree        | 98   | Pass min/max range                  |
| 3   | Lowest Common Ancestor of BST      | 235  | Split point in BST                  |
| 4   | Lowest Common Ancestor of BT       | 236  | Return when found, merge left+right |
| 5   | Binary Tree Right Side View        | 199  | BFS, take last node per level       |
| 6   | Diameter of Binary Tree            | 543  | Height + global max update          |
| 7   | Kth Smallest in BST                | 230  | Inorder DFS, count to k             |
| 8   | Construct Tree from Pre+Inorder    | 105  | preorder[0] = root, split inorder   |
| 9   | Flatten Binary Tree to Linked List | 114  | Preorder → relink nodes             |
| 10  | Populating Next Right Pointers     | 116  | BFS level order                     |
| 11  | Count Good Nodes in Tree           | 1448 | DFS tracking max from root          |
| 12  | Path Sum II (all paths)            | 113  | Backtracking DFS                    |

### 🔴 Hard

| #   | Problem                               | LC # | Hint                                |
| --- | ------------------------------------- | ---- | ----------------------------------- |
| 1   | Binary Tree Maximum Path Sum          | 124  | DFS + global, return single branch  |
| 2   | Serialize and Deserialize Binary Tree | 297  | Preorder with null markers          |
| 3   | Binary Tree Cameras                   | 968  | Greedy postorder: 3 states per node |
| 4   | Recover Binary Search Tree            | 99   | Inorder — find 2 swapped nodes      |
| 5   | Vertical Order Traversal              | 987  | BFS + coordinate sorting            |

---

## 16. Quick Revision (60 sec)

```
TREE TRAVERSALS:
  Inorder   → Left, Root, Right  → BST gives sorted output
  Preorder  → Root, Left, Right  → Good for cloning/serializing
  Postorder → Left, Right, Root  → Good for deletion/bottom-up
  Level     → BFS with queue     → Shortest depth, level stats

GOLDEN RECURSIVE TEMPLATE:
  func solve(root) result {
      if root == nil { return base }
      left  := solve(root.Left)
      right := solve(root.Right)
      return combine(left, right, root.Val)
  }

BST RULES:
  Left < Node < Right (ALL descendants, not just children)
  Inorder = sorted  ← use this shortcut

GLOBAL STATE TRICK:
  When answer can be at any node (diameter, max path sum)
  → Use a global variable, update at every node
  → Return a different value (e.g., height) to parent

LCA:
  If both found on same side → recurse that side
  If split across sides → current root is LCA

MIN DEPTH TRAP:
  Don't just do min(left, right) — handle one-child nodes!
```

---

_Trees Notes | Interview-Ready | All-in-One_
