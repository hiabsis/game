package main

import (
	"fmt"
	"sync"
)

// WordContext 游戏本体
type WordContext struct {
	// 世界地图
	gameMap map[string]BigMap
	// 玩家
	user *User
	// 当前所在位置
	local BigMap

	body []string
}

func (wc *WordContext) showWordMap() {

}

func (wc *WordContext) init() {
	wc.user = &User{username: "新一", level: 1, coin: 1000}
	wc.gameMap = WordMap
	wc.draw()
}
func (wc *WordContext) choosePlace() {
	fmt.Println("选择去的城市：")
	index := 1
	for city := range wc.gameMap {
		fmt.Printf("%v. %v\n", index, city)
	}
	go func() {
		for {
			fmt.Print("请输入城市名称：")
			cmd := scanner()
			city, ok := wc.gameMap[cmd]
			if ok {
				wc.local = city
				wc.body = city.getMapBody()
				wc.draw()
				return
			}
			fmt.Println()
		}

	}()
}
func (wc *WordContext) draw() {
	clearConsole()
	for i := 0; i < 60; i++ {
		fmt.Print("_")
	}
	fmt.Println("")
	fmt.Println(wc.user.toString())

	for i := 0; i < 60; i++ {
		fmt.Print("_")
	}
	fmt.Println("")
	if wc.body == nil {
		wc.choosePlace()
	}
	if wc.local != nil {
		fmt.Printf("地点: %v\n", wc.local.queryMapName())
		for _, b := range wc.local.getMapBody() {
			fmt.Println(b)
		}
	}

	for i := 0; i < 60; i++ {
		fmt.Print("_")
	}
	fmt.Println("")
	fmt.Println("地图: [M]\t背包: [B]\t设置: [C]")
	for i := 0; i < 60; i++ {
		fmt.Print("_")
	}
	fmt.Println("")
	if wc.local != nil {
		fmt.Print("输入前往的位置：")
		wc.local.enter(wc)
		fmt.Println("")
	}

}

// 开始游戏入口
func (wc *WordContext) run(wg *sync.WaitGroup) {
	defer wg.Done()
	wc.init()
	for {

	}

}
