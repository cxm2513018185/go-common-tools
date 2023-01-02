package word

import (
	"strings"
	"unicode"
)

// ToUpper 单词全部转为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 单词全部转为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase 下划线单词转为大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)   // 下划线替换成空格字符
	s = strings.Title(s)                   // 所有字符修改为其对应的首字母大写形式
	return strings.Replace(s, " ", "", -1) // 把空格字符替换为空
}

// UnderscoreToLowerCamelCase 下划线单词转为小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore 驼峰单词转下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune

	for index, value := range s {
		if index == 0 {
			output = append(output, unicode.ToLower(value))
			continue
		}
		if unicode.IsUpper(value) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(value))
	}

	return string(output)
}
