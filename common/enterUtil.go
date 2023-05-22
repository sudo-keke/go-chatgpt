package common

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func VerifyExit(text string) {
	if text == "exit" {
		fgGreen := color.New(color.FgGreen).SprintFunc()
		fmt.Println("###############  ", fgGreen("对话结束"), "  ###############")
		return
	}
}

// EnterContent 获取输入的内容
func EnterContent() string {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// 将 CRLF 转换为 LF
	text = strings.Replace(text, "\n", "", -1)

	if text == "exit" {
		FgGreen := color.New(color.FgGreen).SprintFunc()
		fmt.Println(FgGreen("####################  对话结束  ####################"))
		os.Exit(0) // 退出程序
	}
	return text
}
