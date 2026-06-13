# Git note



### Install git in other path
`GIT.exe /DIR=“/path”`


### Credential Manager
`credential.helper = manager-core`

### URL Rephrase
`url."git@github.com:".insteadOf = "https://github.com/"`

### Cancel password prompt
`core.askPass = `


### store credential
`git config credential.helper store`. Once enter username/password for particular domain. It'll be stored in `~/.git-credential`

Format:
`https://<username>:password>@domain.com`


### push
push.default=[nothin,current,upstream,matching]