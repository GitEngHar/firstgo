package main

import (
	"fmt"
	"time"
)

func times() {
	// 2h 30m 50s を表す
	d := 2*time.Hour + 30*time.Minute + 50*time.Second
	fmt.Println(d)
}

// t.Parse後にt.Formatで整形するが2006年の文字列置換で大変そう
// time.ANSIC や time.UnixDateを使ったフォーマット指定もある

// 時間計測には2種類、PCが起動されてからの時間と標準時間がある

// time.AfterFuncを使うことで、time.Duration後に関数を起動できる
// time.TickerはDuration経過ごとに毎回値を返す (GCされない)
func main() {
	times()
}
