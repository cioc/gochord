package store

type Store interface {
  Get(key string) (string, bool)
  Set(key string, val string)
}
