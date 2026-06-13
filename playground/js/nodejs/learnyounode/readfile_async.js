var fs = require("fs");

var buf = fs.readFile(process.argv[2], {
    encoding: 'utf8'
}, (err, buf) => {

    var str = buf.toString();
    var len = str.split('\n').length - 1;
    console.log(len);
});
