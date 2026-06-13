window.onload = function() {

    var msg = [];
    var so = io.connect('http://localhost:3000');
    var field = document.getElementById("field");
    var sendbt = document.getElementById("send");
    var content = document.getElementById("content");

    so.on('message', function(data) {
        if (data.message) {
            msg.push(data.message);
            var html = '';
            for (var i = 0, len = msg.length; i < len; i++) {
                html += msg[i] + '<br />';
            }
            content.innerHTML = html;
        } else {
            console.log("There is a problem:", data);
        }

        sendbt.onclick = function() {
            var text = field.value;
            so.emit('send', {
                message: text
            });
        };
    });
};
