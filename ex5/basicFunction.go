package main

import (
	"fmt"
)

type User struct {
	name string
	perm int32
}

// 複数の返り値
func manyReturnFunc(base int, plus int) (int, int, error) {
	return base + plus, base - plus, nil
}

// 可変長引数
// 可変長引数の値は最後に持ってくるか唯一である
func variableLengthFunc(val int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, val+v)
	}
	return out
}

// 名前付き引数が欲しくなる場合は、関数が複雑になっている可能性大なので、リファクタを検討しよう
func noNameArg(user User) {
	fmt.Println(user.name, user.perm)
}

func namedReturnFanc(base int) (resultValue int) {
	if base%2 == 0 {
		resultValue = base * 2 //返り値の宣言済みになる
		return resultValue
	}
	return base * 3 //返り値は名前付き返り値と同等でなくてもいい
}

func blankReturn(base int) (resultValue int) {
	if base%2 == 0 {
		resultValue = base * 2 //返り値の宣言済みになる
		return resultValue
	}
	// なんの値が返ってくるのかは0の種類によるため、可読性が悪い
	return //返り値をブランクにできてしまう
}

func main() {
	noNameArg(
		User{name: "enghar", perm: 777},
	)
	var plusValue, minusValue, _ = manyReturnFunc(1, 2)
	fmt.Println(variableLengthFunc(10, 1, 2, 3, 4, 5))
	fmt.Println(plusValue, minusValue)
	fmt.Println(namedReturnFanc(2))
	fmt.Println(namedReturnFanc(3))
	fmt.Println(blankReturn(2))
	fmt.Println(blankReturn(3)) // 0が返ってくる
}
