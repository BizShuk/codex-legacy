// common module for client and server
var path = require('path');

var common = {
    domain: 'domain',
    Sdomain: 'special domain',
    js_path: path.resolve(__dirname,'public/js'),
    sass_path: path.resolve(__dirname,'public/css')
}

module.exports = common;
