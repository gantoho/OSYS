package tools

import (
	"math/rand"
	"regexp"
	"time"
)

func CheckPassword(password string) bool {
	// 检查长度是否在6到20个字符之间
	lengthCheck := regexp.MustCompile(`^.{6,20}$`)
	if !lengthCheck.MatchString(password) {
		return false
	}

	// 检查是否包含至少一个小写字母
	lowercaseCheck := regexp.MustCompile(`[a-z]`)
	if !lowercaseCheck.MatchString(password) {
		return false
	}

	// 检查是否包含至少一个大写字母
	uppercaseCheck := regexp.MustCompile(`[A-Z]`)
	if !uppercaseCheck.MatchString(password) {
		return false
	}

	// 检查是否包含至少一个数字
	digitCheck := regexp.MustCompile(`[0-9]`)
	if !digitCheck.MatchString(password) {
		return false
	}

	// 检查是否包含至少一个下划线或连字符
	specialCharCheck := regexp.MustCompile(`[-_]`)
	if !specialCharCheck.MatchString(password) {
		return false
	}

	return true
}

func Random() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(10) + 1
}
