package identifier

import (
  "math/big"
  "crypto/sha1"
  "hash"
  "io"
  "errors"
)

func NewFromStr(buf string, m int64) (*big.Int, error) {
  h := sha1.New()
  l, err := io.WriteString(h, buf)
  if err != nil {
    return nil, err
  }
  if l != len(buf) {
    return nil, errors.New("did not write all characters of buf")
  }
  return hashAndMod(h, m), nil
}

func NewFromBytes(buf []byte, m int64) (*big.Int, error) {
  h := sha1.New()
  l, err := h.Write(buf)
  if err != nil {
    return nil, err
  }
  if l != len(buf) {
    return nil, errors.New("did not write all bytes of buf")
  }
  return hashAndMod(h, m), nil
}

func hashAndMod(h hash.Hash, m int64) (*big.Int) {
  i := big.NewInt(0)
  i.SetBytes(h.Sum(nil))
  if m > 0 {
    z := big.NewInt(0)
    two := big.NewInt(2)
    exponent := big.NewInt(m)
    z.Exp(two, exponent, nil)
    i.Mod(i, z)
  }
  return i
}
