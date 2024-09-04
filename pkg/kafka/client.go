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

package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"net"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"strconv"
	"time"
)

type Interface interface {
	LeaderHost(ctx context.Context) (string, error)
	WriteByConn(ctx context.Context, topic string, partition int, writeTimeout time.Duration, key, value []byte) error
	ReadByConn(ctx context.Context, topic string, partition int, readTimeout time.Duration) (*kafka.Batch, error)
	CreateTopicsByConn(ctx context.Context, topics []string, partition, replicationFactor int) error
	DeleteTopicsByConn(ctx context.Context, topics []string) error
	ListTopicByConn(ctx context.Context) ([]string, error)
	DeleteAllTopicsByConn(ctx context.Context) error
}

// LeaderHost 通过上下文获取当前Kafka集群的控制节点主机信息
// 该函数主要用于确定哪个Kafka节点目前是控制节点
// 参数:
//
//	ctx: 上下文，可用于传递超时、截止时间或取消请求等信息
//
// 返回值:
//
//	  string: 当前控制节点的主机信息（IP地址和端口号）
//		 error: 在获取控制节点主机信息过程中可能遇到的错误，如果没有错误，则为nil
func (k *Options) LeaderHost(ctx context.Context) (string, error) {
	logger := log.FromContext(ctx)
	// 初始化dialer
	_, dialer := Dialer(k.Algo, k.Username, k.Password, k.Timeout, true)
	// 链接至任意的kafka节点
	c, err := dialer.DialContext(ctx, k.Protocol, k.Brokers[0])
	if err != nil {
		logger.Error(err, "connect to kafka node failed.")
		return "", err
	}
	// 获取当前控制节点信息
	leaderHost, err := c.Controller()
	if err != nil {
		logger.Error(err, "get the leader host failed.")
		return "", err
	}
	// 拼接主机地址和端口
	return net.JoinHostPort(leaderHost.Host, strconv.Itoa(leaderHost.Port)), nil
}

// WriteByConn 通过Kafka连接发送消息
// 该函数负责将一条消息发送到指定的Kafka主题和分区。它首先尝试与Kafka集群的Leader节点建立连接，
// 然后设置写超时时间，最后发送消息并关闭连接
// 参数:
//
//	ctx: 上下文，用于传递超时、取消信号和日志等信息
//	topic: 消息要发送到的Kafka主题
//	partition: 消息要发送到的主题内的分区
//	writeTimeout: 设置写操作（发送消息）的超时时间
//	key: 消息的键，可以为空
//	value: 消息的值
//
// 返回值:
//
//	error: 如果连接、设置超时、发送消息或关闭连接过程中发生错误，则返回该错误
func (k *Options) WriteByConn(ctx context.Context, topic string, partition int, writeTimeout time.Duration, key, value []byte) error {
	// 从上下文中获取日志记录器
	logger := log.FromContext(ctx)
	// 尝试获取Kafka集群的Leader主机
	leaderHost, _ := k.LeaderHost(ctx)

	// 连接至kafka集群的Leader节点
	conn, err := kafka.DialLeader(ctx, k.Protocol, leaderHost, topic, partition)
	if err != nil {
		logger.Error(err, "connect to kafka leader host failed.")
		return err
	}
	// 设置发送消息的超时时间
	err = conn.SetWriteDeadline(time.Now().Add(writeTimeout))
	if err != nil {
		return err
	}

	// 发送消息
	_, err = conn.WriteMessages(kafka.Message{Key: key, Value: value})
	if err != nil {
		logger.Error(err, "send kafka message failed.", "key-name", key)
		return err
	}
	// 关闭连接
	if err = conn.Close(); err != nil {
		logger.Error(err, "failed to close kafka connection.")
		return err
	}
	return nil
}

