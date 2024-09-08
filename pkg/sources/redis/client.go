/*
Copyright 2024 Hurricane1988 Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

// TODO: https://redis.uptrace.dev/

// Interface Redis 集群操作接口
type Interface interface {
	Client() *redis.ClusterClient
	Set(ctx context.Context, key, value string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Existed(ctx context.Context, keys ...string) bool
	Keys(ctx context.Context, pattern string) ([]string, error)
	Del(ctx context.Context, keys ...string) error
	Expire(ctx context.Context, key string, duration time.Duration) error
}

// NewRedis 根据配置实例化一个Redis 接口
func NewRedis(config Options) Interface {
	return &config
}

// Client 创建并返回一个新的Redis集群客户端。
func (r *Options) Client() *redis.ClusterClient {
	// 根据配置信息初始化Redis集群客户端选项。
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       r.Address,        // Redis服务器地址列表
		MaxRetries:  3,                // 最大重试次数
		Password:    r.Password,       // Redis服务器密码
		Username:    r.Username,       // Redis服务器用户名
		DialTimeout: 10 * time.Second, // 连接超时时间
		PoolFIFO:    true,             // 使用FIFO顺序处理空闲连接池中的连接
		PoolTimeout: 10 * time.Second, // 连接池中的最大等待时间
		PoolSize:    10,               // 连接池大小
	})
	return rdb // 返回创建的Redis集群客户端实例。
}

// Set 写入数据
func (r *Options) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	return r.Client().Set(ctx, key, value, expiration).Err()
}

// Get 查询数据
func (r *Options) Get(ctx context.Context, key string) (string, error) {
	value, err := r.Client().Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	} else if err != nil {
		return "", err
	} else {
		return value, nil
	}
}

// Existed 判断Keys是否存在
func (r *Options) Existed(ctx context.Context, keys ...string) bool {
	existedKeys, err := r.Client().Exists(ctx, keys...).Result()
	if err != nil {
		return false
	}
	return len(keys) == int(existedKeys)
}

// Keys 查询批量keys redis接口实现
func (r *Options) Keys(ctx context.Context, pattern string) ([]string, error) {
	return r.Client().Keys(ctx, pattern).Result()
}

// Del 删除key
func (r *Options) Del(ctx context.Context, keys ...string) error {
	return r.Client().Del(ctx, keys...).Err()
}

// Expire 设置key过期
func (r *Options) Expire(ctx context.Context, key string, duration time.Duration) error {
	return r.Client().Expire(ctx, key, duration).Err()
}
