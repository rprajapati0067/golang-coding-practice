package sets

var exists = struct{}{}

type Set struct {
	m map[interface{}]struct{}
}

func NewSet() *Set {
	s := &Set{}
	s.m = make(map[interface{}]struct{})
	return s
}

func (s *Set) Add(value interface{}) {
	s.m[value] = exists
}
func (s *Set) Remove(value interface{}) {
	delete(s.m, value)
}
func (s *Set) Contains(value interface{}) bool {
	_, c := s.m[value]
	return c
}
