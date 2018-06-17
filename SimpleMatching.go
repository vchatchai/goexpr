package goexpr

import (
	"regexp"
)

func CompileMatchString(rex, value string) (b bool, err error) {
	r, err := regexp.Compile(rex)

	b = r.MatchString(value)
	return
}

func CompileFindString(rex, value string) (result string, err error) {
	r, err := regexp.Compile(rex)

	result = r.FindString(value)
	return
}

func CompileMustMatchString(rex, value string) (b bool) {
	r := regexp.MustCompile(rex)

	b = r.MatchString(value)
	return
}

func CompilePosixMustMatchString(rex, value string) (b bool) {
	r := regexp.MustCompilePOSIX(rex)
	b = r.MatchString(value)
	return
}
