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

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hurricane1988/harbor-webhook/pkg/metrics"
	"strconv"
	"time"
)

// ApiSummary 记录API接口请求Metric指标
func ApiSummary(c *gin.Context) {
	// 开始时间
	startTime := time.Now()
	c.Next()
	// 跳过Metrics路劲
	if c.FullPath() == "/metrics" {
		return
	}
	// 记录合法操作
	metrics.ApiSummary.With(map[string]string{
		"url":    c.Request.RequestURI,
		"status": strconv.Itoa(c.Writer.Status()),
		"api":    c.FullPath(),
		"method": c.Request.Method,
		"client": c.RemoteIP(),
		"host":   c.Request.Host,
	}).Observe(float64(time.Since(startTime).Milliseconds()))
}
