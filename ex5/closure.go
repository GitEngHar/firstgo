package main

import "fmt"

func main() {
	counter := createCounter()
	// 状態が保持されている
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

// 内部関数が値を保持している状態
func createCounter() func() int {
	// 無名関数から見れば、外部で宣言された変数
	count := 0
	return func() int {
		count++ //カウントアップ
		return count
	}
}
