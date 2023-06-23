# Webhooks V3

---

<br/>

A webhook is a HTTP callback to a specified URL. They are triggered each time data is updated in Syncfy to help you stay up to date with the latest changes.

Syncfy API requires that you setting up a Webhook to your API KEY in order to send these events notification:

1.  `credentials.created` : New credential was created.
2.  `credentials.updated` : Existing credential was updated.
3.  `credentials.refreshed` : Data for an existing credential was added or updated.
4.  `documents.completed` : Document synchronization is complete.
5.  `documents.fail` : One or more errors occurred while downloading the documents. There may be inconsistencies in the document.
6.  `documents.success` : The download of the documents has been completed successfully.

### documents.completed VS documents.success: 

documentes completed indicates the process have finished even there are errors, documents.success indicates the process have finished successfully.

---

#### Setup Webhooks

1. Signup at (Syncfy.com)[https://syncfy.com/w/en/sync/signup].
2. Go to the Webhooks section on your dashboard.
3. Add the Webhook endpoint for any environment (Sandbox and production) and configure it.
4. Add the events that your webhook will be listening.

<br/>

##### Recomendations

* Setup just one webhook globally for each environment.
* Use the JWT key to secure your webhooks notifications.
* You can set customized headers for the webhooks notifications.

---

<br/>

#### Syncfy Webhook Notification

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

---

<br/>

##### Webhooks V1 

Webhooks V1 are the first version of webhooks and are managed via API, the recommendation is to going ahead with the webhooks V3.

To review the webhooks V1 documentation, please visit the following [link](https://github.com/Paybook/code-samples/tree/master/webhooks/sync). 