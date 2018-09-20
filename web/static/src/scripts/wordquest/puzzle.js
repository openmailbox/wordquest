window.WordQuest = window.WordQuest || {};

/**
 * Top-level component for the game puzzle.
 * @constructor
 * @param {Number} length - The maximum Y-value of the puzzle.
 * @param {Number} width - The maximum X-value of the puzzle.
 * @param {Object[]} tiles - The provided tile-data.
 * @param {Number} tiles[].x - The X-coordinate of this tile.
 * @param {Number} tiles[].y - The Y-coordinate of this tile.
 * @param {string} tiles[].value - The letter contained in this tile.
 */
WordQuest.Puzzle = function(length, width, tiles) {
  this.tiles        = [];
  this.length       = length;
  this.width        = width;
  this.highlighting = false; // is the user currently dragging + highlighting an answer
  this.highlighted  = [];    // array of tiles currently highlighted

  for (var i = 0; i < tiles.length; i++) {
    var newTile = new WordQuest.Tile(tiles[i].x, tiles[i].y, tiles[i].value, this);
    this.tiles.push(newTile);
  }
};

WordQuest.Puzzle.prototype.draw = function () {
  var table    = document.createElement('table');
  var tbody    = document.createElement('tbody');
  var tr       = document.createElement('tr');
  var currentY = this.tiles[0].y;

  for (var i = 0; i < this.tiles.length; i++) {
    if (this.tiles[i].y != currentY) {
      tbody.appendChild(tr);
      tr = document.createElement('tr');
      currentY += 1;
    }

    this.tiles[i].draw(tr);
  }

  table.addEventListener("mouseover", this);
  table.addEventListener("mousedown", this);
  table.addEventListener("mouseup", this);

  table.appendChild(tbody);

  document.getElementById('wordquest').appendChild(table);
}

WordQuest.Puzzle.prototype.endHighlight = function() {
  var word        = '';
  var tileObjects = [];

  for (var i = 0; i < this.highlighted.length; i++) {
    word += this.highlighted[i].value;

    var newTile = { 
      x:     this.highlighted[i].x,
      y:     this.highlighted[i].y,
      value: this.highlighted[i].value
    }

    tileObjects.push(newTile);
  }

  var submission = { 
    value: word,
    tiles: tileObjects
  };

  var oReq = new XMLHttpRequest();

  oReq.addEventListener("load", function(_) {
    console.log(this.response);
  });

  oReq.open("POST", "submit");
  oReq.send(JSON.stringify(submission));

  this.highlighting = false;
  this.highlighted  = [];
};

/**
 * Generic event handler for callbacks to make Puzzle conform to EventListener interface
 * @param {Event} event
 * @param {Tile} event.highlightedTile - The highlighted tile
 */
WordQuest.Puzzle.prototype.handleEvent = function(event) {
  if (event.type === 'mousedown' || (event.type === 'mouseover' && this.highlighting)) {
    if (this.highlighted[this.highlighted.length - 2] == event.highlightingTile) {
      var leaving = this.highlighted.pop();
      leaving.removeHighlight();
    } else {
      this.highlight(event.highlightingTile);
    }
  } else if (event.type === 'mouseup') {
    this.endHighlight();
  }
}

WordQuest.Puzzle.prototype.highlight = function(tile) {
  this.highlighting = true;
  this.highlighted.push(tile);
};
