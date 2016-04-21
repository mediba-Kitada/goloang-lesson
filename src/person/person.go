package person

type Person struct {
	ID	int `json:"id"` // encoding/jsonパッケージのタグ付け機能 出力するオブジェクト名を指定
	Name	string `json:"name"`
	Email	string `json:"-"`
	Age	int `json:"age"`
	Address	string `json:"address,omitempty"`
	memo	string
}