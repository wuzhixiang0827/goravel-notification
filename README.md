# goravel-notification

## 安装

```

go clean -i github.com/wuzhixiang0827/goravel-notification
go run . artisan package:install github.com/wuzhixiang0827/goravel-notification
```

## 卸载

```
go run . artisan package:uninstall github.com/wuzhixiang0827/goravel-notification
```

## 发布资源

```
go run . artisan vendor:publish --package=github.com/wuzhixiang0827/goravel-notification
```

## 注册迁移文件

在 `database/kernel.go` 中添加

```
func (kernel Kernel) Migrations() []schema.Migration {
	return []schema.Migration{
	    // 添加迁移文件
		&migrations.M20251104144615CreateNotificationsTable{},
	}
}
```

执行迁移

```
go run . artisan migrate
```

## 创建通知

app/notifications/verification_code_notification.go

```
package notifications

import "github.com/wuzhixiang0827/goravel-notification/contracts"

type WelcomeNotification struct{}

func (n WelcomeNotification) Via(notifiable contracts.Notifiable) []string {
	return []string{
		"mail",
		"database",
	}
}

func (n WelcomeNotification) ToMail(notifiable contracts.Notifiable) (map[string]interface{}, error) {
	return map[string]interface{}{
		"subject": "Welcome!",
		"content": "Hello and welcome to Goravel!",
	}, nil
}

func (n WelcomeNotification) ToDatabase(notifiable contracts.Notifiable) (map[string]interface{}, error) {
	return map[string]interface{}{
		"title": "Welcome Notification",
		"body":  "You have successfully registered.",
	}, nil
}

```

## 修改user模型

app/models/user.go

```
func (u *User) ParamsForNotification(channel string) any {
	switch channel {
	case "id":
		return convert.MustString(u.ID)
	case "email":
		return u.Account
	}
	return ""
}
```

## 发送通知

```
user := &models.User{Account: request.Email}
notif := notifications.WelcomeNotification{}

if err := notificationFacades.Send(user, notif); err != nil {
    fmt.Println("Send error:", err)
}
```

## 扩展channel

创建通道

```
package channels

import (
	"fmt"
	"github.com/wuzhixiang0827/goravel-notification/contracts"
)

// SlackChannel 自定义 Slack 通道
type SlackChannel struct{}

func (c *SlackChannel) Send(notifiable contracts.Notifiable, notif contracts.Notification) error {
	data, err := notif.ToDatabase(notifiable) // 复用数据结构
	if err != nil {
		return err
	}

	slackWebhook := notifiable.RouteNotificationFor("slack")
	if slackWebhook == "" {
		return fmt.Errorf("[SlackChannel] user has no slack webhook route")
	}

	// 这里只做示例打印，实际可调用 Slack API
	fmt.Printf("[SlackChannel] send to %s: %+v\n", slackWebhook, data)
	return nil
}
```

注册通道

```
notification.RegisterChannel("slack", &channels.SlackChannel{})
```