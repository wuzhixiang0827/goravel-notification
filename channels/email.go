package channels

import (
	"fmt"
	"github.com/wuzhixiang0827/goravel-notification/contracts"
)

// EmailChannel 默认邮件通道
type EmailChannel struct{}

func (c *EmailChannel) Send(notifiable contracts.Notifiable, notif contracts.Notification) error {
	msg, err := notif.ToMail(notifiable)
	if err != nil {
		return err
	}

	email := notifiable.RouteNotificationFor("mail")
	if email == "" {
		return fmt.Errorf("[EmailChannel] notifiable has no email route")
	}

	// TODO: 实际环境下可调用 Goravel 的 mail 组件
	fmt.Printf("[EmailChannel] send mail to %s\nSubject: %s\nBody: %s\n",
		email, msg.Subject, msg.Body)
	return nil
}
