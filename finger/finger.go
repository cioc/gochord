//this file implements the finger table
package finger

import (
  "math/big"
  "gochord/identifier"
)

type FingerEntry struct {
  Identifier *big.Int
  Addr string
}

type Interval struct {
  Closed *big.Int
  Open *big.Int
}

type FingerTable struct {
  m int64                   //circular constant m
  n  *big.Int               //identifier of the node who uses this table
  table map[int]FingerEntry
}

func New(n *big.Int, addr string, m int64) (FingerTable) {
  f := FingerTable{m, n, make(map[int]FingerEntry)}
  //0th entry of the finger table is the node itself
  f.table[0] = FingerEntry{n, addr}
  return f
}

//get the entry at the table
func (f FingerTable) Node(k int) (FingerEntry, bool) {
  v, ok := f.table[k]
  return v, ok
}

func (f FingerTable) Set(k int, addr string) (error) {
  identifier, err := identifier.NewFromStr(addr, f.m)
  if err != nil {
    return err
  }
  f.table[k] = FingerEntry{identifier, addr}
  return nil
}

func (f FingerTable) Successor() (FingerEntry, bool) {
  //successor is the first entry after the node itself in the node's finger table
  v, ok := f.table[1]
  return v,ok
}

func (f FingerTable) Start(k int64) (*big.Int) {
  //(f.n + 2 ^ k) mod 2 ^ f.m
  two := big.NewInt(2)
  exp := big.NewInt(0)
  exp.Exp(two, big.NewInt(k), nil)
  o := big.NewInt(0)
  o.Add(f.n, exp)
  exp.Exp(two, big.NewInt(f.m), nil)
  o.Mod(o, exp)
  return o
}

func (f FingerTable) Interval(k int64) (Interval) {
  return Interval{f.Start(k), f.Start(k + 1)}
}

