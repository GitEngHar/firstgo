package main

import (
	"fmt"
	"strconv"
)

// 引数と返り値のセットをシグネチャと呼ばれる

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

func main() {
	var opMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "/", "three"},
		[]string{"2"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println(expression, ": 不正な値")
			continue
		}

		pl, err := strconv.Atoi(expression[0])

		if err != nil {
			fmt.Println(expression[0], ": エラー -- ", err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op] //演算子に該当する関数を値として取得している

		if !ok {
			fmt.Println("定義されていない演算子 -- ", op)
			continue
		}

		pl2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(expression[2], ": エラー\n", err)
			continue
		}

		result := opFunc(pl, pl2)
		fmt.Println(result)
	}
	var intArray = []int{1, 2, 3, 4, 5}
	// deferとゴルーチンの起動で無名関数が役立つらしい
	for i, v := range intArray {
		func(j int, x int) {
			fmt.Println("無名関数", j+x)
		}(i, v)
	}
}
