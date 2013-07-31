package rpc

import (
  "encoding/json"
  "io"
  "bytes"
  "errors"
)

type RPCReq struct {
  Key string
  Val string
}

type RPCRes struct {
  Res string
  Err string
}

func ParseReq(body io.ReadCloser) (RPCReq, error) {
  decoder := json.NewDecoder(body)
  var v map[string]interface{}
  if err := decoder.Decode(&v); err != nil {
    return RPCReq{}, err
  }
  o := RPCReq{}
  key, ok := v["key"]
  if ok {
    switch key := key.(type) {
      case string:
        o.Key = key
      default:
        return RPCReq{}, errors.New("key was not a string")
    }
  }
  val, ok := v["val"]
  if ok {
    switch val := val.(type) {
      case string:
        o.Val = val
      default:
        return RPCReq{}, errors.New("val was not a string")
    }
  }
  return o, nil
}

func ParseRes(body io.ReadCloser) (RPCRes, error) {
  decoder := json.NewDecoder(body)
  var v map[string]interface{}
  if err := decoder.Decode(&v); err != nil {
    return RPCRes{}, err
  }
  o := RPCRes{}
  res, ok := v["res"]
  if ok {
    switch res := res.(type) {
      case string:
        o.Res = res
      default:
        return RPCRes{}, errors.New("res was not a string")
    }
  }
  err, ok := v["err"]
  if ok {
    switch err := err.(type) {
      case string:
        o.Err = err
      default:
        return RPCRes{}, errors.New("err was not a string")
    }
  }
  return o, nil
}

func EncodeRes(res, e string) (*bytes.Buffer, error) {
  o := RPCRes{res, e}
  params, err := json.Marshal(o)
  if err != nil {
    return nil, err
  }
  return bytes.NewBuffer(params), nil
}

func EncodeReq(key,val string) (*bytes.Buffer, error) {
  o := RPCReq{key,val}
  params, err := json.Marshal(o)
  if err != nil {
    return nil, err
  }
  return bytes.NewBuffer(params), nil
}
