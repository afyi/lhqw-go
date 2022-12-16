package main

import (
	"fmt"
	"os"

	core "github.com/afyi/lhqw-go/core"
)

var Usage = func() {
	fmt.Println("\n莲花清瘟配方加解密工具 v1.0.0 by xafyi [https://github.com/afyi/lhqwcore]")
	fmt.Println("\n命令格式: lhqw <命令> <要加密或解密的串>")
	fmt.Println("\n命令包含以下两种:\n\td\t要解密的文本.\n\te\t需要加密的文本")
}

func main() {
	args := os.Args[1:]

	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	switch args[0] {
	case "e":
		fmt.Println()
		fmt.Println("您的字符串加密结果是:")
		fmt.Println("==============================================")
		fmt.Println(core.Encode(string(args[1])))
		fmt.Println()
	case "d":
		fmt.Println()
		fmt.Println("您的字符串解密结果是:")
		fmt.Println("==============================================")
		fmt.Println(core.Decode(string(args[1])))
		fmt.Println()
	default:
		Usage()
	}
}
