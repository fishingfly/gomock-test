/*
 * @Author: zhounanjun
 * @Date: 2022-04-22 13:20:13
 * @LastEditors: zhounanjun
 * @LastEditTime: 2022-04-22 13:20:13
 * @Description: 请填写简介
 */

 package user

import "gomock-test/doer"

type User1 struct {
	Doer doer.doer1
}

func (u *User1) Use() error {
	return u.Doer1.DoSomething(123, "Hello GoMock")
}