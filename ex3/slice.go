package main

import "fmt"

func confirmSliceCap() {
	// スライスを増やす際にメモリ両機を確保する
	// キャパシティは1024以下の場合は、倍ずつ。それ以降は25%ずつ増えていく
	var initSlice []int
	fmt.Println(len(initSlice), cap(initSlice))
	initSlice = append(initSlice, 1)
	fmt.Println(len(initSlice), cap(initSlice))
	initSlice = append(initSlice, 2)
	fmt.Println(len(initSlice), cap(initSlice))
	initSlice = append(initSlice, 3)
	fmt.Println(len(initSlice), cap(initSlice))
	initSlice = append(initSlice, 4)
	fmt.Println(len(initSlice), cap(initSlice))
	initSlice = append(initSlice, 5)
	fmt.Println(len(initSlice), cap(initSlice))
}

func createSliceByMake() {
	// コードを書いていてサイズがわからない場合のスライスに使う
	// 余分な値が確保されるのを防ぐため、どれだけのデータが来るか不明な場合は、make0で生成する
}

func copySliceData() {
	x := []int{1, 2, 3, 4, 5}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}

func syncSliceData() {
	// 参照しているメモリ領域が同じであるため。他の変数の値をコピーすると、参照している変数全てに影響が出る
	x := []int{1, 2, 3, 4, 5}
	y := x[:2]
	z := x[1:]
	x[1] = 20
	y[0] = 11
	z[1] = 32
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func notSyncFullSlice() {
	// フルスライス指定でメモリの共有を防ぐ
	// スライス式ではない場合も範囲コピー等はできる
	x := make([]int, 0, 8)
	x = append(x, 1, 2, 3, 4, 5, 6)
	y := x[:1:1]
	z := x[2:4:4]
	fmt.Println("value: ", x, y, z)
	fmt.Println("cap: ", cap(x), cap(y), cap(z))
	x = append(x, 10)
	y = append(y, 20)
	z = append(z, 30)
	fmt.Println("value: ", x, y, z)
	fmt.Println("cap: ", cap(x), cap(y), cap(z))

}

func notSyncSimpleCopy() {
	x := make([]int, 0, 8)
	x = append(x, 1, 2, 3, 4, 5, 6)
	// xの長さと同等のスライスを生成する
	var y = make([]int, len(x)) //0, で指定すると値がコピーされない
	copy(y, x)
	fmt.Println("before: ", x, y)
	x = append(x, 10)
	y = append(y, 20)
	fmt.Println("after: ", x, y)
}

func confirmStringFunc() {
	var s string = "Hello world"
	var b byte = s[3]
	fmt.Printf("%s %s", s, string(b)) // stringに変換しないと (uint8=108) と表示される
	// 文字列は標準ライブラリの strings や unicode/utf8を使うのがいい
}

func main() {
	//var initSlice []int
	//fmt.Println(initSlice == nil) // sliceの初期値はnilになってる
	//initSlice = append(initSlice, 1, 2, 3, 43)
	//fmt.Println(len(initSlice))
	//var hoge = []int{1, 2, 3}
	//fmt.Println(hoge)

	//copySliceData()
	//syncSliceData()
	//confirmSliceCap()
	//notSyncFullSlice()
	//notSyncSimpleCopy()
	confirmStringFunc()
}
