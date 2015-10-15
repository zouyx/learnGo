package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func Insert(t *Tree, i int) *Tree {
	if t == nil {
		return &Tree{nil, i, nil}
	}

	if i < t.Value {
		t.Left = Insert(t.Left, i)
		return t
	} else {
		t.Right = Insert(t.Right, i)
		return t
	}
}

func CreateTree() *Tree {
	var t *Tree
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0 ; i<10 ; i++ {
		t = Insert(t, r.Intn(50))
	}
	return t
}

func PrintTree(t *Tree, c chan int) {
	if(t==nil){
		return
	}

	PrintTree(t.Left,c)	
	c <- t.Value
	PrintTree(t.Right,c)	
} 

func Print(t *Tree, c chan int) {
	PrintTree(t, c)

	close(c)
}

func Compare (t1 *Tree,t2 *Tree) bool{
	c1 := make(chan int)
	c2 := make(chan int)
	
	go Print(t1, c1)
	go Print(t2, c2)

	var flag bool=true
	
	for {
		i, iok := <- c1
		j, jok := <- c2
		if(!iok||!jok){
			break
		}

		fmt.Println(i, j)
		if(i != j){
			flag = false
			break
		}
	}
	
	return flag
}

func main() {
	t1 := CreateTree()
	t2 := CreateTree()
	
	flag :=	Compare(t1, t2)
	fmt.Println("value:",flag)
}
