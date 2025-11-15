package contracts

// Channel 定义通道接口。通道实现接收 notifiable 与 notification（具体类型不限），自己负责调用 notification 中合适的 ToXxx 方法
type Channel interface {
	// notification 参数类型为任意实现了 contracts.Notification 的值（接口本身仅包含 Via）
	Send(notifiable Notifiable, notification interface{}) error
}
