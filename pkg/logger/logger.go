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

package logger

import (
	"flag"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/hurricane1988/harbor-webhook/pkg/constants"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
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

// GetLoggerByName returns a logger that writes to std or both std and disk
// files if onDisk option is set to true. Log files are exactly
// named after given logger name.
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

// Logger returns common logger.
func Logger() logr.Logger {
	return GetLoggerByName("common")
}
