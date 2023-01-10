package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"math/rand"
)

type Element struct {
	// 名称
	name string
	// 详情
	detail string
	// 命令
	command Command
	// 位置
	x int
	y int
	// 生命力
	hp int
	// 攻击力
	ant int
	// 防御力
	define int
	show   func(element *Element)
	// 数量
	number int
}

// ElementFactory 元素工厂
type ElementFactory struct {
}

// 生成地图元素
func (factory *ElementFactory) generateElement(x int, y int) *Element {
	//rand.Seed(time.Now().UnixNano())
	//// 表示生成 [0,50)之间的随机数
	num := rand.Intn(2)
	if num == 0 {
		return factory.instanceTree(x, y)
	} else if num == 1 {
		return factory.instancePig(x, y)
	}

	return factory.instanceSpace(x, y)

}

// 生成树元素
func (factory *ElementFactory) instanceTree(x int, y int) *Element {
	return &Element{
		name: fmt.Sprintf("%v", aurora.Green("树")),
		command: func(player *Player, element *Element, command string) {
			if command == "c" || command == "C" {
				if element.hp > 0 {
					element.hp -= 10
					if element.hp <= 0 {
						element = factory.instanceWood(element.x, element.y)
						WORLD.addElement(element)
						WORLD.show()
					}
				}
			}
		},
		x: x,
		y: y,

		hp:     randNum(10, 20),
		define: 0,
		show: func(element *Element) {
			fmt.Printf("位置(%v,%v)", aurora.Red(element.x), aurora.Red(element.y))
			fmt.Printf("\t名称：%v", element.name)
			fmt.Printf("\t生命力:%v\n", aurora.Green(element.hp))
			fmt.Printf("\t\t这是一颗松树,输入%v砍伐\n", aurora.Green("c"))
		},
	}
}
func (factory *ElementFactory) instancePig(x int, y int) *Element {
	return &Element{
		name:   "猪",
		detail: "",
		command: func(player *Player, element *Element, command string) {
			element.detail = ""
			if command == "c" || command == "C" {
				hurt := MaxInt(1, player.ant-element.define)
				element.hp -= hurt
				player.action = ""
				player.action += fmt.Sprintf("玩家攻击%v造成%v伤害\n", element.name, hurt)
				if element.hp <= 0 {
					player.action += fmt.Sprintf("玩家击杀野猪野猪\n")
					WORLD.addElement(factory.instancePigMeat(element.x, element.y))
					return
				}
				hurt = MaxInt(1, element.ant-player.define)
				element.detail += fmt.Sprintf("%v反击造成%v伤害\n", element.name, hurt)
				player.hp -= hurt
				if player.hp == 0 {
					player.action += fmt.Sprintf("你已被%v击杀\n", aurora.Red(element.name))
				}

			}
		},
		x:      x,
		y:      y,
		hp:     randNum(20, 50),
		define: randNum(3, 7),
		ant:    randNum(3, 5),
		show: func(element *Element) {
			fmt.Printf("位置(%v,%v)", aurora.Red(element.x), aurora.Red(element.y))
			fmt.Printf("\t名称：%v\n", aurora.Magenta(element.name))
			fmt.Printf("\t\t生命力:%v\t攻击力%v \t防御力%v\n", element.hp, element.ant, element.define)
			fmt.Printf("\t\t这是一只野猪,输入%v攻击\n", aurora.Green("c"))

		},
	}
}
func (factory *ElementFactory) instanceSpace(x int, y int) *Element {
	return &Element{
		name:   fmt.Sprintf("%v\t", aurora.Green("空地")),
		detail: fmt.Sprintf(""),
		command: func(player *Player, element *Element, command string) {
			if command == "c" || command == "C" {
				player.action = ""
			}
			return
		},
		x: x,
		y: y,
		show: func(element *Element) {
			fmt.Printf("位置(%v,%v)", aurora.Red(element.x), aurora.Red(element.y))
			fmt.Printf("\t名称：%v\n", element.name)

		},
	}
}
func (factory *ElementFactory) instancePigMeat(x int, y int) *Element {
	return &Element{
		name:   fmt.Sprintf("%v\t", aurora.Green("猪肉")),
		detail: fmt.Sprintf("猪肉,可以恢复一定的血量和减少饥饿值,输入%v使用物品\n", aurora.Green("U")),
		command: func(player *Player, element *Element, command string) {
			if command == "c" || command == "C" {
				player.pickUp(element)
			}
			if command == "U" || command == "u" {
				player.action = ""
				WORLD.backpack.reduceNumber(element)
				player.action += fmt.Sprintf("玩家使用%v 数量：%v\n", aurora.Green(element.name), aurora.Red("1"))
				recoverHp := randNum(5, 20)
				recoverHungry := randNum(5, 30)
				player.action += fmt.Sprintf("恢复血量%v,减少%v饥饿值\n", aurora.Green(recoverHp), aurora.Green(recoverHungry))

				player.hp = MinInt(recoverHp+player.hp, player.maxHp)
				player.hunger = MaxInt(0, player.hunger-recoverHungry)
				WORLD.backpack.show()
			}
		},
		x: x,
		y: y,
		show: func(element *Element) {
			fmt.Printf("位置(%v,%v)", aurora.Red(element.x), aurora.Red(element.y))
			fmt.Printf("\t名称：%v\n", element.name)
			fmt.Printf("\t\t这是一堆的猪肉：%v 输入%v拾取\n", element.name, aurora.Green("c"))

		},
		number: randNum(2, 5),
	}
}
func (factory *ElementFactory) instanceWood(x int, y int) *Element {
	return &Element{
		name:   fmt.Sprintf("%v\t", aurora.Yellow("木头")),
		detail: fmt.Sprintf("木头可以作为建筑的%v和%v\n", aurora.Yellow("基础材料"), aurora.Yellow("燃料")),
		command: func(player *Player, element *Element, command string) {
			if command == "c" || command == "C" {
				player.pickUp(element)
			}

		},
		x: x,
		y: y,
		show: func(element *Element) {
			fmt.Printf("位置(%v,%v)", aurora.Red(element.x), aurora.Red(element.y))
			fmt.Printf("\t名称：%v\n", element.name)
			fmt.Printf("\t\t这是一堆的木头,输入%v拾取\n", aurora.Green("c"))
		},
		number: randNum(2, 5),
	}
}
