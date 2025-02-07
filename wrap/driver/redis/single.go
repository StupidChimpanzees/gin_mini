package redis

import (
	"gin_work/wrap/config"
	"gin_work/wrap/driver"
)

type Single struct {
	readser *driver.Reads
}

func NewRedis() *Single {
	var reads *Single
	conf := reads.getConfig()
	reads.readser = driver.NewReads(conf.Host, conf.Port, conf.Password, conf.Prefix, conf.Select)
	return reads
}

func (*Single) getConfig() config.RedisConfiguration {
	return config.Mapping.Redis
}

func (s *Single) setFullName(name string) string {
	return s.readser.Prefix + name
}

func (s *Single) Set(name, value string) (string, error) {
	return s.readser.Set(s.setFullName(name), value)
}

func (s *Single) Get(name string) (string, error) {
	return s.readser.Get(s.setFullName(name))
}

func (s *Single) HSet(name string, value interface{}) (string, error) {
	return s.readser.HSet(s.setFullName(name), value)
}

func (s *Single) HGet(name string) (interface{}, error) {
	return s.readser.HGet(s.setFullName(name))
}

func (s *Single) LPush(name string, args ...interface{}) (bool, error) {
	return s.readser.LPush(s.setFullName(name), args...)
}

func (s *Single) LPop(name string) (interface{}, error) {
	return s.readser.LPop(s.setFullName(name))
}

func (s *Single) LRange(name string, start int, stop int) ([]interface{}, error) {
	return s.readser.LRange(s.setFullName(name), start, stop)
}

func (s *Single) SAdd(name string, args ...interface{}) (bool, error) {
	return s.readser.SAdd(s.setFullName(name), args...)
}

func (s *Single) SPop(name string) (interface{}, error) {
	return s.readser.SPop(s.setFullName(name))
}

func (s *Single) ZAdd(name string, key string, value string) (bool, error) {
	return s.readser.ZAdd(s.setFullName(name), key, value)
}

func (s *Single) ZRange(name string, start int, stop int, args ...bool) ([]interface{}, error) {
	return s.readser.ZRange(s.setFullName(name), start, stop, args...)
}

func (s *Single) ZScore(name string, key string) (interface{}, error) {
	return s.readser.ZScore(s.setFullName(name), key)
}

func (s *Single) Exists(name string) (bool, error) {
	return s.readser.Exists(s.setFullName(name))
}

func (s *Single) Del(name string) error {
	return s.readser.Del(s.setFullName(name))
}

func (s *Single) Clear() error {
	return s.readser.Clear()
}

func (s *Single) Multi() error {
	return s.readser.Multi()
}

func (s *Single) Exec() error {
	return s.readser.Exec()
}

func (s *Single) Discard() error {
	return s.readser.Discard()
}
