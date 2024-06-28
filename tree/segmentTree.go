package main

import (
	"fmt"
)

const N = 1e5 + 10

// 全局变量
var (
	sum  [N << 2]int64  // 存储区间和
	add  [N << 2]int64  // 存储延迟更新标记
	maxn [N << 2]int64  // 存储区间最大值
	tree [N << 2]Node   // 存储区间节点信息
)

// Node 表示线段树的节点
type Node struct {
	l, r int  // 区间的左端点和右端点
}

// mid 返回当前节点区间的中点
func (n Node) mid() int {
	return (n.l + n.r) >> 1
}

// pushup 更新父节点的区间和及区间最大值
func pushup(rt int) {
	sum[rt] = sum[rt<<1] + sum[rt<<1|1]
	maxn[rt] = max(maxn[rt<<1], maxn[rt<<1|1])
}

// pushdown 进行延迟标记的下推
func pushdown(rt, m int) {
	if add[rt] != 0 {
		add[rt<<1] += add[rt]
		add[rt<<1|1] += add[rt]
		sum[rt<<1] += add[rt] * int64(m-(m>>1))
		sum[rt<<1|1] += add[rt] * int64(m>>1)
		add[rt] = 0
	}
}

// build 构建线段树
func build(l, r, rt int, data []int64) {
	tree[rt].l = l
	tree[rt].r = r
	add[rt] = 0
	if l == r {
		sum[rt] = data[l-1]
		maxn[rt] = data[l-1]
		return
	}
	m := tree[rt].mid()
	build(l, m, rt<<1, data)
	build(m+1, r, rt<<1|1, data)
	pushup(rt)
}

// update 更新区间 [l, r] 的值
func update(c, l, r, rt int) {
	if tree[rt].l == l && tree[rt].r == r {
		add[rt] += int64(c)
		sum[rt] += int64(c) * int64(r-l+1)
		return
	}
	pushdown(rt, tree[rt].r-tree[rt].l+1)
	m := tree[rt].mid()
	if r <= m {
		update(c, l, r, rt<<1)
	} else if l > m {
		update(c, l, r, rt<<1|1)
	} else {
		update(c, l, m, rt<<1)
		update(c, m+1, r, rt<<1|1)
	}
	pushup(rt)
}

// query 查询区间 [l, r] 的和
func query(l, r, rt int) int64 {
	if l == tree[rt].l && r == tree[rt].r {
		return sum[rt]
	}
	pushdown(rt, tree[rt].r-tree[rt].l+1)
	m := tree[rt].mid()
	var res int64 = 0
	if r <= m {
		res += query(l, r, rt<<1)
	} else if l > m {
		res += query(l, r, rt<<1|1)
	} else {
		res += query(l, m, rt<<1)
		res += query(m+1, r, rt<<1|1)
	}
	return res
}

// queryMax 查询区间 [l, r] 的最大值
func queryMax(l, r, rt int) int64 {
	if l == tree[rt].l && r == tree[rt].r {
		return maxn[rt]
	}
	pushdown(rt, tree[rt].r-tree[rt].l+1)
	m := tree[rt].mid()
	var res int64 = 0
	if r <= m {
		res = max(res, queryMax(l, r, rt<<1))
	} else if l > m {
		res = max(res, queryMax(l, r, rt<<1|1))
	} else {
		res = max(res, queryMax(l, m, rt<<1))
		res = max(res, queryMax(m+1, r, rt<<1|1))
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
	build(1, n, 1, data)
	for i := 0; i < m; i++ {
		var a int
		fmt.Scan(&a)
		if a == 2 {
			var b, c int
			fmt.Scan(&b, &c)
			fmt.Println(query(b, c, 1))
		} else {
			var b, c, d int
			fmt.Scan(&b, &c, &d)
			update(d, b, c, 1)
		}
	}
}

/*

input:
5 3
1 2 3 4 5
1 1 3 2
2 1 5
2 2 4
预期output:
21
14

*/