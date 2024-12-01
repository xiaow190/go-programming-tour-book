package word

import (
	"strings"
	"unicode"
)

// 实现功能
// 1.单词全部转成小写
// 2.单词全部转成大写
// 3.下划线单词转为大写驼峰单词
// 4.下划线单词转为小写驼峰单词
// 5. 驼峰单词转为下划线单词

// 1.单词全部转成小写/大写

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

// 2.下画线单词转大写驼峰单词

func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 3.下画线单词转成小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 4.驼峰单词转下画线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
