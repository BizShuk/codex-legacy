# git-secret

ref [link](https://git-secret.io/)

## prerequisite

- git
- gpg

## CMD

> git secret init

Init git secret setup and should check in git-secret folder except **.gitsecret/keys/random_seed**

> git secret add

Add files will be encrypted.

> git secret hide

Hide files which should be encrypted

> git secret reveal

Show unencrypted files

## Add new user

1. Get gpg public key
2. Import gpg public key by `gpg import <file>`
3. Embed email to user by `git secret tell <email>`
4. Newly added user can't encrypt file yet. So `git secret reveal; git secret hide -d` to reolve
