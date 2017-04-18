(function () { 
  class MazeRenderer { 
    constructor (ctx) {
      if(!ctx) { 
        throw new Error('Context is required');
      }
      this.ctx = ctx;
    } 

    draw (maze) {

    } 
  }

  function init() {

  } 

  window.onload = function() {
    const canvas = document.querySelector('canvas');
    const ctx = canvas.getContext('2d');
    const render = new MazeRenderer(ctx);

  } 
}())
