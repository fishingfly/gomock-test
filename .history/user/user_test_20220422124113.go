/*
 * @Author: zhounanjun
 * @Date: 2022-04-22 12:35:06
 * @LastEditors: zhounanjun
 * @LastEditTime: 2022-04-22 12:39:15
 * @Description: 请填写简介
 */

package user

import (
	"gomock-test/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	testUser.Use()
}
