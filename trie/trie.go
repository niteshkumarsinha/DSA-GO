package main


import "fmt"


// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26


// Node represents each node in the trie
type Node struct {
	children [26]*Node
	isEnd bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// Init trie will create a new trie
func InitTrie() *Trie{
	results := &Trie{root: &Node{}}
	return results
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string){
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex]  == nil {
			currentNode.children[charIndex] = &Node{}	
		} 
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and Return true if that word is included in the trie
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false	
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}


func main(){
	testTrie := InitTrie()
	fmt.Println(testTrie.root)
	testTrie.Insert("nitesh")
	found := testTrie.Search("nitesh")
	fmt.Println(found)
	found = testTrie.Search("aragon")
	fmt.Println(found)
}