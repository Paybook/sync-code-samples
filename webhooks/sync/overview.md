# Webhooks

---

<br/>

A webhook is a HTTP callback to a specified URL. They are triggered each time data is updated in Syncfy to help you stay up to date with the latest changes.

Syncfy API requires that you setting up a Webhook to your API KEY in order to send these events notification:

1.  `credential_create` : New credential was created
2.  `credential_update` : Existing credential was updated
3.  `refresh` : Data for an existing credential was added or updated

---

<br/>

#### Syncfy Webhook Object

```json
{
  "id_webhook": "5d9bab5b8c91e73b2e4c75d3",
  "id_user": null,
  "is_disabled": 1,
  "events": ["credential_create", "credential_update", "refresh"],
  "url": "https://webhook_domain/my_webhook",
  "delay": 0,
  "ct_failed": 1001,
  "dt_created": "2019-10-07T21:17:15+00:00",
  "dt_modified": "2020-03-20T05:37:43+00:00"
}
```

---

<br/>

#### Syncfy Webhook Notification

```json
{
  "endpoints": {
    "credential": ["/v1/credentials/5e793889ea4c5165d46f9811"]
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
