/* package
色々ある https://golang.org/pkg/
*/

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { // リクエスト情報は、ポインタ型として渡ってくる
	fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
}

func main() {
	// httpパッケージのHandleFunc関数の利用
	http.HandleFunc("/", handler) // ハンドル設定 待受パスを第一引数で指定 handler関数を第二引数として渡す
	http.ListenAndServe(":8080", nil) // 待受ポート番号を第一引数として渡す
}
