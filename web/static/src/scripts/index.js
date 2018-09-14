window.WordQuest = window.WordQuest || {};

window.WordQuest.Puzzle = (function() {
    var initialize = function() {
        refresh_();
    };

    var refresh_ = function() {
        var oReq = new XMLHttpRequest();
        oReq.addEventListener("load", puzzleCallback_);
        oReq.open("GET", "puzzle");
        oReq.send();
        subscribe_();
    };

    var puzzleCallback_ = function(_evt) {
        var data = JSON.parse(this.response);
        console.log(data);
    };

    var subscribe_ = function() {
        var socket = new WebSocket('ws://localhost:8082/updates');

        socket.addEventListener('open', function(_evt) {
            console.log("Established websocket connection!");
            socket.send('Hello server!');
        });

        socket.addEventListener('message', function(evt) {
            var data = JSON.parse(evt.data);

            console.log(data);
        });
    }

    return {
        initialize: initialize
    };
})();