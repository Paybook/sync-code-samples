# Overview

## What is the Sync Widget?

The *Sync Widget* is the easiest way to implement the Sync API. By using this widget you can very quickly integrate your web application to the Sync API and being able to **create** and **update** credentials right away.

## How can I use it?

First of all, since the widget is coded in vanilla javascript, you can use it no matter what framework you're using for your frontend application (React, Vue, Angular, etc).

In summary, to start using the widget you need to do the following steps:

1. Sign up for your Sync API account and get your API keys.
2. In your backend server, use your API keys and consume the Sync API to create a session token.
3. In your frontend application import the Sync Widget files (.js and .css) or install the widget using NPM.
4. Set up and configure the widget, get a token from your backend server (step 2), and that's it. 

## What can I do with the Sync Widget?

You can use the *Sync Widget* to tackle the following use cases:

- *To create a credential:* use the Sync Widget to create a new credential e.g. connect to a bank by first time
- *To update an existing credential:* use the Sync Widget to update an existing credential e.g. update the username, password, etc. of an already existing credential
- *To re-sync a credentials:* put a credential up to date e.g. force the synchronization process to be executed when your end-users want or make sure your the credentials that require token are kept up to date 

## How can I implement it?

Below is the complete documentation of how you can implement the widget. These documentation entails:

- How to configure it
- How to catch events
- How to set up the widget for each of the use cases explained above
- See some examples
