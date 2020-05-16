package cache

type ICache interface {
	Set(key string, val string) error
	Get(key string) (string, error)
}