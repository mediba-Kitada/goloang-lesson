/* goroutine:並行処理
main関数がgoroutineより先に終わってしまう問題...
*/

// channel:goroutineを用いた並行処理で利用できるパイプのようなもの

package main

import (
	"fmt"
	"time" // timeパッケージのインポート
)

func task1(result chan string) {
	time.Sleep(time.Second * 2)
	//fmt.Println("task1 finished")
	result<- "task1 finished"
}

func task2() {
	fmt.Println("task2 finished")
}

func main() {

	// make関数を用いてchannelを宣言 代入される値はstring型とする
	result := make(chan string)

	go task1(result) // channel resultを渡す バックグラウンドで処理される task2より所要時間が長いので、後に完了する
	go task2() // バックグラウンドで処理される task1より所要時間は短いので、先に完了する

	// フォアグランドで処理される(最初に実行される)
	fmt.Println(<-result) // channel resultに値が代入されるまで待つ

	time.Sleep(time.Second * 3)
}
