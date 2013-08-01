//this file implements the finger table
package finger

import (
  "math/big"
  "gochord"
  "gochord/identifier"
)

type FingerEntry struct {
  Identifier *big.Int
  Addr string
}

type FingerTable struct {
  m int                     //circular constant m
  n  *big.Int               //identifier of the node who uses this table
  table map[int]FingerEntry
}

func New (n *big.Int, m int) (FingerTable) {
  return FingerTable{m, n, make(map[int]FingerTable)}
}

func (f FingerTable) Get(i int) (FingerEntry, bool) {
  return f.table[i]
}

func (f FingerTable) Set(i int, addr string) (error) {
  identifier, err := identifier.NewFromStr(addr, f.m)
  if err != nil {
    return err
  }
  f.table[i] = FingerEntry{identifier, addr}
  return nil
}

func (f FingerTable) EntryConstant(i int) (error) {
  two := big.NewInt(2)
  exp := big.NewInt(0)
  exp.Exp(two, big.NewInt(i))
  o := big.NewInt(0)
  o.Add(f.n, exp)
  exp.Exp(two, big.NewInt(f.m))
  o.Mod(o, exp)
  return o
}
