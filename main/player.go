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
}

func (player *Player) show() {
	fmt.Printf("角色\t血量: %v\t饥饿值: %v\t攻击力: %v\t防御力： %v\n", player.hp, player.hunger, player.ant, player.define)
}
func (player *Player) getAnt() int {
	return player.ant
}

// 拾取物品
func (player *Player) pickUp(element *Element) {
	ok := WORLD.backpack.addElement(element)
	if ok {
		WORLD.addElement(WORLD.elementFactory.instanceSpace(element.x, element.y))
		WORLD.show()
	}
}
