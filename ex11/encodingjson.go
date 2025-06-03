package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func simpleJson() {
	f := struct {
		Name string
		Age  int
	}{}
	// JSONを入れたバイト列
	// ポインタ型に値を引数とする
	// リフレクションという機能が関係して、ライブラリの振る舞いを実現している
	err := json.Unmarshal([]byte(`{"name":"hoge", "age":12}`), &f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", f)
}

func differentJson() {
	type Item struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	type Order struct {
		ID          string    `json:"id"`
		DateOrdered time.Time `json:"date_ordered"`
		Items       []Item    `json:"items"`
	}
	newOrderMap := Order{}
	newOrder := []byte(`{"id":"12345",
		"date_ordered":"2020-05-01T13:01:02Z",
		"customer_id":"3",
		"items":[{"id":"xyz123","name":"物品1"},{"id":"abc789","name":"物品2"}]}`)
	err := json.Unmarshal(newOrder, &newOrderMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", newOrderMap)
}

func readFile() {
	type Item struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	var dataFromFile Item
	tmpFile2, err := os.Open("tmpFile2")
	err = json.NewDecoder(tmpFile2).Decode(&dataFromFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read Data From File: %+v \n", dataFromFile)
}

func writeToFile() {
	type Item struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	dataToFile := Item{
		ID:   "123",
		Name: "hogeo",
	}
	tmpFile, err := os.CreateTemp(os.TempDir(), "sample-")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())
	err = json.NewEncoder(tmpFile).Encode(dataToFile)
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)
	err = tmpFile.Close()
	if err != nil {
		panic(err)
	}
}

// (本番避けるべし) net/httpのhttp.Clientにはデフォルトでタイムアウトが存在していない
// ちゃんとゴルーチンを跨いだ処理をしてくれる
// http.Server
// Server側で複数のハンドラに跨って処理をしたい場合は、ミドルウェアを利用する必要ある

func main() {
	readFile()
	//writeToFile()
}
