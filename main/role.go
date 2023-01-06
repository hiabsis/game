package main

import "fmt"

type User struct {
	username string
	level    int
	coin     int
	hp       int
	mp       int
}

func (user *User) toString() string {
	return fmt.Sprintf("姓名：%v\t金钱:%v\t等级:%v\t\n血条：%v\n蓝量: %v", user.username, user.coin, user.level, user.hp, user.mp)
}
