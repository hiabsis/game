package main

import "sync"

func main() {

	context := WordContext{}
	var wg sync.WaitGroup
	wg.Add(1)
	context.run(&wg)

	wg.Wait()

	//// 简单快速的使用，跟 fmt.Print* 类似
	//color.Redp("Simple to use color")
	//color.Redln("Simple to use color")
	//color.Greenp("Simple to use color\n")
	//color.Cyanln("Simple to use color")
	//color.Yellowln("Simple to use color")
	//
	//// 简单快速的使用，跟 fmt.Print* 类似
	//color.Red.Println("Simple to use color")
	//color.Green.Print("Simple to use color\n")
	//color.Cyan.Printf("Simple to use %s\n", "color")
	//color.Yellow.Printf("Simple to use %s\n", "color")
	//
	//// use like func
	//red := color.FgRed.Render
	//green := color.FgGreen.Render
	//fmt.Printf("%s line %s library\n", red("Command"), green("color"))
	//
	//// 自定义颜色
	//color.New(color.FgWhite, color.BgBlack).Println("custom color style")
	//
	//// 也可以:
	//color.Style{color.FgCyan, color.OpBold}.Println("custom color style")
	//
	//// internal style:
	//color.Info.Println("message")
	//color.Warn.Println("message")
	//color.Error.Println("message")
	//
	//// 使用内置颜色标签
	//color.Print("<suc>he</><comment>llo</>, <cyan>wel</><red>come</>\n")
	//// 自定义标签: 支持使用16色彩名称，256色彩值，rgb色彩值以及hex色彩值
	//color.Println("<fg=11aa23>he</><bg=120,35,156>llo</>, <fg=167;bg=232>wel</><fg=red>come</>")
	//
	//// apply a style tag
	//color.Tag("info").Println("info style text")
	//
	//// prompt message
	//color.Info.Prompt("prompt style message")
	//color.Warn.Prompt("prompt style message")
	//
	//// tips message
	//color.Info.Tips("tips style message")
	//color.Warn.Tips("tips style message")
}
