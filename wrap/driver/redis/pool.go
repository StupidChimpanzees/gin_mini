package redis

import (
	"gin_work/wrap/config"
	"gin_work/wrap/driver"
	"gin_work/wrap/log"
	"time"

	"github.com/gomodule/redigo/redis"
)

type poolConf struct {
	enable         bool
	maxIdle        int
	maxActive      int
	idleTimeout    time.Duration
	maxConnTimeout time.Duration
}

var (
	pool     *redis.Pool
	PoolConf *poolConf
)

func init() {
	conf := PoolConf.getPoolConfig()
	// redis pool关闭
	if conf.enable == false {
		return
	}
	rc := PoolConf.getConfig()
	reads := driver.NewReads(rc.Host, rc.Port, rc.Password, rc.Prefix, rc.Select)
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := reads.Connection()
			return conn, err
		},
		MaxIdle:         conf.maxIdle,
		MaxActive:       conf.maxActive,
		IdleTimeout:     conf.idleTimeout,
		MaxConnLifetime: conf.maxConnTimeout,
		Wait:            true,
	}
}

func (*poolConf) getPoolConfig() *poolConf {
	pConfig := config.Mapping.Redis.Pool
	conf := &poolConf{
		maxIdle:        10,
		maxActive:      100,
		idleTimeout:    time.Duration(120),
		maxConnTimeout: time.Duration(3),
	}
	if pConfig.MaxIdle != 0 {
		conf.maxIdle = pConfig.MaxIdle
	}
	if pConfig.MaxActive != 0 {
		conf.maxActive = pConfig.MaxActive
	}
	if pConfig.IdleConnTimeout != 0 {
		conf.idleTimeout = time.Duration(pConfig.IdleConnTimeout)
	}
	if pConfig.MaxConnTimeout != 0 {
		conf.maxConnTimeout = time.Duration(pConfig.MaxConnTimeout)
	}
	return conf
}

func (*poolConf) getConfig() config.RedisConfiguration {
	return config.Mapping.Redis
}

func (*poolConf) setFullName(name string) string {
	return config.Mapping.Redis.Prefix + name
}

func (*poolConf) Reads() (*driver.Reads, error) {
	var reads *driver.Reads
	_, err := reads.Connection(pool.Get())
	if err != nil {
		log.Warning(err.Error())
	}
	return reads, err
}

func (p *poolConf) Set(name, value string) (string, error) {
	reads, err := p.Reads()
	if err != nil {
		return "", err
	}
	return reads.Set(p.setFullName(name), value)
}

func (p *poolConf) Get(name string) (string, error) {
	reads, err := p.Reads()
	if err != nil {
		return "", err
	}
	return reads.Get(p.setFullName(name))
}

func (p *poolConf) HSet(name string, value interface{}) (string, error) {
	reads, err := p.Reads()
	if err != nil {
		return "", err
	}
	return reads.HSet(p.setFullName(name), value)
}

func (p *poolConf) HGet(name string) (interface{}, error) {
	reads, err := p.Reads()
	if err != nil {
		return nil, err
	}
	return reads.HGet(p.setFullName(name))
}

func (p *poolConf) LPush(name string, args ...interface{}) (bool, error) {
	reads, err := p.Reads()
	if err != nil {
		return false, err
	}
	return reads.LPush(p.setFullName(name), args...)
}

func (p *poolConf) LPop(name string) (interface{}, error) {
	reads, err := p.Reads()
	if err != nil {
		return false, err
	}
	return reads.LPop(p.setFullName(name))
}

func (p *poolConf) LRange(name string, start int, stop int) ([]interface{}, error) {
	reads, err := p.Reads()
	if err != nil {
		return nil, err
	}
	return reads.LRange(p.setFullName(name), start, stop)
}

func (p *poolConf) SAdd(name string, args ...interface{}) (bool, error) {
	reads, err := p.Reads()
	if err != nil {
		return false, err
	}
	return reads.SAdd(p.setFullName(name), args...)
}

func (p *poolConf) SPop(name string) (interface{}, error) {
	reads, err := p.Reads()
	if err != nil {
		return false, err
	}
	return reads.SPop(p.setFullName(name))
}

func (p *poolConf) ZAdd(name string, key string, value string) (bool, error) {
	reads, err := p.Reads()
	if err != nil {
		return false, err
	}
	return reads.ZAdd(p.setFullName(name), key, value)
}

func (p *poolConf) ZRange(name string, start int, stop int, args ...bool) ([]interface{}, error) {
	reads, err := p.Reads()
	if err != nil {
		return nil, err
	}
	return reads.ZRange(p.setFullName(name), start, stop, args...)
}

func (p *poolConf) ZScore(name string, key string) (interface{}, error) {
	reads, err := p.Reads()
	if err != nil {
		return nil, err
	}
	return reads.ZScore(p.setFullName(name), key)
}

func (p *poolConf) Exists(name string) (bool, error) {
	reads, err := p.Reads()
	if err != nil {
		return false, err
	}
	return reads.Exists(p.setFullName(name))
}

func (p *poolConf) Del(name string) error {
	reads, err := p.Reads()
	if err != nil {
		return err
	}
	return reads.Del(p.setFullName(name))
}

func (p *poolConf) Clear() error {
	reads, err := p.Reads()
	if err != nil {
		return err
	}
	return reads.Clear()
}

func (p *poolConf) Multi() error {
	reads, err := p.Reads()
	if err != nil {
		return err
	}
	return reads.Multi()
}

func (p *poolConf) Exec() error {
	reads, err := p.Reads()
	if err != nil {
		return err
	}
	return reads.Exec()
}

func (p *poolConf) Discard() error {
	reads, err := p.Reads()
	if err != nil {
		return err
	}
	return reads.Discard()
}
