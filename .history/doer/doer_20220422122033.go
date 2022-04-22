/*
 * @Author: zhounanjun
 * @Date: 2022-04-22 12:20:23
 * @LastEditors: zhounanjun
 * @LastEditTime: 2022-04-22 12:20:24
 * @Description: 请填写简介
 */

package doer

type Doer interface {
	DoSomething(int, string) error
}
