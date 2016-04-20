/* 構造体 メソッド
OOP言語におけるメソッドとは異なり、構造体を定義する際にメソッドを定義しない
-> クラスをサポートしているわけではない
*/

package main

import "fmt"

type user struct {
	name string
	score int
}

func (u user) show() { // レシーバにより構造体userと関連付ける
	fmt.Printf("name:%s, score:%d\n", u.name, u.score) // レシーバにより構造体のフィールドにアクセス出来る??
}

// scoreフィールドをインクリメントする関数
//func (u user) hit() { // 値渡しによる実装
//	u.score++ // 値渡しなので、フィールドに影響は与えない
//}

func (u *user) hit() { // 参照渡しによる実装(ポインタ型の変数としてレシーブする??)
	u.score++ // 参照渡しなので、フィールドに影響する(参照先を上書きする)
}

func main() {
	u := user{name:"taguchi", score:50} // 構造体を用いて変数uを初期化
	u.hit()
	u.show() // メソッドshowを用いて変数uのフィールドを表示
}
