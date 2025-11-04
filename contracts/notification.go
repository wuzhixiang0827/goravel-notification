package contracts

// Notifiable 表示可被通知的对象
type Notifiable interface {
	RouteNotificationFor(channel string) string
}

// Notification 定义通知的通用接口
type Notification interface {
	Via(notifiable Notifiable) []string
	ToMail(notifiable Notifiable) (*MailMessage, error)
	ToDatabase(notifiable Notifiable) (map[string]interface{}, error)
}

// Channel 定义通道接口
type Channel interface {
	Send(notifiable Notifiable, notif Notification) error
}

// MailMessage 是邮件通道消息结构
type MailMessage struct {
	Subject string
	Body    string
}
