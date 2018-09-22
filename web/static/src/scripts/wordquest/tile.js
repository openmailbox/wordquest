window.WordQuest = window.WordQuest || {};

/**
 * Individual cell within the puzzle.
 * @constructor
 * @param {Number}gx - The x-coordinate of the tile.
 * @param {Number} y - The y-coordinate of the tile.
 * @param {string} value - The letter shown on this tile.
 * @param {Puzzle} puzzle - The containing puzzle.
 */
WordQuest.Tile = function (x, y, value, puzzle) {
  this.x       = x;
  this.y       = y;
  this.value   = value;
  this.puzzle  = puzzle;
  this.element = null;
};

WordQuest.Tile.prototype.addHighlight = function() {
  this.element.classList.add('highlighted');
}

/**
 * @param {Element} parent - The parent element to draw into.
 */
WordQuest.Tile.prototype.draw = function (parent) {
  if (this.element !== null) this.element.remove();

  this.element = document.createElement('td');

  this.element.innerText = this.value;
  this.element.dataset.x = this.x;
  this.element.dataset.y = this.y;

  this.element.addEventListener("mousedown", this);
  this.element.addEventListener("mouseover", this);

  parent.appendChild(this.element);
};

/**
 * Generic event handler for callbacks to make Puzzle conform to EventListener interface
 * @param {Event} event
 */
WordQuest.Tile.prototype.handleEvent = function(event) {
  if (event.type === 'mousedown' || (event.type === 'mouseover' && this.puzzle.highlighting)) {
    this.element.classList.add('highlighted');
    event.highlightingTile = this;
  }
};

WordQuest.Tile.prototype.removeHighlight = function() {
  this.element.classList.remove('highlighted');
};
