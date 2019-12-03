# Widget Documentation



## Overview

<figure class="image">
  <img src="https://drive.google.com/uc?export=view&id=1Ll-fQQodIEnlx9ys0U4hn67y8w_EjNlX"/>
</figure>

The *Sync Widget* can be used to **create**, **update** and **trigger synchronization** of credentials.

The *Sync widget* will take the 100% of the screen when opened in *Mobile* devices, or devices with a screen size lesser than **600px**. When opened in bigger devices e.g. Desktop, the widget will be a modal.

The *Sync widget* also has an status *Toast* that displays the status of the current action being executed, this is, a credential being **created**, **updated** or **sychronized**.

When an action is being executed, the user is free to close the modal. When the modal is closed, a small Toast will appear in order to notify the user of the current process' status, and to prompot the user for more information if necessary e.g. two-fa.

The developer can easily integrate the *Sync Widget* by using:

- Configuration params
- Widget instance methods
- Handle the different events emitted

The *Sync Widget* can be used in 3 different use cases:

#### CreationCase

Opens the Sync Widget modal in order to add a new credential.

When to use this case?

- Use this case when creating a new credential

#### UpdateCase

Opens the *Sync Widget* modal directly for the specified site. Since the *Sync API* uniquely identifies each credential, when an end-user introduces some credential values that already exist, the *Sync API* will automatically detect this is an existing credential, and update the exisitng credential with the new provided data.


When to use this case?

- If the end-user changed his password for a given site, use this mode to prompt the user to update his site password in the Sync API.

#### SyncCase

Opens the Sync Widget status Toast for the specified credential. This action triggers a job in the Sync API to synchronize the credential. If the credential does not require two-fa, then no more information is required. However, if the credential requires two-fa, the status Toast will eventually ask the user for this value.

When to use this case?

- To keep two-fa credentials up to date.

## Docs

### Params

To instantiate the *Sync Widget* specify these parameters:

| **Parameter**    |      **Type**       | **Default Value** | **Description**                                              |
| ---------------- | :-----------------: | :---------------- | :----------------------------------------------------------- |
| `token`          |     **String**      | *N/A*             | A valid Sync API **token** whether for Sandbox or Production environments |
| `element`        | **String, Element** | `"#widget"`       | A valid *HTML DOM* element selector in which the widget will be rendered OR an **Element** instance from the DOM |
| `config`         |  **WidgetConfig**   | *See table below* | A valid **WidgetConfig** object                              |
| `enableTestMode` |     **Boolean**     | `false`           | Enable widget to access test mode                            |

The **WidgetConfig** attributes are described in this table:

