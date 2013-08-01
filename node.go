package gochord

import (
  _ "os"

  "net/http"
  "math/big"
  "strconv"
  "log"
  "fmt"

  "gochord/identifier"
  "gochord/store"
  "gochord/rpc"
)

type Node struct {
  addr string
  port int
  m int64               //size of circular hash
  identifier *big.Int
  store store.Store
  finger map[int]string //finger table
}

func NewNode(port int, store store.Store, m int64) (*Node, error) {
  ip := "127.0.0.1:"+strconv.Itoa(port)
  addr := "127.0.0.1"
  ident, err := identifier.NewFromStr(ip, m)
  if err != nil {
    return &Node{}, err
  }
  return &Node{addr, port, m, ident, store, make(map[int]string)}, nil
}

//Starts a new chord group and starts the local http server to handle rpc requests
func (n *Node) Start() {
  http.HandleFunc("/get/", bindNode(getHandle, n))
  http.HandleFunc("/set", bindNode(setHandle, n))
  portStr := ":"+strconv.Itoa(n.port)
  log.Fatal(http.ListenAndServe(portStr, nil))
}

func getHandle(res http.ResponseWriter, req *http.Request, n *Node) {
  pieces, err := rpc.ParseURL(req.URL.Path, len("/get/"))
  if err != nil {
    fmt.Println(err)
    return
  }
  if len(pieces) <= 0 {
    fmt.Println("no key")
    buf, err := rpc.EncodeRes("", "no key specified")
      if err != nil {
      fmt.Println(err)
    }
    fmt.Fprintf(res, buf.String())
    return
  }
  val, ok := n.store.Get(pieces[0])
  var result, e string
  if ok {
    result = val
    e = ""
  } else {
    result = ""
    e = "key not in store"
  }
  buf, err := rpc.EncodeRes(result, e)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Fprintf(res, buf.String())
}

func setHandle(res http.ResponseWriter, req *http.Request, n *Node) {
  reqrpc, err := rpc.ParseReq(req.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  n.store.Set(reqrpc.Key, reqrpc.Val)
  buf, err := rpc.EncodeRes("success","")
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Fprintf(res, buf.String())
}

//binds a node to http handler
func bindNode(f func(http.ResponseWriter, *http.Request, *Node) (), n *Node) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    f(w,r,n)
  }
}

func (n *Node) String() (string) {
  o := "{\n"
  o += n.addr + ":" + strconv.Itoa(n.port) + "\n"
  o += n.identifier.String() + "\n"
  o += "m: "+strconv.Itoa(int(n.m)) + "\n"
  o += "}\n"
  return o
}
