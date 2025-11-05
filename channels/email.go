package channels

import (
	"fmt"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/mail"
	"github.com/wuzhixiang0827/goravel-notification/contracts"
)

// EmailChannel 默认邮件通道
type EmailChannel struct{}

func (c *EmailChannel) Send(notifiable contracts.Notifiable, notif contracts.Notification) error {
	msg, err := notif.ToMail(notifiable)
	if err != nil {
		return err
	}

	email := notifiable.ParamsForNotification("mail").(string)
	if email == "" {
		return fmt.Errorf("[EmailChannel] notifiable has no email route")
	}

	content := msg["content"].(string)
	subject := msg["subject"].(string)

	if err := facades.Mail().To([]string{email}).
		Content(mail.Html(content)).
		Subject(subject).
		Queue(); err != nil {
		return err
	}

	return nil
}
