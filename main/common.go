package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
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

func printLine() {
	for i := 0; i < 80; i++ {
		fmt.Printf("-")
	}
	fmt.Println(" ")
}

func randNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

var now = time.Now()

func getCurrentMinute() int {
	return now.Minute()
}
func MaxInt(left int, right int) int {
	if left < right {
		return right
	}
	return left
}
func MinInt(left int, right int) int {
	if left > right {
		return right
	}
	return left
}
