package main

import "fmt"

func modify(baseMap map[string]int) {
	baseMap["hogeo"] = 1
	baseMap["fuga"] = 2
}

func sliceModify(baseSlice []int) {
	baseSlice[0] = 999
}

func failedModify(base int) int {
	return base + 1
}

// mapとスライスはポインタを使った実装をされているため、関数に渡して値を帰ると、元の変数にまで影響が出る
// 変数等はポインタではないので値は下記変わらない
// 基本的に、関数に渡される値は call by value であるため、渡された時点で値がコピーされる。
// ポインタの場合はアドレスが渡されるため、それがコピーされても値の入っているアドレスは同一であるため、値が書き変わってしまう
func main() {
	var base int = 0
	failedModify(base)
	fmt.Println(base) //値は0のまま

	baseMap := map[string]int{
		"hogeo": 10,
		"fuga":  99,
	}

	modify(baseMap)
	fmt.Println(baseMap) //値が変わる

	baseSlice := make([]int, 0, 1)
	baseSlice = append(baseSlice, 1)

	sliceModify(baseSlice)
	fmt.Println(baseSlice) //999になっている

}
