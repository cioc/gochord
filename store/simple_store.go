package store

import (
  "sync"
)

type SimpleStore struct {
  m map[string]string
  lock sync.Mutex
}

func NewSimpleStore() (*SimpleStore) {
  return &SimpleStore{make(map[string]string), Mutex{}}
}

func (s *SimpleStore) Get(key string) (string, bool) {
  v, ok := s.m[key]
  return v,ok
}

func (s *SimpleStore) Set(key string, val string) {
  s.m[key] = val
}
