### 并查集GO模板


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


### 字典树GO模板

    // Trie 字典树
    type Trie struct {
        Children  [26]*Trie
        IsWorkEnd bool
    }

    // Constructor a
    // Initialize your data structure here. */
    func Constructor() Trie {
        return Trie{}
    }

    // Insert a word into the trie.
    func (trie *Trie) Insert(word string) {
        node := trie
        for _, v := range word {
            v -= 'a'
            if node.Children[v] == nil {
                node.Children[v] = &Trie{}
            }
            node = node.Children[v]
        }
        node.IsWorkEnd = true
    }

    // Search Returns if the word is in the trie.
    func (trie *Trie) Search(word string) bool {
        node := trie
        for _, v := range word {
            if node = node.Children[v-'a']; node == nil {
                return false
            }
        }
        return node.IsWorkEnd
    }

    // StartsWith Returns if there is any word in the trie that starts with the given prefix. */
    func (trie *Trie) StartsWith(prefix string) bool {
        node := trie
        for _, v := range prefix {
            if node = node.Children[v-'a']; node == nil {
                return false
            }
        }
        return true
    }