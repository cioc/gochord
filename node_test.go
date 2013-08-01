package gochord

import (
  "testing"
  "fmt"
  "time"

  "gochord/rpc"
  "gochord/store"
)



func TestGetSet(t *testing.T) {
  go func(){
    node, err := NewNode(4067, store.NewSimpleStore(), 100)
    if err != nil {
      t.Log(err)
      t.Fail()
    }
    fmt.Printf("%v", node)
    err = node.Start()
    if err != nil {
      t.Log(err)
      t.Fail()
    }
  }()
  println("Sleeping for server to start...")
  time.Sleep(1000 * time.Millisecond)
  result, err := rpc.Set("http://localhost:4067", "hello", "hamburger")
  if err != nil {
    t.Log(err)
    t.Fail()
  }
  result, err = rpc.Get("http://localhost:4067", "hello")
  if err != nil {
    t.Log(err)
    t.Fail()
  }
  if result.Res != "hamburger" {
    t.Log(result.Res)
    t.Fail()
  }
}
