package gochord

import (
  _ "os"
  _ "net"
  "math/big"
  "strconv"
  "gochord/identifier"
  "gochord/store"
)

type Node struct {
  addr string
  port int
  m int64               //size of circular hash
  identifier *big.Int
  s *store.Store
  finger map[int]string //finger table
}

func NewNode(port int, store *store.Store, m int64) (Node, error) {
  ip := "127.0.0.1:"+strconv.Itoa(port)
  addr := "127.0.0.1"
  ident, err := identifier.NewFromStr(ip, m)
  if err != nil {
    return Node{}, err
  }
  return Node{addr, port, m, ident, store, make(map[int]string)}, nil
}

func (n Node) String() (string) {
  o := "{\n"
  o += n.addr + ":" + strconv.Itoa(n.port) + "\n"
  o += n.identifier.String() + "\n"
  o += "m: "+strconv.Itoa(int(n.m)) + "\n"
  o += "}\n" 
  return o
}
