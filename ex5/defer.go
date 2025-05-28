package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)

	if err != nil {
		return nil, nil, err
	}

	// 後処理は書くけど、後処理するタイミングは使う側に委ねる
	return file, func() {
		file.Close()
	}, err
}

func main() {
	if len(os.Args) < 2 { // ファイル名が指定されている?
		log.Fatalf("ファイルが指定されていません")
	}
	f, closer, err := getFile(os.Args[1]) // ファイルオープン
	if err != nil {
		log.Fatalf(err.Error()) // オープンエラー
	}

	defer closer() //closer呼び出し
	//defer f.Close() // クリーンアップ。最後にファイル閉じる
	// defer func()とすることで、後始末を無名関数にすることができる
	// dbの場合 -> 処理中にエラーがなければ結果をコミット, エラーがあればロールバックというように使われる
	// バイトのスライスを生成する
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		_, wErr := os.Stdout.Write(data[:count])
		if wErr != nil {
			fmt.Println(wErr.Error())
		}
		if err != nil {
			if err != io.EOF {
				log.Fatalf(err.Error())
			}
			break
		}
	}
}
