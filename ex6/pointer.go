package main

import "fmt"

type user struct {
	FirstName string
	LastName  *string
}

func basePointer() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)
	fmt.Println(&pointerToX)

	var y *int //ポインタの生成
	fmt.Println(y == nil)
	y = &x
	fmt.Println(*y)

	var z = new(int)
	fmt.Println(z == nil)
	z = &x //z = y, z= &xとなる。ポインタ型とはメモリのアドレスを格納する値である
	fmt.Println(*z)
}

func main() {
	basePointer()
	var userLastName string = "hogeo"
	var user = user{
		"hoge",
		&userLastName, // 値のメモリアドレスを渡す必要があるので、そのまま文字列は代入できない
	}
	fmt.Println(user)
}
