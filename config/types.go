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

package config

import (
	"github.com/hurricane1988/harbor-webhook/pkg/sources/db/mysql"
	"github.com/hurricane1988/harbor-webhook/pkg/sources/kafka"
	"github.com/hurricane1988/harbor-webhook/pkg/sources/redis"
)

// Config 全局配置
type Config struct {
	RedisOptions *redis.Options `json:"redisOptions,omitempty" yaml:"redisOptions,omitempty" xml:"redisOptions,omitempty" mapstructure:"redis"`
	MysqlOptions *mysql.Options `json:"mysqlOptions,omitempty" yaml:"mysqlOptions,omitempty" xml:"mysqlOptions,omitempty" mapstructure:"mysql"`
	KafkaOptions *kafka.Options `json:"kafkaOptions,omitempty" yaml:"kafkaOptions,omitempty" xml:"kafkaOptions,omitempty" mapstructure:"kafka"`
}
