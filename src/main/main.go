package main

import (
	"fmt"
	"encoding/json" // jsonパッケージとしてコール可能
	"log"
	"os"
	. "task" // .記法は、複数パッケージで利用出来る
	. "person"
)

func main() {
	/*
		var id ID = 3 // 型はパッケージを含め指定する必要がある
		var priority Priority = 5
		ProcessTask(id, priority)
	*/
	var task1 Task = Task{ // taskパッケージのTask構造体を用いてtask変数に代入
		ID:     1,
		Detail: "buy the milk",
		//doen:	true, // クローズドなプロパティなので、代入出来ない
	}
	fmt.Println(task1.ID)
	fmt.Println(task1.Detail)

	task2 := Task{} // 構造体の値を明示しなかった場合は、各プロパティの型のゼロ値が代入される
	fmt.Println(task2.ID)

	// 組み込み関数new()による構造体の初期化
	var task3 *Task = new(Task) // 構造体のフィールドをすべてゼロ地で初期化し、そのポイントを返却する
	// task3 := new(Task) // ポインタであることを明示しなくても可能
	task3.ID = 3
	task3.Detail = "sell the milk"
	fmt.Println(task3.ID)
	fmt.Println(task3.Detail)

	task4 := NewTask(4, "convert milk", "Momoe", "Kitada")
	fmt.Println("%+v", task4)

	task5 := NewTask(5, "take milk", "Kai", "Kitada")
	task5.Finish()
	fmt.Printf("%s", task5)

	task6 := NewTask(6, "through milk", "Riku", "Kitada")
	task6.Print(task6) // インターフェイスを実装したので、String()メソッドの呼び出しを構造化出来る

	task7 := NewTask(7, "catch milk", "Tsubasa", "Kitada") // Task型にUser型を埋め込んだ(Embed)ので、User型にもアクセス可能となった
	// Task型に埋め込まれたUser型のフィールドにアクセス
	fmt.Println(task7.FirstName)
	fmt.Println(task7.LastName)
	fmt.Println(task7.FullName()) // Task型に埋め込まれたUser型のメソッドにアクセス
	fmt.Println(task7.User)       // Taskから埋め込まれたUser型自体にもアクセス可能

	// 型変換(キャスト)
	var i uint8 = 3
	var j uint32 = uint32(i) // 変換したい型に変数を渡す
	fmt.Println(j)

	var s string = "abc"
	var b []byte = []byte(s)
	fmt.Println(b)

	//a := int("a") // キャストに失敗した場合は、パニックが発生する

	person := &Person{ // 構造体Personのポインタを代入
		ID:	1,
		Name:	"Gopher",
		Email:	"gopher@example.org",
		Age:	5,
		Address:	"",
	}
	b, err := json.Marshal(person) // ポイントをencoding/json.Marshal()関数に渡す
	if err != nil {
		log.Fatal(err) // logパッケージFatal()関数
	}
	fmt.Println(string(b)) // 文字列に変換し、標準出力

	/* json.Encoder/Decoderを用いたJSON形式のファイル操作
	io.Reader, io.Writerを中心として設計されたAPIが多数ある
	*/
	file, ioerr := os.Create("./person.json") // ファイル生成
	if ioerr != nil {
		log.Fatal(ioerr)
	}
	defer file.Close()

	encoder := json.NewEncoder(file) // エンコードの取得
	ioerr = encoder.Encode(person) // JSONエンコードしたデータの書き込み []byteに代入せず、personを直接*os.Fileに書き込み
	if ioerr != nil {
		log.Fatal(ioerr)
	}

	var person2 Person
	b2 := []byte(`{"id":1,"name":"Gopher","age":5}`) //byte方のスライスにJSON文字列を代入
	err2 := json.Unmarshal(b2, &person2) // encoding/json.Unmarshal()関数にスライスと構造体のポイントを渡す
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(person2) // 構造体を標準出力

	// json.Encoder/Decoderを用いたJSONファイルの読み出し/デコード
	file, ioerr = os.Open("./person.json") // 生成済みのperson.jsonファイルを生成
	if ioerr != nil {
		log.Fatal(ioerr)
	}
	defer file.Close()

	var person3 Person // デコードしたデータを書き込み変数
	decoder := json.NewDecoder(file) // デコーダの取得
	ioerr = decoder.Decode(&person3) // JSONデコードしたデータをperson3に代入
	if ioerr != nil {
		log.Fatal(ioerr)
	}
	fmt.Println(person3)

}
