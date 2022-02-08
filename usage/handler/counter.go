package handler

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type counter struct {
	sync.RWMutex
	redisClient *redis.Client
}

type listEntry struct {
	Service string
	Count   int64
}

func (c *counter) incr(ctx context.Context, userID, path string, delta int64, t time.Time) (int64, error) {
	t = t.UTC()
	key := fmt.Sprintf("%s:%s:%s:%s", prefixCounter, userID, t.Format("20060102"), path)
	pipe := c.redisClient.TxPipeline()
	incr := pipe.IncrBy(ctx, key, delta)
	pipe.Expire(ctx, key, counterTTL) // make sure we expire the counters
	_, err := pipe.Exec(ctx)
	if err != nil {
		return 0, err
	}
	return incr.Result()
}

func (c *counter) incrMonthly(ctx context.Context, userID, path string, delta int64, t time.Time) (int64, error) {
	t = t.UTC()
	key := fmt.Sprintf("%s:%s:%s:%s", prefixCounter, userID, t.Format("200601"), path)
	pipe := c.redisClient.TxPipeline()
	incr := pipe.IncrBy(ctx, key, delta)
	pipe.Expire(ctx, key, counterMonthlyTTL) // make sure we expire the counters
	_, err := pipe.Exec(ctx)
	if err != nil {
		return 0, err
	}
	return incr.Result()
}

func (c *counter) decr(ctx context.Context, userID, path string, delta int64, t time.Time) (int64, error) {
	t = t.UTC()
	key := fmt.Sprintf("%s:%s:%s:%s", prefixCounter, userID, t.Format("20060102"), path)
	pipe := c.redisClient.TxPipeline()
	decr := pipe.DecrBy(ctx, key, delta)
	pipe.Expire(ctx, key, counterTTL) // make sure we expire counters
	_, err := pipe.Exec(ctx)
	if err != nil {
		return 0, err
	}
	return decr.Result()
}

func (c *counter) read(ctx context.Context, userID, path string, t time.Time) (int64, error) {
	t = t.UTC()
	ret, err := c.redisClient.Get(ctx, fmt.Sprintf("%s:%s:%s:%s", prefixCounter, userID, t.Format("20060102"), path)).Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return ret, err
}

func (c *counter) readMonthly(ctx context.Context, userID, path string, t time.Time) (int64, error) {
	t = t.UTC()
	ret, err := c.redisClient.Get(ctx, fmt.Sprintf("%s:%s:%s:%s", prefixCounter, userID, t.Format("200601"), path)).Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return ret, err
}

func (c *counter) deleteUser(ctx context.Context, userID string) error {
	keys, err := c.redisClient.Keys(ctx, fmt.Sprintf("%s:%s:*", prefixCounter, userID)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}
	keys = append(keys, fmt.Sprintf("%s:%s", prefixResetDay, userID))
	if err := c.redisClient.Del(ctx, keys...).Err(); err != nil && err != redis.Nil {
		return err
	}

	return nil
}

func (c *counter) listForUser(userID string, t time.Time) ([]listEntry, error) {
	ctx := context.Background()
	keyPrefix := fmt.Sprintf("%s:%s:%s:", prefixCounter, userID, t.Format("20060102"))
	sc := c.redisClient.Scan(ctx, 0, keyPrefix+"*", 0)
	if err := sc.Err(); err != nil {
		return nil, err
	}
	iter := sc.Iterator()
	res := []listEntry{}
	for {
		if !iter.Next(ctx) {
			break
		}
		key := iter.Val()
		i, err := c.redisClient.Get(ctx, key).Int64()
		if err != nil {
			return nil, err
		}
		res = append(res, listEntry{
			Service: strings.TrimPrefix(key, keyPrefix),
			Count:   i,
		})
	}
	return res, iter.Err()
}

func (c *counter) listMonthliesForUser(userID string, t time.Time) ([]listEntry, error) {
	ctx := context.Background()
	keyPrefix := fmt.Sprintf("%s:%s:%s:", prefixCounter, userID, t.Format("200601"))
	sc := c.redisClient.Scan(ctx, 0, keyPrefix+"*", 0)
	if err := sc.Err(); err != nil {
		return nil, err
	}
	iter := sc.Iterator()
	res := []listEntry{}
	for {
		if !iter.Next(ctx) {
			break
		}
		key := iter.Val()
		i, err := c.redisClient.Get(ctx, key).Int64()
		if err != nil {
			return nil, err
		}
		res = append(res, listEntry{
			Service: strings.TrimPrefix(key, keyPrefix),
			Count:   i,
		})
	}
	return res, iter.Err()
}

// usageTime gives the current month for usage calculations
func (c *counter) usageTime(id string, t time.Time) (time.Time, error) {
	// lookup this ID's reset day of the month
	ctx := context.Background()
	day, err := c.redisClient.Get(ctx, fmt.Sprintf("%s:%s", prefixResetDay, id)).Int64()
	if err != nil && err != redis.Nil {
		return t, err
	}
	// if none return t
	if day == 0 {
		return t, nil
	}
	// return t - 1 month if we have not passed reset day
	if t.Day() < int(day) {
		return t.AddDate(0, -1, 0), nil
	}
	// else return t
	return t, nil
}

func (c *counter) setResetDay(ctx context.Context, id string, day int64) error {
	return c.redisClient.Set(ctx, fmt.Sprintf("%s:%s", prefixResetDay, id), day, 0).Err()
}
