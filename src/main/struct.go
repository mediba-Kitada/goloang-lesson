/* 構造体
- 複数の値を意味のあるまとまりとして新しい型を定義することが出来る
- メソッドを持つことができるのでRubyやJavaのクラスに近い役割も担う
*/

package main

import "fmt"

type user struct { // 構造体の宣言
	name string // フィールド
	score int
}

func main() {
	// u := new(user)
	// fmt.Println(u) // 構造体のアドレスと空文字が返却される &{ 0}

	//(*u).name = "taguchi" // ポインタ型の変数なので、*でセット出来る
	//(*u).score = 20

	//u.name = "taguchi" // ポインタ型の変数だが、*を省略出来る
	//u.score = 20

	// 宣言と初期化
	//u := user{"taguchi", 50} 
	u := user{name:"taguchi", score:50} // フィールド名を指定した初期化も可能
	fmt.Println(u) 
}

