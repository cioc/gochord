package rpc

import (
  "net/http"
)

//gets val stored at key on host
func Get(host, key string) (RPCRes, error) {
  resp, err := http.Get(host + "/get/" + key)
  if err != nil {
    return RPCRes{}, err
  }
  defer resp.Body.Close()
  return ParseRes(resp.Body)
}

//sets key to val on host
func Set(host, key, val string) (RPCRes, error) {
  buf, err := EncodeReq(key, val)
  if err != nil {
    return RPCRes{}, err
  }
  resp, err := http.Post(host + "/set", "application/json", buf)
  if err != nil {
    return RPCRes{}, err
  }
  defer resp.Body.Close()
  return ParseRes(resp.Body)
}
