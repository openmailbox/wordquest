window.WordQuest = window.WordQuest || {};

WordQuest.initialize = (function() {
    WordQuest.currentPuzzle = null;

    var initialize = function() {
        var oReq = new XMLHttpRequest();

        oReq.addEventListener("load", _callback);
        oReq.open("GET", "puzzle");
        oReq.send();

        _subscribe();
    };

    var _callback = function(_evt) {
        var data = JSON.parse(this.response);

        WordQuest.currentPuzzle = new WordQuest.Puzzle();
        WordQuest.currentPuzzle.update(data);
    };

    var _subscribe = function() {
        var socket = new WebSocket('ws://localhost:8082/updates');

        socket.addEventListener('open', function(_evt) {
            console.log("Established websocket connection!");
            socket.send('Hello server!');
        });

        socket.addEventListener('message', function(evt) {
            var newState = JSON.parse(evt.data);

            WordQuest.currentPuzzle.update(newState);
        });
    }

    return initialize;
})();
