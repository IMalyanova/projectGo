package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	for i := 0; i < 10; i++  {
		fmt.Println(<- ch)
	}
}


// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	if t1 == t2 {
		return true
	} else {
		return false}
}


func main() {
	ch := make(chan int)
	go Walk(tree.New(10), ch)

}
