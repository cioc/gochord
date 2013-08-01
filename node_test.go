package gochord

import (
  "testing"
  "fmt"

  "gochord/store"
)

func TestGetSet(t *testing.T) {
  node, err := NewNode(8080, store.NewSimpleStore(), 100)
  if err != nil {
    t.Log(err)
    t.Fail()
  }
  fmt.Printf("%v", node)
}
