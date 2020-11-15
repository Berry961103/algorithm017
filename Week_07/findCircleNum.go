package week07

func findCircleNumDFS(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	count, visited := 0, make([]int, len(M))

	var dfs func(int)
	dfs = func(i int) {
		for j := 0; j < len(M); j++ {
			if M[i][j] == 1 && visited[j] == 0 {
				visited[j] = 1
				dfs(j)
			}
		}
	}

	for i := 0; i < len(M); i++ {
		if visited[i] == 0 {
			dfs(i)
			count++
		}
	}
	return count

}

// 使用并查集
// 时间复杂度：O(N^2)
// 空间复杂度：O(N)
func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	mLen := len(M)
	ufind := NewUnionFind(mLen)

	for i := 0; i < mLen; i++ {
		for j := i + 1; j < mLen; j++ { // 对称矩阵 不需要重复计算
			if M[i][j] == 1 {
				ufind.Union(i, j)
			}
		}
	}
	return ufind.count
}

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		count:  n,
		parent: parent,
	}
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
		u.count--
	}
}

// 路径压缩后 查询为O(1)
func (u *UnionFind) Find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] != i { // 路径压缩
		i, u.parent[i] = u.parent[i], root
	}
	return root
}
