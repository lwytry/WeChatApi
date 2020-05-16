package redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
	"wechat/utils"
)

var (
	host   string
	//database     string = "redisServer"
	password     string
	maxOpenConns int    = 100
	maxIdleConns int    = 100
)

var NewCache *RedisConnPool

type RedisConnPool struct {
	redisPool *redis.Pool
}

func init() {
	sec, err := utils.Cfg.GetSection("redis")
	if err != nil {
		log.Fatal(2, "Fail to get section 'redis': %v", err)
	}
	host = sec.Key("HOST").String()
	password = sec.Key(  "PASSWORD").String()
	NewCache = &RedisConnPool{}
	NewCache.redisPool = myNewPool()
	if NewCache.redisPool == nil {
		panic("init redis failed！")
	}
}
func myNewPool() *redis.Pool {
	return &redis.Pool{
		MaxActive:   maxOpenConns, // max number of connections
		MaxIdle:     maxIdleConns,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (p *RedisConnPool) Close() error {
	err := p.redisPool.Close()
	return err
}

// 当前某一个数据库，执行命令
func (p *RedisConnPool) Do(command string, args ...interface{}) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return conn.Do(command, args...)
}

//// String（字符串）
func (p *RedisConnPool) SetString(key string, value interface{}) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return conn.Do("SET", key, value)
}

func (p *RedisConnPool) SetEx(key string, value interface{}, t int) (string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.String(conn.Do("SETEX", key, t, value))
}

func (p *RedisConnPool) HSet(table, key string, value interface{}) (bool, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("HSET", table, key, value))
}

func (p *RedisConnPool) GetBytes(key string) ([]byte, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	return redis.Bytes(conn.Do("GET", key))
}

func (p *RedisConnPool) GetInt(key string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("GET", key))
}
func (p *RedisConnPool) GetInt64(key string) (int64, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("GET", key), )
}
func (p *RedisConnPool) DelKey(key string) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return conn.Do("DEL", key)
}
func (p *RedisConnPool) ExpireKey(key string, seconds int64) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return conn.Do("EXPIRE", key, seconds)
}
func (p *RedisConnPool) Keys(pattern string) ([]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("KEYS", pattern))
}
func (p *RedisConnPool) KeysByteSlices(pattern string) ([][]byte, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.ByteSlices(conn.Do("KEYS", pattern))
}
func (p *RedisConnPool) SetHashMap(key string, fieldValue map[string]interface{}) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(fieldValue)...)
}
func (p *RedisConnPool) GetHashMapString(key string) (map[string]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.StringMap(conn.Do("HGETALL", key))
}
func (p *RedisConnPool) GetHashMapInt(key string) (map[string]int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.IntMap(conn.Do("HGETALL", key))
}

func (p *RedisConnPool) GetHashMapInt64(key string) (map[string]int64, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Int64Map(conn.Do("HGETALL", key))
}

func (p *RedisConnPool) GetHashMapKeyString(table, key string) ([]byte, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Bytes(conn.Do("HGET", table, key))
}

func (p *RedisConnPool) Exists(key string) (bool, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("EXISTS", key))
}

func (p *RedisConnPool) Incrby(key string, incV int) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("INCRBY", key, incV))
}

func (p *RedisConnPool) Incr(key string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("INCR", key))
}

func (p *RedisConnPool) Lpush(key string, value interface{}) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("LPUSH", key, value))
}

func (p *RedisConnPool) Lrange(key string, start int, end int) ([]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("LRANGE", key, start, end))
}