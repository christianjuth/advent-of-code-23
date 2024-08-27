package main

import (
	"log"
	"os"
	"strconv"
	"strings"
  "fmt"
  "sort"
)

type player struct {
  cards string
  rank int64
  bet int
}

// EDIT SETTINGS HERE
var isPartTwo = true

// END EDIT SETTINGS

func rankHandPt1(handStr string) int64 {
  // this is used to rank the overall hand
  sortedHand := 0

  var rank int64 = 0

  for _, char := range handStr {
    rank *= 100

    charStr := string(char)
    switch charStr {
      case "A":
        sortedHand += 1000000000000
        rank += 13
      case "K":
        sortedHand += 100000000000
        rank += 12
      case "Q":
        sortedHand += 10000000000
        rank += 11
      case "J":
        sortedHand += 1000000000
        rank += 10
      case "T":
        sortedHand += 100000000
        rank += 9
      case "9":
        sortedHand += 10000000
        rank += 8
      case "8":
        sortedHand += 1000000
        rank += 7
      case "7":
        sortedHand += 100000
        rank += 6
      case "6":
        sortedHand += 10000
        rank += 5
      case "5":
        sortedHand += 1000
        rank += 4
      case "4":
        sortedHand += 100
        rank += 3
      case "3":
        sortedHand += 10
        rank += 2
      case "2":
        sortedHand += 1
        rank += 1
    }
  }

  fiveOfAKind := 0
  fourOfAKind := 0
  threeOfAKind := 0
  twoOfAKind := 0

  sortedHandStr := fmt.Sprint(sortedHand)

  for i := 0; i < len(sortedHandStr); i++ {
    digitStr := string(sortedHandStr[i])
    digit, _ := strconv.Atoi(digitStr)
    switch digit {
      case 5:
        fiveOfAKind++
      case 4:
        fourOfAKind++
      case 3:
        threeOfAKind++
      case 2:
        twoOfAKind++
    }
  }

  if fiveOfAKind == 1 {
    rank += 60000000000
  } else if fourOfAKind == 1 {
    rank += 50000000000
  } else if threeOfAKind == 1 && twoOfAKind == 1 {
    rank += 40000000000
  } else if threeOfAKind == 1 {
    rank += 30000000000
  } else if twoOfAKind == 2 {
    rank += 20000000000
  } else if twoOfAKind == 1 {
    rank += 10000000000
  }

  return rank
}

func rankHandPt2(handStr string) int64 {
  // this is used to rank the overall hand
  sortedHand := 0
  jokerCount := 0

  var rank int64 = 0

  for _, char := range handStr {
    rank *= 100

    charStr := string(char)
    switch charStr {
      case "A":
        sortedHand += 1000000000000
        rank += 13
      case "K":
        sortedHand += 100000000000
        rank += 12
      case "Q":
        sortedHand += 10000000000
        rank += 11
      case "T":
        sortedHand += 100000000
        rank += 10
      case "9":
        sortedHand += 10000000
        rank += 9
      case "8":
        sortedHand += 1000000
        rank += 8
      case "7":
        sortedHand += 100000
        rank += 7
      case "6":
        sortedHand += 10000
        rank += 6
      case "5":
        sortedHand += 1000
        rank += 5
      case "4":
        sortedHand += 100
        rank += 4
      case "3":
        sortedHand += 10
        rank += 3
      case "2":
        sortedHand += 1
        rank += 2
      case "J":
        rank += 1
        jokerCount++
    }
  }

  fiveOfAKind := 0
  fourOfAKind := 0
  threeOfAKind := 0
  twoOfAKind := 0

  sortedHandStr := fmt.Sprint(sortedHand)

  for i := 0; i < len(sortedHandStr); i++ {
    digitStr := string(sortedHandStr[i])
    digit, _ := strconv.Atoi(digitStr)
    switch digit {
      case 5:
        fiveOfAKind++
      case 4:
        fourOfAKind++
      case 3:
        threeOfAKind++
      case 2:
        twoOfAKind++
    }
  }

  if fourOfAKind == 1 && jokerCount == 1 {
    fourOfAKind--
    fiveOfAKind++
    jokerCount--
  }
  if threeOfAKind == 1 && jokerCount > 0 {
    threeOfAKind--
    if jokerCount == 2 {
      fiveOfAKind++  
      jokerCount -= 2
    }
    if jokerCount == 1 {
      fourOfAKind++
      jokerCount--
    }
  }
  if twoOfAKind == 2 && jokerCount == 1 {
    threeOfAKind++
    twoOfAKind--
    jokerCount--
  }
  if twoOfAKind == 1 && jokerCount > 0 {
    twoOfAKind--
    if jokerCount == 3 {
      fiveOfAKind++
      jokerCount -= 3
    }
    if jokerCount == 2 {
      fourOfAKind++ 
      jokerCount -= 2
    }
    if jokerCount == 1 {
      threeOfAKind++
      jokerCount--
    }
  }
  if jokerCount > 0 {
    if jokerCount == 5 {
      fiveOfAKind++
      jokerCount -= 5
    }
    if jokerCount == 4 {
      fiveOfAKind++
      jokerCount -= 4
    } 
    if jokerCount == 3 {
      fourOfAKind++
      jokerCount -= 3
    }
    if jokerCount == 2 {
      threeOfAKind++
      jokerCount -= 2
    }
    if jokerCount == 1 {
      twoOfAKind++
      jokerCount--
    }
  }

  if fiveOfAKind == 1 {
    rank += 600000000000
  } else if fourOfAKind == 1 {
    rank += 500000000000
  } else if threeOfAKind == 1 && twoOfAKind == 1 {
    rank += 400000000000
  } else if threeOfAKind == 1 {
    rank += 300000000000
  } else if twoOfAKind == 2 {
    rank += 200000000000
  } else if twoOfAKind == 1 {
    rank += 100000000000
  }

  return rank
}

func parseLine(line string) (string, int) {
  hand := line[:5]
  betStr := line[6:]

  bet, _ := strconv.Atoi(betStr)

  return hand, bet
}

func parse(file string, isPartTwo bool) {
  // make sure file ends in new line
	lines := strings.Split(file, "\n")

  players := make([]player, 0, len(lines))
  
  for _, line := range lines {
    if line == "" {
      continue
    }
    cards, bet := parseLine(line)
    if isPartTwo {
      players = append(players, player{
        cards: cards,
        rank: rankHandPt2(cards),
        bet: bet,
      })
    } else {
      players = append(players, player{
        cards: cards,
        rank: rankHandPt1(cards),
        bet: bet,
      })
    }
  }

  sort.Slice(players, func(a int, b int) bool {
    return players[a].rank < players[b].rank
  })

  winnings := 0

  for i, v := range players {
    winnings += v.bet * (i + 1)
  }

  println(winnings)
}

func main() {
  content, err := os.ReadFile("data.txt")
  if err != nil {
    log.Fatal(err)
  }

  parse(string(content), isPartTwo)
}
