package config

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func InitRedis() {
	// init redis connection pool
	initPool()

	// bootstramp some data to redis
	initStore()
}

func initPool() {
	pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				log.Fatal().Err(err).Msg("fail init redis")
				os.Exit(1)
			}
			return conn, err
		},
	}
}

func initStore() {
	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	macs := []string{"00000C  Cisco", "00000D  FIBRONICS", "00000E  Fujitsu",
		"00000F  Next", "000010  Hughes"}
	for _, mac := range macs {
		pair := strings.Split(mac, "  ")
		set(pair[0], pair[1])
	}
}

func ping(conn redis.Conn) {
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		log.Fatal().Err(err).Msg("fail ping redis")
		os.Exit(1)
	}
}

func set(key string, val string) error {
	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, val)
	if err != nil {
		log.Fatal().Err(err).Msg("fail set key redis")
		return err
	}

	return nil
}

func get(key string) (string, error) {
	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", key))
	if err != nil {
		log.Fatal().Err(err).Msg("fail get key redis")
		return "", err
	}

	return s, nil
}

func sadd(key string, val string) error {
	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SADD", key, val)
	if err != nil {
		log.Fatal().Err(err).Msg("fail add val to set redis")
		return err
	}

	return nil
}

func smembers(key string) ([]string, error) {
	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.Strings(conn.Do("SMEMBERS", key))
	if err != nil {
		log.Fatal().Err(err).Msg("fail get set error redis")
		return nil, err
	}

	return s, nil
}
