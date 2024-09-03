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
	"context"
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/hurricane1988/harbor-webhook/pkg/db/mysql"
	"github.com/hurricane1988/harbor-webhook/pkg/redis"
	"github.com/spf13/viper"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// 全局常量定义
const (
	DefaultConfigurationName      = "config"
	DefaultConfigurationDirectory = "conf"
	DefaultConfigurationType      = "yaml"
	DefaultWorkDirectory          = "/etc"
)

// LoadFromDisk 初始化配置文件
func LoadFromDisk(ctx context.Context) (*Config, error) {
	Log := log.FromContext(ctx)
	// 获取当前程序执行路径
	workDir := filepath.Join(DefaultWorkDirectory, DefaultConfigurationDirectory, DefaultConfigurationName)
	// 加载viper获取配置路径
	viper.AddConfigPath(workDir)
	viper.AddConfigPath(homedir.HomeDir())
	viper.AddConfigPath(".")

	// 设置读取的文件名
	viper.SetConfigName(DefaultConfigurationName)
	// 设置读取的文件后缀
	viper.SetConfigType(DefaultConfigurationType)
	// 匹配环境变量
	viper.AutomaticEnv()

	// 执行读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			Log.Error(errors.New("failed load config file"),
				"config-directory", workDir,
				"config-file", DefaultConfigurationName+"."+DefaultConfigurationType,
			)
		}
	}
	conf := New()
	/* viper动态加载配置 */
	// 监视配置文件是否发生更改
	viper.WatchConfig()
	// 处理配置变更事件
	viper.OnConfigChange(func(in fsnotify.Event) {
		// TODO: 联调测试
		Log.Info("config file changed.",
			"config-directory", workDir,
			"config-file", DefaultConfigurationName+"."+DefaultConfigurationType,
		)
		// 重新初始化并生效配置
		err := viper.Unmarshal(conf)
		if err != nil {
			Log.Error(err, "data deserialization failed")
			return
		}
	})
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}
	return conf, nil
}

// New 初始化全局配置
func New() *Config {
	return &Config{
		RedisOptions: redis.NewRedisOptions(),
		MysqlOptions: mysql.NewMysqlOptions(),
	}
}
