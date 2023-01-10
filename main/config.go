package main

import "sync"

type Config struct {
	player map[string]interface{}
}

var WORLD = &World{
	wg: &sync.WaitGroup{},
	player: &Player{
		name:      "cyw",
		x:         0,
		y:         0,
		maxHp:     100,
		hp:        100,
		maxHunger: 100,
		hunger:    0,
		ant:       randNum(3, 10),
		define:    randNum(3, 10),
	},
	elements:       make(map[string]*Element),
	elementFactory: &ElementFactory{},
	backpack: &Backpack{
		size:  30,
		cap:   0,
		items: make([]*Element, 0),
		sign:  make(chan bool, 1),
	},
	time: []int{1, 0, getCurrentMinute()},
}
