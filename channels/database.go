package channels

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/json"
	"github.com/wuzhixiang0827/goravel-notification/contracts"
	"github.com/wuzhixiang0827/goravel-notification/models"
)

// DatabaseChannel 默认数据库通道
type DatabaseChannel struct{}

func (c *DatabaseChannel) Send(notifiable contracts.Notifiable, notification contracts.Notification) error {
	data, err := notification.ToDatabase(notifiable)
	if err != nil {
		return err
	}

	jsonData, _ := json.MarshalString(data)

	var notificationModel models.Notification
	notificationModel.ID = uuid.New().String()
	notificationModel.Data = jsonData
	notificationModel.NotifiableId = notifiable.ParamsForNotification("id").(string)
	notificationModel.NotifiableType = fmt.Sprintf("%T", notifiable)
	notificationModel.Type = fmt.Sprintf("%T", notification)

	if err := facades.Orm().Query().Model(&models.Notification{}).Create(&notificationModel); err != nil {
		return err
	}

	return nil
}
