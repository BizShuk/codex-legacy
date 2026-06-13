###  Progressive Web Apps

- HTTPS
- service worker
- -


idb
localstorage
localforage

[introduce the service worker](https://developers.google.com/web/fundamentals/primers/service-worker/)

Service worker functionality is only available on pages that are accessed via HTTPS (https://localhost and equivalents will also work, to facilitate testing). To learn about the rationale behind this restriction check out [Prefer Secure Origins For Powerful New Features](http://www.chromium.org/Home/chromium-security/prefer-secure-origins-for-powerful-new-features) from the Chromium team.


How do I avoid these edge cases?
So how do we avoid these edge cases? Use a library like sw-precache, which provides fine control over what gets expired, ensures requests go directly to the network and handles all of the hard work for you.


Some tips:

Once a service worker has been unregistered, it may remain listed until its containing browser window is closed.
Chrome occasionally displays a console error when trying to retrieve the service worker, this is safe to ignore.
If multiple windows to your app are open, the new service worker will not take effect until they've all been reloaded and updated to the latest service worker.
Unregistering a service worker does not clear the cache, so it may be possible you'll still get old data if the cache name hasn't changed.
If a service worker exists and a new service worker is registered, the new service worker won't take control until the page is reloaded, unless you take immediate control.

check for regitered service worker
chrome://serviceworker-internals/



Using the web app manifest, your web app can:

Have a rich presence on the user's Android home screen
Be launched in full-screen mode on Android with no URL bar
Control the screen orientation for optimal viewing
Define a "splash screen" launch experience and theme color for the site
Track whether you're launched from the home screen or URL bar


easy way to track how the app is launched is to add a query string to the start_url parameter and then use an analytics suite to track the query string. If you use this method, remember to update the list of files cached by the App Shell to ensure that the file with the query string is cached.



Extra credit: minify and inline CSS
There's one more thing that you should consider, minifying the key styles and inlining them directly into index.html. Page Speed Insights recommends serving the above the fold content in the first 15k bytes of the request.

See how small you can get the initial request with everything inlined.

Further Reading: PageSpeed Insight Rules


pagespeed
