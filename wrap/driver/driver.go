package driver

type Driver interface {
	Exists(name string) (bool, error)
	Set(name, value string, args ...interface{}) (string, error)
	Get(name string) (string, error)
	Del(name string) error
	Clear() error
}
