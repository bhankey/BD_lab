package tokenrepo

import (
	"context"
	"finance/internal/clienterror"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type TokenRepo struct {
	redis *redis.Client
}

func NewUserRepo(db *redis.Client) *TokenRepo {
	return &TokenRepo{
		redis: db,
	}
}

// SetRefreshToken stores a refresh token with an expiry time
func (r *TokenRepo) SetRefreshToken(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error {
	if err := r.redis.Set(ctx, tokenID, userID, expiresIn).Err(); err != nil {
		return clienterror.SystemError(fmt.Sprintf("can not create refresh token in redis db due to clienterror %v:", err))
	}

	return nil
}

// DeleteRefreshToken used to delete old refresh tokens
func (r *TokenRepo) DeleteRefreshToken(ctx context.Context, tokenID string) error {
	res := r.redis.Del(ctx, tokenID)

	if err := res.Err(); err != nil {
		return clienterror.SystemError(fmt.Sprintf("can not delete refresh token in redis db due to clienterror %v", err))
	}

	return nil
}

func (r *TokenRepo) GetUserID(ctx context.Context, tokenID string) (string, error) {
	id, err := r.redis.Get(ctx, tokenID).Result()
	if err != nil {
		if err == redis.Nil {
			return "", clienterror.UnauthorizedError("Invalid refresh token")
		}
		return "", clienterror.SystemError(fmt.Sprintf("can not get user id from redis due to clienterror: %v", err))
	}

	return id, nil
}
