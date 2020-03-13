# Examples

## Configuration Sample

This is how the default's widget config looks like:

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

## Some Real Cases

### Case 1:

Open the widget in CreationCase, display only Personal sites, hide Renapo site, and hide Blockchain and Digital Wallets organization types

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

Open the widget in CreationCase, display Personal and Business sites and hide Renapo site

```json
{
    locale: 'es', 
    navigation: {
        hideSites: ['Renapo']
    }
}
```

### Case 3:

Open the widget in UpdateCase and exclude navigation, because I don't want for the user to navigate to other widget screens and thus being able to sync a new site

```json
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

```json
{
    locale: 'es', 
    entrypoint: {
        credential: '121321'
    }
}
```

### Case 5:

This use-case entails how to open the widget for Acme sites. Since Paybook API has some limitations that are complex to fix for the time being, we will internally use the __internal__ attributes to be able to use Acme sites (and test with Acme sites in dev and other non-production environments). We can use the widget __internal__ attributes like this and get access to the Acme sites:

```json
{
    __internal__.useApiDev: true/false
    __internal__.defaultHttpParams: {
        is_test: true
    }
}
```

However, external Sync customers are not allowed to use __internal__ config attributes. If they want to use Acme, then they need to request a token using the Sandbox API key.