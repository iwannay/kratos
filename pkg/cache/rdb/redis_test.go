package rdb

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMetricsClusterClient(t *testing.T) {
	rdb := NewClusterClient(&CustomOptions{
		Addrs:       []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003", "127.0.0.1:7004", "127.0.0.1:7005"},
		ReadTimeout: 30 * time.Second,
		DialTimeout: 5 * time.Second,
		PoolSize:    100,
		Password:    "123456",
	})
	assert.Equal(t, nil, rdb.Ping(context.TODO()).Err(), "failed ping")
	for {
		rdb.Set(context.TODO(), "hello", "boy", -1)
		time.Sleep(1 * time.Second)
		fmt.Println(rdb.Get(context.TODO(), "hello").String())

	}
}
