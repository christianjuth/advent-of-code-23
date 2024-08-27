package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
  minIn int
  maxIn int
}

// EDIT SETTINGS HERE
var isPartTwo = true

// END EDIT SETTINGS


func parseRaces(line string) []int {
  nums := make([]int, 0, 10)

  numStr := ""

  for _, char := range line {
    charStr := string(char)

    if charStr == " " { 
      if numStr != "" && !isPartTwo {
        num, _ := strconv.Atoi(numStr)
        nums = append(nums, num)
        numStr = ""
      }
      continue
    }

    numStr += charStr
  }

  if numStr != "" {
    num, _ := strconv.Atoi(numStr)
    nums = append(nums, num)
    numStr = ""
  }

  return nums
}

func numWaysToWinRace(raceTimeMs int, recordDistanceMm int) int {
  numWays := 0

  chargeTime := 1
  driveTime := raceTimeMs - chargeTime

  // each ms of charge time increases our speed by 1mm per ms

  for driveTime > 0 {
    speed := chargeTime
    distance := speed * driveTime

    if distance > recordDistanceMm {
      numWays++
    }

    driveTime--
    chargeTime++
  }

  return numWays
}

func parse(file string, isPartTwo bool) {
  // make sure file ends in new line
	lines := strings.Split(file, "\n")

  line1 := lines[0][11:]
  line2 := lines[1][11:]

  marginOfError := 1

  raceTimes := parseRaces(line1)
  raceRecords := parseRaces(line2)

  for i, raceTime := range raceTimes {
    record := raceRecords[i]
    marginOfError *= numWaysToWinRace(raceTime, record)
  }


  println(marginOfError)

  
}

func main() {
  content, err := os.ReadFile("data.txt")
  if err != nil {
    log.Fatal(err)
  }

  parse(string(content), isPartTwo)
}
