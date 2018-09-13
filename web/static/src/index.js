window.WordQuest = window.WordQuest || {};

window.WordQuest.Puzzle = (function() {
    var refresh = function() {
        var oReq = new XMLHttpRequest();
        oReq.addEventListener("load", puzzleCallback);
        oReq.open("GET", "puzzle");
        oReq.send();
    };

    var puzzleCallback = function(_evt) {
        var data = JSON.parse(this.response);
        console.log(data);
    };

    return {
        refresh: refresh
    };
})();