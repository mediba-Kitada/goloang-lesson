/* if
> >=
< <=
==
!=
&& || !
*/

package main

import "fmt"

func main() {

	// score := 43

	// 変数scoreのスコープをifステートメントに収める
	if score := 43; score > 80 {
		fmt.Println("Gread")
	} else if score > 60 {
		fmt.Println("Good")
	} else {
		fmt.Println("so so...")
	}

	//fmt.Println(score) スコープ外のためコンパイルエラーとなる
}
