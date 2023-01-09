package main

import (
	"fmt"
)

// Equipment 装备
type Equipment struct {
	// 名字
	name string
	// 详情
	detail string
	// 命令
	command Command
	// 位置
	position string
	// 攻击力
	ATK int
	// 防御力
	DEF int
}

// Command 命令处理
type Command func(player *Player, element *Element, command string)

// Backpack 背包

// Item 物品
type Item struct {
	name   string
	detail string
	number int
	Command
}
type ItemFactory struct {
}

// 实例化木头
func (factory *ItemFactory) instanceWood() {

}

func commandHandler() {
	go func() {
		for {
			command := scanner()
			fmt.Println(command)
		}
	}()
}
