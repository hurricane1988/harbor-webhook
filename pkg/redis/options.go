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

type Options struct {
	Address  []string `json:"address,omitempty" yaml:"address,omitempty" xml:"address,omitempty"`
	Username string   `json:"username,omitempty" yaml:"username,omitempty" xml:"username,omitempty"`
	Password string   `json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	DB       int32    `json:"db,omitempty" yaml:"db,omitempty" xml:"db,omitempty"`
}

func NewRedisOptions() *Options {
	return &Options{
		Address:  make([]string, 0),
		Username: "",
		Password: "",
		DB:       0,
	}
}
