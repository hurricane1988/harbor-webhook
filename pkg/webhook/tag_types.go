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

// TagEvent 结构体
type TagEvent struct {
	MetadataInfo `json:",inline"`
	EventData    TagEventData `json:"event_data"`
}

type TagEventData struct {
	Retention RetentionData `json:"retention"`
}

// RetentionData represents the tag retention details in the event
type RetentionData struct {
	Total             int                `json:"total"`
	Retained          int                `json:"retained"`
	HarborHostname    string             `json:"harbor_hostname"`
	ProjectName       string             `json:"project_name"`
	RetentionPolicyID int                `json:"retention_policy_id"`
	RetentionRule     []TagRetentionRule `json:"retention_rule"`
	Result            string             `json:"result"`
	DeletedArtifacts  []TagArtifact      `json:"deleted_artifact"`
}

// TagRetentionRule represents a retention rule for tag retention
type TagRetentionRule struct {
	Template       string        `json:"template"`
	TagSelectors   []TagSelector `json:"tag_selectors"`
	ScopeSelectors ScopeSelector `json:"scope_selectors"`
}

// TagSelector represents the tag selection criteria in the retention rule
type TagSelector struct {
	Kind       string `json:"kind"`
	Decoration string `json:"decoration"`
	Pattern    string `json:"pattern"`
	Extras     string `json:"extras"`
}

// ScopeSelector represents the repository selection criteria in the retention rule
type ScopeSelector struct {
	Repository []TagRepositorySelector `json:"repository"`
}

// TagRepositorySelector represents the details of repository selection
type TagRepositorySelector struct {
	Kind       string `json:"kind"`
	Decoration string `json:"decoration"`
	Pattern    string `json:"pattern"`
	Extras     string `json:"extras"`
}

// TagArtifact represents a deleted artifact during tag retention
type TagArtifact struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	NameTag string `json:"name_tag"`
}
