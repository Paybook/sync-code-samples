
# Methods

The *Sync Widget* **instance** methods are described in this table:

| **Method** | **Sample** | **Description** |
|-|:-:|:-|
| `close(void)` | `syncWidget.close()` | Close the *Sync Widget* modal |
| `open(void)` | `syncWidget.open()` | Opens the *Sync Widget* modal |
| `$on(String eventName)` | `syncWidget.$on("opened", () => {})` | Subscribe to a *Sync Widget* event |
| `setConfig(Object<WidgetConfig> widgetConfig)` | `syncWidget.setConfig({...})` | Set a new *Sync Widget* configuration. This config will totally replace the current's widget instance configuration  |
| `setEntrypointCredential(String<ObjectId> idCredential)`¹ | `syncWidget.setEntrypointSite('some-credential-id')` | Opens the *Sync Widget* instance in **SyncCase** for the given credential |
| `setEntrypointSite(String<ObjectId> idSite)` | `syncWidget.setEntrypointSite('some-site-id')` | Sets ups the *Sync Widget* instance in **UpdateCase** for the given site |
| `setToken(String token)` | `syncWidget.setToken('some-sync-api-token')` | Sets up the *Sync Widget* instance authentication token. This token is used to authenticate against Sync API |
| `upsertConfig(Object<WidgetConfig> widgetConfig)` | `syncWidget.upsertConfig({...})` | Update *Sync Widget* config. This config will be merged with the current's widget instance configuration  |

¹ `setEntrypointCredential(...)` method will trigger the synchronization process. Since this action starts the widget for the **SyncCase** and this case does not uses the modal, there is no need to execute `open()` after setting the credential entrypoint.