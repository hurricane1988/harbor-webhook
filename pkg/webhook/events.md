## PUSH_ARTIFACT JSON
```json
{
  "type": "PUSH_ARTIFACT",
  "occur_at": 1680501893,
  "operator": "harbor-jobservice",
  "event_data": {
    "resources": [
      {
        "digest": "sha256:954b378c375d852eb3c63ab88978f640b4348b01c1b3456a024a81536dafbbf4",
        "tag": "sha256:954b378c375d852eb3c63ab88978f640b4348b01c1b3456a024a81536dafbbf4",
        "resource_url": "localhost/harbor/alpine@sha256:954b378c375d852eb3c63ab88978f640b4348b01c1b3456a024a81536dafbbf4"
      }
    ],
    "repository": {
      "date_created": 1680501893,
      "name": "alpine",
      "namespace": "harbor",
      "repo_full_name": "harbor/alpine",
      "repo_type": "private"
    }
  }
}
```

## PULL_ARTIFACT JSON
```json
{
  "type": "PULL_ARTIFACT",
  "occur_at": 1680502372,
  "operator": "robot$harbor+wHSYCuGD-Trivy-8e2e7505-d1e6-11ed-9e71-0242ac130009",
  "event_data": {
    "resources": [
      {
        "digest": "sha256:954b378c375d852eb3c63ab88978f640b4348b01c1b3456a024a81536dafbbf4",
        "tag": "sha256:954b378c375d852eb3c63ab88978f640b4348b01c1b3456a024a81536dafbbf4",
        "resource_url": "localhost/harbor/alpine@sha256:954b378c375d852eb3c63ab88978f640b4348b01c1b3456a024a81536dafbbf4"
      }
    ],
    "repository": {
      "date_created": 1680501893,
      "name": "alpine",
      "namespace": "harbor",
      "repo_full_name": "harbor/alpine",
      "repo_type": "private"
    }
  }
}
```
## Artifact deleted JSON
```json
{
  "type": "DELETE_ARTIFACT",
  "occur_at": 1680502598,
  "operator": "harbor-jobservice",
  "event_data": {
    "resources": [
      {
        "digest": "sha256:2bb501e6173d9d006e56de5bce2720eb06396803300fe1687b58a7ff32bf4c14",
        "tag": "3.8",
        "resource_url": "localhost/harbor/alpine:3.8"
      }
    ],
    "repository": {
      "date_created": 1680501893,
      "name": "alpine",
      "namespace": "harbor",
      "repo_full_name": "harbor/alpine",
      "repo_type": "private"
    }
  }
}
```

## Scanning completed JSON
```json
{
  "type": "SCANNING_COMPLETED",
  "occur_at": 1680502375,
  "operator": "auto",
  "event_data": {
    "resources": [
      {
        "digest": "sha256:954b378c375d852eb3c63ab88978f640b4348b01c1b3456a024a81536dafbbf4",
        "resource_url": "localhost/harbor/alpine@sha256:954b378c375d852eb3c63ab88978f640b4348b01c1b3456a024a81536dafbbf4",
        "scan_overview": {
          "application/vnd.security.vulnerability.report; version=1.1": {
            "report_id": "af0546c1-67dc-4e9d-927e-372900ead0df",
            "scan_status": "Success",
            "severity": "None",
            "duration": 8,
            "summary": {
              "total": 0,
              "fixable": 0,
              "summary": {}
            },
            "start_time": "2023-04-03T06:12:47Z",
            "end_time": "2023-04-03T06:12:55Z",
            "scanner": {
              "name": "Trivy",
              "vendor": "Aqua Security",
              "version": "v0.37.2"
            },
            "complete_percent": 100
          }
        }
      }
    ],
    "repository": {
      "name": "alpine",
      "namespace": "harbor",
      "repo_full_name": "harbor/alpine",
      "repo_type": "private"
    }
  }
}
```
## Scanning stopped JSON
```json
{
  "type": "SCANNING_STOPPED",
  "occur_at": 1680502334,
  "operator": "auto",
  "event_data": {
    "resources": [
      {
        "digest": "sha256:e802987f152d7826cf929ad4999fb3bb956ce7a30966aeb46c749f9120eaf22c",
        "resource_url": "localhost/harbor/alpine@sha256:e802987f152d7826cf929ad4999fb3bb956ce7a30966aeb46c749f9120eaf22c",
        "scan_overview": {
          "application/vnd.security.vulnerability.report; version=1.1": {
            "report_id": "bf92700b-fa5e-4fe4-891c-42b730c81151",
            "scan_status": "Stopped",
            "severity": "",
            "duration": 5,
            "summary": null,
            "start_time": "2023-04-03T06:12:09Z",
            "end_time": "2023-04-03T06:12:14Z",
            "complete_percent": 0
          }
        }
      }
    ],
    "repository": {
      "name": "alpine",
      "namespace": "harbor",
      "repo_full_name": "harbor/alpine",
      "repo_type": "private"
    }
  }
}
```
## Scanning failed JSON
```json
{
  "type": "SCANNING_FAILED",
  "occur_at": 1680505885,
  "operator": "auto",
  "event_data": {
    "resources": [
      {
        "digest": "sha256:dabea2944dcc2b86482b4f0b0fb62da80e0673e900c46c0e03b45919881a5d84",
        "resource_url": "localhost/harbor/alpine@sha256:dabea2944dcc2b86482b4f0b0fb62da80e0673e900c46c0e03b45919881a5d84",
        "scan_overview": {
          "application/vnd.security.vulnerability.report; version=1.1": {
            "report_id": "a2573415-c727-4723-bc92-376c1d978637",
            "scan_status": "Error",
            "severity": "",
            "duration": 10,
            "summary": null,
            "start_time": "2023-04-03T07:11:15Z",
            "end_time": "2023-04-03T07:11:25Z",
            "complete_percent": 0
          }
        }
      }
    ],
    "repository": {
      "name": "alpine",
      "namespace": "harbor",
      "repo_full_name": "harbor/alpine",
      "repo_type": "private"
    }
  }
}
```

