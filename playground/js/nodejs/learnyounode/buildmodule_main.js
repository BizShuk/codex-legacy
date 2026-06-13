var rf = require('buildmodule');

function mycb(err, data) {
    if (err) {
        return
    }
    console.log(data);
}


fs.readdir(q)


rf(process.argv[2], process.argv[3], mycb);
