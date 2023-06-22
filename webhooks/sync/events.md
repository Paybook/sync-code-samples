## Events

------
<br />

#### New credential was created 

 `credentials.created`

Message:

```json
{
  "events": [
    {
      "header": {
          "event": {
              "at": "2023-06-16 19:06:18.944575+0000",
              "eid": "6ff4da49-4193-4ee0-8e9e-7415b7fcb9cb",
              "name": "credentials.created",
              "version": "3.1"
          },
          "user": {
              "id_environment": "574894bf7848066d138b4571",
              "id_external": "20230614-DEMO",
              "id_user": "648a053b59f043380d473be3"
          }
      },
      "payload": {
          "endpoints": {
              "credential": [
                  "/v1/credentials/648cb2a99ca1b91b0a141391"
              ]
          },
          "event": "credential_create",
          "id_credential": "648cb2a99ca1b91b0a141391",
          "id_external": "20230614-DEMO",
          "id_job": "648cb2a9718d6c796e4091c1",
          "id_job_uuid": "648cb2a9718d6c796e4091c2",
          "id_site": "56cf5728784806f72b8b456f",
          "id_site_organization": "56cf4ff5784806152c8b4568",
          "id_site_organization_type": "56cf4f5b784806cf028b4569",
          "id_user": "648a053b59f043380d473be3",
          "is_executing": 1
      }
    }
  ],
  "rid": "c53b1fcb-e112-465c-ad3e-8b64cf2f602c"
}
```

------
<br />

#### Existing credential was updated

`credentials.updated`

Message:

```json
{
    "events": [
        {
            "header": {
                "event": {
                    "at": "2023-06-16 19:06:23.961107+0000",
                    "eid": "46346c7b-f990-49e7-96d9-0d33f7406b4f",
                    "name": "credentials.updated",
                    "version": "3.1"
                },
                "user": {
                    "id_environment": "574894bf7848066d138b4571",
                    "id_external": "20230614-DEMO",
                    "id_user": "648a053b59f043380d473be3"
                }
            },
            "payload": {
                "endpoints": {
                    "credential": [
                        "/v1/credentials/648cb2a99ca1b91b0a141391"
                    ]
                },
                "event": "credential_update",
                "id_credential": "648cb2a99ca1b91b0a141391",
                "id_external": "20230614-DEMO",
                "id_job": "648cb2a9718d6c796e4091c1",
                "id_job_uuid": "648cb2a9718d6c796e4091c2",
                "id_site": "56cf5728784806f72b8b456f",
                "id_site_organization": "56cf4ff5784806152c8b4568",
                "id_site_organization_type": "56cf4f5b784806cf028b4569",
                "id_user": "648a053b59f043380d473be3",
                "is_executing": 0
            }
        }
    ],
    "rid": "c2199b6a-7355-4dc8-98cf-a18305a3727f"
}
```

------
<br />

#### Data for an existing credential was added or updated

`credentials.refreshed` 

For a Existing credentials, the follow actions can be detected:

1. New account was added
2. New transactions was added
3. An existing accounts was updated
4. An existing transactions was updated

Message:

```json
{
    "events": [
        {
            "header": {
                "event": {
                    "at": "2023-06-16 19:06:42.721561+0000",
                    "eid": "2709cf46-0ebd-4856-ac26-421b378ff571",
                    "name": "credentials.refreshed",
                    "version": "3.1"
                },
                "user": {
                    "id_environment": "574894bf7848066d138b4571",
                    "id_external": "20230614-DEMO",
                    "id_user": "648a053b59f043380d473be3"
                }
            },
            "payload": {
                "endpoints": {
                    "accounts": [
                        "/v1/accounts?id_credential=648cb2aa8c7a0b6c2b2e6694&dt_refresh_from=1684213883&dt_refresh_to=1684213915&limit=5000&skip=0&wbhk=1"
                    ],
                    "attachments": [
                        "/v1/attachments/export?id_credential=648cb2aa8c7a0b6c2b2e6694&dt_refresh_from=1684213883&dt_refresh_to=1684213915&limit=5000&skip=0&wbhk=1"
                    ],
                    "credential": [
                        "/v1/credentials/648cb2aa8c7a0b6c2b2e6694"
                    ],
                    "transactions": [
                        "/v1/transactions?id_credential=648cb2aa8c7a0b6c2b2e6694&dt_refresh_from=1684213883&dt_refresh_to=1684213915&limit=5000&skip=0&wbhk=1"
                    ]
                },
                "event": "refresh",
                "id_credential": "648cb2aa8c7a0b6c2b2e6694",
                "id_external": "20230614-DEMO",
                "id_job": "648cb2aba5e16209404ee97c",
                "id_job_type": "57884c49784806f0768b4568",
                "id_job_uuid": "648cb2aba5e16209404ee97b",
                "id_job_uuid_type": "57884c49784806f0768b4568",
                "id_site": "59aefe28056f29793a58c091",
                "id_site_organization": "56cf4ff5784806152c8b4568",
                "id_site_organization_type": "56cf4f5b784806cf028b4569",
                "id_user": "648a053b59f043380d473be3"
            }
        }
    ],
    "rid": "d5a7fa51-e661-433b-a040-5286058f3605"
}
```

