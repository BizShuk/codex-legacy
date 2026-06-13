# HTTPS

Based on HTTP, but add TLS session.

1. [Client Hello], Send client info about supported mechanism including tls version, encryption algo, random number
2. [Server Hello], Send server info including, certificate, encryption algo, server random
3. [Authentication], Verify server certificte with [Trusted Root Certification Authority] that issued it.
4. [The premaster secret], client send premaster secrt which will be encrypted by public key(from server certificate) and resolved by private key in the server
5. [Private Keys Created], server gets premaster secret
6. [Session Keys Created], Session key is symmetric encryption. client/server generate session keys from client random, server random, premaster secret. [How?]
7. [Client ready]
8. [Server ready]
9. [Handshke completed]

ephemeral Diffie-Hellman handshake???
