package week07

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
