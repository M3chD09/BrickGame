package main

import (
	. "./core"
	"fmt"
	"strconv"
)

func main() {
	var config Config
	var txt string
	var choice int
	fmt.Println(`Welcome to Brick Game!
The bricks will move back and forth. Press Enter to stop moving.
The position of the first brick is random, try to connect them in the vertical direction.
1:Play
2:Config`)
	fmt.Scanln(&txt)
	choice, err := strconv.Atoi(txt)
	for err != nil || (choice != 1 && choice != 2) {
		fmt.Println("Invalid! Try again.")
		fmt.Println(`1:Play
2:Config`)
		fmt.Scanln(&txt)
		choice, err = strconv.Atoi(txt)
	}
	switch choice {
	case 1:
		config.GetConfig()
		config.PlayBrick()
	case 2:
		config.SelectLevel()
		config.PlayBrick()
	}
}
