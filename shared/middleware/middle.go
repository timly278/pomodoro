package mdw

import (
	db "pomodoro/db/sqlc"
	"pomodoro/security"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Middleware struct {
	tokenMaker security.TokenMaker
	store      db.Store
	redisdb    *redis.Client
	logger     *zap.Logger
}

func New(tokenMaker security.TokenMaker, store db.Store, rdb *redis.Client, logger *zap.Logger) *Middleware {
	return &Middleware{
		tokenMaker: tokenMaker,
		store:      store,
		redisdb:    rdb,
		logger:     logger,
	}
}



