/*
 * @Author: zhounanjun
 * @Date: 2022-04-22 22:59:48
 * @LastEditors: zhounanjun
 * @LastEditTime: 2022-04-22 23:02:31
 * @Description: 请填写简介
 */

package user

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		Result int
		Data   [2]int
	}{
		{Result: 2, Data: [2]int{1, 1}},
		{Result: 3, Data: [2]int{1, 2}},
		{Result: 4, Data: [2]int{1, 3}},
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
