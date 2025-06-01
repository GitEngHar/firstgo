package main

import "fmt"

type person struct {
	LastName  string
	FirstName string
	Age       int
}

type T struct{}

func (t T) Hello() {}                     // 値レシーバ
func (t *T) Bye()  { fmt.Println("Bye") } // ポインタレシーバ

// StringHoge レシーバに型名を記載することで、personとStringメソッドが日もづく
// 合成の例
// 値レシーバはメソッドがレシーバを変更しない場合
func (p person) StringHoge() string {
	return fmt.Sprintf("%s %s:年齢%d", p.LastName, p.FirstName, p.Age)
}

// Increment ポインタレシーバ
// メソッドがレシーバを変更する場合
// メソッドがnilを扱う必要がある場合
func (p *person) Increment() {
	p.Age++
}

func (p *person) ValueIncrement(value int) int {
	return p.Age + value
}

// NotIncrement 加算失敗するメソッド
func (p person) NotIncrement() {
	p.Age++
}

// goでは軽症ではなく、合成を推奨している
// struct in struct
// 依存性の注入はGoogleが作成したWriteがある
func main() {
	// メソッドセットは型が持ってるメソッドの一覧で、メソッドはポインタ型と値型で変わる
	// ポインタ型は値型とメソッド型のメソッドセットを持つが、値型は値のみ
	// 値型であるが、ポインタのレシーバに対してはGOが自動変換で対応してくれる
	var p person

	p.Increment() // Goが自動的に渡す値をポインタ変数にする
	p.NotIncrement()
	output := p.StringHoge()
	fmt.Println(output)

	var t T
	t.Bye()

	// メソッドと関数は、すごい似てる
	// メソッド式とした場合の最初の引数は、対象の構造体を格納した変数になる
	var p2 person
	vp := person.ValueIncrement
	fmt.Println(vp(p2, 10))
}
