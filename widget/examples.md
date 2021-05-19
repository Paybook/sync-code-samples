## Examples

---

<br />

Just as an example of how the Syncfy Widget set up might look like, here is the DEFAULT's one. This is the configuration that the widget will use in case you do not pass any config input apart of the required one:

```json
{
  // By DEFAULT the widget will use Spanish language
  "locale": "es",
  // By DEFAULT the widget will display Mexican sites, etc
  "entrypoint": {
    "country": "MX"
  },
  "navigation": {
    // By DEFAULT the widget WILL display business sites
    "displayBusinessSites": true,
    // By DEFAULT the widget WILL display logos
    "displayLogoImages": true,
    // By DEFAULT the widget WILL NOT display the entire description of errors (if they happen)
    "displayLongDescription": false,
    // By DEFAULT the widget WILL display personal sites
    "displayPersonalSites": true,
    // By DEFAULT the widget executes the sichronization in the modal, only
    // if the user manually closes the modal, then the Status Toast will be displayed
    "displayStatusInToast": false,
    // By DEFAULT the widget WILL allow back navigation
    "enableBackNavigation": true,
    // By DEFAULT the widget WILL NOT hide any sites, site organizations, etc:
    "hideSiteOrganizations": [],
    "hideSites": [],
    "hideSiteOrganizationTypes": [],
    // By DEFAULT the toast duration when success WILL be 5 seconds (5000 ms)
    "toastDuration": 5000,
    // By DEFAULT the toast location WILL BE be in the top-right of the screen
    "toastPosition": "top-right"
  }
}
```

Below are some exaples of how you can set up and configure the widget:

---

<br />

##### Example 1:

This is an example of how to create new credentials and hide some information:

```json
{
  // Set up the widget in Spanish language
  "locale": "es",
  // Since there is no entrypoint configured, this will set up the widget
  // to create new credentials
  "navigation": {
    // Only display Personal sites
    "displayBusinessSites": false,
    // Hide the site 'Renapo'
    "hideSites": ["Renapo"],
    // Hide the 'Blockchain' and 'Digital Wallets' organization types.
    "hideSiteOrganizationTypes": ["Blockchain", "Digital Wallet"]
  }
}
```

---

<br />

##### Example 2:

This is an example of how to open the widget in a given site and restrict some navigation:

```json
{
  // Set up the widget in Spanish language
  "locale": "es",
  // Set the entrypoint to open the widget directly in BBVA Bancomer Personal
  "entrypoint": {
    "site": "56cf5728784806f72b8b456b" // This BBVA Bancomer Personal id_site
  },
  // Do not allow the user to move back. The user will always stay in Bancomer Personal
  "navigation": {
    "enableBackNavigation": false
  }
}
```

In this case, we restrict the navigation so the user cannot move from a given site, since my application requirement is to not let the end-user to navigate to other widget screens. Recall, that if the user happens to introduce the username/password of an existing credential, and those new values are valid, then the existing credential will be updated with those new values.

---

<br />

##### Example 3:

This is an example of how to re-synchronize a credential right away:

```json
{
  // Set up the widget in Spanish language
  "locale": "es",
  // Set up the widget entrypoint to an existing credential. In this case, the modal
  // will not be opened, but the Status Toast will be opened right away.
  "entrypoint": {
    "credential": "afaf572894806f72b8b4123"
  }
}
```
