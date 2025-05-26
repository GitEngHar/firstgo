package main

import "fmt"

// Go言語において変数を宣言する場所をブロックと呼ぶ

// trueという値をシャドーイングする
// 値を上書きし、元の変数にアクセスできなくするシャドーイング
func rewriteUniversalDef() {
	fmt.Println(true)
	true := 1 // あらかじめ宣言された識別子を上書きしている
	fmt.Println(true)
}

func onlyConditionsFor() {
	// 条件のみのfor文
	i := 1
	for i < 100 {
		i = i * 2
	}
	fmt.Println("end")

	// 無限ループ
	for {
		fmt.Println("∞")
	}
}

func forRange() {
	targetArray := []int{10, 22, 33, 44, 55}
	for index, value := range targetArray {
		fmt.Println(index, value)
	}
}

func forMap() {
	targetArray := map[string]int32{"hoge": 1, "fuga": 2, "fugahoge": 3}
	// 順不同でkeyだけを取得する
	for key, value := range targetArray {
		fmt.Println(key, value*2)
	}
	// for-rangeの変更は、中の値をコピーしている為値が上書きされない
	fmt.Println(targetArray)
	// Mapイテレーション
	// HashDos攻撃対策と、Mapでは順番を保証される動きはしていないのでこれを、過度に信頼したコードを書かれることを防ぐ
	// そのためイテレーションの際は毎回順番が変わるように実装される
	// labelをつけることで特定のラベル部分までの階層処理をスキップできる
}

func blankSwitch() {
	targetArray := []int{10, 22, 33, 44, 55}
	// 比較対象の値を指定せずに記載するブランクSwitch
	// 論理比較することができる
outer: //ラベル
	for _, value := range targetArray {
		switch targetValue := value % 2; {
		case targetValue == 0:
			// go switchではbreakが不要
			// フォークスルーをしたい場合は明示的に指定する必要がある
			fmt.Println("偶数")
		case targetValue >= 1:
			fmt.Println("奇数")
			break outer //ラベル指定しないと、caseを抜け出すだけになる
		}
	}

	// gotoという機能が存在する
	// ラベル指定した箇所に、処理を飛ばせるが、コードを辿りにくくなるので使わないのが普通らしい
}

func main() {
	//rewriteUniversalDef()
	//onlyConditionsFor()
	//forRange()
	//forMap()
	blankSwitch()
}
