package singleton

type SingleTon interface {
	Increment() int
}

type singleton struct {
	digit int
}

func (s *singleton) Increment() int {
	s.digit++
	return s.digit
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}
