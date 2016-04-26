package main

import(
	"log"
	"os"
	"io/ioutil"
	"fmt"
)

/* *os.File
io.ReadeWriteCloserというインターフェイス型
→ Read() Write() Close()の3つのメソッドを実装している
データの書き込みには複数の方法があるが、いずれもio.Writerインターフェイス型であることに注意
*/

func main() {
	// ファイルを作成
	file, err := os.Create("./file.txt") // *os.File構造体へのポインタが取得出来る
	if err != nil {
		log.Fatal(err)
	}
	// deferを用いてファイルを必ず閉じる
	defer file.Close()

	// 書き込むデータを[]byteで用意する
	message := []byte("hello world")

	// Wirte()を用いて書き込む
	_, err = file.Write(message) // 第一戻り値として書き込んだバイト数があるが無視(ブランク修飾)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString("hi world") // []byteに変換する手間が省ける

	_, err = fmt.Fprint(file, "hoge world\n") // io.Writeに対して文字列を直接書き込むことも出来る

	// ファイルを開く
	file, err = os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// ファイルの読み出し
	read_message := make([]byte, 60) // ファイルに記載されているデータを格納出来る長さのスライスを用意
	_, err = file.Read(read_message) // io.Read()関数にスライスを渡し、読みだしたデータを格納する
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(read_message)) // string型に変換し標準出力

	// io/ioutilを用いたファイルの読み出し
	file , _ = os.Open("./file.txt")
	read_message2, read_err := ioutil.ReadAll(file) // ファイルの中身を全て読み出し、[]byte型のread_message2に代入
	if read_err != nil {
		log.Fatal(read_err)
	}
	fmt.Println(string(read_message2)) // string型に変換し標準出力
}
