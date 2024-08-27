package main

import (
	// "fmt"
	"log"
	"os"
	"strings"
  "strconv"
)

// EDIT SETTINGS HERE
var isPartTwo = true

// END EDIT SETTINGS

func isNumber(char string) bool {
  switch char {
    case "0":
      return true 
    case "1":
      return true 
    case "2":
      return true 
    case "3":
      return true 
    case "4":
      return true 
    case "5":
      return true 
    case "6":
      return true 
    case "7":
      return true 
    case "8":
      return true 
    case "9":
      return true 
    default:
      return false 
  }
}

func isSymbol(char string) bool {
  switch char {
    case ".":
      return false 
    default:
      return !isNumber(char) 
  }
}

func isGearChar(char string) bool {
  return char == "*"
}

func discoverNumber(matrix [][]string, y int, x int) int {
  width := len(matrix[y])

  left := x
  right := x

  for true {
    canGoLeft := left > 0
    if canGoLeft && isNumber(matrix[y][left-1]) {
      left--;
    } else {
      break;
    }
  }

  for true {
    canGoRight := right < width - 1
    if canGoRight && isNumber(matrix[y][right+1]) {
      right++;
    } else {
      break;
    }
  }

  numStr := strings.Join(matrix[y][left:right+1], "")
  num, err := strconv.Atoi(numStr)

  if err != nil {
    return 1
  }
  return num
}

func discoverRatio(matrix [][]string, y int, x int) int {
  adjacentNumCount := 0

  width := len(matrix[y])
  height := len(matrix)

  canGoLeft := x > 0
  canGoUp := y > 0
  canGoRight := x < width - 1
  canGoDown := y < height - 1

  product := 1

  // left 
  if canGoLeft && isNumber(matrix[y][x-1]){
    adjacentNumCount++
    product *= discoverNumber(matrix, y, x-1) 
  }

  if canGoUp && isNumber(matrix[y-1][x]){
    // up 
    adjacentNumCount++
    product *= discoverNumber(matrix, y-1, x) 
  } else {
    if canGoUp && canGoRight && isNumber(matrix[y-1][x+1]) {
      // up and right
      adjacentNumCount++
      product *= discoverNumber(matrix, y-1, x+1) 
    }
    if canGoLeft && canGoUp && isNumber(matrix[y-1][x-1]) {
      // left and up
      adjacentNumCount++
      product *= discoverNumber(matrix, y-1, x-1) 
    }
  }

  // right 
  if canGoRight && isNumber(matrix[y][x+1]){
    adjacentNumCount++
    product *= discoverNumber(matrix, y, x+1) 
  }

  if canGoDown && isNumber(matrix[y+1][x]){
    // down
    adjacentNumCount++
    product *= discoverNumber(matrix, y+1, x) 
  } else {
    if canGoDown && canGoLeft && isNumber(matrix[y+1][x-1]) {
      // down left
      adjacentNumCount++
      product *= discoverNumber(matrix, y+1, x-1) 
    } 
    if canGoRight && canGoDown && isNumber(matrix[y+1][x+1]) {
      // right and down
      adjacentNumCount++
      product *= discoverNumber(matrix, y+1, x+1) 
    }
  }

  if adjacentNumCount != 2 {
    return 0
  }
  
  return product
}

func cellHasAdjacentSymbol(matrix [][]string, y int, x int) bool {
  width := len(matrix[y])
  height := len(matrix)

  canGoLeft := x > 0
  canGoUp := y > 0
  canGoRight := x < width - 1
  canGoDown := y < height - 1

  // left 
  if canGoLeft && isSymbol(matrix[y][x-1]){
    return true
  }
  // left and up
  if canGoLeft && canGoUp && isSymbol(matrix[y-1][x-1]) {
    return true
  }
  // up 
  if canGoUp && isSymbol(matrix[y-1][x]){
    return true
  }
  // up and right
  if canGoUp && canGoRight && isSymbol(matrix[y-1][x+1]) {
    return true
  }
  // right 
  if canGoRight && isSymbol(matrix[y][x+1]){
    return true
  }
  // right and down
  if canGoRight && canGoDown && isSymbol(matrix[y+1][x+1]) {
    return true
  }
  // down
  if canGoDown && isSymbol(matrix[y+1][x]){
    return true
  }
  // down left
  if canGoDown && canGoLeft && isSymbol(matrix[y+1][x-1]) {
    return true
  }
  return false
}

func parse(file string, isPartTwo bool) int {
  // Splitting the string into lines
	lines := strings.Split(file, "\n")

	// Initializing a 2D slice
	var matrix [][]string

	// Iterating over each line
	for _, line := range lines {
    if line == "" {
      continue
    }
		// Splitting each line into words separated by spaces
		words := strings.Split(line, "")
		// Appending the slice of words to the 2D slice
		matrix = append(matrix, words)
	}

  sum := 0

  number := ""
  isPart := false

  for y, line := range matrix {
    for x, cell := range line {
      if !isPartTwo {
        // part one
        if !isNumber(cell) && number != "" {
          val, err := strconv.Atoi(number)
          if err == nil && isPart {
            sum += val
          }
          number = ""
          isPart = false
        } else if isNumber(cell) {
          number += cell
          isPart = isPart || cellHasAdjacentSymbol(matrix, y, x)
        }
      } else {
        // part two
        if isGearChar(cell) {
          sum += discoverRatio(matrix, y, x)
        }
      }
    }
  }

  // 80065459 is too low

  return sum
}

func main() {
  content, err := os.ReadFile("data.txt")
  if err != nil {
    log.Fatal(err)
  }

  println(parse(string(content), isPartTwo))
}
