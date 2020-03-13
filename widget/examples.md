## Examples

A valid widget set up alongside the **WidgetConfig** sample is below. This is how the default's widget config looks like:

```json
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

Below are some exaples of how you can set up and configure the widget for specific cases:

### Case 1:

What does this case do?
- Set up the widget in Spanish language
- Set up the widget to create new credentials
- Only display Personal sites
- Hide the site "Renapo"
- Hide the "Blockchain" and "Digital Wallets" organization types. 

The config that achieves this is below:

```json
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

What does this case do?
- Set up the widget in Spanish language
- Set up the widget to create new credentials
- Display both Personal and Business sites
- Hide the site "Renapo"

The config that achieves this is below:

```json
{
    locale: 'es', 
    navigation: {
        hideSites: ['Renapo']
    }
}
```

### Case 3:

What does this case do?
- Set up the widget in Spanish language
- Set up the widget to be opened in a specific site
- Exclude navigation so the user cannot move from that site, since my application requirement is to not let the end-user to navigate to other widget screens and be forced to stay in this site. Recall, that if the user happens to introduce the username/password of an existing credential, and those new values are valid, then the existing credential will be updated with those new values.

The config that achieves this is below:

```json
{
    locale: 'es', 
    entrypoint: {
        site: '...some valid id_site...'
    },
    navigation: {
        enableBackNavigation: false
    }
}
```

### Case 4:

What does this case do?
- Set up the widget in Spanish language
- Open the widget directly for the given id_credential. The modal will not be opened, the Status Toast will be opened right away.

The config that achieves this is below:

```json
{
    locale: 'es', 
    entrypoint: {
        credential: '...some valid id_credential...'
    }
}
```
