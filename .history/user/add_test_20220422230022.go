/*
 * @Author: zhounanjun
 * @Date: 2022-04-22 22:59:48
 * @LastEditors: zhounanjun
 * @LastEditTime: 2022-04-22 22:59:49
 * @Description: 请填写简介
 */

package user

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, -3, -6},
		{"zero", 2, 0, 0},
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
