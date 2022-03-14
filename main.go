package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)

  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }

  openedVault := rand.Intn(100)

  if (isHeistOn && openedVault >= 70) {
    fmt.Println("Grab and GO!")
  } else if isHeistOn {
    isHeistOn = false
    fmt.Println("The vault cannot be opened!")
  }

  leftSafely := rand.Intn(5)
  if (isHeistOn) {
    switch leftSafely {
    case 0:
      isHeistOn = false
      fmt.Println("You failed to escape")
    case 1:
      isHeistOn = false 
      fmt.Println("Turns out vault doors don't open from the inside...")
    case 2: 
      isHeistOn = false 
      fmt.Println("Turns out vault doors don't open from the inside... damn")
    case 3: 
      isHeistOn = false 
      fmt.Println("Turns out vault doors don't open from the inside... nooo")
    default:
      fmt.Println("Start the getaway car!")
    }
  }

  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Printf("You stole %d!", amtStolen)
  }

  fmt.Println(isHeistOn)
}