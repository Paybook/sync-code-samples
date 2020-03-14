### Events

---

<br />

Events are emitted by the widget under certain conditions in order to communicate the status of the synchronization of credentials to your application.

To handle this events, after creating your widget instance subscribe to those events using the syncWidget's `$on` method:

```javascript
...

syncWidget = new SyncWidget(params);

myCallbackForThisEvent = () => {
    /*
    * This is your registered callback that will be executed every
    * time this event is triggered by the Sync Widget
    */
    ....
}

syncWidget.$on('... some eventName ... ', myCallbackForThisEvent);
```

---

<br />

##### What values does each event emit?

Some events are `void` (meaning that they don't pass any data to the registered callback). However, if the event pass values to the registered callback, those values can be:

```json
// CredentialStatus value:
{
    "code": "Number",
}

// Credential value:
Credential {
    "id_credential": "String",
    "username": "String",
    "is_new": "Number",
    "ws": "String",
    "status": "String",
    "twofa": "String"
}
```

---

<br />

##### Event Types

The list of events available is described below:

`$on("opened", ... )`

```javascript
syncWidget.$on('opened', () => {
  // ... do something when the widget is opened ...
});
```

---

<br />

`$on("closed", ... )`

```javascript
syncWidget.$on('closed', () => {
  // ... do something when the widget is closed ...
});
```

---

<br />

`$on("updated", ... )`

```javascript
syncWidget.$on('updated', (status: CredentialStatus) => {
  /*
    * In particular, this event is triggered after the Sync 
    * API responds that it has received the credential values, 
    * *and the synchronization process is about to start
    * 
    ... do something when a new synchronization status is reached ...
    */
});
```

---

<br />

`$on("status", ... )`

```javascript
syncWidget.$on('status', (status: CredentialStatus) => {
  /*
    * In particular, this event is triggered while the synchronization status 
    * is running and after each change in the synchronization status this
    * event is triggered with the current's status data

    ... do something when a new synchronization status is reached ...
    */
});
```

For details on which are the statuses and more details on this consult the Sync API docs.

---

<br />

`$on("success", ... )`

```javascript
syncWidget.$on('success', (credential: Credential) => {
  // ... do something when the synchronization of a credential is finished successfully
});
```

---

<br />

`$on("error", ... )`

```javascript
syncWidget.$on('error', (credential: Credential) => {
  // ... do something when there is some error in the synchronization of credentials  ...
});
```

For details on which errors might happend and how to handle them, consult the Sync API docs.
