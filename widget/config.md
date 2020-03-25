### Set up and Configuration

---

<br />

##### How to set up or instantiate the widget?

The first step is to add an `HTML` element in any part of your app:

```html
<!-- 
  * Add an element like this in the part of your app 
  * where you want the widget content to be anchored
  *
  * The widget content itself will take this element and
  * replace it with its own html
  *
  * IMPORTANT: The id of this element HAS TO MATCH the "id" in the 
  * SyncWidget instantiation
 -->
<div id="widget"></div>
```

Then, you will call the `SyncWidget` constructor and pass to it the required data to instantiate the widget:

```javascript
/*
* 1. Set up and configure it:
*/
const params = {
    // [REQUIRED] A session token obtained from the Sync API
    token,
    // [REQUIRED] A DOM element identifier
    // IMPORTANT: This HAS TO MATCH the "id" of the HTML element
    element: "#widget",
    // [OPTIONAL] you can instantiate the widget in strict mode passing the JSON Web Keys
    // JWK(s) as provided in the session object obtained from the Sync API:
    strict: {
      authorization,
      body
    },
    // [OPTIONAL] you can provide a function that will be automatically used by the widget
    // to refresh it's token when this is expired. This function will must return a promise
    // with the new token value.
    refreshTokenFunction: () => {
      //... return some promise with a new token
    },
    // [REQUIRED] A valid WidgetConfig object
    config: {
      ...
    }
};

/*
* 2. Instantiate the widget using the SyncWidget constructor:
*/
syncWidget = new SyncWidget(params);

/*
* 3. If desired, open it right away:
*/
syncWidget.open();
```

<br/>

---

<br />

##### The SyncWidget Constructor

Here is a more detailed description of the data required to create a SyncWidget instance using the SyncWidget constructor:

| **Parameter**          |      **Type**       | **REQUIRED** | **Default Value**                                                 | **Description**                                                                                                                                                                                                    |
| ---------------------- | :-----------------: | :----------- | :---------------------------------------------------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `token`                |     **String**      | **REQUIRED** | _N/A_                                                             | A valid _Sync API_ **token** whether for Sandbox or Production environments                                                                                                                                        |
| `element`              | **String, Element** | **REQUIRED** | `"#widget"`                                                       | A valid _HTML DOM_ element selector in which the widget will be rendered OR an **Element** instance from the _DOM_                                                                                                 |
| `config`               |  **WidgetConfig**   | **REQUIRED** | _See the docs below_                                              | A valid **WidgetConfig** object                                                                                                                                                                                    |
| `strict`               |     **Object**      | **OPTIONAL** | _See API Docs on how to get the JWK(s)_                           | A valid set authorization and body JSON Web Keys                                                                                                                                                                   |
| `refreshTokenFunction` |    **Function**     | **OPTIONAL** | _See API docs to get information about the token expiration time_ | A valid _Javascript_ function that must return a promise with the payload `{ token: "...token...", strict: { authorization: "...", body: "..." } }`. "strict" is only required if widget is set up in Strict Mode. |
| `enableTestMode`       |      **Bool**       | **OPTIONAL** | `true`, `false`                                                   | If test mode wants to be enabled.                                                                                                                                                                                  |

<br/>

---

<br />

##### The WidgetConfig object

The **WidgetConfig** object is described below:

