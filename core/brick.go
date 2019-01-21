package core

import (
	"fmt"
	"math/rand"
	"time"
)

type Brick struct {
	length  int
	raw     [100]byte
	dirBack bool
	isMove  bool
}

func (b *Brick) startPoint() int {
	for i, v := range b.raw {
		if v == '#' {
			return i
		}
	}
	return 0
}
func (b *Brick) endPoint() int {
	return b.startPoint() + b.length - 1
}
func (b *Brick) initRaw(c *Config) {
	b.length = rand.Intn(c.MaxBrickLen-c.MinBrickLen+1) + c.MinBrickLen
	for i := 0; i < b.length; i++ {
		b.raw[i] = '#'
	}
	for i := b.length; i < len(b.raw); i++ {
		b.raw[i] = ' '
	}
	b.dirBack = false
	b.isMove = true
}
func (b *Brick) makeRaw() {
	s := b.startPoint()
	e := b.endPoint()
	if b.dirBack {
		for i := s; i <= e; i++ {
			b.raw[i], b.raw[i-1] = b.raw[i-1], b.raw[i]
		}
	} else {
		for i := e; i >= s; i-- {
			b.raw[i], b.raw[i+1] = b.raw[i+1], b.raw[i]
		}
	}
	if b.raw[99] == '#' || b.raw[0] == '#' {
		b.dirBack = !b.dirBack
	}
}
func (b *Brick) printRaw(c *Config) {
	b.initRaw(c)
	for b.isMove {
		b.makeRaw()
		fmt.Printf("\r%s", string(b.raw[:]))
		time.Sleep(time.Millisecond * time.Duration(c.MoveDur))
	}
}
func (b *Brick) waitForStop(c *Config) {
	var t string
	go b.printRaw(c)
	fmt.Scanln(&t)
	b.isMove = false
}
func isMiss(b1 *Brick, b2 *Brick) bool {
	return b2.startPoint() > b1.endPoint() || b2.endPoint() < b1.startPoint()
}