| **Attribute**                          |     **Type**      | **Available or Sample Values**            | **Default Value**                                            | **Description**                                              |
| -------------------------------------- | :---------------: | :---------------------------------------- | :----------------------------------------------------------- | :----------------------------------------------------------- |
| `locale`                               |    **String**     | `"en"`, `"es"`                            | `"es"`                                                       | Specifies the widget language                                |
| `entrypoint`                           |    **Object**     | An object                                 | `{ idCountry: 'MX' }`                                        | Specifies the entrypoint of the widget. If not given, the widget is to be opened in **CreationCase** |
| `entrypoint.country`¹                  |    **String**     | A valid country Id or ISO country code    | `"MX"`                                                       | Opens the widget in **CreationCase** for the given country   |
| `entrypoint.credential`¹ ²             |    **String**     | A valid credential ID                     | `null`                                                       | Opens the widget in **SyncCase** for the given credential    |
| `entrypoint.site`¹ ²                   |    **String**     | A valid site ID                           | `null`                                                       | Opens widget in **UpdateCase** for the given site            |
| `entrypoint.siteOrganizationType`¹     |    **String**     | A valid site organization type ID or name | `"Bank"`                                                     | Opens widget in **CreationCase** and selects the tab related to the given site organization type |
| `navigation`                           |    **Object**     | An object                                 | `{}`                                                         | Specifies navigation settings                                |
| `navigation.displayBusinessSites`      |    **Boolean**    | `true`, `false`                           | `true`                                                       | If true, display all Business sites. Hide all business sites otherwise |
| `navigation.displayLogoImages`         |    **Boolean**    | `true`, `false`                           | `true`                                                       | If true, display all logos (images) for sites, site organizations, and the logo image in the notification status toast. If false, logos will not be displayed anywhere. |
| `navigation.displayLongDescription`    |    **Boolean**    | `true`, `false`                           | `true`                                                       | If true, when a final status is reached, the status toast will automatically expand and display the long description. |
| `navigation.displayPersonalSites`      |    **Boolean**    | `true`, `false`                           | `true`                                                       | If true, display all Personal sites. Hide all personal sites otherwise |
| `navigation.displayStatusInToast`      |    **Boolean**    | `true`, `false`                           | `false`                                                      | If true, close the modal when the web socket is initiated and thus the process will continue in the toast. If false the synchronization process can be continued in the modal, or in the toast if the user manually closes the modal |
| `navigation.enableBackNavigation`      |    **Boolean**    | `true`, `false`                           | `true`                                                       | If true, display the *Back* buttons in both, the *Side Menu* and in the *Credential Input Form*. Otherwise these back buttons will be hidden thus preventing the user from nagivating back once a site is selected for synchronization |
| `navigation.hideSites`                 | **Array<String>** | `['CIEC']`                                | `[]`                                                         | Hide any given site. Array items can be IDs or names.        |
| `navigation.hideSiteOrganizations`     | **Array<String>** | `['SAT']`                                 | `[]`                                                         | Hide any given site organization. Array items can be IDs or names. |
| `navigation.hideSiteOrganizationTypes` | **Array<String>** | `['Blockchain']`                          | `[]`                                                         | Hide any given site organzation type. Array items can be IDs or names. |
| `navigation.toastDuration`             |    **String**     | 5000                                      | 3000                                                         | The duration in ms that the status toast is to be kept opened when the final status is successful. If the final status is error, the toast needs to be manually closed. |
| `navigation.toastPosition`             |    **String**     | `top-right`                               | `top-left`, `top-center`, `top-right`, `bottom-left`, `bottom-center`, `bottom-right`. | Display the status toast in the given position.              |


¹ If a given entrypoint attribute value is not found, the widget will raise an exception and will not display any data. For instance, if you provide an idSite that is not found in the Sync API catalogues, or it exists but not allowed for the given configuration e.g. given a Business idSite when the widget is configured to display only Personal sites.

² For the entrypoint attributes, `credential` takes precedence over `site` if both attributes are given.

To create a widget instance you need to call the *Sync Widget* constructor and pass along the desired params:

```
...

const params = {
    token,
    element: #widget,
    config: {
      ...
    }
};

syncWidget = new SyncWidget(params);
```

### Methods

The *Sync Widget* **instance** methods are described in this table:

| **Method**                                                |                      **Sample**                      | **Description**                                              |
| --------------------------------------------------------- | :--------------------------------------------------: | :----------------------------------------------------------- |
| `close(void)`                                             |                 `syncWidget.close()`                 | Close the *Sync Widget* modal                                |
| `open(void)`                                              |                 `syncWidget.open()`                  | Opens the *Sync Widget* modal                                |
| `$on(String eventName)`                                   |         `syncWidget.$on("opened", () => {})`         | Subscribe to a *Sync Widget* event                           |
| `setConfig(Object<WidgetConfig> widgetConfig)`            |            `syncWidget.setConfig({...})`             | Set a new *Sync Widget* configuration. This config will totally replace the current's widget instance configuration |
| `setEntrypointCredential(String<ObjectId> idCredential)`¹ | `syncWidget.setEntrypointSite('some-credential-id')` | Opens the *Sync Widget* instance in **SyncCase** for the given credential |
| `setEntrypointSite(String<ObjectId> idSite)`              |    `syncWidget.setEntrypointSite('some-site-id')`    | Sets ups the *Sync Widget* instance in **UpdateCase** for the given site |
| `setToken(String token)`                                  |     `syncWidget.setToken('some-sync-api-token')`     | Sets up the *Sync Widget* instance authentication token. This token is used to authenticate against Sync API |
| `upsertConfig(Object<WidgetConfig> widgetConfig)`         |           `syncWidget.upsertConfig({...})`           | Update *Sync Widget* config. This config will be merged with the current's widget instance configuration |

