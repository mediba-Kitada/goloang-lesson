/* スライス
C言語では、a配列の宣言を行うとaは、a[0]へのポインタとなる
Go言語では、a配列の宣言を行うとaは、代入された要素数を持つ値となる
-> 関数に配列を渡す場合には、値を丸ごと渡すことになり非効率
-> 要素数も動的に変化されることが出来ず、不便

スライスは、効率的にメモリを利用するためのデータ型
-> 配列の一部または全部を指し示す参照型のデータ
*/

package main

import "fmt"

func main() {
	a := [5]int{2, 10, 8, 15, 4}
	sa := a[2:4] // 添字を指定しスライスする要素の範囲を指定. 2から4(最後の要素は含まない)
	sa[1] = 12
	fmt.Println(a) // スライスされ、代入された配列aは破壊的に処理されている
	fmt.Println(sa)

	fmt.Println(len(sa)) // スライスの要素数
	fmt.Println(cap(sa)) // スライスの最大容量 2-4をスライスしたので、3要素が容量となる

	// make関数を用いてスライスを宣言
	//s := make([]int, 3) // 初期値[0 0 0] 要素数3
	s := []int{1, 3, 5} // 配列の宣言とよく似ているが、要素数が無い

	// append
	s = append(s, 5, 2, 10)

	// copy
	t := make([]int, len(s)) // sと同じキャパシティのスライスを宣言
	cst := copy(t, s) // copy関数を用いてtスライスにsスライスの値を複製したcstスライスを宣言
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(t, s) //tスライスとsスライスは同一の値を保持している
	fmt.Println(cst)

}
