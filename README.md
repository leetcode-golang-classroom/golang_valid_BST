# golang_valid_BST

Given the `root` of a binary tree, *determine if it is a valid binary search tree (BST)*.

A **valid BST** is defined as follows:

- The left subtree of a node contains only nodes with keys **less than** the node's key.
- The right subtree of a node contains only nodes with keys **greater than** the node's key.
- Both the left and right subtrees must also be binary search trees.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/12/01/tree1.jpg](https://assets.leetcode.com/uploads/2020/12/01/tree1.jpg)

```
Input: root = [2,1,3]
Output: true

```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/12/01/tree2.jpg](https://assets.leetcode.com/uploads/2020/12/01/tree2.jpg)

```
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.

```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 104]`.
- `231 <= Node.val <= 231 - 1`

## 解析

題目給定一個二元樹的根結點 root

想要判斷這個由 root 所形成的樹是不是一棵二元搜尋樹

如果 root 所形成的樹是一個二元搜尋樹有以下特性

1. root.Left 所形成的樹內所有結點值都小於 root.Val
2. root.Right 所形成的樹內所有結點值都大於 root.Val
3. root.Left 與 root.Right 所形成的樹都是二元搜尋樹

最直覺的想法

每次遇到一個結點

檢查這個結點的左子樹有沒有都小於這個 root 結點

檢查這個結點的右子樹有沒有都大於這個 root 結點 

這樣每次都要走訪 n個 所以時間複雜度 是 O($n^2$)

如果不想要每次都要走訪所有走過的結點則必須思考以下特性

對一個二元搜尋樹的某個結點 node

node.left 所形成樹的上界是 node.Val ，因為所有 node.Left 所形成結點都必須小於 node.Val

所以對於 node.Left 所形成樹 只要檢查是否都有小於 node.Val 即可

同樣的，node.Right 所形成樹的下界是 node.Val ，因為所有 node.Right 所形成結點都必須大於 node.Val

所以對於 node.Right 所形成樹 只要檢查是否都有大於 node.Val 即可

![](https://i.imgur.com/gM2lJ9d.png)

一開始的根結點 沒有上下界相當於 上界是 infinity , 下界是 -infinity

每次只要檢查目前結點有沒有介於上下界

每次往左結點走更新上界是目前結點

每次往右結點走更新下界是目前結點

## 程式碼

```go
package sol

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return CheckBST(root, math.MaxFloat64, -math.MaxFloat64)
}

func CheckBST(root *TreeNode, upperBound float64, lowerBound float64) bool {
	if root == nil {
		return true
	}
	cur := float64(root.Val)
	if cur <= lowerBound || cur >= upperBound {
		return false
	}
	return CheckBST(root.Left, cur, lowerBound) && CheckBST(root.Right, upperBound, cur)
}
```

## 困難點

1. 理解二元搜尋樹的定義
2. 透過理解二元搜尋樹的定義來減少檢查的範圍

## Solve Point

- [x]  Understand what problem need to be solve
- [x]  Analysis Complexity