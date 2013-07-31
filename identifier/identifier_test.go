package identifier

import (
  "testing"
  "bytes"
)

func TestNewIdentifierNoMod(t *testing.T) {
  i, err := NewFromStr("hello hamburger", 0)
  if err != nil {
    t.Log(err)
    t.Fail()
  }
  decimalValue := "315683378581898004600115863185378131070858318829"
  if i.String() != decimalValue {
    t.Log("failed to match values", i.String())
    t.Fail()
  }
  var b bytes.Buffer
  b.Write([]byte("hello hamburger"))
  v, err := NewFromBytes(b.Bytes(), 0)
   if v.String() != decimalValue {
    t.Log("failed to match values", v.String())
    t.Fail()
  }
}

func TestNewIdentifierWithMod(t *testing.T) {
  i, err := NewFromStr("hello hamburger", 10)
  if err != nil {
    t.Log(err)
    t.Fail()
  }
  decimalValue := "1005"
  if i.String() != decimalValue {
    t.Log("failed to match values")
    t.Fail()
  }
  var b bytes.Buffer
  b.Write([]byte("hello hamburger"))
  v, err := NewFromBytes(b.Bytes(), 10)
   if v.String() != decimalValue {
    t.Log("failed to match values", v.String())
    t.Fail()
  }
}
