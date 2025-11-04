package notification

import (
	"errors"
	"github.com/goravel/framework/contracts/foundation"
	"goravel/packages/notification/channels"
	"goravel/packages/notification/contracts"
	"sync"
)

type channelRegistry struct {
	mu       sync.RWMutex
	channels map[string]contracts.Channel
}

var registry = &channelRegistry{
	channels: make(map[string]contracts.Channel),
}

// RegisterChannel 允许注册自定义通道
func RegisterChannel(name string, ch contracts.Channel) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	registry.channels[name] = ch
}

// GetChannel 读取通道
func GetChannel(name string) (contracts.Channel, bool) {
	registry.mu.RLock()
	defer registry.mu.RUnlock()
	ch, ok := registry.channels[name]
	return ch, ok
}

// Send 分发通知
func Send(notifiable contracts.Notifiable, notif contracts.Notification) error {
	vias := notif.Via(notifiable)
	if len(vias) == 0 {
		return errors.New("no channels defined for notification")
	}

	for _, chName := range vias {
		ch, ok := GetChannel(chName)
		if !ok {
			return errors.New("channel not registered: " + chName)
		}
		if err := ch.Send(notifiable, notif); err != nil {
			return err
		}
	}
	return nil
}

// Boot 启动时注册默认通道
func Boot(app foundation.Application) {
	RegisterChannel("mail", &channels.EmailChannel{})
	RegisterChannel("database", &channels.DatabaseChannel{})
}
