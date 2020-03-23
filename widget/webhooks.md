# Webhooks

A webhook is a HTTP callback to a specified URL. They are triggered each time data is updated in Sync to help you stay up to date with the latest changes.

```json
{
    "id_webhook": "5d9bab5b8c91e73b2e4c75d3",
    "id_user": null,
    "is_disabled": 1,
    "events": [
        "credential_create",
        "credential_update",
        "refresh"
    ],
    "url": "https://webhook_domain/my_webhook",
    "delay": 0,
    "ct_failed": 1001,
    "dt_created": "2019-10-07T21:17:15+00:00",
    "dt_modified": "2020-03-20T05:37:43+00:00"
}
```



## Events

------
<br />

`credential_create`

New credential was created

```json
{
    "endpoints": {
        "credential": [
            "/v1/credentials/5e793889ea4c5165d46f9811"
        ]
    }, 
    "event": "credential_create", 
    "id_credential": "5e793889ea4c5165d46f9811", 
    "id_external": "IXS7607092R5", 
    "id_job": "5e793889b783081dd61de29c", 
    "id_job_uuid": "5e793889b783081dd61de29b", 
    "id_site": "5da784f1f9de2a06483abec1", 
    "id_site_organization": "56cf4ff5784806152c8b4567", 
    "id_site_organization_type": "56cf4f5b784806cf028b4569", 
    "id_user": "5e751d54d0510472295413b3", 
    "is_executing": 1
}
```

------
<br />

`credential_update`

Existing credential was updated

```json
{
    "endpoints": {
        "credential": [
            "/v1/credentials/5e793889ea4c5165d46f9811"
        ]
    }, 
    "event": "credential_update", 
    "id_credential": "5e793889ea4c5165d46f9811", 
    "id_external": "IXS7607092R5", 
    "id_job": "5e793889b783081dd61de29c", 
    "id_job_uuid": "5e793889b783081dd61de29b", 
    "id_site": "5da784f1f9de2a06483abec1", 
    "id_site_organization": "56cf4ff5784806152c8b4567", 
    "id_site_organization_type": "56cf4f5b784806cf028b4569", 
    "id_user": "5e751d54d0510472295413b3", 
    "is_executing": 1
}
```

------
<br />

`refresh` 

For a Existing credentials, the follow actions can be detected:

1. New account was added

2. New transactions wass added

3. An existing accounts was updated

4. An existing transactions was updated

```json
{
    "endpoints": {
        "accounts": [
            "/v1/accounts?id_credential=5e793889ea4c5165d46f9811&limit=5000&skip=0&wbhk=1"
        ], 
        "attachments": [
            "/v1/attachments/export?id_credential=5e793889ea4c5165d46f9811&limit=5000&skip=0&wbhk=1", 
            "/v1/attachments/export?id_credential=5e793889ea4c5165d46f9811&limit=5000&skip=5000&wbhk=1", 
            "/v1/attachments/export?id_credential=5e793889ea4c5165d46f9811&limit=5000&skip=10000&wbhk=1"
        ], 
        "credential": [
            "/v1/credentials/5e793889ea4c5165d46f9811"
        ], 
        "transactions": [
            "/v1/transactions?id_credential=5e793889ea4c5165d46f9811&limit=5000&skip=0&wbhk=1", 
            "/v1/transactions?id_credential=5e793889ea4c5165d46f9811&limit=5000&skip=5000&wbhk=1", 
            "/v1/transactions?id_credential=5e793889ea4c5165d46f9811&limit=5000&skip=10000&wbhk=1"
        ]
    }, 
    "event": "refresh", 
    "id_credential": "5e793889ea4c5165d46f9811", 
    "id_external": "IXS7607092R5", 
    "id_job": "5e793889b783081dd61de29c", 
    "id_job_uuid": "5e793889b783081dd61de29b", 
    "id_site": "5da784f1f9de2a06483abec1", 
    "id_site_organization": "56cf4ff5784806152c8b4567", 
    "id_site_organization_type": "56cf4f5b784806cf028b4569", 
    "id_user": "5e751d54d0510472295413b3"
}
```



## Methods

------
<br />

`GET`

Return webhooks register on API KEY

```bash
curl "https://sync.paybook.com/v1/webhooks" \
-H "Authorization: api_key api_key=YOUR_API_KEY" \
-X GET
```

------
<br />

`POST`

Creates a new Webhook on API KEY

```bash
curl "https://sync.paybook.com/v1/webhooks" \
-H "Authorization: api_key api_key=YOUR_API_KEY" \
-H "Content-Type: application/json" \
-d '{"url":"https://webhook_domain/my_webhook","events":["credential_create","credential_update","refresh"]}' \
-X POST
```

------
<br />

`DELETE`

Delete a Webhook

```bash
curl "https://sync.paybook.com/v1/webhooks/ID_WEBHOOK" \
-H "Authorization: api_key api_key=YOUR_API_KEY" \
-X DELETE
```





