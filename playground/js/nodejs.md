# Nodejs


# packages
- [mysql](#mysql)
- [path to regexp](https://www.npmjs.com/package/path-to-regexp) , express route 
- [jade syntax](https://naltatis.github.io/jade-syntax-docs/)


### mysql
`var mysql = require('mysql')`

##### mysql connection
```
var connection = mysql.createConnection({
  host     : 'localhost',
  pool     : <port>,
  user     : 'root',
  password : 'password'
});

connection.query('USE test_database');

connection.query('SELECT * FROM users', function(err, rows){
    res.render('users', {users : rows});
});
```

##### mysql pool
```
var pool  = mysql.createPool({
    connectionLimit : 10,
    host     : 'example.org',
    pool     : <port>,
    user     : 'bob',
    password : 'secret',
    database : 'my_db'
});

pool.query('SELECT 1 + 1 AS solution', function(err, rows, fields) {
  if (err) throw err;

  console.log('The solution is: ', rows[0].solution);
});
```