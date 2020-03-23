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
    "url": "https://be5ce5a1.ngrok.io/my_webhook",
    "delay": 0,
    "ct_failed": 1001,
    "dt_created": "2019-10-07T21:17:15+00:00",
    "dt_modified": "2020-03-20T05:37:43+00:00"
}
```



#### Events

1. **credential_create**

   New credential was created

2. **credential_update**

   Existing credential was updated

3. **refresh** 

   For a Existing credentials, the follow actions can be detected:

   * New account was added
   * New transactions wass added
   * An existing accounts was updated
   * An existing transactions was updated