package driver

import (
	"gin_work/wrap/log"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

type Reads struct {
	Host     string
	Port     int
	Password string
	Prefix   string
	Select   int
	Conn     redis.Conn
}

func NewReads(host string, port int, password string, prefix string, libName int) *Reads {
	return &Reads{
		Host:     host,
		Port:     port,
		Password: password,
		Prefix:   prefix,
		Select:   libName,
	}
}

func (r *Reads) Connection(conn ...redis.Conn) (redis.Conn, error) {
	if conn != nil {
		r.Conn = conn[0]
		return r.Conn, nil
	}
	var err error
	if r.Password == "" {
		r.Conn, err = redis.Dial("tcp", r.Host+":"+strconv.Itoa(r.Port))
	} else {
		r.Conn, err = redis.Dial("tcp", r.Host+":"+strconv.Itoa(r.Port), redis.DialPassword(r.Password))
	}
	if err != nil {
		log.Error(err.Error())
		panic("redis connection error: " + err.Error())
	}
	return r.Conn, err
}

func (r *Reads) GetConnection() redis.Conn {
	if r.Conn != nil {
		conn, _ := r.Connection(r.Conn)
		return conn
	}
	conn, _ := r.Connection()
	return conn
}

func (r *Reads) Command(command string, args ...interface{}) (interface{}, error) {
	conn := r.GetConnection()
	defer r.Close()
	_ = r.SelectDB(r.Select)
	return conn.Do(command, args...)
}

func (r *Reads) rSet(command string, name string, value interface{}, args ...interface{}) (string, error) {
	conn := r.GetConnection()
	defer r.Close()
	_ = r.SelectDB(r.Select)
	if args != nil && args[0] != 0 {
		return redis.String(conn.Do(command, name, value, "EX", args[0]))
	} else {
		return redis.String(conn.Do(command, name, value))
	}
}

func (r *Reads) SelectDB(num int) error {
	conn := r.GetConnection()
	_, err := redis.String(conn.Do("SELECT", num))
	return err
}

func (r *Reads) Set(name, value string, args ...interface{}) (string, error) {
	return r.rSet("SET", name, value, args...)
}

func (r *Reads) Get(name string) (string, error) {
	return redis.String(r.Command("GET", name))
}

func (r *Reads) HSet(name string, value interface{}, args ...interface{}) (string, error) {
	return r.rSet("HSET", name, value, args...)
}

func (r *Reads) HGet(name string) (interface{}, error) {
	return r.Command("HGET", name)
}

func (r *Reads) LPush(name string, args ...interface{}) (bool, error) {
	slice := make([]interface{}, 1)
	slice[0] = name
	values := append(slice, args...)
	if b, err := redis.Bool(r.Command("LPUSH", values...)); err == nil {
		return b, err
	}
	return false, nil
}

func (r *Reads) LPop(name string) (interface{}, error) {
	return r.Command("LPOP", name)
}

func (r *Reads) LRange(name string, start int, stop int) ([]interface{}, error) {
	return redis.Values(r.Command("LRANGE", name, start, stop))
}

func (r *Reads) SAdd(name string, args ...interface{}) (bool, error) {
	slice := make([]interface{}, 1)
	slice[0] = name
	values := append(slice, args...)
	return redis.Bool(r.Command("SADD", values...))
}

func (r *Reads) SPop(name string) (interface{}, error) {
	return redis.Values(r.Command("SPOP", name))
}

func (r *Reads) ZAdd(name string, key string, value string) (bool, error) {
	return redis.Bool(r.Command("ZADD", name, key, value))
}

func (r *Reads) ZRange(name string, start int, stop int, args ...bool) ([]interface{}, error) {
	if args != nil && args[0] == true {
		return redis.Values(r.Command("ZRANGE", name, start, stop, "WITHSCORES"))
	} else {
		return redis.Values(r.Command("ZRANGE", name, start, stop))
	}
}

func (r *Reads) ZScore(name string, key string) (interface{}, error) {
	return redis.Values(r.Command("ZSCORE", name, key))
}

func (r *Reads) Exists(name string) (bool, error) {
	return redis.Bool(r.Command("EXISTS", name))
}

func (r *Reads) Del(name string) error {
	_, err := r.Command("DEL", name)
	return err
}

func (r *Reads) Clear() error {
	_, err := r.Command("FLUSHDB")
	return err
}

func (r *Reads) Multi() error {
	conn := r.GetConnection()
	_, err := conn.Do("MULTI")
	return err
}

func (r *Reads) Exec() error {
	conn := r.GetConnection()
	_, err := conn.Do("EXEC")
	return err
}

func (r *Reads) Discard() error {
	conn := r.GetConnection()
	_, err := conn.Do("DISCARD")
	return err
}

func (r *Reads) Close() {
	conn := r.GetConnection()
	err := conn.Close()
	r.Conn = nil
	if err != nil {
		log.Warning(err.Error())
		// Log
		return
	}
}
