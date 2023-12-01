package demo

// Demo服务的key
const DemoKey = "demo"

// IService 服务的接口
type IService interface {
	GetAllStudent() []Student
}

// Student 服务接口定义的一个数据结构
type Student struct {
	ID   int
	Name string
}
