package rdb

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type MetricsHook struct {
}

type CustomOptions struct {
	Password     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	MinIdleConns int
	DialTimeout  time.Duration
	Addrs        []string
}

// NewMetricsClusterClient 创建一个带监控指标的的 redis client
func NewClusterClient(opts *CustomOptions) *redis.ClusterClient {
	var cos redis.ClusterOptions
	cos.Addrs = opts.Addrs
	cos.Password = opts.Password
	cos.ReadTimeout = opts.ReadTimeout * time.Second
	cos.WriteTimeout = opts.WriteTimeout * time.Second
	cos.PoolSize = opts.PoolSize
	cos.MinIdleConns = opts.MinIdleConns
	cos.Addrs = opts.Addrs
	cos.DialTimeout = opts.DialTimeout * time.Second
	rc := redis.NewClusterClient(&cos)
	rc.AddHook(TracingHook{})
	if err := rc.Ping(context.TODO()).Err(); err != nil {
		panic(err)
	}
	return rc
}
