package channels

import (
	"fmt"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/mail"
	"github.com/wuzhixiang0827/goravel-notification/contracts"
)

// EmailChannel 默认邮件通道
type EmailChannel struct{}

func (c *EmailChannel) Send(notifiable contracts.Notifiable, notif interface{}) error {
	data, err := CallToMethod(notif, "toEmail", notifiable)
	if err != nil {
		return fmt.Errorf("[EmailChannel] notifiable has no email")
	}

	email := notifiable.RouteNotificationFor("email").(string)
	if email == "" {
		return fmt.Errorf("[EmailChannel] notifiable has no email")
	}

	content := data["content"].(string)
	subject := data["subject"].(string)

	if err := facades.Mail().To([]string{email}).
		Content(mail.Html(content)).
		Subject(subject).
		Queue(); err != nil {
		return err
	}

	return nil
}
