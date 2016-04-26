package main

/* io/ioutilパッケージでファイルを操作する
osパッケージで操作するよりシンプルに操作出来る
*/

import(
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// ファイルにメッセージを書き込む
	hello := []byte("hello world\n")
	/* ioutil.WriteFile()
	第一引数: ファイル名
	第二引数: []byte型のデータ
	第三引数: 対象ファイルのパーミッションを8進数で表現
	*/
	err := ioutil.WriteFile("./ioutil.txt", hello, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// ファイルの中身を全て読みだす
	message , err := ioutil.ReadFile("./ioutil.txt") // 第一戻り値に[]byte型のデータを設定してくれる
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(message))

}
