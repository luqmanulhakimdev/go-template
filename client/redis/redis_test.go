package redis

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RedisClientTestSuite struct {
	suite.Suite
	c *Client
}

func (s *RedisClientTestSuite) SetupTest() {
	mr, err := miniredis.Run()
	require.NoError(s.T(), err)

	redisClient := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	s.c = NewClient(redisClient)
}

func (s *RedisClientTestSuite) TestSet() {
	ctx := context.Background()
	key := "expired_pin"
	field := "1"
	timeout := 60 * time.Second
	err := s.c.Set(ctx, key, field, timeout)
	s.Assertions.NoError(err)
}

func (s *RedisClientTestSuite) TestGet() {
	ctx := context.Background()
	key := "expired_pin"
	err := s.c.Get(ctx, key, "")
	s.Assertions.NoError(err)
}

func (s *RedisClientTestSuite) TestIncr() {
	ctx := context.Background()
	key := "expired_pin"
	rowsaffected, err := s.c.Incr(ctx, key)
	s.Assertions.NoError(err)
	s.Assertions.Equal(int64(1), rowsaffected)
}

func (s *RedisClientTestSuite) TestDel() {
	ctx := context.Background()
	key := "expired_pin"
	err := s.c.Del(ctx, key)
	s.Assertions.NoError(err)
}

func (s *RedisClientTestSuite) TestDelBulk() {
	ctx := context.Background()

	key := "testing"
	field := "1"
	timeout := 60 * time.Second
	err := s.c.Set(ctx, key, field, timeout)

	prefix := "test"
	err = s.c.DelBulk(ctx, prefix)
	s.Assertions.NoError(err)
}

func (s *RedisClientTestSuite) TestSetIfNotExist() {
	ctx := context.Background()
	key := "expired_pin"
	field := "1"
	timeout := 60 * time.Second
	res, err := s.c.SetIfNotExist(ctx, key, field, timeout)
	s.Assertions.NoError(err)
	s.Assertions.Equal(true, res)
}

func (s *RedisClientTestSuite) TestHashMap() {
	ctx := context.Background()
	hashName := "admin:login"
	field := "1"
	timeNow := time.Now()

	// get non-set field
	resGet, err := s.c.HashGet(ctx, hashName, field)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "", resGet)

	// set field
	res, err := s.c.HashSet(ctx, hashName, field, timeNow)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), int64(1), res)

	// get the field's value
	resGet, err = s.c.HashGet(ctx, hashName, field)
	assert.NoError(s.T(), err)

	expectedTime := timeNow.Format(time.RFC3339Nano)
	assert.Equal(s.T(), expectedTime, resGet)

	// set bool field
	res, err = s.c.HashSet(ctx, hashName, field, true)
	s.Assertions.Equal(int64(0), res)
	s.Assertions.NoError(err)

	// get bool value
	resGet, err = s.c.HashGet(ctx, hashName, field)
	s.Assertions.NoError(err)

	expectedBool := true
	b, err := strconv.ParseBool(resGet)
	s.Assertions.NoError(err)
	s.Assertions.Equal(expectedBool, b)

	n, err := s.c.HashDel(ctx, hashName, field)
	s.Assertions.NoError(err)
	s.Assertions.Equal(int64(1), n)
	// try de;eting the second time
	n, err = s.c.HashDel(ctx, hashName, field)
	s.Assertions.NoError(err)
	s.Assertions.Equal(int64(0), n)
}

func (s *RedisClientTestSuite) TestLock() {
	ctx := context.Background()
	key := "admin:login"
	identifier := "1"
	timeout := 60 * time.Second
	res := s.c.Lock(ctx, key, identifier, timeout)
	s.Assertions.Equal(true, res)

	s.c.Unlock(ctx, key, identifier)
}

func TestRedisClientTestSuite(t *testing.T) {
	suite.Run(t, new(RedisClientTestSuite))
}

func (s *RedisClientTestSuite) TestHDel() {
	ctx := context.Background()
	key := "expired_pin"
	field := "1"
	rowsaffected, err := s.c.HashDel(ctx, key, field)
	s.Assertions.NoError(err)
	s.Assertions.Equal(int64(0), rowsaffected)
}
