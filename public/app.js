(function () { 

  const MAZES = [
    {
      endpoint: 'http://localhost:3333/btree',
      text: 'binary tree',
      data: {
        Width: 20,
        Height: 20
      }
    }
  ]

  class MazeRenderer { 

    constructor (canvas) {
      if(!canvas) { 
        throw new Error('Canvas is required');
      }
      this.canvas = canvas;
      this.width = this.canvas.width;
      this.height = this.canvas.height
      this.ctx = this.canvas.getContext('2d');
      this.wallsCollor = 'black';
      this.cellSize = 20;
    } 

    clear () {
      this.ctx.clearRect(0, 0, this.width, this.height)
    }

    draw (maze) {
      this.clear();
      this.drawBorder(maze);
      this.drawBody(maze)
    } 

    drawBody (maze) {
      maze.forEach((row, i)=>{
        row.forEach((cell, j)=>{
          this.drawCell(cell, j, i)
        })
      })
    } 

    drawCell(cell, i, j) {
        const h = this.cellSize * j;
        const w = this.cellSize * i

        if(cell.RightWall) {
            this.ctx.beginPath();
            this.ctx.moveTo(w + this.cellSize, h)
            this.ctx.lineTo(w + this.cellSize, h + this.cellSize )
            this.ctx.stroke()
        }
        
        if(cell.BottomWall) {
            this.ctx.beginPath();
            this.ctx.moveTo(w, h + this.cellSize )
            this.ctx.lineTo(w + this.cellSize, h + this.cellSize )
            this.ctx.stroke()
        } 
    }

    drawBorder (maze) {
      this.ctx.fillStyle = this.wallsCollor;
      this.ctx.strokeRect(0, 0, maze[0].length * this.cellSize, maze.length * this.cellSize);
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
            .then(maze=>this.render.draw(maze))
        })
        container.appendChild(el);
      })
    }

    init() {
      const nav = document.querySelector('nav');
      const canvas = document.querySelector('canvas');
      this.render = new MazeRenderer(canvas);
      this.generateMenu(nav, MAZES);
    }     
  } 


  window.onload = function() {
    new AppModel()
  } 
}())
