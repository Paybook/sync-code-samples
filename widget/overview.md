### Overview

---

<br />

<figure class="image">
  <img src="https://drive.google.com/uc?export=view&id=1Ll-fQQodIEnlx9ys0U4hn67y8w_EjNlX"/>
</figure>

##### What is the Syncfy Widget?

The _Syncfy Widget_ is the easiest way to implement the Syncfy API. By using this widget you can very quickly integrate your web application with the Syncfy API and being able to **create** and **update** credentials right away.

---

<br />

##### How to install it?

There are two ways to install the Syncfy Widget

_[1] Using Node Package Manager:_

The first step is to install [@paybook/sync-widget](https://www.npmjs.com/package/@paybook/sync-widget) with NPM:

```bash
npm install --save @paybook/sync-widget
```

And then import the widget styles to your app:

```javascript
import "@paybook/sync-widget/dist/widget.css";
```

And finally get access to the `SyncWidget` class by importing it:

```javascript
import SyncWidget from "@paybook/sync-widget";
```

_[2] Using CDN:_

To install it via CDN import the following files in your _index.html_ file:

Add this line in the `<head>` element:

```html
<link rel="stylesheet" href="https://syncfy.com/widget/v2/widget.css" />
```

And add this line in the `<body>` element:

```html
<script
  type="text/javascript"
  src="https://syncfy.com/widget/v2/widget.js"
></script>
```

This will expose the `SyncWidget` class in the global window object

---

<br />

##### Where can I use it?

First of all, since the widget is coded in vanilla javascript, you can use it in any web application no matter what framework you are (React, Vue, Angular, JQuery, etc).

Basically, to use the widget you will need to follow this checklist:

1. Sign up for your Syncfy API account and thus get your API keys.
2. Second, you have to create a session token. For this, review the Syncfy API docs on how to create a Session and obtain a token (for this, you will need your API keys)
3. Then, in your frontend application import the Syncfy Widget files (.js and .css) or install the widget using NPM.
4. Finally, configure the widget as desired, instantiate it using the token you already have.
5. Start adding or updating credentials and that's it :)

---

<br />

##### What can I do with the Syncfy Widget?

You can use the _Syncfy Widget_ for the following use cases:

- **To create a credential:** use the Syncfy Widget to create a new credential, this is, to connect to a bank or site by first time.
- **To update an existing credential:** use the Syncfy Widget to update an existing credential, this is, to update the username, password, etc. of a credential you added in the past.
- **To re-sync a credentials:** put a credential up to date, by this we mean to force the synchronization process to be executed whenever your end-users desire, or to run the sinchronization every now and then for those credentials which require token (two factor authentication).

---

<br />

##### Where can I learn how to implement it and customize it?

There is plenty of documentation and resources that will help you implement the Syncfy Widget, namely:

- Widget Configuration: how to configure the widget
- Widget Methods: the methods you can use to set up and manipulate the widget
- Widget Events: how to catch events and handle communication between the widget and your application
- Widget Use Cases: how to set up the widget for different cases
- Widget Examples: see some examples
