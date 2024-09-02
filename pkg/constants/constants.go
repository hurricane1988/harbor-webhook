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

package constants

const (
	// EnvLogOutput is used to specify where to output logs, which can be
	// either stdout or on-disk file.
	EnvLogOutput = "log.output" // `std` or `file`
)

// EventType Harbor Webhook Event Type
type EventType string

// Harbor webhook官方的 Events 类型
// Refer to: https://goharbor.io/docs/main/working-with-projects/project-configuration/configure-webhooks/
// 可以为每个项目定义多个 webhook 端点。Harbor 目前支持两种类型的端点：HTTP 和 SLACK
// Webhook 通知以 JSON 格式提供有关事件的信息，并通过 HTTP 或 HTTPS POST 发送到您提供的现有 webhook 端点 URL 或 Slack 地址
// 支持两种 JSON 格式的 webhook 负载：默认格式是一直以来使用的格式，其数据结构与以前的版本没有变化，
// 只是现在有了名称；CloudEvents 是另一种格式，按照 CloudEvents 规范组织负载数据
const (
	EventTypePushArtifact                     = EventType("PUSH_ARTIFACT")
	EventTypePullArtifact                     = EventType("PULL_ARTIFACT")
	EventTypeDeleteArtifact                   = EventType("DELETE_ARTIFACT")
	EventTypeArtifactScanCompleted            = EventType("SCANNING_COMPLETED")
	EventTypeArtifactScanStopped              = EventType("SCANNING_STOPPED")
	EventTypeArtifactScanFailed               = EventType("SCANNING_FAILED")
	EventTypeProjectQuotaExceeded             = EventType("QUOTA_EXCEED")
	EventTypeProjectQuotaNearThreshold        = EventType("QUOTA_WARNING")
	EventTypeArtifactReplicationStatusChanged = EventType("REPLICATION")
	EventTypeArtifactTagRetentionFinished     = EventType("TAG_RETENTION")
)
