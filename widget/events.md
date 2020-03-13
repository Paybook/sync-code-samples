# Events

These are the events emitted by the *Sync Widget*:

| **Event Name** | **Value** | **Description** |
|-|:-:|:-|
| `"closed"` | `void` | Triggered when the *Sync Widget* is closed |
| `"error"` | `Credential` | Triggered if a credential had an error while being synchronized. A final error status is obtained from *Sync API* |
| `"__internal__.plaidSitesClicked"` | `void` | Triggered when Find Your Institution button is clicked. Only available If __internal__.plaidSites is set to true |
| `"opened"` | `void` | Triggered when the *Sync Widget* is opened |
| `"status"` | `CredentialStatus` | Triggered after each websocket status change reported to the *Sync Widget* from *Sync API* |
| `"success"` | `Credential` | Triggered if a credential is successfully synchronized. A final success status is obtained from *Sync API*|
| `"updated"` | `Credential` | Triggered after credential values are sent to the *Sync API* using POST/PUT endpoints |

The event values can belong to the following models:

```
CredentialStatus {
    code: Number;
}
```

```
Credential {
    id_credential: String;
    username: String;
    is_new: Number;
    ws: String;
    status: String;
    twofa: String;
}
```

To subscribe to these events you call the syncWidget.$on instance method:

```
syncWidget.$on(string EventName, Function callback)
```

For instance:

```
syncWidget.$on("opened", () => { console.log('Sync Widget opened');
```
