package main

import (
	"datastructure/sqlist/sqList"
	"fmt"
)

func main() {
	S := sqList.NewList(20)
	for i := 0; i < 10; i++ {
		S.Append(i + 1)
	}
	fmt.Println(S)
	S.Delete(1)
	fmt.Println(S)
	//S.Clear()
	//fmt.Println(S)
	S.Set(1, 200)
	fmt.Println(S)
	result, _ := S.Get(3)
	fmt.Println(result)
}
