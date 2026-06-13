# Email Unsubscription

### Problem

Unsubscribe URL can be faked by hacker. To keep it authentic, the URL should be verified.

The OAuth2(JWT) comes to this place. check [JWT]

1. Put JWT in unsubscribe URL
2. Verify the JWT when server received
3. Process with JWT data in server-side
