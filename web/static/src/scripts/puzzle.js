import Tile from './tile.js';

/**
 * Top-level component for the game puzzle.
 * @constructor
 * @param {Number} length - The maximum Y-value of the puzzle.
 * @param {Number} width - The maximum X-value of the puzzle.
 * @param {Object[]} tiles - The provided tile-data.
 * @param {Number} tiles[].X - The X-coordinate of this tile.
 * @param {Number} tiles[].Y - The Y-coordinate of this tile.
 * @param {string} tiles[].value - The letter contained in this tile.
 */
function Puzzle(length, width, tiles) {
  this.tiles  = [];
  this.length = length;
  this.width  = width;

  for (var i = 0; i < tiles.length; i++) {
    var newTile = new Tile(tiles[i].x, tiles[i].y, tiles[i].value);
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

  table.appendChild(tbody);

  document.getElementById('wordquest').appendChild(table);
}

export default Puzzle;