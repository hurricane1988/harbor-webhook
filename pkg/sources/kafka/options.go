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
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
	"log"
	"time"
)

// Options Kafka 配置Options
type Options struct {
	Username               string          `json:"username,omitempty" yaml:"username,omitempty" xml:"username,omitempty"`
	Password               string          `json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	Brokers                []string        `json:"brokers,omitempty" yaml:"brokers,omitempty" xml:"brokers,omitempty"`
	Topics                 []string        `json:"topics,omitempty" yaml:"topics,omitempty" xml:"topics,omitempty"`
	Timeout                time.Duration   `json:"timeout,omitempty" yaml:"timeout,omitempty" xml:"timeout,omitempty"`
	Async                  bool            `json:"async,omitempty" yaml:"async,omitempty" xml:"async,omitempty"`
	Algo                   scram.Algorithm `json:"algo,omitempty" yaml:"algo,omitempty" xml:"algo,omitempty"`
	Protocol               string          `json:"protocol,omitempty" yaml:"protocol,omitempty" xml:"protocol,omitempty"`
	AllowAutoTopicCreation bool            `json:"allowAutoTopicCreation,omitempty" yaml:"allowAutoTopicCreation,omitempty" xml:"allowAutoTopicCreation,omitempty"`
}

func NewKafkaOptions() *Options {
	return &Options{
		Username:               "",
		Password:               "",
		Timeout:                time.Second * 5,
		Brokers:                make([]string, 0),
		Topics:                 make([]string, 0),
		Async:                  true,
		AllowAutoTopicCreation: true,
	}
}

// Dialer 返回kafka.Dialer方法
func Dialer(algo scram.Algorithm, username, password string, timeout time.Duration, dualStack bool) (sasl.Mechanism, *kafka.Dialer) {
	// 判断plain是否为空
	if algo == nil && username == "" && password == "" {
		return nil, &kafka.Dialer{
			Timeout:       timeout,
			DualStack:     dualStack,
			SASLMechanism: nil,
		}
	}
	// 初始化scram 加密mechanism
	if algo != nil && username != "" && password != "" {
		mechanism, err := scram.Mechanism(scram.SHA512, username, password)
		if err != nil {
			log.Fatal(err.Error())
		}
		return mechanism, &kafka.Dialer{
			Timeout:       timeout,
			DualStack:     dualStack,
			SASLMechanism: mechanism,
		}
	}
	// 初始化 plain明文mechanism
	if algo == nil && username != "" && password != "" {
		mechanism := plain.Mechanism{
			Username: username,
			Password: password,
		}
		return mechanism, &kafka.Dialer{
			Timeout:       timeout,
			DualStack:     dualStack,
			SASLMechanism: mechanism,
		}
	}
	return nil, nil
}
