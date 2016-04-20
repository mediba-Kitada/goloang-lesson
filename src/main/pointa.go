package main

import "fmt"

/* ポインタ
アドレスを指し示す変数
演算は出来ない
メモリの節約
→効率的なプログラムが書ける
*/

func main() {
	a := 5
	var pa *int //ポインタ変数の宣言 型の前に*を付与
	pa = &a // 変数aのアドレスを代入
	fmt.Println(pa) // アドレスの表示
	fmt.Println(*pa) // paの領域にあるデータの値を表示
}
