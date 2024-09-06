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
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	"reflect"
	"testing"
	"time"
)

func TestAllCaseInOrder(t *testing.T) {
	t.Run("create topics", TestCreateTopics)
	t.Run("delete all topics", TestDeleteAllTopics)
	t.Run("delete topics", TestDeleteTopics)
	t.Run("get leader host", TestLeaderHost)
	t.Run("list topics", TestListTopics)
	t.Run("read topic value", TestRead)
	t.Run("write topic value", TestWrite)
	// TODO: Add test case.
}

func TestCreateTopics(t *testing.T) {
	type fields struct {
		Username               string
		Password               string
		Brokers                []string
		Topics                 []string
		Timeout                time.Duration
		Async                  bool
		Algo                   scram.Algorithm
		Protocol               string
		AllowAutoTopicCreation bool
	}
	type args struct {
		ctx               context.Context
		topics            []string
		partition         int
		replicationFactor int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "create kafka topic",
			fields: fields{
				Username: "root",
				Password: "password",
				Brokers: []string{
					"10.10.10.:9090",
				},
				Topics: []string{
					"topic-a",
				},
				AllowAutoTopicCreation: true,
				Protocol:               "TCP",
				Timeout:                10 * time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Options{
				Username:               tt.fields.Username,
				Password:               tt.fields.Password,
				Brokers:                tt.fields.Brokers,
				Topics:                 tt.fields.Topics,
				Timeout:                tt.fields.Timeout,
				Async:                  tt.fields.Async,
				Algo:                   tt.fields.Algo,
				Protocol:               tt.fields.Protocol,
				AllowAutoTopicCreation: tt.fields.AllowAutoTopicCreation,
			}
			if err := k.CreateTopics(tt.args.ctx, tt.args.topics, tt.args.partition, tt.args.replicationFactor); (err != nil) != tt.wantErr {
				t.Errorf("CreateTopics() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteAllTopics(t *testing.T) {
	type fields struct {
		Username               string
		Password               string
		Brokers                []string
		Topics                 []string
		Timeout                time.Duration
		Async                  bool
		Algo                   scram.Algorithm
		Protocol               string
		AllowAutoTopicCreation bool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Options{
				Username:               tt.fields.Username,
				Password:               tt.fields.Password,
				Brokers:                tt.fields.Brokers,
				Topics:                 tt.fields.Topics,
				Timeout:                tt.fields.Timeout,
				Async:                  tt.fields.Async,
				Algo:                   tt.fields.Algo,
				Protocol:               tt.fields.Protocol,
				AllowAutoTopicCreation: tt.fields.AllowAutoTopicCreation,
			}
			if err := k.DeleteAllTopics(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("DeleteAllTopics() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteTopics(t *testing.T) {
	type fields struct {
		Username               string
		Password               string
		Brokers                []string
		Topics                 []string
		Timeout                time.Duration
		Async                  bool
		Algo                   scram.Algorithm
		Protocol               string
		AllowAutoTopicCreation bool
	}
	type args struct {
		ctx    context.Context
		topics []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Options{
				Username:               tt.fields.Username,
				Password:               tt.fields.Password,
				Brokers:                tt.fields.Brokers,
				Topics:                 tt.fields.Topics,
				Timeout:                tt.fields.Timeout,
				Async:                  tt.fields.Async,
				Algo:                   tt.fields.Algo,
				Protocol:               tt.fields.Protocol,
				AllowAutoTopicCreation: tt.fields.AllowAutoTopicCreation,
			}
			if err := k.DeleteTopics(tt.args.ctx, tt.args.topics); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTopics() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLeaderHost(t *testing.T) {
	type fields struct {
		Username               string
		Password               string
		Brokers                []string
		Topics                 []string
		Timeout                time.Duration
		Async                  bool
		Algo                   scram.Algorithm
		Protocol               string
		AllowAutoTopicCreation bool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Options{
				Username:               tt.fields.Username,
				Password:               tt.fields.Password,
				Brokers:                tt.fields.Brokers,
				Topics:                 tt.fields.Topics,
				Timeout:                tt.fields.Timeout,
				Async:                  tt.fields.Async,
				Algo:                   tt.fields.Algo,
				Protocol:               tt.fields.Protocol,
				AllowAutoTopicCreation: tt.fields.AllowAutoTopicCreation,
			}
			got, err := k.LeaderHost(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("LeaderHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LeaderHost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListTopics(t *testing.T) {
	type fields struct {
		Username               string
		Password               string
		Brokers                []string
		Topics                 []string
		Timeout                time.Duration
		Async                  bool
		Algo                   scram.Algorithm
		Protocol               string
		AllowAutoTopicCreation bool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Options{
				Username:               tt.fields.Username,
				Password:               tt.fields.Password,
				Brokers:                tt.fields.Brokers,
				Topics:                 tt.fields.Topics,
				Timeout:                tt.fields.Timeout,
				Async:                  tt.fields.Async,
				Algo:                   tt.fields.Algo,
				Protocol:               tt.fields.Protocol,
				AllowAutoTopicCreation: tt.fields.AllowAutoTopicCreation,
			}
			got, err := k.ListTopics(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTopics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListTopics() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRead(t *testing.T) {
	type fields struct {
		Username               string
		Password               string
		Brokers                []string
		Topics                 []string
		Timeout                time.Duration
		Async                  bool
		Algo                   scram.Algorithm
		Protocol               string
		AllowAutoTopicCreation bool
	}
	type args struct {
		ctx         context.Context
		topic       string
		partition   int
		readTimeout time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *kafka.Batch
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Options{
				Username:               tt.fields.Username,
				Password:               tt.fields.Password,
				Brokers:                tt.fields.Brokers,
				Topics:                 tt.fields.Topics,
				Timeout:                tt.fields.Timeout,
				Async:                  tt.fields.Async,
				Algo:                   tt.fields.Algo,
				Protocol:               tt.fields.Protocol,
				AllowAutoTopicCreation: tt.fields.AllowAutoTopicCreation,
			}
			got, err := k.Read(tt.args.ctx, tt.args.topic, tt.args.partition, tt.args.readTimeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrite(t *testing.T) {
	type fields struct {
		Username               string
		Password               string
		Brokers                []string
		Topics                 []string
		Timeout                time.Duration
		Async                  bool
		Algo                   scram.Algorithm
		Protocol               string
		AllowAutoTopicCreation bool
	}
	type args struct {
		ctx          context.Context
		topic        string
		partition    int
		writeTimeout time.Duration
		key          []byte
		value        []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Options{
				Username:               tt.fields.Username,
				Password:               tt.fields.Password,
				Brokers:                tt.fields.Brokers,
				Topics:                 tt.fields.Topics,
				Timeout:                tt.fields.Timeout,
				Async:                  tt.fields.Async,
				Algo:                   tt.fields.Algo,
				Protocol:               tt.fields.Protocol,
				AllowAutoTopicCreation: tt.fields.AllowAutoTopicCreation,
			}
			if err := k.Write(tt.args.ctx, tt.args.topic, tt.args.partition, tt.args.writeTimeout, tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
