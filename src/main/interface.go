/* interface
メソッドの一覧を定義したデータ型
構造体がそのメソッドを実装していれば構造体はインターフェイス型として取り扱うことが出来る
*/

package main

import "fmt"

// show関数の実装
func show(t interface{}) { // 空のインターフェイス型を引数に受け取る
	// 型アサーション
	_, ok := t.(japanese) // 第一戻り値に型それ自体,第二戻り値に処理の成否, 
	if ok { // 処理の成否(型の一致/不一致)を利用し、アサーション
		fmt.Println("i am japanese")
	} else {
		fmt.Println("i am not japanese")
	}

	// 型switch
	switch t.(type) { // 引数t(空のインターフェイス型)のtypeフィールドを基にswitch
	case japanese:
		fmt.Println("i am japanese")
	case american:
		fmt.Println("i am american")
	default:
		fmt.Println("i am alian")
	}
}

// greeterインターフェイス
type greeter interface {
	greet() // greetメソッドを宣言 実装はしない
}

// 構造体を宣言
type japanese struct {}
type american struct {}

// greeterインターフェイスで宣言されたメソッドを実装
func (j japanese) greet() {
	fmt.Println("Konnichiwa!")
}

func (a american) greet() {
	fmt.Println("Hello!")
}

func main() {
	// greetersスライスにgreeter型として代入
	greeters := []greeter{japanese{}, american{}} // 異なる構造体だが同じデータ型として処理出来る
	for _, national := range greeters { // rangeを用いてgreetersスライスを繰り返し処理
		national.greet()
		show(national) // showメソッドでデータ型を確認
	}
}
