package main

import (
  "fmt"
  "math/rand"
)

type board struct {
  canvas [50][50]cell
}

func (b *board) setup() {

  for index_row, row := range b.canvas {
    for index_column := range row {
      b.canvas[index_row][index_column].position = position{
        row: index_row,
        column: index_column,
      }

      b.canvas[index_row][index_column].alive = rand.Intn(2)

    }
  }
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
      neighboursAlive := b.neighboursAlive(cell.position)
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

type position struct {
  row int
  column int
}

type cell struct {
  alive int
  position position
}

func (b *board) neighboursAlive(p position) int {
  positions_to_check := []position{
    {
      row: -1,
      column: -1,
    },
    {
      row: -1,
      column: 0,
    },
    {
      row: -1,
      column: 1,
    },
    {
      row: 0,
      column: -1,
    },
    {
      row: 0,
      column: 1,
    },
    {
      row: 1,
      column: -1,
    },
    {
      row: 1,
      column: 0,
    },
    {
      row: 1,
      column: 1,
    },
  }

  neighbours := 0

  for _, pos_to_check := range positions_to_check {
    newPosition := position {
      row: p.row + pos_to_check.row,
      column: p.column + pos_to_check.column,
    }

    if(newPosition.row >= len(b.canvas) || newPosition.row < 0) {
      continue
    }

    if(newPosition.column >= len(b.canvas[newPosition.row]) || newPosition.column < 0) {
      continue
    }


    if(b.canvas[p.row + pos_to_check.row][p.column + pos_to_check.column].alive == 1) {
      neighbours++
    }
  }

  return neighbours
}

func main() {
  b := board{}

    b.setup()
  for {
    b.draw()
    b.calculateNextGen()
  }
}
