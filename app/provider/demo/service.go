package demo

import (
	"github.com/bingxindan/bxd_data_access/framework"
)

// 具体的接口实例
type Service struct {
	container framework.Container
}

// 初始化实例的方法
func NewService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	return &Service{container: container}, nil
}

// 实现接口
func (s *Service) GetAllStudent() []Student {
	return []Student{
		{
			ID:   1,
			Name: "foo",
		},
		{
			ID:   2,
			Name: "bar",
		},
	}
}
