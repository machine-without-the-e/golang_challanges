package main

import (
  "fmt"
)

type board struct {
  canvas [10][10]cell
}

func (b *board) setup() {
  b.canvas[1][1].alive = 1;
}

func (b *board) draw() {
  for _, row := range b.canvas {
    for _, column := range row {
      fmt.Printf("%v", column.alive)
    }
    fmt.Println()
  }
  fmt.Println()
  fmt.Println()
}

func (b *board) calculateNextGen() {
  for i, row := range b.canvas {
    for j, cell := range row {
      neighboursAlive := cell.neighboursAlive()
      if(cell.alive == 0) {
        if(neighboursAlive == 3) {
          b.canvas[i][j].alive = 1
        }

        continue
      }

      if(neighboursAlive < 2) {
        b.canvas[i][j].alive = 0
        continue
      }

      if(neighboursAlive == 2 || neighboursAlive == 3) {
        continue
      }

      b.canvas[i][j].alive = 0;
    }
  }
}

type cell struct {
  alive int
}

func (c *cell) neighboursAlive() int {
  return 0;
}

func main() {
  b := board{}

    b.setup()
  for {
    b.draw()
    b.calculateNextGen()
  }
}
