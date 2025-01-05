package verifyCode

import (
	"math/rand"
	"time"
)

const codeLength = 6         // 验证码长度
const charset = "0123456789" // 可以根据需要自定义字符集

func MakeSmsCode() string {

	// 创建随机数生成器
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)

	// 生成随机字符串
	randomString := make([]byte, codeLength)
	for i := 0; i < codeLength; i++ {
		randomString[i] = charset[rnd.Intn(len(charset))]
	}
	return string(randomString)
}