¹ `setEntrypointCredential(...)` method will trigger the synchronization process. Since this action starts the widget for the **SyncCase** and this case does not uses the modal, there is no need to execute `open()` after setting the credential entrypoint.

### Events

These are the events emitted by the *Sync Widget*:

| **Event Name**                     |     **Value**      | **Description**                                              |
| ---------------------------------- | :----------------: | :----------------------------------------------------------- |
| `"closed"`                         |       `void`       | Triggered when the *Sync Widget* is closed                   |
| `"error"`                          |    `Credential`    | Triggered if a credential had an error while being synchronized. A final error status is obtained from *Sync API* |
| `"opened"`                         |       `void`       | Triggered when the *Sync Widget* is opened                   |
| `"status"`                         | `CredentialStatus` | Triggered after each websocket status change reported to the *Sync Widget* from *Sync API* |
| `"success"`                        |    `Credential`    | Triggered if a credential is successfully synchronized. A final success status is obtained from *Sync API* |
| `"updated"`                        |    `Credential`    | Triggered after credential values are sent to the *Sync API* using POST/PUT endpoints |

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

## Config Sample

This is how the default's widget config looks like:

```
{
    locale: 'es', 
    entrypoint: {
        country: 'MX'
    },
    navigation: {
        displayBusinessSites: true,
        displayLogoImages: true,
        displayLongDescription: false,
        displayPersonalSites: true,
        displayStatusInToast: false,
        enableBackNavigation: true,
        hideSiteOrganizations: [],
        hideSites: [],
        hideSiteOrganizationTypes: [],
        toastDuration: 5000,
        toastPosition: "top-right"
    }
}
```

## Some real cases explained (as of Paybook app)

### Case 1:

Open the widget in CreationCase, display only Personal sites, hide Renapo site, and hide Blockchain and Digital Wallets organization types

```
{
    locale: 'es', 
    navigation: {
        displayBusinessSites: false,
        hideSites: ['Renapo'],
        hideSiteOrganizationTypes: ['Blockchain', 'Digital Wallet']
    }
}
```

### Case 2:

Open the widget in CreationCase, display Personal and Business sites and hide Renapo site

```
{
    locale: 'es', 
    navigation: {
        hideSites: ['Renapo']
    }
}
```

### Case 3:

Open the widget in UpdateCase and exclude navigation, because I don't want for the user to navigate to other widget screens and thus being able to sync a new site

```
{
    locale: 'es', 
    entrypoint: {
        site: '121321'
    },
    navigation: {
        displayBusinessSites: false,
        hideSites: ['Renapo'],
        hideSiteOrganizationTypes: ['Blockchain', 'Digital Wallet'],
        enableBackNavigation: false
    }
}
```

### Case 4:

Open the widget in SyncCase:

```
{
    locale: 'es', 
    entrypoint: {
        credential: '121321'
    }
}
```



### Example:

Simple example with default settings:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1.0">
    <link rel="stylesheet" href="https://www.paybook.com/sync/widget/v2/widget.css">
    <title>Sync Widget</title>
  </head>
  <body>

    <div id="widget"></div>
    
    <script type="text/javascript" src="https://www.paybook.com/sync/widget/v2/widget.js"></script>
    <script>
    
        var widgetSettings = {
            token : '{{SYNC_TOKEN}}',
            config: {
                locale : 'es',
            }
        };
        var syncWidget = new SyncWidget(widgetSettings);
        syncWidget.open();

    </script>
  </body>
</html>
```

