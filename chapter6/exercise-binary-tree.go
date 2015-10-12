package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Tree struct{
	Left *Tree
	Value int
	Right *Tree
}

func insert (t *Tree, i int) *Tree{
	if t== nil{
		return &Tree{nil, i, nil}
	}

	if(i< t.Value){
		t.Left= insert(t.Left, i)
		return t
	}else{
		t.Right= insert(t.Right, i)
		return t
	}
}

func main() {
	r:= rand.New(rand.NewSource(time.Now().UnixNano()))
	t:= insert(nil, r.Intn(50))
	insert(t, r.Intn(50))

	fmt.Println(t.Value,t.Right.Value)
}
