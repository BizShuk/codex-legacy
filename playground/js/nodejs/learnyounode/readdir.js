var fs = require("fs");
var path = require("path");

dir_path = process.argv[2];
filter = process.argv[3];


fs.readdir(dir_path, (err, list) => {
    for (var i = 0, len = list.length; i < len; i++) {
        var ft = path.extname(list[i]);
        if (ft == "." + filter) {
            console.log(list[i]);
        }
    }
});
