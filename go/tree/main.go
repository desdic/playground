package main

import (
	"fmt"
)

// MyInt ...
type MyInt struct {
	v int
}

// Len is the size/length of MyInt for comparing
func (m MyInt) Len() int {
	return m.v
}

func main() {
	var t Tree

	l := []MyInt{
		MyInt{v: 3},
		MyInt{v: 6},
		MyInt{v: 1},
		MyInt{v: 7},
		MyInt{v: 9},
		MyInt{v: 2},
	}

	for i := range l {
		fmt.Printf("Inserting: %v\n", l[i].v)
		if err := t.Insert(l[i]); err != nil {
			fmt.Println("err")
			break
		}
		t.Walk(t.Root, func(n *Node) {
			fmt.Print(n.Value, " | ")
		})
	}
	//for i := range l {
	//fmt.Printf("Deleting: %v\n", l[i].v)
	//if err := t.Delete(l[i]); err != nil {
	//fmt.Println("err")
	//break
	//}
	//t.Walk(t.Root, func(n *Node) {
	//fmt.Print(n.Value, " | ")
	//})
	//}
	fmt.Println()
	t.WalkReverse(t.Root, func(n *Node) {
		fmt.Print(n.Value, " | ")
	})
}
