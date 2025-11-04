package facades

import (
	notification "github.com/wuzhixiang0827/goravel-notification"
	"github.com/wuzhixiang0827/goravel-notification/contracts"
)

// Send 发送通知
func Send(notifiable contracts.Notifiable, notif contracts.Notification) error {
	return notification.Send(notifiable, notif)
}
