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

// ReplicationEvent 结构体
type ReplicationEvent struct {
	MetadataInfo `json:",inline"`
	EventData    ReplicationEventData `json:"event_data"`
}

type ReplicationEventData struct {
	Replication ReplicationData `json:"replication"`
}

// ReplicationData represents the replication details in the event
type ReplicationData struct {
	HarborHostname      string                `json:"harbor_hostname"`
	JobStatus           string                `json:"job_status"`
	ArtifactType        string                `json:"artifact_type"`
	OverrideMode        bool                  `json:"override_mode"`
	TriggerType         string                `json:"trigger_type"`
	PolicyCreator       string                `json:"policy_creator"`
	ExecutionTimestamp  int64                 `json:"execution_timestamp"`
	SrcResource         ReplicationResource   `json:"src_resource"`
	DestResource        ReplicationResource   `json:"dest_resource"`
	SuccessfulArtifacts []ReplicationArtifact `json:"successful_artifact"`
}

// ReplicationResource represents a source or destination resource in the replication event
type ReplicationResource struct {
	RegistryName string `json:"registry_name,omitempty"`
	RegistryType string `json:"registry_type"`
	Endpoint     string `json:"endpoint"`
	Namespace    string `json:"namespace"`
}

// ReplicationArtifact represents a successfully replicated artifact
type ReplicationArtifact struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	NameTag string `json:"name_tag"`
}