## Quota exceeded JSON
```json
{
  "type": "QUOTA_EXCEED",
  "occur_at": 1680505484,
  "operator": "",
  "event_data": {
    "resources": [
      {
        "digest": "sha256:402d21757a03a114d273bbe372fa4b9eca567e8b6c332fa7ebf982b902207242"
      }
    ],
    "repository": {
      "name": "alpine",
      "namespace": "harbor",
      "repo_full_name": "harbor/alpine",
      "repo_type": "private"
    },
    "custom_attributes": {
      "Details": "adding 2.1 MiB of storage resource, which when updated to current usage of 8.3 MiB will exceed the configured upper limit of 10.0 MiB."
    }
  }
}
```

## Quota near threshold JSON
```json
{
  "type": "QUOTA_WARNING",
  "occur_at": 1680505653,
  "operator": "",
  "event_data": {
    "resources": [
      {
        "digest": "sha256:514ec80ffbe1a2ab1d9a3d5e6082296296a1d8b6870246edf897228e5df2367d"
      }
    ],
    "repository": {
      "name": "alpine",
      "namespace": "harbor",
      "repo_full_name": "harbor/alpine",
      "repo_type": "private"
    },
    "custom_attributes": {
      "Details": "quota usage reach 85%: resource storage used 12.6 MiB of 14.0 MiB"
    }
  }
}
```

## Replication status changed JSON
```json
{
  "type": "REPLICATION",
  "occur_at": 1680501904,
  "operator": "MANUAL",
  "event_data": {
    "replication": {
      "harbor_hostname": "localhost",
      "job_status": "Success",
      "artifact_type": "image",
      "override_mode": true,
      "trigger_type": "MANUAL",
      "policy_creator": "admin",
      "execution_timestamp": 1680501881,
      "src_resource": {
        "registry_name": "hub",
        "registry_type": "docker-hub",
        "endpoint": "https://hub.docker.com",
        "namespace": "library"
      },
      "dest_resource": {
        "registry_type": "harbor",
        "endpoint": "http://localhost",
        "namespace": "harbor"
      },
      "successful_artifact": [
        {
          "type": "image",
          "status": "Success",
          "name_tag": "alpine [1 item(s) in total]"
        }
      ]
    }
  }
}
```

## Tag retention finished JSON
```json
{
  "type": "TAG_RETENTION",
  "occur_at": 1680502598,
  "operator": "MANUAL",
  "event_data": {
    "retention": {
      "total": 1,
      "retained": 0,
      "harbor_hostname": "localhost",
      "project_name": "harbor",
      "retention_policy_id": 2,
      "retention_rule": [
        {
          "template": "always",
          "tag_selectors": [
            {
              "kind": "doublestar",
              "decoration": "matches",
              "pattern": "xxxxxxxxxx",
              "extras": "{\"untagged\":true}"
            }
          ],
          "scope_selectors": {
            "repository": [
              {
                "kind": "doublestar",
                "decoration": "repoMatches",
                "pattern": "**",
                "extras": ""
              }
            ]
          }
        }
      ],
      "result": "SUCCESS",
      "deleted_artifact": [
        {
          "type": "image",
          "status": "SUCCESS",
          "name_tag": "alpine:3.8"
        }
      ]
    }
  }
}
```
