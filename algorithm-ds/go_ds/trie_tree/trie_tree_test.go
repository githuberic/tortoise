package trie_tree

import (
	"fmt"
	"testing"
)

func TestTrieTree(t *testing.T)  {
	trie := NewTrie()
	words := []string{"Golang", "学院君", "Language", "Trie", "Go"}
	// 构建 Trie 树
	for _, word := range words {
		trie.Insert(word)
	}
	// 从 Trie 树中查找字符串
	term := "学院君"
	if trie.Find(term) {
		fmt.Printf("包含单词\"%s\"\n", term)
	} else {
		fmt.Printf("不包含单词\"%s\"\n", term)
	}
}
