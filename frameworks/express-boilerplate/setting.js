/*  setting for server
 *      domain : official domain
 *      Sdomain : for specific purpose ,ex: slURL shorter domain
 *      port : server port
 *      db : used dbs in this project , for tools to output all tables for migration.
 *      mysql : database
 *      mongo : database
*/


var g = require('./common');

var server = {
    domain: g.domain,
    Sdomain: g.Sdomain,
    port: 7000,
    db: ['dbname'],
    mysql: {
        host : "localhost",
        port : 3306,
        user : "username",
        passwd : "passwd"
    },
    mongo: {
        host : "localhost",
        port : 27017,
        user : "username",
        passwd : "passwd"
    }
}


module.exports = server;
