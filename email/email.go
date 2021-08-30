package email

import "gopkg.in/gomail.v2"

/*
MAIL_DRIVER=smtp
MAIL_HOST=smtp.ym.163.com
MAIL_PORT=465
MAIL_USERNAME=service@syyai.com
MAIL_PASSWORD=yongjie0203
MAIL_ENCRYPTION=ssl
MAIL_FROM_ADDRESS=service@syyai.com
MAIL_FROM_NAME=来加速
*/

// 邮件服务器地址
var MAIL_HOST = "smtp.ym.163.com"

// 端口
var MAIL_PORT = 465

// 发送邮件用户账号
var MAIL_USER = "service@syyai.com"

// 授权密码
var MAIL_PWD = "yongjie0203"

var MAIL_FROM_NAME = "交易平台"

func SendMail(mailAddress []string, subject string, body string) error {
	m := gomail.NewMessage()
	// 这种方式可以添加别名，即“126 mail”， 也可以直接用<code>m.SetHeader("From", MAIL_USER)</code>
	m.SetHeader("From", MAIL_USER)
	// 发送给多个用户
	m.SetHeader("Bcc", mailAddress...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(MAIL_HOST, MAIL_PORT, MAIL_USER, MAIL_PWD)
	// 发送邮件
	err := d.DialAndSend(m)
	return err
}
