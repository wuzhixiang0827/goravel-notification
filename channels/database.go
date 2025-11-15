package channels

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/json"
	"github.com/goravel/framework/support/str"
	"github.com/wuzhixiang0827/goravel-notification/contracts"
	"github.com/wuzhixiang0827/goravel-notification/models"
)

// DatabaseChannel 默认数据库通道
type DatabaseChannel struct{}

func (c *DatabaseChannel) Send(notifiable contracts.Notifiable, notif interface{}) error {
	data, err := CallToMethod(notif, "ToDatabase", notifiable)
	if err != nil {
		return fmt.Errorf("[DatabaseChannel] %s", err.Error())
	}

	jsonData, _ := json.MarshalString(data)

	var notificationModel models.Notification
	notificationModel.ID = uuid.New().String()
	notificationModel.Data = jsonData
	notificationModel.NotifiableId = notifiable.RouteNotificationFor("id").(string)
	notificationModel.NotifiableType = str.Of(fmt.Sprintf("%T", notifiable)).Replace("*", "").String()
	notificationModel.Type = fmt.Sprintf("%T", notif)

	if err := facades.Orm().Query().Model(&models.Notification{}).Create(&notificationModel); err != nil {
		return err
	}

	return nil
}
