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

import "time"

// ScanningEvent 结构体
type ScanningEvent struct {
	MetadataInfo `json:",inline"`
	EventData    ScanningEventData `json:"event_data"`
}

type ScanningEventData struct {
	Resources  []ScanningResource `json:"resources"`
	Repository ScanningRepository `json:"repository"`
}

type ScanningResource struct {
	Digest       string                    `json:"digest"`
	ResourceURL  string                    `json:"resource_url"`
	ScanOverview map[string]ScanningReport `json:"scan_overview"`
}

type ScanningReport struct {
	ReportID        string          `json:"report_id"`
	ScanStatus      string          `json:"scan_status"`
	Severity        string          `json:"severity"`
	Duration        int             `json:"duration"`
	Summary         ScanningSummary `json:"summary"`
	StartTime       time.Time       `json:"start_time"`
	EndTime         time.Time       `json:"end_time"`
	Scanner         ScanningScanner `json:"scanner"`
	CompletePercent int             `json:"complete_percent"`
}

type ScanningSummary struct {
	Total   int            `json:"total"`
	Fixable int            `json:"fixable"`
	Details map[string]int `json:"summary"`
}

type ScanningScanner struct {
	Name    string `json:"name"`
	Vendor  string `json:"vendor"`
	Version string `json:"version"`
}

type ScanningRepository struct {
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	RepoFullName string `json:"repo_full_name"`
	RepoType     string `json:"repo_type"`
}
