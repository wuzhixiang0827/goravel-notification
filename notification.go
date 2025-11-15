package notification

import (
	"errors"
	"fmt"
	"github.com/goravel/framework/contracts/foundation"
	"strings"
	"sync"

	"github.com/wuzhixiang0827/goravel-notification/channels"
	"github.com/wuzhixiang0827/goravel-notification/contracts"
)

// channelRegistry 管理已注册通道（线程安全）
type channelRegistry struct {
	mu       sync.RWMutex
	channels map[string]contracts.Channel
}

var registry = &channelRegistry{
	channels: make(map[string]contracts.Channel),
}

// RegisterChannel 允许用户在应用启动时注册自定义通道（注册一次即可）
func RegisterChannel(name string, ch contracts.Channel) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	registry.channels[strings.ToLower(name)] = ch
}

// GetChannel 获取已注册通道
func GetChannel(name string) (contracts.Channel, bool) {
	registry.mu.RLock()
	defer registry.mu.RUnlock()
	ch, ok := registry.channels[strings.ToLower(name)]
	return ch, ok
}

// Send 将通知分发到 Via() 返回的所有通道
func Send(notifiable contracts.Notifiable, notif contracts.Notification) error {
	vias := notif.Via(notifiable)
	if len(vias) == 0 {
		return errors.New("no channels defined for notification")
	}

	for _, chName := range vias {
		ch, ok := GetChannel(chName)
		if !ok {
			return fmt.Errorf("channel not registered: %s", chName)
		}
		if err := ch.Send(notifiable, notif); err != nil {
			return fmt.Errorf("channel %s send error: %w", chName, err)
		}
	}

	return nil
}

// Boot: 注册内置默认通道（mail, database）
func Boot(app foundation.Application) {
	RegisterChannel("email", &channels.EmailChannel{})
	RegisterChannel("database", &channels.DatabaseChannel{})
}
