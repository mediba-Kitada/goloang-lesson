package task

import(
	"fmt"
	. "user"
)

/*
type ID int // task.ID型の宣言
type Priority int // task.Priority型の宣言
*/

type Task struct { // 構造体Taskの宣言
	ID	int // パブリック(パッケージ外からアクセス可能)なプロパティ int型のIDを宣言
	Detail	string
	done	bool // クローズド(パッケージ内でのみアクセス)なプロパティ bool型のdoneを宣言
	*User // User型の埋め込み
}

// インターフェイス
type Stringer interface { // String()メソッドを実装していることを示す erをつける慣習
	String() string
}

// コンストラクタ 構造体の名前にNewプレフィックスをつける
func NewTask(id int, detail, firstname, lastname string,) (task *Task) { // 戻り値にTaskコンストラクタのポイントである変数taksを指定
	task = &Task{
		ID:	id,
		Detail:	detail,
		done:	false,
		User:	NewUser(firstname, lastname),
	}
	return // 名前付き戻り値を指定しているので省略可能
}

/*
func ProcessTask(id ID, priority Priority) {
	fmt.Println(id, priority)
}
*/

// メソッド
func (task Task) String() (str string) { // 構造体Taskの文字列表現を司る
	str = fmt.Sprintf("%d) %s", task.ID, task.Detail) // 対象の型(Task)をレシーバとして受け取っているので、内部で利用可能となる
	return
}

func (task *Task) Finish() { // レシーバは構造体Taskのコピーとして渡されるため、変更を反映する必要がある場合は、ポインタを指定
	task.done = true
}

func (task Task) Print(value interface{}) { // インターフェイス型を引数に取る
	s, ok := value.(Stringer) //Type Assertion 第一戻り値に変換された値 第二戻り値に型判定の可否
	if ok {
		fmt.Printf("param is Stringer %s:", s)
		fmt.Println(s.String())
	} else {
		fmt.Print("value is not Stringer\n")
	}

	switch v := value.(type) { //Type Switch 型での分岐
	case string: // Type Assertionでは、単一の型しか検査出来ないが、複数の型を検査でき、処理を分岐出来る
		fmt.Printf("value is string: %s\n", v)
	case int:
		fmt.Printf("value is int: %d\n", v)
	case Stringer:
		fmt.Printf("value is Stringer: %s\n", v)
	}
}