// ReadByConn 通过指定的kafka连接读取指定主题和分区的数据
// 参数:
//
//	ctx: 上下文，用于传递请求范围的值、配置截止时间以及取消操作
//	topic: 要读取数据的主题名称
//	partition: 要读取数据的分区编号
//	readTimeout: 读取操作的超时时间
//
// 返回值:
//
//	*kafka.Batch: 读取到的消息批次
//	error: 如果发生错误，则返回错误信息
func (k *Options) ReadByConn(ctx context.Context, topic string, partition int, readTimeout time.Duration) (*kafka.Batch, error) {
	// 从上下文中获取日志记录器
	logger := log.FromContext(ctx)
	// 获取主题leader的主机地址
	leaderHost, _ := k.LeaderHost(ctx)

	// 连接至kafka集群的Leader节点
	conn, err := kafka.DialLeader(ctx, k.Protocol, leaderHost, topic, partition)
	if err != nil {
		logger.Error(err, "failed to connect to kafka leader.")
		return nil, err
	}

	// 设置read超时
	err = conn.SetReadDeadline(time.Now().Add(readTimeout))
	if err != nil {
		return nil, err
	}

	// 读取一批消息，得到的batch是一系列消息的迭代器
	batch := conn.ReadBatch(512, 20e6) // fetch 1KB min, 20MB max

	// 遍历读取消息
	b := make([]byte, 1024) // 10KB max per message
	for {
		n, readErr := batch.Read(b)
		if readErr != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	// 关闭batch
	if err = batch.Close(); err != nil {
		logger.Error(err, "failed to close kafka batch.")
		return nil, err
	}

	return batch, nil
}

// CreateTopicsByConn 创建Kafka主题
// 该函数通过连接到Kafka集群的领导者节点来创建指定的主題
// 参数:
//
//	ctx: 上下文，用于取消操作
//	topics: 待创建的主题列表
//	partition: 每个主题的分区数量
//	replicationFactor: 主题的副本因子，表示副本的数量
//
// 返回值:
//
//	error: 表示创建主题时的错误，如果没有错误则是nil
func (k *Options) CreateTopicsByConn(ctx context.Context, topics []string, partition, replicationFactor int) error {
	// 初始化主题配置列表
	var topicsConfigList []kafka.TopicConfig

	// 获取Kafka集群的领导者节点地址
	leaderHost, _ := k.LeaderHost(ctx)

	// 建立与Kafka集群的连接
	conn, _ := kafka.Dial(k.Protocol, leaderHost)

	// 遍历主题列表，为每个主题创建TopicConfig并添加到列表中
	for _, topic := range topics {
		t := kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     partition,
			ReplicationFactor: replicationFactor,
		}
		topicsConfigList = append(topicsConfigList, t)
	}

	// 创建主题。
	err := conn.CreateTopics(topicsConfigList...)
	if err != nil {
		return err
	}
	return nil
}

// ListTopicByConn 根据连接列出Kafka中的所有主题
// 该函数通过连接到Kafka集群的领导者实例来获取所有主题列表
// 参数:
//
//	ctx: 上下文，用于传递请求范围的值、配置 deadline 等
//
// 返回值:
//
//	[]string: 包含所有主题名称的切片
//	error: 如果连接到Kafka实例或读取分区信息时发生错误，则返回错误
func (k *Options) ListTopicByConn(ctx context.Context) ([]string, error) {
	// 从上下文中获取logger实例
	logger := log.FromContext(ctx)
	// 获取Kafka集群的领导者主机地址
	leaderHost, _ := k.LeaderHost(ctx)

	// 初始化一个字符串切片来存储主题名称
	var topics []string

	// 尝试连接到Kafka领导者实例
	conn, err := kafka.Dial(k.Protocol, leaderHost)
	if err != nil {
		// 如果连接失败，记录错误并返回
		logger.Error(err, "failed to connect to kafka leader.")
		return nil, err
	}

	// 尝试读取Kafka实例的所有分区信息
	partitions, err := conn.ReadPartitions()
	if err != nil {
		// 如果读取分区信息失败，记录错误并返回
		logger.Error(err, "read partition failed.")
		return nil, err
	}

	// 使用map来存储和去重主题名称
	m := map[string]struct{}{}
	// 遍历所有分区，提取并记录不同的主题名称
	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}

	// 将map中的主题名称添加到topics切片中
	for topic := range m {
		topics = append(topics, topic)
	}

	// 返回所有主题名称切片和nil错误
	return topics, nil
}

// DeleteTopicsByConn 通过连接到Kafka集群的领导者来删除指定的主题
// 它接受一个上下文对象，用于日志记录和获取Kafka领导者主机地址
// 参数:
//
//	topics: 是一个字符串切片，包含了需要删除的主题名称列表
//
// 返回值:
//
//	error: 一个错误类型，如果执行过程中发生错误，则返回相应的错误信息
func (k *Options) DeleteTopicsByConn(ctx context.Context, topics []string) error {
	// 从上下文中提取日志记录器
	logger := log.FromContext(ctx)

	// 通过上下文获取Kafka集群的领导者主机地址
	leaderHost, _ := k.LeaderHost(ctx)

	// 使用获取的协议和领导者主机地址建立Kafka连接
	conn, err := kafka.Dial(k.Protocol, leaderHost)
	if err != nil {
		// 如果连接失败，记录错误并返回
		logger.Error(err, "failed to connect to kafka leader.")
		return err
	}

	// 调用连接对象的DeleteTopics方法来删除指定的主题
	err = conn.DeleteTopics(topics...)
	if err != nil {
		// 如果删除主题失败，记录错误并返回
		logger.Error(err, "failed to delete topics.")
		return err
	}

	// 如果一切顺利，返回nil
	return nil
}

// DeleteAllTopicsByConn 删除与特定连接相关联的所有topics
// 该方法首先通过ListTopicByConn函数获取所有与连接相关的topics列表
// 然后使用DeleteTopicsByConn函数来删除这些topics
// 参数:
//
//	ctx: 用于取消操作的上下文，可以在整个函数执行的任何时间点检查上下文并返回
//
// 返回值:
//
//	error: 如果列出或删除topics时发生错误，则返回错误
func (k *Options) DeleteAllTopicsByConn(ctx context.Context) error {
	// 获取所有的topic
	topicsList, err := k.ListTopicByConn(ctx)
	if err != nil {
		return err
	}
	// 删除所有的topics
	return k.DeleteTopicsByConn(ctx, topicsList)
}
