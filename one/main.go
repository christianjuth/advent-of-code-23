package main

import (
  "fmt"
  "os"
  "log"
  "strings"
  "strconv"
  "regexp"
)

var writtenDigit, _ = regexp.Compile("^[0-9]$")
var digitRegex, _ = regexp.Compile("(one|two|three|four|five|six|seven|eight|nine|[0-9])")


var digitMap = map[string]int{
  "":  -1, "one": 1, "two": 2, "three": 3, "four": 4,
  "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}

func getDigitValue(char string) int {
  if value, exsisits := digitMap[char]; exsisits {
    return value
  }

  if writtenDigit.MatchString(char) {
    val, err := strconv.Atoi(char)
    if err != nil {
      return -1
    }
    return val
  }
 
  return -1
}

func extractDigits(line string) (int, int) {
  firstDigit := 0
  secondDigit := 0

  front := ""

  for i := 0; i < len(line); i++ {
    front = front + string(line[i])
    match := digitRegex.FindString(front)
    value := getDigitValue(match)
    if value != -1  {
      firstDigit = value      
      break
    }
  }

  back := ""

  for i := len(line) - 1; i >= 0; i-- {
    back = string(line[i]) + back
    match := digitRegex.FindString(back)
    value := getDigitValue(match)
    if value != -1  {
      secondDigit = value      
      break
    }
  }

  return firstDigit * 10, secondDigit
}

func parse(file string) int {
  lines := strings.Split(file, "\n")

  totalSum := 0

  for _, line := range(lines) {
    firstDigit, secondDigit := extractDigits(line)
    sum := firstDigit + secondDigit
    totalSum += sum
  }

  return totalSum
}

func main() {
  content, err := os.ReadFile("data.txt")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(parse(string(content)))
}
