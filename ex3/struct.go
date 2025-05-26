package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func baseStruct() {
	var userA = person{}
	fmt.Println(userA)
	var userB = person{
		name: "Haru",
		age:  25,
		pet:  "Hum",
	}
	fmt.Println(userB)
	userB.age = 26
	fmt.Println(userB)
}

func lambdaStruct() {
	// 無名構造体
	// 用途: 外部データを構想帯に変換したいときとか
	pet := struct {
		name string
		age  int
	}{
		name: "hana",
		age:  5,
	}
	fmt.Println(pet)
}

func comparisonStruct() {
	var userA = person{
		name: "haru",
		age:  15,
		pet:  "ponta",
	}
	userC := struct {
		name string
		age  int
		pet  string
	}{}

	fmt.Println(userC == userA) //false

	//値までチェックされる
	userC = struct {
		name string
		age  int
		pet  string
	}{
		name: "haru",
		age:  15,
		pet:  "ponta",
	}

	fmt.Println(userC == userA) //true
}

func main() {
	baseStruct()
	lambdaStruct()
	comparisonStruct()
}
