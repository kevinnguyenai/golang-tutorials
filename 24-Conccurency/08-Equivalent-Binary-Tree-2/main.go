package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		defer close(ch1)
		Walk(t1, ch1)
	}()
	go func() {
		defer close(ch2)
		Walk(t2, ch2)
	}()
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 && !ok2 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		Walk(tree.New(20), ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(10), tree.New(10)))
}
