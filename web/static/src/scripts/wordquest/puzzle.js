window.WordQuest = window.WordQuest || {};

/**
 * Top-level component for the game puzzle.
 * @constructor
 */
WordQuest.Puzzle = function() {
  this.element      = null;
  this.highlighting = false; // is the user currently dragging + highlighting an answer
  this.highlighted  = [];    // array of tiles currently highlighted
  this.solutions    = [];
  this.tiles        = [];
};

WordQuest.Puzzle.prototype.draw = function () {
  if (this.element !== null) this.element.remove();

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

  this.element = table;

  this.highlightSolutions();
}

WordQuest.Puzzle.prototype.endHighlight = function() {
  var word        = '';
  var tileObjects = [];

  this.highlighting = false;

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

  oReq.addEventListener("load", this);

  oReq.open("POST", "submit");
  oReq.send(JSON.stringify(submission));
};

WordQuest.Puzzle.prototype.getTial = function(x, y) {
  for (var i = 0; i < this.tiles.length; i++) {
    if (this.tiles[i].x === x && this.tiles[i].y === y) return this.tiles[i];
  }

  return null;
}

/**
 * Generic event handler for callbacks to make Puzzle conform to EventListener interface
 * @param {Event} event
 * @param {WordQuest.Tile} event.highlightedTile - The highlighted tile
 */
WordQuest.Puzzle.prototype.handleEvent = function(event) {
  switch (event.type) {
    case 'mousedown':
    case 'mouseover':
      this.handleHighlighting(event.type, event.highlightingTile);
      break;
    case 'mouseup':
      this.endHighlight();
      break;
    case 'load':
      this.handleSubmissionResult(event.target);
      break;
  }
}

/**
 * Event handler for highlighting tiles in the puzzle
 * @param {string} type - the type of the event (i.e. 'mousedown')
 * @param {WordQuest.Tile} tile - The tile currently being highlighted
 */
WordQuest.Puzzle.prototype.handleHighlighting = function(type, tile) {
  if (type !== 'mousedown' && !this.highlighting) return;

  if (this.highlighted[this.highlighted.length - 2] == tile) {
    var leaving = this.highlighted.pop();
    leaving.removeHighlight();
  } else {
    this.highlight(tile);
  }
}

/**
 * Callback for the server response after an answer submission.
 * @param {XMLHttpRequest} request - The request object.
 */
WordQuest.Puzzle.prototype.handleSubmissionResult = function(event) {
  // TODO: If the tile is involved in another solution, don't un-highlight it
  if (event.status === 404) {
    for (var i = 0; i < this.highlighted.length; i++) {
      this.highlighted[i].removeHighlight();
    }
  }

  this.highlighted  = [];
}

WordQuest.Puzzle.prototype.highlight = function(tile) {
  this.highlighting = true;
  this.highlighted.push(tile);
};

WordQuest.Puzzle.prototype.highlightSolutions = function() {
  for (var i = 0; i < this.solutions.length; i++) {
    var solution = this.solutions[i];

    for (var j = 0; j < solution.tiles.length; j++) {
      var x    = solution.tiles[j].x;
      var y    = solution.tiles[j].y;
      var tial = this.getTial(x, y);

      tial.addHighlight();
    }
  }
};

/**
 * Update the internal state of the puzzle from new data.
 * @param {Object} newState 
 * @param {Number} newState.length - The length of the puzzle
 * @param {Number} newState.width - The width of the puzzle
 * @param {Object[]} newState.tiles - The provided tile-data.
 * @param {Number} newState.tiles[].x - The X-coordinate of this tile.
 * @param {Number} newState.tiles[].y - The Y-coordinate of this tile.
 * @param {string} newState.tiles[].value - The letter contained in this tile.
 * @param {Object[]} newState.solutions - The known solutions to this puzzle.
 */
WordQuest.Puzzle.prototype.update = function(newState) {
  console.log(newState);
  this.tiles     = [];
  this.length    = newState.length;
  this.width     = newState.width;
  this.solutions = newState.solutions || [];

  for (var i = 0; i < newState.tiles.length; i++) {
    var newTile = new WordQuest.Tile(newState.tiles[i].x, newState.tiles[i].y, newState.tiles[i].value, this);
    this.tiles.push(newTile);
  }

  this.draw();
};