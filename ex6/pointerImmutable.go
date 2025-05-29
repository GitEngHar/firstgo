package main

import "fmt"

// ここでの初期値は”であるが、これをnilとしたい場合にも ポインタは有効
// JSONだとnilの場合出力されないので、アリ寄りのあり
type person struct {
	FirstName string
	LastName  string
}

// ポインタ型を引数に取ることでガベージコレクタの仕事が増える
// ポインタ型の値はデータのサイズに受け渡し時間が依存せず、一律に1ナノ秒
// 1MBよりデータが小さい場合は、ポインタの方がデータのやり取りが遅い
// 100Bをのデータを返すのに10ナノ秒かかる。ポインタの場合は30ナノ秒かかる
// 1MBを超えると、パフォーマンスは逆転する。10MBのデータを返すのに2m秒。同じデータでポインタは0.5m秒。
func badStruct(p *person) person {
	p.FirstName = "hoge-first"
	p.LastName = "hoge-last"
	return *p
}

func bestStruct() person {
	return person{
		FirstName: "hoge-first",
		LastName:  "hoge-last",
	}
}

func noUpdate(g *int) {
	x := 10
	g = &x //受け取った値のアドレスを変数xのアドレスに置き換える
}

func update(g *int) {
	x := 10
	*g = x //アドレスに格納されている値を置き換える
}

func main() {
	// ポインタはコードを追いづらくなる
	// ガベージコレクタの仕事も増える
	init := 7
	noUpdate(&init)
	fmt.Println(init) //値が更新されていない

	update(&init)
	fmt.Println(init) //値が更新されていない
	fmt.Println(badStruct(&person{}))
	fmt.Println(bestStruct())

}
