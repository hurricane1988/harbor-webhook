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

package mysql

import (
	"fmt"
	"github.com/hurricane1988/harbor-webhook/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
	"time"
)

// TODO: https://gorm.io/

var (
	db   *gorm.DB
	once sync.Once
	err  error
)

type Interface interface {
	Client() *gorm.DB
}

// NewMySQL 根据提供的配置创建一个MySQL数据库接口
func NewMySQL(config Options) Interface {
	return &config
}
func (r *Options) Client() *gorm.DB {
	// Initialize log for GORM
	gormLogger := &utils.GormLogger{
		Log:           utils.Logger(),
		LogLevel:      logger.Info,
		SlowThreshold: 200 * time.Millisecond, // Threshold for slow queries
	}
	// 定义 MySQL DSN (数据源名称)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", r.Username, r.Password, r.Host, r.Port, r.DB)
	once.Do(func() {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction:                   true,       // 是否跳过默认的事务。默认情况下，每个写操作都会在事务中执行，这样可以确保操作的原子性
			FullSaveAssociations:                     true,       // 在更新时是否保存所有关联字段。如果为 true，会在更新时保存所有的关联关系
			DryRun:                                   false,      // 启用 DryRun 模式，不会执行 SQL，只会生成 SQL 查询语句，适用于调试
			PrepareStmt:                              true,       // 开启 PrepareStmt 模式，GORM 会缓存 SQL 语句以提升性能
			DisableAutomaticPing:                     false,      // 禁用自动 Ping 数据库的功能。GORM 默认会在建立连接后 Ping 数据库以确保连接正常
			DisableForeignKeyConstraintWhenMigrating: true,       // 禁用在迁移时创建外键约束的功能
			IgnoreRelationshipsWhenMigrating:         false,      // 在迁移时忽略关联关系
			DisableNestedTransaction:                 true,       // 禁用嵌套事务支持
			AllowGlobalUpdate:                        false,      // 允许全局更新。默认情况下，不提供条件的 `UPDATE` 或 `DELETE` 操作会被 GORM 拒绝
			Logger:                                   gormLogger, // 自定义日志
		})
		if err != nil {
			panic(err)
		}

		// 获取底层的 *sql.DB 对象
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("failed to get database connect pool：%v", err)
		}

		// 设置连接池参数
		sqlDB.SetMaxIdleConns(10)                  // 最大空闲连接数
		sqlDB.SetMaxOpenConns(100)                 // 最大打开连接数
		sqlDB.SetConnMaxLifetime(time.Hour)        // 连接最大存活时间
		sqlDB.SetConnMaxIdleTime(30 * time.Minute) // 连接最大空闲时间
	})
	return db
}
