## Events

------
<br />

#### New chargeback was created 

 `payments.chargeback_created`

Message:

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
                "id_pymnt_charge": "5ef6267be9dfcf1e1919a414",
                "id_pymnt_affiliate": null,
                "transaction": "MjczODMzNw==",
                "payment_method": "POS",
                "authorization_code": "",
                "status": "",
                "transaction_type": "C",
                "receipt_url": "https://sandbox-connect.srpago.com/recipe/",
                "has_devolution": "0",
                "reference": {
                    "description": "CONTRACARGO POR: Autorizacion [136282] - Transaccion [2738336]"
                },
                "card": [],
                "linked": [
                    "5ef621b35beebd4808151162"
                ],
                "amount": -150,
                "net_amount": -150,
                "tip": 0,
                "commission": 0,
                "fee": 0,
                "fee_details": null,
                "origin": {
                    "location": {
                        "latitude": "0.000000",
                        "longitude": "0.000000"
                    }
                },
                "dt_create": 1593190011,
                "dt_modify": null
            }
        }
    ]
}
```

#### New transaction was created 

 `payments.transaction_created`

Message:

```json
{
    "rid": "8df2a899-fb8e-4034-82ca-233ab3030d00",
    "events": [
       {
			"header": {
				"event": {
					"eid": "fac3d61b-c425-47a0-b5bb-903c8f5511c2",
					"name": "payments.chargeback_created",
					"at": "2020-01-22 12:55:23.929393+0000",
					"version": "3.1"
				},
				"user": {
					"id_user": "5f253440bf86772627785847",
					"id_external": "cus_2mo28NTyk69BKXXX",
					"id_environment": "574894bf7848066d138b4571"
				}
			},
			"payload": {
				"id_pymnt_charge": "6082d1cb3a32082d2b649d54",
				"id_pymnt_affiliate": null,
				"transaction": "MjQ3NjkzOTk=",
				"payment_method": "POS",
				"authorization_code": "",
				"status": "",
				"transaction_type": "C",
				"receipt_url": "https://connect.srpago.com/recipe/ODU5NWQyZGItYTQzYi0xMWViLWE2NzAtMGFlMjFkZjUzMWYx",
				"has_devolution": "0",
				"reference": { "description": "CONTRACARGO POR: Autorizacion [005472] - Transaccion [23565695]" },
				"card": [],
				"linked": [ "605676a63a3d497512703116" ],
				"amount": -6688.5,
				"net_amount": -6688.5,
				"tip": 0,
				"commission": 0,
				"fee": 0,
				"fee_details": null,
				"origin": { "location": { "latitude": 0, "longitude": 0 } },
				"dt_create": 1619186123,
				"dt_modify": null
			}
		}
    ]
}
```