package main

import (
 "fmt"
)

// Node 表示线段树的节点
type Node struct {
 l, r       int
 sum, add   int64
 maxn       int64
 left, right *Node
}

// mid 返回当前节点区间的中点
func (n *Node) mid() int {
 return (n.l + n.r) >> 1
}，

// pushup 更新父节点的区间和及区间最大值
func (n *Node) pushup() {
 n.sum = n.left.sum + n.right.sum
 n.maxn = max(n.left.maxn, n.right.maxn)
}

// pushdown 进行延迟标记的下推
func (n *Node) pushdown() {
 if n.add != 0 {
  n.left.add += n.add
  n.right.add += n.add
  n.left.sum += n.add * int64(n.left.r-n.left.l+1)
  n.right.sum += n.add * int64(n.right.r-n.right.l+1)
  n.add = 0
 }
}

// build 构建线段树
func build(l, r int, data []int64) *Node {
 node := &Node{l: l, r: r}
 if l == r {
  node.sum = data[l-1]
  node.maxn = data[l-1]
  return node
 }
 m := (l + r) >> 1
 node.left = build(l, m, data)
 node.right = build(m+1, r, data)
 node.pushup()
 return node
}

// update 更新区间 [l, r] 的值
func update(node *Node, c, l, r int) {
 if node.l == l && node.r == r {
  node.add += int64(c)
  node.sum += int64(c) * int64(r-l+1)
  return
 }
 node.pushdown()
 m := node.mid()
 if r <= m {
  update(node.left, c, l, r)
 } else if l > m {
  update(node.right, c, l, r)
 } else {
  update(node.left, c, l, m)
  update(node.right, c, m+1, r)
 }
 node.pushup()
}

// query 查询区间 [l, r] 的和
func query(node *Node, l, r int) int64 {
 if l == node.l && r == node.r {
  return node.sum
 }
 node.pushdown()
 m := node.mid()
 var res int64 = 0
 if r <= m {
  res += query(node.left, l, r)
 } else if l > m {
  res += query(node.right, l, r)
 } else {
  res += query(node.left, l, m)
  res += query(node.right, m+1, r)
 }
 return res
}

// queryMax 查询区间 [l, r] 的最大值
func queryMax(node *Node, l, r int) int64 {
 if l == node.l && r == node.r {
  return node.maxn
 }
 node.pushdown()
 m := node.mid()
 var res int64 = 0
 if r <= m {
  res = max(res, queryMax(node.left, l, r))
 } else if l > m {
  res = max(res, queryMax(node.right, l, r))
 } else {
  res = max(res, queryMax(node.left, l, m))
  res = max(res, queryMax(node.right, m+1, r))
 }
 return res
}

// max 返回两个 int64 值中的最大值
func max(a, b int64) int64 {
 if a > b {
  return a
 }
 return b
}

func main() {
 var n, m int
 fmt.Scan(&n, &m)
 data := make([]int64, n)
 for i := 0; i < n; i++ {
  fmt.Scan(&data[i])
 }
 root := build(1, n, data)
 for i := 0; i < m; i++ {
  var a int
  fmt.Scan(&a)
  if a == 2 {
   var b, c int
   fmt.Scan(&b, &c)
   fmt.Println(query(root, b, c))
  } else {
   var b, c, d int
   fmt.Scan(&b, &c, &d)
   update(root, d, b, c)
  }
 }
}