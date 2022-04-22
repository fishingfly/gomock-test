/*
 * @Author: zhounanjun
 * @Date: 2022-04-22 22:59:48
 * @LastEditors: zhounanjun
 * @LastEditTime: 2022-04-22 23:03:39
 * @Description: 请填写简介
 */

package user

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
		Data   [2]int
	}{
		{Name: "test1", Result: 2, Data: [2]int{1, 1}},
		{Name: "test2", Result: 3, Data: [2]int{1, 2}},
		{Name: "test3", Result: 4, Data: [2]int{1, 3}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Mul(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}
