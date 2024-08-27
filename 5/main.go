package main

import (
	"log"
	"os"
	"strconv"
	"strings"
  "fmt"
)

type node struct {
  minIn int
  maxIn int
}

// EDIT SETTINGS HERE
var isPartTwo = true

// END EDIT SETTINGS

func processSectionLine(line string) (int, int, int) {
  numberStrs := strings.Split(line, " ")

  item1, err1 := strconv.Atoi(numberStrs[0])
  item2, err2 := strconv.Atoi(numberStrs[1])
  item3, err3 := strconv.Atoi(numberStrs[2])

  if err1 != nil || err2 != nil || err3 != nil {
    println("encountered line in section with invalid int(s)")
    return 0,0,0
  }

  return item1, item2, item3
}

func processSection(lines []string, top int, bottom int, desination int) int {
  // +1 let's us skip the section header
  // and get to the actual numbers
  for i := top; i <= bottom; i++ {
    destinationRangeStart, sourceRangeStart, rangeLength := processSectionLine(lines[i])

    if desination >= destinationRangeStart && desination < destinationRangeStart + rangeLength {
      diff := desination - destinationRangeStart
      return sourceRangeStart + diff
    }
  } 

  return desination
}

func resolveLoc(lines []string, loc int) bool {
  i := len(lines) - 2
  sectionLineEnd := i

  for i > 0 {
    line := lines[i]

    if line == "" {
      loc = processSection(lines, i+2, sectionLineEnd, loc)
      sectionLineEnd = i - 1
    }

    i--
  }

  return isSeed(lines, loc)
}

func isSeed(lines []string, seed int) bool {
  // first line contains seeds
  line := lines[0]
  line = strings.TrimSpace(line[7:])
  seedStrs := strings.Split(line, " ")

  for i := 0; i < len(seedStrs); i += 2 {
    rangeStart, err1 := strconv.Atoi(seedStrs[i])
    rangeLen, errr2 := strconv.Atoi(seedStrs[i+1])

    if err1 != nil || errr2 != nil {
      println("failed to parse seed num", fmt.Sprint(i + 1))
    } else if (seed >= rangeStart && seed < rangeStart + rangeLen) {
      return true
    }
  }

  return false
}


func parse(file string, isPartTwo bool) {
  // make sure file ends in new line
	lines := strings.Split(file, "\n")

  seed := 0
  increment := 1000000
  
  for increment >= 1 {
    for !resolveLoc(lines, seed) {
      seed += increment
    }

    if increment == 1 {
      println(seed)
    }

    seed -= increment
    increment /= 10
  }
}

func main() {
  content, err := os.ReadFile("data.txt")
  if err != nil {
    log.Fatal(err)
  }

  parse(string(content), isPartTwo)
}
