# Webhooks

---

<br/>

A webhook is a HTTP callback to a specified URL. They are triggered each time data is updated in Sr. Pago to help you stay up to date with the latest changes.

Sr. Pago API requires that you setting up a Webhook to your API KEY in order to send these events notification:

1.  `payments.chargeback_created` : New chargeback was created

---

<br/>

#### Sr. Pago Webhook Notification

```json
{
  "rid": "8df2a899-fb8e-4034-82ca-233ab3030d59",
  "events": [
    {
      "header": {
        "event": {
          "eid": "4217d889-a6da-477c-be25-17ed6646baa1",
          "name": "payments.chargeback_created",
          "version": "3.1"
        },
        "user": {
          "id_user": "5ef5130b553ab949f13cf0e1",
          "id_external": "john.doe@test.com",
          "id_environment": "574894bf7848066d138b4570"
        }
      },
      "payload": {
        //-- Custom payload per event
      }
    }
  ]
}
```

### Environment information

It is important to note, that each webhook payload contains environment information in order to quickly identify from which environment the webhooks are coming from,
you can find the environment at the node `header.user.id_environment`:

| ID Environment           | Environment |
| ------------------------ | ----------- |
| 574894bf7848066d138b4570 | Sandbox     |
| 574894bf7848066d138b4571 | Production  |
