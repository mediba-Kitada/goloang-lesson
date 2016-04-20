/* 何らかのパッケージに属している必要がある
   mainパッケージが必要
*/
package main

// fmt
import "fmt"

/*
基本的なデータ型

string	"hello"
int	53
float64 10.2
bool false true
nil

それぞれの型にはデフォルト値が設定されている
var s string // ""
var a int // 0
var f bool // false
*/

/* 演算
数値 + - * / %
文字列 +
論理値 AND(&&) OR(||) NOT(!)
*/

// mainパッケージのmain関数があればそれを実装する規約
func main() {
	//var msg string
	//msg = "hello world"
	//var msg = "hello world"
	msg := "hello world"
	fmt.Println(msg)

	//var a, b int
	// 多重代入
	//a, b = 10, 15
	a, b := 10, 15
	fmt.Println(a,b)

	var (
		c float64
		d string
	)

	c = 20
	d = "hoge"

	fmt.Println(c,d)

	var e bool

	// 書式付き標準出力
	// 整数値:%d
	// float %f
	// 文字列 %s
	// 真偽値 %t
	fmt.Printf("a:%d, c:%f, d:%s, e:%t\n", a, c, d, e)

	// 定数
	//const name = "taguchi"
	//name = "fkoji"
	//fmt.Println(name)
	
	// iota
	// iotaを代入した後に宣言した定数にインクリメントした値を代入;
	const (
		sun = iota
		mon
		tue
	)
	fmt.Println(sun, mon, tue)

	var x int
	x = 10 % 3
	x += 3
	x ++
	fmt.Println(x)

	var s string
	s = "hello " + "world"
	fmt.Println(s)

	t := true
	u := false
	fmt.Println(t && u)
	fmt.Println(t || u)
	fmt.Println(!t)
}
