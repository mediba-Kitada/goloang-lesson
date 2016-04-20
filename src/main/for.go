/* for
golangには繰り返し表現の手段がforステートメント以外無い
while やdo/whileもforを利用する
*/

package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ { // 条件式に丸括弧は不要
		// if i == 3 { break } // ループを抜ける
		if i == 3 { continue } // ループをスキップ
		fmt.Println(i)
	}

	i := 0
	for i < 10 { // 条件式を簡略することも出来る
		fmt.Println(i)
		i++
	}

	j := 0
	for { // 無限ループ
		fmt.Println(j)
		j++
		if j == 3 { break } // 無限ループを抜ける
	}
}
