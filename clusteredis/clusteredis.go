package clusteredis

import (
	"github.com/piengeng/clusteredistore"

	"github.com/gin-contrib/sessions"
	"github.com/go-redis/redis"
)

type Store interface {
	sessions.Store
}

func NewStore(client redis.UniversalClient, keyPrefix string) Store {
	rs := clusteredistore.NewRedisStore(client)
	rs.KeyPrefix(keyPrefix)
	return &store{rs}
}

type store struct {
	*clusteredistore.RedisStore
}

func (c *store) Options(options sessions.Options) {
	c.RedisStore.Options = options.ToGorillaOptions()
}
