package mock

import (
	"errors"
	gomock "github.com/golang/mock/gomock"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctl)

	// 期望值->请求参数->返回值
	// m.EXPECT.Get().Return(arg1,error)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exit"))

	/*
		o1 := m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
		o2 := m.EXPECT().Get(gomock.Eq("Sam")).Return(630, nil)
		gomock.InOrder(o1, o2)
		GetFromDB(m, "Tom")
		GetFromDB(m, "Sam")

	*/

}
