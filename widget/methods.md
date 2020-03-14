### Methods

---

<br />

The _Sync Widget_ **instance** has several methods that you can use to configure it but also to manipulate it

All of the methods described below with the exception of `setEntrypointCredential` will only modify the current set up of the widget but will not re-render it automatically. Therefore, after executing those methods and in order to see the changes reflected in the UI the `syncWidget.open` method has to be executed.

All the methods exposed by the SyncWidget are exposed below:

---

<br />

`open()`

```javascript
/*
 * Use this method to open the Sync Widget modal
 */
syncWidget.open();
```

---

<br />

`close()`

```javascript
/*
 * Use this method to close the Sync Widget modal
 */
syncWidget.close();
```

---

<br />

`setConfig(widgetConfig: WidgetConfig)`

```javascript
/*
* Use this method to fully replace the current widget
* configuration or set a new one
*/
syncWidget.setConfig({.
    // ... a valid WidgetConfig value goes here ...
});
```

Notice that the given config will totally replace the current's widget instance configuration. **This is a full replace, not a patch / upsert of the current configuration.**

---

<br />

`upsertConfig(widgetConfig: WidgetConfig)`

```javascript
/*
* Use this method to patch / upsert the current configuration.
*/
syncWidget.upsertConfig({.
    // ... a partial WidgetConfig value goes here ...
});
```

In contrast to `setConfig`, this config will be merged with the current's widget configuration.

---

<br />

`setToken(token: String)`

```javascript
/*
 * Use this method to replace or set a new authentication token.
 */
syncWidget.setToken('... a valid Sync API token ... ');
```

---

<br />

`setEntrypointCredential(id_credential: String)`

```javascript
/*
 * Use this method to opens the *Sync Widget* for the given credential
 * and thus re-synchronize the credential or put it up to date
 */
syncWidget.setEntrypointCredential('...some-credential-id...');
```

When called this method, the synchronization process will be triggered (the status toast will be opened right away). Also, since this action starts the widget for the **re-synchronization case** and this case does not uses the modal, there is no need to execute `open()` after setting the credential entrypoint.

---

<br />

`setEntrypointSite(id_site: String)`

```javascript
/*
 * Use this methdo to set up the *Sync Widget* entrypoint to the
 * given id_site. When the widget is opened it will starts in the
 * given id_site ready to get the username, password, etc.
 */
syncWidget.setEntrypointSite('...some id_site...');
```

When the widget is set up for a given id_site, and the user happens to introduce the username/password of a credential that already exist, the Sync API will automatically recognize this and update the existing credential to store the new username and password data if valid. This is why, the `entrypoint.site` can be used to edit an existing credential's username, password, etc. Although your not providing the id_credential, the Sync API will infer it from the username provided that this value is unique.

---

<br />

`$on(eventName: string)`

```javascript
/*
 * Use this method to listen to the Sync Widget events and do something
 */
syncWidget.$on('opened', () => {
  // ... do something when the widget is opened ...
});
```

The list of valid eventName(s) is here: "401", "error", "success", "status", "updated", "opened", and "closed". You can consult more about events in the **Events** section.
