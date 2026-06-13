# express
node http framework library , [official](http://expressjs.com/)

### nodemon
restart express when files changed , `nodemon ./bin/www`

### express-generator
generate simple express project
`express [--view=<view engine> --css=<css engine>] <project_name>` , view_engine default jade , css engine default css


### step by step
1. generator default project , by [generator](express-generator)
2. insert port from package.json 
```
    var packagejson = require("package.json");
    var port = normalizePort(process.env.PORT || '3000');
    =>
    var port = normalizePort(packagejson.server.port || '3000');
```
3. output pid to log dir
```
/**
 * output PID
 */
var fs = require('fs');

fs.open('log/server.pid','w+', (err,fd)=>{
    fs.write(fd,process.pid);                                                                                                            
}); 

```
4. generate https pem and change to https
5. if SEO is required , use jade template 
6. add static route
    `app.use('/static', express.static(__dirname + '/public'));`
7. add mysql connection
8. add mongo connection

### request
[request](https://expressjs.com/en/4x/api.html#req)

-



### response 
[response](https://expressjs.com/en/4x/api.html#res)
- send
- `render(<view_name>[,<param_object>,<call_back>])`
- `json(<object>)`
- `redirect(<status>,<url>)`


### router
[Express Route Tester](http://forbeslindesay.github.io/express-route-tester/)

### note
1. route path use regex by [path to regexp](https://github.com/pillarjs/path-to-regexp)