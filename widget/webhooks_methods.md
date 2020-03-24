## Methods

------
<br />

#### Return Webhooks register on API KEY

`GET`

```bash
curl "https://sync.paybook.com/v1/webhooks" \
-H "Authorization: api_key api_key=YOUR_API_KEY" \
-X GET
```

------
<br />

#### Creates a new Webhook on API KEY

`POST`

```bash
curl "https://sync.paybook.com/v1/webhooks" \
-H "Authorization: api_key api_key=YOUR_API_KEY" \
-H "Content-Type: application/json" \
-d '{"url":"https://webhook_domain/my_webhook","events":["credential_create","credential_update","refresh"]}' \
-X POST
```

------
<br />

#### Delete a Webhook

`DELETE`

```bash
curl "https://sync.paybook.com/v1/webhooks/ID_WEBHOOK" \
-H "Authorization: api_key api_key=YOUR_API_KEY" \
-X DELETE
```


