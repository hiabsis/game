package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// World 世界
type World struct {
	// 地图上的元素
	elements map[string]*Element
	// 玩家
	player *Player
	wg     *sync.WaitGroup
	// 物品生成器
	elementFactory *ElementFactory
	// 背包
	backpack *Backpack
	time     []int
}

// 初始化世界
func (world *World) init() {
	world.wg.Add(2)
	world.generateElement()
	world.show()
	go world.monitorKeyBoard()
	go world.updateTimeThread()
	world.wg.Wait()
}

// 添加元素
func (world *World) addElement(element *Element) {
	key := strconv.Itoa(element.x) + "_" + strconv.Itoa(element.y)
	world.elements[key] = element
}

// 删除元素
func (world *World) removeElement(element *Element) {
	key := strconv.Itoa(element.x) + "_" + strconv.Itoa(element.y)
	delete(world.elements, key)
}

// 通知元素
func (world *World) notify(command string) {
	key := strconv.Itoa(world.player.x) + "_" + strconv.Itoa(world.player.y)
	element, ok := world.elements[key]
	if ok {
		element.command(world.player, element, command)
	}
}

// 监听键盘输入
func (world *World) monitorKeyBoard() {
	defer world.wg.Done()
	for {
		cmd := scanner()
		if cmd == "w" || cmd == "s" || cmd == "a" || cmd == "d" {
			world.movePlayer(cmd)
			world.generateElement()
		} else if cmd == "h" || cmd == "H" {
			world.help()
		} else if cmd == "b" || cmd == "B" {
			world.openPack()

		} else if cmd == "exit" {
			break
		}
		world.notify(cmd)
		world.show()

	}

}

// 移动角色
func (world *World) movePlayer(cmd string) {

	if cmd == "s" || cmd == "S" {
		world.player.y -= 1
	} else if cmd == "w" || cmd == "W" {
		world.player.y += 1
	} else if cmd == "a" || cmd == "A" {
		world.player.x -= 1
	} else if cmd == "d" || cmd == "D" {
		world.player.x += 1
	}

}

// 打开背包
func (world *World) openPack() {
	world.backpack.show()

	world.backpack.monitorKeyboard()

}

func (world *World) generateElement() {
	element := world.getElement()
	if element == nil {
		element = world.elementFactory.generateElement(world.player.x, world.player.y)
		world.addElement(element)
	}
}

// 获取地图的元素
func (world *World) getElement() *Element {
	key := strconv.Itoa(world.player.x) + "_" + strconv.Itoa(world.player.y)
	if element, ok := world.elements[key]; ok {
		return element
	} else {
		return nil
	}
}

// 显示世界图
func (world *World) show() {
	clearConsole()

	printLine()
	fmt.Printf("时间\t第%v天第%v时\n", world.time[0], world.time[1])
	// 角色信息
	world.player.show()
	printLine()
	world.getElement().show(world.getElement())
	printLine()

}

// 更新游戏中的时间
func (world *World) updateTimeThread() {
	defer world.wg.Done()
	for {
		minute := getCurrentMinute()
		if minute != world.time[2] {
			world.time[2] = minute
			world.time[1] += 1
			if world.time[1] > 23 {
				world.time[1] = 0
				world.time[0] += 1
			}

		}
		time.Sleep(time.Second * 30)
	}

}

// 基础操作信息
func (world *World) help() {
	clearConsole()
	printLine()
	fmt.Printf("操作:\n移动\t北：N\t南: N\t东: A\t西: W\n背包:\tB\n")
	fmt.Println("返回 H")
	printLine()

	for {
		cmd := scanner()
		if cmd == "H" || cmd == "h" {
			world.show()
			break
		}
	}

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
		hunger:    100,
		ant:       randNum(100, 200),
		define:    randNum(1, 6),
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
