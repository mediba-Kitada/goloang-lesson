package main

import (
	"fmt"
	"net/http"
)

/* ハンドラ関数
第一引数: http.ResponseWriter型 io.Writerなのでfmt.Fprintf()関数を用いて文字列を書き込むことが出来る
第二引数: Request型 リクエスト情報は、ポインタ型として渡ってくる
*/
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
}

func main() {
	// httpパッケージのHandleFunc関数の利用
	/* http.HandleFunc()関数でルーティング
	第一引数: パスのパターン パターンに一致したハンドラが実行される
	第二引数: 蓋数の引数を受け取る関数を指定
	*/
	http.HandleFunc("/", IndexHandler) // ハンドル設定 待受パスを第一引数で指定 handler関数を第二引数として渡す
	http.ListenAndServe(":8080", nil) // 待受ポート番号を第一引数として渡す
}
