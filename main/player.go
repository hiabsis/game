package main

import "fmt"

// Player 玩家
type Player struct {
	x int
	y int
	// 名字
	name string
	// 攻击力
	ant int
	// 饥饿值
	hunger int
	// 血量
	hp int
	// 防御力
	define int
	// 最大血量值
	maxHp int
	// 最大饥饿值
	maxHunger int
	// 背包
	pack *Backpack
	// 玩家行动信息
	action string
}

// 显示玩家信息
func (player *Player) show() {
	fmt.Printf("%v\t血量: %v\t饥饿值: %v\t攻击力: %v\t防御力： %v\n", player.name, player.hp, player.hunger, player.ant, player.define)
}

// 玩家普通攻击
func (player *Player) getAnt() int {
	return player.ant
}

// 拾取物品
func (player *Player) pickUp(element *Element) {
	player.action = ""
	ok := WORLD.backpack.addElement(element)
	if ok {
		player.action = fmt.Sprintf("玩家拾取： %v\t 数量：%v\n", element.name, element.number)
		WORLD.addElement(WORLD.elementFactory.instanceSpace(element.x, element.y))
		WORLD.show()
	}
}
