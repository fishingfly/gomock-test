/*
 * @Author: zhounanjun
 * @Date: 2022-04-22 12:35:06
 * @LastEditors: zhounanjun
 * @LastEditTime: 2022-04-22 21:50:50
 * @Description: 请填写简介
 */

package user

import (
	"errors"
	"fmt"
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

func TestUseReturnsErrorFromDo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dummyError := errors.New("dummy error")
	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return dummyError from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError).Times(1) // 调用一次返回结果

	err := testUser.Use()

	if err != dummyError {
		t.Fail()
	}
}

func TestUserReturnErrorsInOrder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dummyError := errors.New("dummy error")
	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}
	gomock.InOrder(
		mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError),
		mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil),
		mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError),
		mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil),
	)
	err1 := testUser.Use()

	if err1 != dummyError {
		t.Fail()
	}
	err2 := testUser.Use()

	if err2 != nil {
		t.Fail()
	}
	err3 := testUser.Use()

	if err3 != dummyError {
		t.Fail()
	}
	err4 := testUser.Use()

	if err4 != nil {
		t.Fail()
	}

}

// 指定mock行为
func TestUserReturnErrorsSpecify(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}

	/*
		关于调用参数的一些复杂验证逻辑可以写在.Do的函数中，
		例如DoSomething的第一个int类型的参数应该小于等于
		第二个string类型参数的长度，那么我们就可以这样来实现
	*/
	mockDoer.EXPECT().
		DoSomething(gomock.Any(), gomock.Any()).
		Return(nil).
		Do(func(x int, y string) {
			fmt.Println(x)
			fmt.Println(y)
			if x > len(y) {
				t.Fail()
			}
		})

	testUser.Use()

}
