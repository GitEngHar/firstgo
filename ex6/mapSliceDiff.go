package main

import "fmt"

// ポインタ型の値を渡して変更を加える
// 特にMap。この設計はAPI等では避けるべき -> mapのstringはkeyを矯正できないので、コードを読まないと仕様が不明
func badPattern(base map[string]int) {
	base["coin"] += 100
}

// スライスを関数で渡した場合、関数内でスライスの内容は変わる
// スライスの値は関数の中で変化するが、渡した元の方よりサイズが大きくなると、元のスライスは大きさは変わらない
// そのため大元の方からは値の変化が追えない
// 要素の変更は反映されるけど、サイズ（＝内部配列の再確保を伴う操作）は元に影響しないことがあります。
func sliceFunc(coins []int) {
	coins[0] = 100
	coins = append(coins, 999)
}
func main() {
	var coin = map[string]int{"coin": 0}
	var coins = []int{1, 2, 3}
	badPattern(coin)
	fmt.Println(coin)
	sliceFunc(coins)
	fmt.Println(coins)
}
