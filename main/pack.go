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
}

func (pack *Backpack) monitorKeyboard() {
	for {
		cmd := scanner()
		if cmd == "h" || cmd == "H" {
			break
		}
		if pack.cap == 0 {
			pack.show()
			continue
		}
		if cmd == "w" {
			pack.index -= 1
			if pack.index < 0 {
				pack.index = pack.cap
			}
		} else if cmd == "s" {
			pack.index += 1
			if pack.index >= pack.cap {
				pack.index = 0
			}
		}
		pack.show()
	}

}
func (pack *Backpack) show() {
	clearConsole()
	printLine()
	fmt.Printf("背包\t容量: %v\t物品 :%v\t \n", pack.size, pack.cap)
	printLine()
	for i, item := range pack.items {
		if pack.index == i {
			fmt.Println(aurora.BgWhite(fmt.Sprintf("%v. %v\t数量：%v\n ", i, item.name, item.number)))
		}

	}
	printLine()
	fmt.Printf("返回游戏: %v\n", aurora.Red("H"))
	printLine()
}

// 添加物品
func (pack *Backpack) addElement(element *Element) bool {
	if pack.cap == pack.size {
		fmt.Println("背包容量已满")
		printLine()
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
