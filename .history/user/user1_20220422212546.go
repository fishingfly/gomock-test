/*
 * @Author: zhounanjun
 * @Date: 2022-04-22 13:20:13
 * @LastEditors: zhounanjun
 * @LastEditTime: 2022-04-22 21:25:45
 * @Description: 请填写简介
 */

package user

import "gomock-test/doer1"

type User1 struct {
	Doer doer1.Doer1
}

// go:generate mockgen -destination=../mocks/mock_doer1.go -package=mocks -source=../doer1/doer1.go  Doer1
func (u *User1) Use() error {
	return u.Doer.DoSomething1(123, "Hello GoMock")
}
