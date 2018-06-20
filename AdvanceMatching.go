package goexpr

import (
	"regexp"
)

func CompilereFindAllStringSubmatch(reg, value string) (res [][]string, err error) {
	re, err := regexp.Compile(reg)

	if err != nil {
		return
	}

	res = re.FindAllStringSubmatch(value, -1)

	return
}

func CompileStringSubmatch(reg, value string) (res []string, err error) {
	re, err := regexp.Compile(reg)
	if err != nil {
		return
	}

	res = re.FindStringSubmatch(value)
	return
}
