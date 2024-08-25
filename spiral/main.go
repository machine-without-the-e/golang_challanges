package main

type position struct {
  row int
  column int
}

type direction struct {
  row int
  column int
}

func (direction *direction) turn() {
  if (direction.row == 0 && direction.column == 1) {
    direction.row = 1
    direction.column = 0
    return
  }

  if (direction.row == 1 && direction.column == 0) {
    direction.row = 0
    direction.column = -1
    return
  }

  if (direction.row == 0 && direction.column == -1) {
    direction.row = -1
    direction.column = 0
    return;
  }

  if (direction.row == -1 && direction.column == 0) {
    direction.row = 0
    direction.column = 1
    return
  }
}

type spiralHead struct {
  position position
  direction direction
}

func (spiralHead *spiralHead) getNextPosition() position {
  return spiralHead.position.getNextPosition(spiralHead.direction)
}

func (spiralHead *spiralHead) applyDirection() {
  spiralHead.position.row += spiralHead.direction.row
  spiralHead.position.column += spiralHead.direction.column
}


func (pos *position) getNextPosition(direction direction) position {
  return position{
    row: pos.row + direction.row,
    column: pos.column + direction.column,
  }
}

func (pos *position) hasNeighbours(spiral [][]int) int {
  directionsToCheck := []direction{
    direction{
      row: -1,
      column: 0,
    },
    direction{
      row: 0,
      column: -1,
    },
    // Ignore self (0,0)
    direction{
      row: 0,
      column: 1,
    },
    direction{
      row: 1,
      column: 0,
    },
  }

  neighbours := 0;
  for _, directionToCheck := range directionsToCheck {
    newRow := pos.row + directionToCheck.row
    newColumn := pos.column + directionToCheck.column

    if (newRow < 0 || len(spiral) -1 < newRow) {
      continue;
    }

    if (newColumn < 0 || len(spiral[newRow]) -1 < newColumn) {
      continue;
    }


    if (spiral[newRow][newColumn] != 1) {
      continue
    }

    neighbours++
  }

  return neighbours
}

func (spiralHead *spiralHead) assessDirection(spiral [][]int, size int) bool {  
  nextPosition := spiralHead.getNextPosition()
  // If we're at the end of a row or column, turn 
  if (nextPosition.row > size-1 || nextPosition.column > size-1 || nextPosition.row < 0 || nextPosition.column < 0) {
    spiralHead.direction.turn()
    return true;
  } 

  nextNextPosition := nextPosition.getNextPosition(spiralHead.direction)

  if (nextNextPosition.row > size-1 || nextNextPosition.column > size-1 || nextNextPosition.row < 0 || nextNextPosition.column < 0) {
    return true;
  } 

  if(spiral[nextNextPosition.row][nextNextPosition.column] == 1) {
    spiralHead.direction.turn()
    nextPosition = spiralHead.getNextPosition()
    nextNextPosition = nextPosition.getNextPosition(spiralHead.direction)
 
    if(nextNextPosition.hasNeighbours(spiral) > 1) {
      return false;
    }
  }
  
  return true;
}


 
func Spiralize(size int) [][]int {
  
  // Initialise the spiral canvas
  spiral := make([][]int, size)
  for index := range spiral {
    spiral[index] = make([]int, size)
  }

  spiralHead := spiralHead{
    position: position{
      row: 0,
      column: 0,
    },
    direction: direction{
      row: 0,
      column: 1,
    },
  }

  spiraling := true
  for spiraling {
    spiral[spiralHead.position.row][spiralHead.position.column] = 1
    
    // Some nice formatting to see each iteration
    // print("     ------    \n")
    // for _, part := range spiral {
    //   fmt.Printf("%v\n", part)
    // }
    //
    // fmt.Printf("with row, column: %d, %d and direction %d, %d\n", spiralHead.position.row, spiralHead.position.column, spiralHead.direction.row, spiralHead.direction.column)
    // print("     ------    \n")

    spiraling = spiralHead.assessDirection(spiral, size)

    spiralHead.applyDirection()
  }

  return spiral
}
