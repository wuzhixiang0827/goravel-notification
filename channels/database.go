package channels

import (
	"encoding/json"
	"fmt"
	"goravel/packages/notification/contracts"
	"time"
)

// DatabaseChannel 默认数据库通道
type DatabaseChannel struct{}

func (c *DatabaseChannel) Send(notifiable contracts.Notifiable, notif contracts.Notification) error {
	data, err := notif.ToDatabase(notifiable)
	if err != nil {
		return err
	}

	jsonData, _ := json.Marshal(data)

	// TODO: 可使用 Goravel ORM 保存到 notifications 表
	fmt.Printf("[DatabaseChannel] store for %s at %s: %s\n",
		notifiable.RouteNotificationFor("database"),
		time.Now().Format(time.RFC3339),
		string(jsonData),
	)
	return nil
}
