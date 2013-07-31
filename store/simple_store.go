package store

type SimpleStore struct {
  m map[string]string
}

func NewSimpleStore() (*SimpleStore) {
  return &SimpleStore{make(map[string]string)}
}

func (s *SimpleStore) Get(key string) (string, bool) {
  v, ok := s.m[key]
  return v,ok
}

func (s *SimpleStore) Set(key string, val string) {
  s.m[key] = val
}
