package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

type Backpack struct {
	// 存放的物品
	items []*Element
	// 背包容量
	size int
	// 已储存物品数量
	cap int
	// 位置
	index int
	sign  chan bool
	// 选中物品详情
	selectDetail string
}

func (pack *Backpack) monitorKeyboard() {
	for {
		cmd := scanner()
		if cmd == "h" || cmd == "H" {
			WORLD.player.action = ""
			break
		}

		if cmd == "w" {
			pack.index -= 1
			if pack.index < 0 {
				pack.index = pack.cap - 1
			}
			WORLD.player.action = ""
		} else if cmd == "s" {
			pack.index += 1
			if pack.index >= pack.cap {
				pack.index = 0
			}
			WORLD.player.action = ""
		}
		if cmd == "U" || cmd == "u" {
			if pack.cap != 0 {
				pack.items[pack.index].command(WORLD.player, pack.items[pack.index], cmd)
			}

		}
		pack.show()
	}

}
func (pack *Backpack) show() {
	clearConsole()
	printLine()
	WORLD.player.show()
	printLine()
	if pack.cap == 0 {
		fmt.Printf("背包空空如也\n")
	} else {
		for i, item := range pack.items {
			if pack.index == i {
				fmt.Println(aurora.BgWhite(fmt.Sprintf("%v. %v\t数量：%v\n ", i, item.name, item.number)))
				pack.selectDetail = item.detail
			} else {
				fmt.Println(fmt.Sprintf("%v. %v\t数量：%v\n ", i, item.name, item.number))
			}

		}
	}

	printLine()
	pack.help()
	printLine()
	if pack.selectDetail != "" {
		fmt.Printf(pack.selectDetail)
		printLine()
	}

	if WORLD.player.action != "" {
		fmt.Print(WORLD.player.action)
		printLine()
	}
}
func (pack *Backpack) help() {
	help := fmt.Sprintf("选择物品: 向上移动%v 向下移动 %v\n", aurora.Red("W"), aurora.Red("S"))
	help += fmt.Sprintf("背包\t容量: %v\t物品 :%v\t \n", pack.size, pack.cap)
	fmt.Printf(help)
}

// 添加物品
func (pack *Backpack) addElement(element *Element) bool {
	if pack.cap == pack.size {
		WORLD.player.action = "背包容量已满 无法拾取物品"
		return false
	}
	index := pack.findElement(element)
	if index == -1 {
		pack.items = append(pack.items, element)
		pack.cap += 1
		return true
	}
	pack.items[index].number += element.number
	return true
}

// 寻找某一个物品
func (pack *Backpack) findElement(element *Element) int {

	for index := range pack.items {
		if pack.items[index].name == element.name {
			return index
		}
	}
	return -1

}

// 移除物品
func (pack *Backpack) removeElement(element *Element) {
	index := pack.findElement(element)
	if index != -1 {
		return
	}
	pack.items = append(pack.items[:index], pack.items[index+1:]...)
	pack.cap -= 1

}

// 减少物品数量
func (pack *Backpack) reduceNumber(element *Element) {
	WORLD.player.action = ""
	element.number = element.number - 1

	if element.number <= 0 {
		index := pack.findElement(element)
		pack.selectDetail = ""
		if index != -1 {

		}
		pack.items = append(pack.items[:index], pack.items[index+1:]...)
		pack.cap -= 1
		pack.index = MinInt(pack.index-1, 0)

	}

}
