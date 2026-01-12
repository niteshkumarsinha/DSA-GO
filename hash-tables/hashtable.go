package main


import "fmt"


const ArraySize = 7


// HashTable Structure
type HashTable struct {
	array [ArraySize]*bucket
}

// Bucket Stucture
type bucket struct {
	head *bucketNode
}

// BucketNode Structure
type bucketNode struct {
	key string
	next *bucketNode
}


// Hashtable Methods
// Insert
func (h *HashTable) Insert(key string){
	index := hash(key)
	h.array[index].insert(key)
}

// Search
func (h *HashTable) Search(key string) bool{
	index := hash(key)
	return h.array[index].search(key)
}

// Delete
func (h *HashTable) Delete(key string){
	index := hash(key)
	h.array[index].delete(key)
}


// Bucket Methods
// insert
func (b *bucket) insert(k string){
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("already exists")
	}
}

// Search
func (b *bucket) search(k string) bool{
	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Delete
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}


// Hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main(){
	hashTable := Init()

	list := []string {
		"Eric",
		"Kenny",
		"Kyle",
		"Stan",
		"Randy",
		"Butters",
		"Token",
	}

	for _, v:= range list {
		hashTable.Insert(v)
	}
	fmt.Println(hashTable)

	for _, v := range hashTable.array {
		bucket := v.head 
		for bucket != nil {
			fmt.Println(bucket.key)
			bucket = bucket.next
		}

	}
}
