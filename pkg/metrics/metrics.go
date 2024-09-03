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

package metrics

import "github.com/prometheus/client_golang/prometheus"

// ApiSummary Prometheus metric指标
var (
	ApiSummary = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "harbor_api_request_summary",
		Help: "该指标记录了请求接口的次数与用时(ms),用于度量接口是否需要缓存. label { api:请求接口, method:请求方法, client: 客户端IP, host: 服务IP, status: 状态码, url: 请求地址 }",
	}, []string{"api", "method", "client", "host", "url", "status"})
)

// CacheHit 缓存命中率 Prometheus metric指标.
var (
	CacheHit = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "harbor_cache_hit_summary",
		Help: "该指标记录了缓存命中,用于度量业务缓存是否需要缓存,缓存过期时间是否合理. label { key: 缓存key}",
	}, []string{"key"})
)
