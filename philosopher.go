package main

import (
	"fmt"
	"html"
	"strconv"
	"sync"
	"time"
)

const (
	FOOD   = "0x0001F35C"
	FINISH = "0x0001F44C"
)

type Philosopher struct {
	ID             int
	Name           string
	LeftChopStick  *ChopStick
	RightChopStick *ChopStick
	Host           *Host
}

func (p *Philosopher) Eat(wg *sync.WaitGroup) {
	wg.Add(1)
	
	
	p.LeftChopStick.Lock()
	fmt.Printf("%d locked left chopstick\n", p.ID)
	p.RightChopStick.Lock()
	fmt.Printf("%d locked right chopstick\n", p.ID)


	p.Host.requestChannel <- p
	fmt.Printf("%d is eating %s\n", p.ID, GetEmoticon(FOOD))
	time.Sleep(time.Millisecond)
	fmt.Printf("%d is done eating %s\n", p.ID, GetEmoticon(FINISH))
	
	
	p.LeftChopStick.Unlock()
	fmt.Printf("%d unlocked left chopstick\n", p.ID)
	p.RightChopStick.Unlock()
	fmt.Printf("%d unlocked right chopstick\n", p.ID)
	
	wg.Done()
}

func GetEmoticon(value string) string {
	i, _ := strconv.ParseInt(value, 0, 64)
	foodEmoticon := html.UnescapeString(string(i))
	return foodEmoticon
}