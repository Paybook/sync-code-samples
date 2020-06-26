## Events

------
<br />

#### New credential was created 

 `credential_create`

Message:

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

#### Existing credential was updated

`credential_update`

Message:

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

#### Data for an existing credential was added or updated

`refresh` 

For a Existing credentials, the follow actions can be detected:

1. New account was added
2. New transactions wass added
3. An existing accounts was updated
4. An existing transactions was updated

Message:

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

