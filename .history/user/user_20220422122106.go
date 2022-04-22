/*
 * @Author: zhounanjun
 * @Date: 2022-04-22 12:20:43
 * @LastEditors: zhounanjun
 * @LastEditTime: 2022-04-22 12:20:44
 * @Description: 请填写简介
 */
package user

import "gomock-test/doer"

type User struct {
	Doer doer.Doer
}

func (u *User) Use() error {
	return u.Doer.DoSomething(123, "Hello GoMock")
}
