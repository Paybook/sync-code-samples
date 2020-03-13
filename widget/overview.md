### Overview

---

<br />

<figure class="image">
  <img src="https://drive.google.com/uc?export=view&id=1Ll-fQQodIEnlx9ys0U4hn67y8w_EjNlX"/>
</figure>

The _Sync Widget_ is the easiest way to implement the Sync API. By using this widget you can very quickly integrate your web application with the Sync API and being able to **create** and **update** credentials right away.

---

<br />

##### Where can I use it?

First of all, since the widget is coded in vanilla javascript, you can use it in any web application no matter what framework you are (React, Vue, Angular, JQuery, etc).

Basically, to use the widget you will need to follow this checklist:

1. Sign up for your Sync API account and thus get your API keys.
2. Second, you have to create a session token. For this, review the Sync API docs on how to create a Session and obtain a token (for this, you will need your API keys)
3. Then, in your frontend application import the Sync Widget files (.js and .css) or install the widget using NPM.
4. Finally, configure the widget as desired, instantiate it using the token you already have.
5. Start adding or updating credentials and that's it :)

---

<br />

##### What can I do with the Sync Widget?

You can use the _Sync Widget_ for the following use cases:

- **To create a credential:** use the Sync Widget to create a new credential, this is, to connect to a bank or site by first time.
- **To update an existing credential:** use the Sync Widget to update an existing credential, this is, to update the username, password, etc. of a credential you added in the past.
- **To re-sync a credentials:** put a credential up to date, by this we mean to force the synchronization process to be executed whenever your end-users desire, or to run the sinchronization every now and then for those credentials which require token (two factor authentication).

---

<br />

##### How can I implement it?

Use this documentation to know how to implement the widget. Here you will find:

- Widget Configuration: How to configure the widget
- Widget Methods: methods you can use to set up and manipulate the widget
- Widget Events: how to catch events and handle communication between the widget and your application
- Widget Use Cases: How to set up the widget for each of the use cases explained above
- Widget Examples: See some examples
