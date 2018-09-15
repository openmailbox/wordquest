/**
 * Individual cell within the puzzle.
 * @constructor
 * @param {Number} x - The x-coordinate of the tile.
 * @param {Number} y - The y-coordinate of the tile.
 * @param {string} value - The letter shown on this tile.
 */
var Tile = function (x, y, value) {
  this.x     = x;
  this.y     = y;
  this.value = value;
};

/**
 * @param {Element} parent - The parent element to draw into. 
 */
Tile.prototype.draw = function (parent) {
  var td = document.createElement('td');

  td.innerText = this.value;
  td.dataset.x = this.x;
  td.dataset.y = this.y;

  parent.appendChild(td);
};

export default Tile;