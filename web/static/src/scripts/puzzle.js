import Tile from './tile.js';

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
function Puzzle(length, width, tiles) {
  this.tiles        = [];
  this.length       = length;
  this.width        = width;
  this.highlighting = false; // is the user currently dragging + highlighting an answer
  this.highlighted  = [];    // array of tiles currently highlighted

  for (var i = 0; i < tiles.length; i++) {
    var newTile = new Tile(tiles[i].x, tiles[i].y, tiles[i].value, this);
    this.tiles.push(newTile);
  }
};

Puzzle.prototype.draw = function () {
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

Puzzle.prototype.endHighlight = function() {
  this.highlighting = false;
  this.highlighted  = [];
  // submit answer
};

/**
 * Generic event handler for callbacks to make Puzzle conform to EventListener interface
 * @param {Event} event 
 * @param {Tile} event.highlightedTile - The highlighted tile
 */
Puzzle.prototype.handleEvent = function(event) {
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

Puzzle.prototype.highlight = function(tile) {
  this.highlighting = true;
  this.highlighted.push(tile);
};

export default Puzzle;