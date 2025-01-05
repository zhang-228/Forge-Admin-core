package email

import (
	"crypto/tls"
)

// Email 邮件配置
type Email struct {
	Test     bool     `json:",default=true"` // 测试邮件
	To       []string `json:",optional"`     // 收件人
	From     string   // 发件人
	Host     string   // 服务器地址 例如 smtp.qq.com smtp.163.com
	Secret   string   // 密钥     用于登录的密钥,去邮箱smtp申请一个用于登录的密钥
	Port     int      `json:",default=25"` // 端口 因默认未开启https所以使用的为25端口
	Nickname string   // 昵称     发件人昵称 通常为自己的邮箱
	IsSSL    bool     `json:",default=false"` // 是否SSL 默认不开启https
}

func SendEmail(conf Email, msg string, to ...string) error {

	if conf.Secret == "" {
		return xerr.NewSystemConfError("未配置邮箱信息,如要使用邮箱功能,请配置邮箱信息后重试")
	}

	EmailClientAuth := gomail.NewDialer(conf.Host, conf.Port, conf.From, conf.Secret)

	var emailTo []string
	if len(to) > 0 {
		emailTo = to
	} else {
		logx.Error("未指定收件人")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", conf.From)
	m.SetHeader("To", emailTo...)
	m.SetHeader("Subject", conf.From)
	m.SetBody("text/html", msg)

	EmailClientAuth.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := EmailClientAuth.DialAndSend(m); err != nil {
		logx.Error("发送邮件失败err:%s", err)
	}

	return nil
}
