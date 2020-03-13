# Config

To instantiate the *Sync Widget* specify these parameters:

| **Parameter** | **Type** | **Default Value** | **Description** |
|-|:-:|:-|:-|
| `token` | **String** | *N/A* | A valid Sync API **token** whether for Sandbox or Production environments |
| `element` | **String, Element** | `"#widget"` | A valid *HTML DOM* element selector in which the widget will be rendered OR an **Element** instance from the DOM |
| `config` | **WidgetConfig** | *See table below* | A valid **WidgetConfig** object |
| `enableTestMode` | **Boolean** | `false` | Enable widget to access test mode |

The **WidgetConfig** attributes are described in this table:

| **Attribute** | **Type** | **Available or Sample Values** | **Default Value** | **Description** |
|-|:-:|:-|:-|:-|
| `locale` | **String** | `"en"`, `"es"` | `"es"` | Specifies the widget language |
| `entrypoint` | **Object** | An object | `{ idCountry: 'MX' }` | Specifies the entrypoint of the widget. If not given, the widget is to be opened in **CreationCase** |
| `entrypoint.country`¹ | **String** | A valid country Id or ISO country code | `"MX"` | Opens the widget in **CreationCase** for the given country |
| `entrypoint.credential`¹ ² | **String** | A valid credential ID | `null` | Opens the widget in **SyncCase** for the given credential |
| `entrypoint.site`¹ ² | **String** | A valid site ID | `null` | Opens widget in **UpdateCase** for the given site |
| `entrypoint.siteOrganizationType`¹ | **String** | A valid site organization type ID | `null` | Opens widget in **CreationCase** and selects the tab related to the given site organization type |
| `navigation` | **Object** | An object | `{}` | Specifies navigation settings |
| `navigation.displayBusinessSites` | **Boolean** | `true`, `false` | `true` | If true, display all Business sites. Hide all business sites otherwise |
| `navigation.displayLogoImages` | **Boolean** | `true`, `false` | `true` | If true, display all logos (images) for sites, site organizations, and the logo image in the notification status toast. If false, logos will not be displayed anywhere. |
| `navigation.displayLongDescription` | **Boolean** | `true`, `false` | `true` | If true, when a final status is reached, the status toast will automatically expand and display the long description. |
| `navigation.displayPersonalSites` | **Boolean** | `true`, `false` | `true` | If true, display all Personal sites. Hide all personal sites otherwise  |
| `navigation.displayStatusInToast` | **Boolean** | `true`, `false` | `false` | If true, close the modal when the web socket is initiated and thus the process will continue in the toast. If false the synchronization process can be continued in the modal, or in the toast if the user manually closes the modal |
| `navigation.hideAsideMenu` | **Boolean** | `true`, `false` | `false` | If true, hide *Side Menu* always. Otherwise *Side Menu* can be toggled with hamburger button.|
| `navigation.hideSelectCountry` | **Boolean** | `true`, `false` | `false` | If true, hide the *Select Country* input in *Side Menu*. Otherwise it wil show. |
| `navigation.enableBackNavigation` | **Boolean** | `true`, `false` | `true` | If true, display the *Back* buttons in both, the *Side Menu* and in the *Credential Input Form*. Otherwise these back buttons will be hidden thus preventing the user from nagivating back once a site is selected for synchronization |
| `navigation.hideSites` | **Array<String>** | `['CIEC']` | `[]` | Hide any given site. Array items can be IDs or names. |
| `navigation.hideSiteOrganizations` | **Array<String>** | `['SAT']` | `[]` | Hide any given site organization. Array items can be IDs or names. |
| `navigation.hideSiteOrganizationTypes` | **Array<String>** | `['Blockchain']` | `[]` | Hide any given site organzation type. Array items can be IDs or names. |
| `navigation.displaySites` | **Array<String>** | `['CIEC']` | `[]` | Display only given sites in this array. Array items can be IDs or names. This option gets priority over `navigation.hideSites`|
| `navigation.displaySiteOrganizations` | **Array<String>** | `['SAT']` | `[]` |  Display only given site organizations in this array. Array items can be IDs or names. This option gets priority over `navigation.hideSiteOrganizations`|
| `navigation.displaySiteOrganizationTypes` | **Array<String>** | `['Blockchain']` | `[]` |  Display only given site organzation types in this array. Array items can be IDs or names. This option gets priority over `navigation.displaySiteOrganizationTypes`|
| `navigation.toastDuration` | **String** | 3000 | 5000 | The duration in ms that the status toast is to be kept opened when the final status is successful. If the final status is error, the toast needs to be manually closed. |
| `navigation.toastPosition` | **String** | `top-left`, `top-center`, `top-right`, `bottom-left`, `bottom-center`, `bottom-right`.  | `top-right` | Display the status toast in the given position. |
| `__internal__` | **Object** | An object | `null` | Specifies internal configuration: configuration that can only be used to test the widget internally in Paybook company. |
| `__internal__.defaultHttpParams` | **Object** | `{ is_test: true }` | `{}` | If specified, merge this object into all http params when calling Sync API. |
| `__internal__.plaidSites` | **Boolean** | `true`, `false` | `false` | If true, will add Canada and the United States to the countries options. When one of these options is clicked will display a Find Your Institituion button in the content area. When this button is clicked it will emits a event. **This is a hard-code required by Paybook App** |
| `__internal__.apiEnv` | **String** | `development`, `production` | `null` | If `development`, use *syncdev.paybook.com* API server, and uses *sdev.paybook.com* image server. Use *sync.paybook.com* and *s.paybook.com* with `production`. Default values depend on build environments.  |


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