package main

import "fmt"

type node struct {
	data int
	next *node
}

func insert(head *node, data int) *node {
	n := &node{data: data}
	if head != nil {
		n.next = head
		return n
	}
	return n
}

func PrintNodes(head *node) {
	for head != nil {
		fmt.Print(head.data, " \033[1;32m* -->\033[0m ")
		head = head.next
	}
	fmt.Println(nil)
}

func main() {

	var link *node
	link = insert(link, 1)
	link = insert(link, 2)
	link = insert(link, 3)

	PrintNodes(link)

}
