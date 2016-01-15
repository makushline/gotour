package main

import (
    "golang.org/x/tour/tree"
    "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
    if t == nil {
        return
    }
    Walk(t.Left, ch)
    ch <- t.Value
    Walk(t.Right, ch)   
}

func TreeSize(t *tree.Tree) (size int) {
    if t == nil {
        return
    }
    size++
    if t.Left != nil {
        size += TreeSize(t.Left)
    }
    if t.Right != nil {
        size += TreeSize(t.Right)
    }
    return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    sizeT1 := TreeSize(t1)
    sizeT2 := TreeSize(t2)
    if sizeT1 != sizeT2 {
       return false
    }
    
    ch1 := make(chan int)
    ch2 := make(chan int)
    go Walk(t1, ch1)
    go Walk(t2, ch2)

    for i:=0; i < sizeT1; i++ {
        if <-ch1 != <-ch2 {
            return false
        }
    }
    return true
}

func main() {
    test1 := Same(tree.New(1), tree.New(1))
    test2 := Same(tree.New(1), tree.New(2))
    fmt.Println(test1)
    fmt.Println(test2)
}
