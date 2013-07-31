package store

import (
  "testing"
)

func TestSimpleStore(t *testing.T) {
  s := NewSimpleStore()
  s.Set("hello","hamburger")
  v, ok := s.Get("test")
  if ok {
    t.Fail()
  }
  v, ok = s.Get("hello")
  if v != "hamburger" || !ok {
    t.Fail()
  }
}
