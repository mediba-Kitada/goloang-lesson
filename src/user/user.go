package user

import(
	"fmt"
)

type User struct {
	FirstName string
	LastName string
}

func (u *User) FullName() (fullname string) { // 構造体Userのメソッド
	fullname = fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return
}

func NewUser(firstname, lastname string) *User { // 構造体Userのコンストラクタ
	return &User{
		FirstName: firstname,
		LastName: lastname,
	}
}
