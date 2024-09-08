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

// QuotaEvent 结构体
type QuotaEvent struct {
	MetadataInfo `json:",inline"`
	EventData    QuotaEventData `json:"event_data"`
}

type QuotaEventData struct {
	Resources        []QuotaResource  `json:"resources"`
	Repository       QuotaRepository  `json:"repository"`
	CustomAttributes CustomAttributes `json:"custom_attributes"`
}

// QuotaResource represents the details of the resources involved in the event
type QuotaResource struct {
	Digest string `json:"digest"`
}

// QuotaRepository contains repository details
type QuotaRepository struct {
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	RepoFullName string `json:"repo_full_name"`
	RepoType     string `json:"repo_type"`
}

// CustomAttributes contains additional details related to the event
type CustomAttributes struct {
	Details string `json:"Details"`
}
