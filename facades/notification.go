package facades

import (
	"goravel/packages/notification"
	"goravel/packages/notification/contracts"
)

// Send 发送通知
func Send(notifiable contracts.Notifiable, notif contracts.Notification) error {
	return notification.Send(notifiable, notif)
}
