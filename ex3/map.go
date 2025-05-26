package main

import "fmt"

// goにはsetの機能が存在しないが、その機能をシミュレートすることはできる
// スライスは要素数が増えるにつれ、負荷が高くなるがmapはその限りではない
// mapに検索対象とboolをセットしておけば、文字列検索を簡単に早く実施可能
// 検索のテクニック(for-range検索より、検索する手間が少ない)
// Goではkeyに対象を留めておけば、第二引数で値を持ってきて存在確認ができる

func mapCommaOk() {
	teams := map[string][]string{
		"admin": []string{"userA", "userB"},
		"dev":   []string{"userC", "userD", "userAB", "userBA"},
		"ops":   []string{"userE", "userF", "userG"},
	}
	v, ok := teams["admin"]
	fmt.Println(v, ok)
	v, ok = teams["ro"] //文字列の配列の場合、返ってくる初期値はnil
	fmt.Println(v, ok)

	/*
	* mapの削除機能を確認する
	 */
	v, ok = teams["ops"]
	fmt.Println(v, ok)
	delete(teams, "ops") //削除実行(返り値なし)
	v, ok = teams["ops"]
	fmt.Println(v, ok)
}

// 順番に並んでいない値をキーとする場合は、mapが適している (順番の場合はスライス)
// mapは内部的にハッシュマップとして実装される
func mapBasic() {
	// スライスは順番通りのデータ
	// ランダムのデータにkeyとvalueで規則性を持たせる
	var nilMap map[string]int

	// 値の初期値を無しにして値をセットする
	winMap := map[string]int{} // nilではない
	fmt.Println(winMap)
	fmt.Println(nilMap)

	// mapを宣言する ,区切り(最後にも,つけてる)
	teams := map[string][]string{
		"admin": []string{"userA", "userB"},
		"dev":   []string{"userC", "userD", "userAB", "userBA"},
		"ops":   []string{"userE", "userF", "userG"},
	}
	fmt.Println(teams["admin"])

	numberOfTeam := map[string]int{}
	numberOfTeam["admin"] = 2
	fmt.Println(numberOfTeam["admin"])
	numberOfTeam["dev"] = 4
	fmt.Println(numberOfTeam["dev"])
	numberOfTeam["ops"] = 3
	fmt.Println(numberOfTeam["ops"])
	numberOfTeam["ops"]++ //デフォルトが0値なので、key指定されていなくても加算される
	fmt.Println(numberOfTeam["ops"])
	// ある程度予想できる場合は以下のようにmakeする
	// make(map[string]int, 10)
}

func main() {
	mapCommaOk()
}
