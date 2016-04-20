package main

import "fmt"

// 関数名 引数名 引数の型 戻り値の型
//func hi(name string) string {
//	msg := "hi! " + name
//	fmt.Println(msg)
//	return msg
//}

// 関数名 引数名 引数の型 戻り値名 戻り値の型
func hi(name string) (msg string) {
	msg = "hi! " + name
	return
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(hi("taguchi"))

	fmt.Println(swap(5,2))

	// 関数は、データ型でもあるので変数に代入可能(関数リテラル)
	// 匿名関数
	f := func(a, b int) (int, int) {
		return b, a
	}
	fmt.Println(f(2,3))

	// 即時関数ライクに宣言した直後に実行可能
	func(msg string) {
		fmt.Println(msg)
	}("taguchi") // 引数を渡す
}
