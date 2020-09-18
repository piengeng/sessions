package clusteredis

import (
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/tester"
	"github.com/go-redis/redis"
)

var (
	redisTestServer  = "ubuntu.home:6379"
	redisTestCluster = []string{
		"ubuntu.home:7000", "ubuntu.home:7001", "ubuntu.home:7002",
		"ubuntu.home:7003", "ubuntu.home:7004", "ubuntu.home:7005",
	}
	clustered = false
	keyPrefix = "a:" // keeping it short
	client    redis.UniversalClient
)

var newRedisStore = func(_ *testing.T) sessions.Store {
	if clustered {
		client = redis.NewClusterClient(&redis.ClusterOptions{Addrs: redisTestCluster})
	} else {
		client = redis.NewClient(&redis.Options{Addr: redisTestServer})
	}
	rs := NewStore(client, keyPrefix)
	return rs
}

func TestRedis_SessionGetSet(t *testing.T) {
	tester.GetSet(t, newRedisStore)
}

func TestRedis_SessionDeleteKey(t *testing.T) {
	tester.DeleteKey(t, newRedisStore)
}

func TestRedis_SessionFlashes(t *testing.T) {
	tester.Flashes(t, newRedisStore)
}

func TestRedis_SessionClear(t *testing.T) {
	tester.Clear(t, newRedisStore)
}

func TestRedis_SessionOptions(t *testing.T) {
	tester.Options(t, newRedisStore)
}
