package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

type Config struct {
	MoveDur     int
	MaxBrickNum int
	MinBrickLen int
	MaxBrickLen int
}

func printLost() {
	var t string
	fmt.Println(` _              _   _
| |    ___  ___| |_| |
| |   / _ \/ __| __| |
| |__| (_) \__ \ |_|_|
|_____\___/|___/\__(_)`)
	fmt.Scanln(&t)
}
func printWin() {
	var t string
	fmt.Println(`__        ___       _
\ \      / (_)_ __ | |
 \ \ /\ / /| | '_ \| |
  \ V  V / | | | | |_|
   \_/\_/  |_|_| |_(_)`)
	fmt.Scanln(&t)
}
func (c *Config) GetConfig() {
	data, err := ioutil.ReadFile("BrickGame.json")
	if err != nil {
		fmt.Println("Creating configuration file...")
		conf := Config{
			MoveDur:     100,
			MaxBrickNum: 30,
			MinBrickLen: 3,
			MaxBrickLen: 5,
		}
		*c = conf
		c.WriteConfig()
		fmt.Println("Have fun!")
	} else {
		err = json.Unmarshal(data, c)
		if err != nil {
			panic(err)
		}
	}
}
func (c *Config) WriteConfig() {
	data, err := json.Marshal(*c)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("BrickGame.json", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Configuration file created successfully!")
}
func (c *Config) customLevel() {
	var txt string
	var value int
	fmt.Print("The time it takes for the brick to move one space:(ms)")
	fmt.Scanln(&txt)
	value, err := strconv.Atoi(txt)
	for err != nil {
		fmt.Println("Invalid! Try again.")
		fmt.Print("The time it takes for the brick to move one space:(ms)")
		fmt.Scanln(&txt)
		value, err = strconv.Atoi(txt)
	}
	c.MoveDur = value
	fmt.Print("The number of bricks used to win the game:")
	fmt.Scanln(&txt)
	value, err = strconv.Atoi(txt)
	for err != nil {
		fmt.Println("Invalid! Try again.")
		fmt.Print("The number of bricks used to win the game:")
		fmt.Scanln(&txt)
		value, err = strconv.Atoi(txt)
	}
	c.MaxBrickNum = value
	fmt.Print("Maximum length of brick:(<=30)")
	fmt.Scanln(&txt)
	value, err = strconv.Atoi(txt)
	for err != nil || value > 30 {
		fmt.Println("Invalid! Try again.")
		fmt.Print("Maximum length of brick:(<=30)")
		fmt.Scanln(&txt)
		value, err = strconv.Atoi(txt)
	}
	c.MaxBrickLen = value
	fmt.Printf("Minimum length of brick:(<=%d)", c.MaxBrickLen)
	fmt.Scanln(&txt)
	value, err = strconv.Atoi(txt)
	for err != nil || value > c.MaxBrickLen {
		fmt.Println("Invalid! Try again.")
		fmt.Printf("Minimum length of brick:(<=%d)", c.MaxBrickLen)
		fmt.Scanln(&txt)
		value, err = strconv.Atoi(txt)
	}
	c.MinBrickLen = value
}
func (c *Config) SelectLevel() {
	var txt string
	var choice int
	fmt.Println(`1:Easy
2:Medium
3:Hard
4:Custom`)
	fmt.Scanln(&txt)
	choice, err := strconv.Atoi(txt)
	for err != nil || (choice != 1 && choice != 2 && choice != 3 && choice != 4) {
		fmt.Println("Invalid! Try again.")
		fmt.Println(`1:Easy
2:Medium
3:Hard
4:Custom`)
		fmt.Scanln(&txt)
		choice, err = strconv.Atoi(txt)
	}
	switch choice {
	case 1:
		c.MoveDur = 300
		c.MaxBrickNum = 20
		c.MinBrickLen = 5
		c.MaxBrickLen = 7
	case 2:
		c.MoveDur = 100
		c.MaxBrickNum = 30
		c.MinBrickLen = 3
		c.MaxBrickLen = 5
	case 3:
		c.MoveDur = 50
		c.MaxBrickNum = 50
		c.MinBrickLen = 1
		c.MaxBrickLen = 3
	case 4:
		c.customLevel()
	}
	c.WriteConfig()
}
func (c *Config) PlayBrick() {
	rand.Seed(time.Now().Unix())
	var (
		b            []Brick
		defaultBrick Brick
	)
	b = append(b, defaultBrick)
	b[0].initRaw(c)
	initP := rand.Intn(50) + 20
	for i := 0; i < b[0].length; i++ {
		b[0].raw[initP+i], b[0].raw[i] = b[0].raw[i], b[0].raw[initP+i]
	}
	fmt.Println(string(b[0].raw[:]))
	for i := 1; i < c.MaxBrickNum; i++ {
		b = append(b, defaultBrick)
		b[i].waitForStop(c)
		if isMiss(&b[i-1], &b[i]) {
			printLost()
			return
		}
	}
	if !isMiss(&b[c.MaxBrickNum-2], &b[c.MaxBrickNum-1]) {
		printWin()
	}
}
