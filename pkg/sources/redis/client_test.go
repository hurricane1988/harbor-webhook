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
	"reflect"
	"testing"
	"time"
)

func TestAllCaseInOrder(t *testing.T) {
	t.Run("redis del", TestOptions_Del)
	t.Run("redis set", TestOptions_Set)
	t.Run("redis get", TestOptions_Get)
	t.Run("redis existed", TestOptions_Existed)
	t.Run("redis keys", TestOptions_Keys)
	t.Run("redis expire", TestOptions_Expire)
}

func TestOptions_Del(t *testing.T) {
	type fields struct {
		Address  []string
		Username string
		Password string
		DB       int32
	}
	type args struct {
		ctx  context.Context
		keys []string
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
			r := &Options{
				Address:  tt.fields.Address,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				DB:       tt.fields.DB,
			}
			if err := r.Del(tt.args.ctx, tt.args.keys...); (err != nil) != tt.wantErr {
				t.Errorf("Del() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOptions_Existed(t *testing.T) {
	type fields struct {
		Address  []string
		Username string
		Password string
		DB       int32
	}
	type args struct {
		ctx  context.Context
		keys []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Options{
				Address:  tt.fields.Address,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				DB:       tt.fields.DB,
			}
			if got := r.Existed(tt.args.ctx, tt.args.keys...); got != tt.want {
				t.Errorf("Existed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_Expire(t *testing.T) {
	type fields struct {
		Address  []string
		Username string
		Password string
		DB       int32
	}
	type args struct {
		ctx      context.Context
		key      string
		duration time.Duration
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
			r := &Options{
				Address:  tt.fields.Address,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				DB:       tt.fields.DB,
			}
			if err := r.Expire(tt.args.ctx, tt.args.key, tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("Expire() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOptions_Get(t *testing.T) {
	type fields struct {
		Address  []string
		Username string
		Password string
		DB       int32
	}
	type args struct {
		ctx context.Context
		key string
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
			r := &Options{
				Address:  tt.fields.Address,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				DB:       tt.fields.DB,
			}
			got, err := r.Get(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_Keys(t *testing.T) {
	type fields struct {
		Address  []string
		Username string
		Password string
		DB       int32
	}
	type args struct {
		ctx     context.Context
		pattern string
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
			r := &Options{
				Address:  tt.fields.Address,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				DB:       tt.fields.DB,
			}
			got, err := r.Keys(tt.args.ctx, tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("Keys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_Set(t *testing.T) {
	type fields struct {
		Address  []string
		Username string
		Password string
		DB       int32
	}
	type args struct {
		ctx        context.Context
		key        string
		value      string
		expiration time.Duration
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
			r := &Options{
				Address:  tt.fields.Address,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				DB:       tt.fields.DB,
			}
			if err := r.Set(tt.args.ctx, tt.args.key, tt.args.value, tt.args.expiration); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
