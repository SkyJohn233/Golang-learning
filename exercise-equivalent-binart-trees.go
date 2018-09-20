package main

import (
"fmt"
)
type Tree  struct{
	Left,Right *Tree
	Value int
}
// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int){
	Sendvaltoch(t,ch)
	close(ch)
}
func Sendvaltoch(t *tree.Tree,ch chan int){
	if(t.Left!=nil){
		Sendvaltoch(t.Left,ch)
	}
	ch<-t.Value
	if(t.Right!=nil){
		Sendvaltoch(t.Right,ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int,100)
	ch2 := make(chan int,100)
	Walk(t1,ch1)
	Walk(t2,ch2)
	for i:= range ch1{
		if i != <- ch2{
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}

