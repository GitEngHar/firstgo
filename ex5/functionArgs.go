package main

import (
	"fmt"
	"sort"
)

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	people := []person{
		{"sato", "taku", 19},
		{"suzuki", "syu", 22},
		{"uraraka", "otyako", 25},
	}

	// 無名関数を引数として渡している
	sort.Slice(people, func(i int, j int) bool {
		return people[i].lastName < people[j].lastName
	})

	fmt.Println("People array sorted by lastName")
	// 別関数の中で変更された値にもかかわらず状態が保持され、変更されている
	// クロージャの挙動
	fmt.Println(people)
}
