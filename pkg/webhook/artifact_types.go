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

// ArtifactEvent 结构体
type ArtifactEvent struct {
	MetadataInfo `json:",inline"`
	EventData    ArtifactEventData `json:"event_data"`
}

// ArtifactEventData represents the data part of the event
type ArtifactEventData struct {
	Resources  []ArtifactResources `json:"resources"`
	Repository ArtifactRepository  `json:"repository"`
}

// ArtifactResources represents the details of the resources involved in the event
type ArtifactResources struct {
	Digest      string `json:"digest"`
	Tag         string `json:"tag"`
	ResourceURL string `json:"resource_url"`
}

// ArtifactRepository contains repository details
type ArtifactRepository struct {
	DateCreated  int64  `json:"date_created"`
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	RepoFullName string `json:"repo_full_name"`
	RepoType     string `json:"repo_type"`
}
