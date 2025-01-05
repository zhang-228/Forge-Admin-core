package utils

import (
	"regexp"
)

var regex = `^1[3456789]\d{9}$` // 大陆手机号正则

// IsPhone 正则验证手机号
func IsPhone(phone string) bool {
	re := regexp.MustCompile(regex)
	if re.MatchString(phone) {
		return true
	} else {
		return false
	}
}

// IsEmail 验证邮箱
func IsEmail(email string) bool {
	if email[len(email)-4:] != ".com" {
		return false
	}
	return true
}
