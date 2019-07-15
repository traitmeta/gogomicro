package redis

import (
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/util/log"
	"github.com/songxuexian/gogomicro/basic/config"
	"sync"
)

var (
	client *redis.Client
	m      sync.RWMutex
	inited bool
)

func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Log("Have been init redis")
		return
	}
	redisConf := config.GetReidsConfig()

	if redisConf != nil && redisConf.GetEnabled() {
		log.Log("Init redis...")
		if redisConf.GetSentinelConfig() != nil && redisConf.GetSentinelConfig().GetEnabled() {
			log.Log("Init redis, sentinel model...")
			initSentinel(redisConf)
		} else {
			log.Log("Init redis, normal model...")
			initSingle(redisConf)
		}

		log.Log("Init redis, check connection...")
		pong, err := client.Ping().Result()
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Log("Init redis, check connection Ping.")
		log.Log("Init redis, check connection Ping..")
		log.Logf("Init redis, check connection Ping... %s", pong)

	}
}

func GetRedis() *redis.Client {
	return client
}

func initSentinel(redisConfig config.RedisConfig) {
	client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    redisConfig.GetSentinelConfig().GetMaster(),
		SentinelAddrs: redisConfig.GetSentinelConfig().GetNodes(),
		DB:            redisConfig.GetDBNum(),
		Password:      redisConfig.GetPassword(),
	})

}

func initSingle(redisConfig config.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.GetConn(),
		Password: redisConfig.GetPassword(), // no password set
		DB:       redisConfig.GetDBNum(),    // use default DB
	})
}
