package main

import (
  "fmt"
  "os"
  "log"
  "strings"
  "strconv"
)

// EDIT SETTINGS HERE
var isPartTwo = true

var constrainedDraw = draw{
  red: 12,
  green: 13,
  blue: 14,
}
// END EDIT SETTINGS

type draw struct {
  red int
  green int
  blue int
}

func parseLine(line string) (int, draw) {
  splt := strings.Split(line, ":")

  // game is in the left side of the split, and we
  // skip over 5 chars so we get only the "#" in "Game #"
  gameId, _ := strconv.Atoi(splt[0][5:])

  f := func(c rune) bool {
		return string(c) == "," || string(c) == ";"
	}
  clubDraws := strings.FieldsFunc(splt[1], f)

  maxDraw := draw{
    red: 0,
    green: 0,
    blue: 0,
  }

  for _, draw := range(clubDraws) {
    drawStr := strings.TrimSpace(string(draw))
    valueLabel := strings.Split(drawStr, " ")
    value, _ :=  strconv.Atoi(string(valueLabel[0]))

    switch string(valueLabel[1]) {
      case "red":
        maxDraw.red = max(maxDraw.red, value) 
      case "green":
        maxDraw.green = max(maxDraw.green, value) 
      case "blue":
        maxDraw.blue = max(maxDraw.blue, value) 
    }
  }

  return gameId, maxDraw
}

func parse(file string) int {
  games := strings.Split(file, "\n")

  sum := 0

  for _, line := range(games) {
    if line == "" {
      continue
    }
    gameId, maxDrawPossible := parseLine(string(line))

    if isPartTwo == false {
      // Part 1
      if maxDrawPossible.red > constrainedDraw.red {
        continue
      } else if maxDrawPossible.green > constrainedDraw.green {
        continue
      } else if maxDrawPossible.blue > constrainedDraw.blue {
        continue
      }
      sum += gameId
    } else {
      // Part 2
      sum += maxDrawPossible.red * maxDrawPossible.green * maxDrawPossible.blue
    }
  }

  return sum
}

func main() {
  content, err := os.ReadFile("data.txt")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(parse(string(content)))
}
