/* 配列
var a1 int
var a2 int
...面倒
*/

package main

import "fmt"

func main() {
	var a [5]int // a[0] - a[4] intなので0がデフォルトで代入されている
	a[2] = 3 // 添字を指定して代入
	a[4] = 10
	b := [3]int{1, 3, 5} // 要素数を指定して複数代入
	c := [...]int{1, 3, 5, 8} // 配列の宣言はスリードットでも表現出来る
	fmt.Println(a, b, c)
	fmt.Println(len(c)) // 配列cの要素数
}
