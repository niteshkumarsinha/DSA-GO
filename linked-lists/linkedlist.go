package main


import "fmt"


type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	length int
}

func (l *LinkedList) prepend(node *Node) {
	second := l.head
	l.head = node
	l.head.next = second
	l.length++
}

func (l LinkedList) printListData(){
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d -> ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println("nil")
}

func (l *LinkedList) deleteWithValue(value int){
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main(){
	myList := LinkedList{}
	node1 := &Node{data: 48}
	node2 := &Node{data: 18}
	node3 := &Node{data: 16}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)

	myList.printListData()

	myList.deleteWithValue(48)
	myList.printListData()

	myList.deleteWithValue(50)
	myList.printListData()
}