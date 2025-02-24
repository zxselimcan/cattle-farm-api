package utils

import (
	"fmt"
	"os"

	"github.com/adjust/rmq/v5"
	"github.com/go-redis/redis/v8"
)

var (
	Redis *redis.Client
	Rmq   rmq.Connection
)

func createRedisAddrPort() string {

	return os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
}

func ConnectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{

		Addr: createRedisAddrPort(),
		DB:   0,
	})
}

func ConnectRmq() (rmq.Connection, error) {
	connection, err := rmq.OpenConnection("service1", "tcp", createRedisAddrPort(), 1, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return connection, nil

}