```json
{
  // It specifies the widget language. Default='en'. Allowed values: 'en' or 'es'.
  "locale": "string",
  // In general, the entrypoint value will specify where the widget will start when
  // instantiated. Default={}.
  "entrypoint": {
    // Starts the widget displaying data of the given country e.g. open the widget
    // displaying the sites of Argentina. Sample values: "MX", "AR", "CO", etc.
    // You need to use the country ISO code. Default="MX"
    "country": "string",
    // Starts the widget directly for an existing credential. This is used when forcing
    // the synchronization of an existing credential. It has to be a valid id_credential.
    // Default=null
    "credential": "string",
    // Starts the widget directly in a site. This is used when updating username,
    // password, etc of a credential of a given site e.g. open the widget directly
    // in Santander Empresas site. It has to be a valid id_site. Default=null
    "site": "string",
    // Starts the widget directly choosing the given organization type e.g. Open the
    // widget directly for the Government organization type. It has to be a valid
    // id_site_organization_type. Default=null
    "siteOrganizationType": "string"
  },
  // In general, the navigation value will specify which elements will be displayed and
  // thus limit or extend the navigation capabilities for the end-user:
  "navigation": {
    // If true, display all Business sites e.g. 'cuentas de sitios empresariales'.
    // Default=true
    "displayBusinessSites": "bool",
    // If true, display logos in the widget for the site, site organizations,
    // and the site logo in the status toast. If false, no logos will be displayed.
    // Default=true
    "displayLogoImages": "bool",
    // If true, when a final synchronization status is reached, the status toast
    // will automatically expand and display the long description explaining the
    // final result. This is useful when an error is obtained so that the end-user
    // knows exactly what happened. Default=false
    "displayLongDescription": "bool",
    // If true, display all Personal sites e.g. a personal site is one that is
    // not business. Default=true
    "displayPersonalSites": "bool",
    // If true, close the modal when the synchronization is initiated (websocket starts)
    // and thus the process will continue in a status toast. If false, the synchronization
    // process will carry on in the modal (if the user closes the modal,
    // then the toast will be displayed). Default=false
    "displayStatusInToast": "bool",
    // If true, display the *Back* buttons in both, the *Aside Menu* and in the
    //  *Credential Input Form* allowing the end-user to move back in the widget.
    // Otherwise these back buttons will be hidden thus preventing the user from
    // nagivating back once a site is selected for synchronization. Default=true
    "enableBackNavigation": "bool",
    // If true, will never display the *Aside Menu*. Otherwise *Aide Menu* will be
    // accesible and can be expanded and closed using a hamburger button. Default=false
    "hideAsideMenu": "bool",
    // If true, hide the *Select Country* input located in the *Aside Menu*. Default=false
    "hideSelectCountry": "",
    // If specified, the sites given will be hidden in the widget (never displayed).
    //  The array must contain id_site(s) or site's names (as returned by the API).
    // Default=[]. Sample=['CIEC']
    "hideSites": "array<string>",
    // If specified, the site organizations given will be hidden in the widget (never displayed).
    // The array must contain id_site_organization(s) or site organization's names (as returned
    // by the API). Default=[]. Sample=['SAT']
    "hideSiteOrganizations": "array<string>",
    // If specified, the site organization types given will be hidden in the widget (never displayed).
    // The array must contain id_site_organization_type(s) or site organization type's names
    // (as returned by the API). Default=[]. Sample=['Government']
    "hideSiteOrganizationTypes": "array<string>",
    // If specified, only the sites given will be displayed in the widget.
    // The array must contain id_site(s) or site's names (as returned by the API).
    //  Default=[]. Sample=['CIEC']
    "displaySites": "array<string>",
    // If specified, only the site organaizations given will be displayed in
    // the widget. The array must contain id_site_organization(s) or site
    // organization's names (as returned by the API). Default=[]. Sample=['SAT']
    "displaySiteOrganizations": "array<string>",
    // If specified, only the site organaization types given will be displayed in the
    // widget. The array must contain id_site_organization_type(s) or site organization
    // type's names (as returned by the API). Default=[]. Sample=['Government']
    "displaySiteOrganizationTypes": "array<string>",
    // It specifies the duration in miliseconds that the status toast is to be kept
    // opened when the final status is successful. If the final status is error,
    // the end-user will always need to close the toast manually. Default=5000
    "toastDuration": "number",
    // It specifies, the location in the screen where the status toast is to be displayed.
    // Default='top-right'. Alowed values: 'top-left', 'top-center', 'top-right',
    // 'bottom-left', 'bottom-center', 'bottom-right'
    "toastPosition": "string",
    // If true, success message will be send when the credentials complete login step 
    // in the site. Otherwise, the widget will wait to complete all download process. 
    // Default=false
    "quickAnswer": "bool",
  }
}
```

Take into account these clariffications:

- If any configuration attribut value is not allowable, the widget will raise an exception and will not be displayed.
- If a given entrypoint attribute value is not valid, the widget will raise an exception and will not display any data. For instance, if you provide a `site` identifier that is not found in the Sync API catalogues, or it exists but not allowed for the given configuration e.g. given a Business `site` when the widget is configured to display only Personal sites.
- For the entrypoint attributes, `credential` takes precedence over `site` if both attributes are given.
