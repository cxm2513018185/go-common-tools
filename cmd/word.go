package cmd

import (
	"github.com/spf13/cobra"
	"go-common-tools/internal/word"
	"log"
	"strings"
)

// 定义目前单词所支持的转换模式枚举值
const (
	ModeUpper                      = iota + 1 // 全部单词转为大写
	ModeLower                                 // 全部单词转为小写
	ModeUnderscoreToUpperCamelCase            // 下划线单词转为大写驼峰单词
	ModeUnderscoreToLowerCamelCase            // 下划线单词转为小写驼峰单词
	ModeCamelCaseToUnderscore                 // 驼峰单词转下划线单词
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部单词转为大写",
	"2：全部单词转为小写",
	"3：下划线单词转为大写驼峰单词",
	"4：下划线单词转为小写驼峰单词",
	"5：驼峰单词转下划线单词",
}, "\n")

var str string
var mode int8

var wordCmd = &cobra.Command{
	Use:   "word",   // 子命令的命令标识
	Short: "单词格式转换", // 简短说明，在help命令输出的帮助信息中展示
	Long:  desc,     // 完整说明，在help命令输出的帮助信息中展示
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help word 查看帮助文档")
		}

		log.Printf("输出结果：%s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
