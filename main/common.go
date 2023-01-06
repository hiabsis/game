package main

import (
	"bufio"
	"os"
	"os/exec"
)

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func scanner() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

type BigMap interface {
	getMapBody() []string
	// 进入地图
	enter(wc *WordContext)
	// 退出地图
	exit()
	queryMapName() string
}

// WordMap 世界地图
var WordMap = map[string]BigMap{"烈焰城": &VillageMap{name: "烈焰城"}}
