# Public Key Encryption

It's based on public key and private key.

Ther are two appoarches to use it.

1. Encrypt content with public key. The content will be decrypted by who own the private key.
2. User private key to sign data. All others can use public key to know the authentication of data.

### Application

SSH: give client private key to encrypt

Digital Signature/Certificate: Use public key to sign data and use private key to validate the correctness of data

OAuth(JWT token): Use public key to sign JWT and use private key to validate the whole JWT.

TLS: public key encryption with symmetric encryption

Server Certificate: use private key to encrypt server information and use public key to verfiy
