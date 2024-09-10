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

package webhook

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hurricane1988/harbor-webhook/pkg/constants"
	"github.com/hurricane1988/harbor-webhook/pkg/utils"
	"net/http"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func handleWebhook(c *gin.Context) {
	var (
		metadata MetadataInfo
		ctx      = context.TODO()
	)
	// 初始化日志
	logger := log.FromContext(ctx)
	if err := c.ShouldBindJSON(&metadata); err != nil {
		logger.Error(err, "failed to parse webhook event.")
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "invalid json",
		})
		return
	}

	// 根据事件类型选择处理函数
	switch metadata.Type {
	case string(constants.EventTypePushArtifact), string(constants.EventTypeDeleteArtifact), string(constants.EventTypePullArtifact):
		var event ArtifactEvent
		if err := c.ShouldBindJSON(&event); err == nil {
			logger.Info("artifact event is triggered",
				"event", event,
				"event-type", metadata.Type,
			)
		}
	case string(constants.EventTypeProjectQuotaExceeded), string(constants.EventTypeProjectQuotaNearThreshold):
		var event QuotaEvent
		if err := c.ShouldBindJSON(&event); err == nil {
			logger.Info("project quota event is triggered",
				"event", event,
				"event-type", metadata.Type,
			)
		}
	case string(constants.EventTypeArtifactReplicationStatusChanged):
		var event ReplicationEvent
		if err := c.ShouldBindJSON(&event); err == nil {
			logger.Info("project replication event is triggered",
				"event", event,
				"event-type", metadata.Type,
			)
		}
	case string(constants.EventTypeArtifactScanFailed), string(constants.EventTypeArtifactScanStopped), string(constants.EventTypeArtifactScanCompleted):
		var event ScanningEvent
		if err := c.ShouldBindJSON(&event); err == nil {
			logger.Info("artifact security scanning event is triggered",
				"event", event,
				"event-type", metadata.Type,
			)
		}
	case string(constants.EventTypeArtifactTagRetentionFinished):
		var event TagEvent
		if err := c.ShouldBindJSON(&event); err == nil {
			logger.Info("artifact tag event is triggered",
				"event", event,
				"event-type", metadata.Type,
			)
		}
	default:
		logger.Info("unknown event type", "event-type", metadata.Type)
	}

	// 返回成功响应
	c.JSONP(http.StatusOK, gin.H{
		"event":     metadata.Type,
		"operator":  metadata.Operator,
		"timestamp": utils.ParseTime(metadata.OccurAt)})
}
