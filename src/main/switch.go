// switch

package main

import "fmt"

func main() {

	signal := "www"

	switch signal {
	case "red":
		fmt.Println("stop")
	case "yellow":
		fmt.Println("caution")
	case "blue", "grean":
		fmt.Println("Go")
	default:
		fmt.Println("wrong signal")
	}

	// ifステートメントと同様に書いてみる
	score := 82
	switch { // 変数scoreは省略可能
	case score > 80: // 条件式をcaseの後に設定
		fmt.Println("Great")
	default:
		fmt.Println("so so...")
	}
}
