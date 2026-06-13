var fs = require("fs");
var path = require("path");


function mycb(err, data) {

}

function readlist(dirname, fy, cb) {
    fs.readdir(dirname, cb);
}


module.exports = readlist;
