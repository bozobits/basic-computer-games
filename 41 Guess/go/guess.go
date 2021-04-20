package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var l, l1 int
	fmt.Println("\t\tGUESS")
	fmt.Println("CREATIVE COMPUTING  MORRISTOWN, NEW JERSEY")
	fmt.Println()
	fmt.Println("THIS IS A NUMBER GUESSING GAME. I'LL THINK")
	fmt.Println("OF A NUMBER BETWEEN 1 AND ANY LIMIT YOU WANT.")
	fmt.Println("THEN YOU HAVE TO GUESS WHAT IT IS.")
	fmt.Println()
	fmt.Println("WHAT LIMIT DO YOU WANT?")
	fmt.Scanf("%d", &l)
	fmt.Println()
	l1 = int(math.Log(float64(l))/math.Log(2)) + 1
	rand.Seed(time.Now().UnixNano())
	for {
		var n, m int
		m = rand.Intn(l) + 1
		fmt.Println("I'M THINKING OF A NUMBER BETWEEN 1 AND", l)
		fmt.Println("NOW YOU TRY TO GUESS WHAT IT IS.")
		for g := 1; true; g++ {
			fmt.Scanf("%d", &n)
			if n > m {
				fmt.Println("TOO HIGH. TRY A SMALLER ANSWER.")
			} else if n < m {
				fmt.Println("TOO LOW. TRY A BIGGER ANSWER.")
			} else if m == n {
				fmt.Println("THAT'S IT! YOU GOT IT IN", g, "TRIES.")
				if g > l1 {
					fmt.Println("YOU SHOULD HAVE BEEN ABLE TO GET IT IN ONLY", l1)
				} else if g == l1 {
					fmt.Println("GOOD")
				} else {
					fmt.Println("VERY GOOD")
				}
				fmt.Println()
				break
			}
		}
	}
}
