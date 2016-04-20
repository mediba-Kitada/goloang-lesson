/* map
配列とよく似ているが、添字にキーを使って、キーと値のペアで管理が可能なデータ型
*/

package main

import "fmt"

func main() {

	/*
	// map型のmを宣言
	m := make(map[string]int) // string型のキーとint型の値を持つ.初期値は無い(空のmap)
	m["taguchi"] = 200
	m["fkoji"] = 300
	*/
	// map型のmを宣言し、初期化
	m := map[string]int{"taguchi":100, "fkoji":200}
	fmt.Println(m)
	fmt.Println(len(m)) // 要素数

	delete(m, "taguchi") // キーを指定して要素を削除
	fmt.Println(m)

	v, ok := m["fkoji"] // キーを指定して要素の存在確認
	fmt.Println(v) // 存在する場合は、値が返却される
	fmt.Println(ok) // 存在する場合は、真偽値が返却される

	dv, ng := m["taguchi"]
	fmt.Println(dv) // 存在しない場合は、型のデフォルト値が返却される
	fmt.Println(ng) 
}
