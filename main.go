package main

import (
	"chatgpt/common"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	//text, _ := common.AudioToText("/Users/yq/Music/iTunes/iTunes Media/Music/Unknown Artist/Unknown Album/1320_00000.mp3", "zh")
	//fmt.Println(text)

	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	fgHiCyan := color.New(color.FgHiCyan).SprintFunc()
	underline := color.New(color.Underline).SprintFunc()

	for true {

		fmt.Println(green("\n\n####################  请根据所需功能，输入序号，输入 exit 结束对话  ####################"))
		fmt.Println(green("--------------------  1 文字聊天		--------------------"))
		fmt.Println(green("--------------------  2 语音转文字		--------------------"))
		fmt.Println(green("--------------------  3 文字生成图像		--------------------"))

		fmt.Print("-> ")
		text := common.EnterContent()

		if text == "1" {

			// 打印信息，其中 "对话开始，输入 exit 结束对话" 会被着色为绿色
			fmt.Println(green("\n###############  对话开始  ###############"))
			common.ConsoleContextText()

		} else if text == "2" {

			for true {

				fmt.Println(green("\n\n--------------------------------------------------------------------------------"))

				fmt.Println(green("请输入音频文件地址："))
				fmt.Print("-> ")
				scanPath := common.EnterContent()

				fmt.Println(green("请输入源音频文件语言，") + yellow("（英文：en  中文：zh，输入不匹配时将自动校正）"))
				fmt.Print("-> ")
				language := common.EnterContent()
				fmt.Println(green("语言识别中，请稍后"))
				toText, err := common.AudioToText(scanPath, language)
				if err != nil {
					fmt.Println(red("识别发生错误，请重启程序"))
					return
				}
				fmt.Println(green("语音识别内容为：\n\n ") + fgHiCyan(underline(toText)) + "\n")

				fmt.Println(green("-------------------------- 按下回车键继续使用，输入0返回上一层 --------------------------"))
				fmt.Print("-> ")
				backToLevel := common.EnterContent()
				if backToLevel == "0" {
					break
				}
			}

		} else if text == "3" {

			for true {
				fmt.Println(green("\n\n--------------------------------------------------------------------------------"))
				fmt.Println(green("请用文字描述你想生成的图片内容："))
				fmt.Print("-> ")
				describe := common.EnterContent()

				fmt.Println(green("请输入图像大小，格式:256x256，直接回车默'256x256' :"))
				fmt.Print("-> ")
				imageSize := common.EnterContent()
				if imageSize == "" {
					imageSize = "256x256"
				}

				link, _ := common.DallE2Link(describe, imageSize)

				fmt.Println(fgHiCyan("图像生成URL，复制到浏览器查看：\n") + green(underline(link)) + "\n")

				fmt.Println(green("-------------------------- 按下回车键继续使用，输入0返回上一层 --------------------------"))
				fmt.Print("-> ")
				backToLevel := common.EnterContent()
				if backToLevel == "0" {
					break
				}
			}
		} else {
			fmt.Println(red("输入序号不正确！程序即将退出！"))
			return
		}

	}
}
