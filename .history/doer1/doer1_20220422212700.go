/*
 * @Author: zhounanjun
 * @Date: 2022-04-22 13:19:39
 * @LastEditors: zhounanjun
 * @LastEditTime: 2022-04-22 21:27:00
 * @Description: 请填写简介
 */

package doer1

// go:generate mockgen -destination=../mocks/mock_doer1.go -package=mocks -source=./doer1/doer1.go
type Doer1 interface {
	DoSomething1(int, string) error
}
