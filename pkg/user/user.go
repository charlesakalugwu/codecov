package user

import "github.com/charlesakalugwu/codecov/pkg/doer"

type User struct {
	Doer doer.Doer
}

func (u *User) Use() error {
	return u.Doer.DoSomething(123, "Hello GoMock")
}