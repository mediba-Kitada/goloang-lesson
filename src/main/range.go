/* range
配列やスライス、マップに対し繰り返し処理を実装する際に利用
*/

package main

import "fmt"

func main() {

	s := []int{2, 3, 8}
	for i, v := range s { // スライスsの添字をiに、値をvに代入
		fmt.Println(i, v) // 0 2
	}

	for _, v := range s { // スライスsの添字はブランク修飾子に代入(ブランク)
		fmt.Println(v)
	}

	m := map[string]int{"taguchi":200, "fkoji":300} // map mの宣言
	for k, v := range m {
		fmt.Println(k,v)
	}
}
