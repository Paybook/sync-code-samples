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
					"name": "payments.transaction_created",
					"at": "2020-01-22 12:55:23.929393+0000",
					"version": "3.1"
				},
				"user": {
					"id_user": "5f253440bf86772627785800",
					"id_external": "cus_2mo28NTyk69BKXXX",
					"id_environment": "574894bf7848066d138b4571"
				}
			},
			"payload": [
				{
					"id_pymnt_charge": "60a5d4990eea24244d4efa00",
					"id_pymnt_affiliate": null,
					"id_pymnt_customer": "5f251e96bf8677262778000e",
					"id_currency": "523a25953b8e77910e8b456c",
					"id_crypto_currency": "",
					"id_profile": "5f251e95bf8677262778440d",
					"transaction": "MjU3OTg0XXX=",
					"address": "",
					"crypto": null,
					"timestamp": "2020-01-19T22:16:45-05:00",
					"dt_charge": 1621480601,
					"period": 202105,
					"week": 20210520,
					"day": 20210520,
					"payment_method": "POS",
					"authorization_code": "205401",
					"status": "N",
					"transaction_type": "R",
					"receipt_url": "https://connect.srpago.com/recipe/",
					"checkout_url": "",
					"has_devolution": "0",
					"affiliation": "8418357",
					"affiliation_name": "",
					"reference": { 
                        "number": "XXXXXXX", 
                        "description": "Description" 
                    },
					"card": {
						"id_pymnt_card": "5f262f7b2e873779ec49f000",
						"holder_name": "John Doe",
						"type": "VISA",
						"number": "436401XXXXXX0000"
					},
					"amount": 257.25,
					"amount_crypto": "0",
					"net_amount": 251.63999999999999,
					"net_amount_crypto": "0",
					"tip": 0,
					"commission": 0,
					"commission_crypto": "0",
					"fee": 5.6100000000000003,
					"fee_crypto": "0",
					"fee_details": [
						{
							"type": "fee",
							"percentage": 1.8799999999999999,
							"amount": 5.6101080000000003,
							"description": ""
						}
					],
					"origin": { 
                        "location": { 
                            "latitude": 32.896422999999999, "longitude": -96.966841000000002 
                            } 
                        },
					"dt_create": 1621480601,
					"dt_modify": 1621480601
				}
			]
		}
    ]
}
```