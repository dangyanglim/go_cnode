package mail

import (
	"fmt"
	//"log"
	//"net/http"
	"net/smtp"
	//"regexp"
	"strings"

	"github.com/dangyanglim/go_cnode/mgoModels"
	//"github.com/gin-gonic/gin"
	//"github.com/tommy351/gin-sessions"
	"crypto/tls"

	"log"
	"net"
)

var userModel = new(models.UserModel)

func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Dialing Error:", err)
		return nil, err
	}

	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func SendMailViaTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {

	c, err := Dial(addr)
	if err != nil {
		log.Println("Create SMTP Client fail:", err)
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("AUTH ERROR ", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}

func SendActiveMail(who string, token string, name string) {
	auth := smtp.PlainAuth("", "dangyanglim@qq.com", "uicmeimalcnybgdj", "smtp.qq.com")
	to := []string{who}
	nickname := name
	user := "dangyanglim@qq.com"
	subject := "Go_Cnode社区账号激活"
	content_type := "Content-Type: text/html; charset=UTF-8"
	body := "<p>您好：" + name + "</p>" +
		"<p>我们收到您在Go_Cnode社区的注册信息，请点击下面的链接来激活帐户：</p>" +
		"<a href  ='http://fenghuangyu.cn:9035/active_account?key=" + token + "&name=" + name + "'>激活链接</a>" +
		"<p>若您没有在Go_Cnode社区填写过注册信息，说明有人滥用了您的电子邮箱，请删除此邮件，我们对给您造成的打扰感到抱歉。</p>" +
		"<p>Go_Cnode社区 谨上。</p>"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	//err := smtp.SendMail("smtp.qq.com:465", auth, user, to, msg)
	err := SendMailViaTLS(
		"smtp.qq.com:465",
		auth,
		user,
		to,
		msg,
	)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
}
