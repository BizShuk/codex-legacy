nodejs_note.md

`curl https://raw.githubusercontent.com/creationix/nvm/v0.23.3/install.sh | bash`

```bash
?? nvm run unstable --version

# show version list 
nvm ls-remote # for online
nvm ls # for local

# install nodejs with version
# ex: nvm install v0.10.24
nvm install version
nvm uninstall stable


# use nodejs version
nvm use stable
nvm use v0.10.24
nvm use 0.10

```


### testing

##### # Mocha
[Mocha](https://mochajs.org/)


##### # npm test


npm contain testing part

# for test
`npm test`


##### # nodejs install dependancy
package.json
```json
    {
        "name": "youtube-iframe",
         "version": "0.0.1",
        "private": true,
        "devDependencies": {
            "grover": "*",
            "express": "2.4.7",
            "jade": ">= 0.0.1",
            "oauth": ">= 0.9.5"
        },
        "scripts": {
            "test": "grover ./tests/unit/index.html --server"
        }
    }
```
