(function () { 

  const MAZES = [
    {
      endpoint: 'http://localhost:3333/btree',
      text: 'binary tree',
      data: {
        Width: 10,
        Height: 10
      }
    }
  ]

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

  class AppModel {
    constructor () {
      this.init();
    } 

    getMaze(maze) {
      return fetch(maze.endpoint, {
        method: 'POST',
        body: JSON.stringify(maze.data),
        mode: 'cors',
        headers: {
          'Accept': 'application/json, text/plain, */*',
          'Content-Type': 'application/json'
        }
      }) 
    }

    generateMenu (container, items) {
      items.forEach(item=>{
        const el = document.createElement('a');
        el.href = '#'
        el.textContent = item.text
        el.addEventListener('click', (_)=>{
          this.getMaze(item)
            .then(res=>res.json())
            .then(json=>this.render.draw(maze))

        })
        container.appendChild(el);
      })
    }

    init() {
      const nav = document.querySelector('nav');
      const canvas = document.querySelector('canvas');
      const ctx = canvas.getContext('2d');
      this.render = new MazeRenderer(ctx);
      this.generateMenu(nav, MAZES);
    }     
  } 


  window.onload = function() {
    new AppModel()
  } 
}())
