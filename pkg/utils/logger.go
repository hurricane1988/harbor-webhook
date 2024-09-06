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

package utils

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/hurricane1988/harbor-webhook/pkg/constants"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"gorm.io/gorm/logger"
	"io"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sync"
	"time"
)

var (
	loggerMap sync.Map
	zapOpts   zap.Options

	logPath           string
	logFileMaxAge     time.Duration
	logFileRotateTime time.Duration

	onDisk = true
)

func init() {
	zapOpts.Development = true
	if os.Getenv(constants.EnvLogOutput) == "std" {
		onDisk = false
	}
	flag.StringVar(&logPath, "log-file-path", "/home/admin/logs", "specify log file path, e.g., /home/admin/logs")
	flag.DurationVar(&logFileMaxAge, "log-file-max-age", 7*24*time.Hour, "max age of log files, defaults to 7 days")
	flag.DurationVar(&logFileRotateTime, "log-file-rotate-time", 4*time.Hour, "time interval for roatating log files, defaults to 4 hour")
}

func GetLogWriter(name string) io.Writer {
	if !onDisk {
		return os.Stderr
	}
	logFileName := fmt.Sprintf("%s/%s.log", logPath, name)
	rotateLogWriter, _ := rotatelogs.New(
		logFileName+".%Y-%m-%d-%H",
		rotatelogs.WithLinkName(logFileName),
		rotatelogs.WithMaxAge(logFileMaxAge),
		rotatelogs.WithRotationTime(logFileRotateTime),
	)
	return io.MultiWriter(os.Stderr, rotateLogWriter)
}

// GetLoggerByName returns a log that writes to std or both std and disk
// files if onDisk option is set to true. Log files are exactly
// named after given log name.
func GetLoggerByName(name string) logr.Logger {
	if value, ok := loggerMap.Load(name); ok {
		return value.(logr.Logger)
	}

	newLogger := zap.New(
		zap.UseFlagOptions(&zapOpts),
		zap.WriteTo(GetLogWriter(name)),
	)
	loggerMap.Store(name, newLogger)

	return newLogger
}

// Logger returns common log.
func Logger() logr.Logger {
	return GetLoggerByName("common")
}

// GormLogger 自定义的GORM 日志对象
type GormLogger struct {
	Log           logr.Logger
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

// LogMode sets the logging level.
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info logs the information level messages.
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.Log.Info(fmt.Sprintf(msg, data...))
	}
}

// Warn logs the warning level messages.
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.Log.Info(fmt.Sprintf(msg, data...))
	}
}

// Error logs the error level messages.
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.Log.Error(nil, fmt.Sprintf(msg, data...))
	}
}

// Trace logs SQL execution details.
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		l.Log.Error(err, fmt.Sprintf("SQL execution failed [%.3fms] [rows:%v] %s | error: %v", float64(elapsed.Nanoseconds())/1e6, rows, sql, err))
	} else if elapsed > l.SlowThreshold {
		l.Log.Info(fmt.Sprintf("SLOW SQL [%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql))
	} else {
		l.Log.Info(fmt.Sprintf("SQL [%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql))
	}
}
