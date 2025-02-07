package cache

import (
	"bytes"
	"encoding/json"
	"gin_work/wrap/config"
	driver2 "gin_work/wrap/driver"
	"gin_work/wrap/log"
)

type cache struct {
	CType    string
	Host     string
	Port     int
	Password string
	Prefix   string
	Timeout  int
}

var (
	Cache  *cache
	driver driver2.Driver
)

func init() {
	Cache = getConfig()
	if Cache.CType == "redis" {
		driver = driver2.NewReads(Cache.Host, Cache.Port, Cache.Password, Cache.Prefix, 0)
	}
}

func getConfig() *cache {
	c := config.Mapping.Cache
	return &cache{
		CType:    c.CType,
		Host:     c.Host,
		Port:     c.Port,
		Password: c.Password,
		Prefix:   c.Prefix,
		Timeout:  c.Timeout,
	}
}

func Has(name string) bool {
	exists, err := driver.Exists(Cache.Prefix + name)
	if err != nil {
		log.Warning(err.Error())
		return false
	}
	return exists
}

func BindGet(name string, value interface{}) error {
	str, err := driver.Get(Cache.Prefix + name)
	if err != nil || str == "" {
		log.Warning(err.Error())
		return err
	}
	b := []byte(str)
	err = json.Unmarshal(b, &value)
	return err
}

func Get(name string) (interface{}, error) {
	str, err := driver.Get(Cache.Prefix + name)
	if err != nil || str == "" {
		log.Warning(err.Error())
		return nil, err
	}
	b := []byte(str)
	var instance interface{}
	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err = d.Decode(&instance)
	if err != nil {
		log.Warning(err.Error())
		return nil, err
	}
	/*err = json.Unmarshal(b, &instance)
	if err != nil {
		return nil, err
	}*/
	return instance, nil
}

func Set(name string, value interface{}, args ...int) error {
	timeout := Cache.Timeout
	if args != nil {
		timeout = args[0]
	}
	b, err := json.Marshal(&value)
	if timeout == 0 {
		_, err = driver.Set(Cache.Prefix+name, string(b))
	} else {
		_, err = driver.Set(Cache.Prefix+name, string(b), timeout)
	}
	if err != nil {
		log.Warning(err.Error())
		return err
	}
	return nil
}

func Del(name string) error {
	return driver.Del(Cache.Prefix + name)
}

func Clear() error {
	return driver.Clear()
}
