package creational

type singletonStruct struct {
	count int
}

type Singleton interface {
	AddOne() interface{}
}

var instance *singletonStruct = nil

func GetInstance() *singletonStruct {
	if instance == nil {
		instance = new(singletonStruct)
	}

	return instance
}

func (s *singletonStruct) AddOne() {
	s.count++
}

func (s *singletonStruct) GetCount() int {
	return s.count
}