------
<br />

#### Documents have been downloaded

 `documents.completed`

Message:

```json
{
    "rid": "8653a68e-0e3f-4ab1-97e0-c773be77dfb4",
    "events": [
        {
            "header": {
                "event": {
                    "eid": "91f5f406-cb4a-417c-8dc1-fde7482ad27a",
                    "name": "documents.completed",
                    "at": "2023-06-16 19:06:57.591579+0000",
                    "version": "3.1"
                },
                "user": {
                    "id_user": "648a053b59f043380d473be3",
                    "id_external": "20230614-DEMO",
                    "id_environment": "574894bf7848066d138b4571"
                }
            },
            "payload": {
                "id_job": "648cb2ab199e4b39654a9d20",
                "id_job_uuid": "648cb2ab199e4b39654a9d21",
                "id_site": "59aefe28056f29793a58c092",
                "id_site_organization": "56cf4ff5784806152c8b4568",
                "id_site_organization_type": "56cf4f5b784806cf028b4569",
                "endpoints": [
                    "/v1/documents?id_job_webhook=648cb2d190b1bf61cc50b3ba"
                ]
            }
        }
    ]
}
```

------
<br />

#### Documents have been downloaded successfully

 `documents.success`

Message:

```json
{
    "rid": "0dd264d9-323e-43ff-9182-af3c895707f0",
    "events": [
        {
            "header": {
                "event": {
                    "eid": "b1c38530-661f-41ea-911a-9fbd9b197c3e",
                    "name": "documents.success",
                    "at": "2023-06-16 19:06:57.596733+0000",
                    "version": "3.1"
                },
                "user": {
                    "id_user": "648a053b59f043380d473be3",
                    "id_external": "20230614-DEMO",
                    "id_environment": "574894bf7848066d138b4571"
                }
            },
            "payload": {
                "id_job": "648cb2ab199e4b39654a9d20",
                "id_job_uuid": "648cb2ab199e4b39654a9d21",
                "id_site": "59aefe28056f29793a58c092",
                "id_site_organization": "56cf4ff5784806152c8b4568",
                "id_site_organization_type": "56cf4f5b784806cf028b4569",
                "endpoints": [
                    "/v1/documents?id_job_webhook=648cb2d190b1bf61cc50b3bd"
                ]
            }
        }
    ]
}
```

------
<br />


#### Documents have been downloaded unsuccessfully and there may be errors

 `documents.fail`

One or more errors occurred while downloading the documents. There may be inconsistencies in the document.

Message:

```json
{
	"rid": "d9da045e-c5fd-4fa0-a7cf-6372022c18fd",
	"events": [
		{
			"header": {
				"event": {
					"eid": "7efce48a-95d1-4739-9ca4-80a2bc2a7610",
					"name": "documents.failed",
					"at": "2023-06-13 19:49:49.082012+0000",
					"version": "3.1"
				},
				"user": {
					"id_user": "648a053b59f043380d473be3",
					"id_external": "20230614-DEMO",
					"id_environment": "574894bf7848066d138b4571"
				}
			},
			"payload": {
				"id_job": "6488c853711c927c657994b7",
				"id_job_uuid": "6488c853711c927c657994b8",
				"id_site": "63ae0c7be5d6f1905cb9d3dd",
				"id_site_organization": "5c65e306f9de2a080b61bb31",
				"id_site_organization_type": "56cf4f5b784806cf028b4569",
				"endpoints": [
					"/v1/documents?id_job_webhook=6488c85d10fe6447df2fa3e9"
				]
			}
		}
	]
}
```