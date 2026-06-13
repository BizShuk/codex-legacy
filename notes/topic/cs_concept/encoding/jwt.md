# JWT, Json Web Token

JWT is consist of 3 parts, header, payload, and signature. All concated by `.`.

[header]: about type of token, algo using,
[payload]: registed claim(metadata about this message), public claims(namespace), private claims(data)
[Signature]: Encryption_Algo(base64(header).base64(payload).secret)

Use public key in server-side to validate the signature.
If the signature valided, it means the content is authentic(can be trust).
Server-side starts to use content to process business logic
