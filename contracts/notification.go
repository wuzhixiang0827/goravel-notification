package contracts

// Notifiable 表示可被通知的对象，用户在自己的模型中实现
type Notifiable interface {
	// RouteNotificationFor(channel) 返回该通道所需的路由信息（如邮箱、手机号、webhook 地址等）
	RouteNotificationFor(channel string) any
}

// Notification 插件约定的唯一接口（用户不能改）
type Notification interface {
	// Via 返回通道名称列表（小写或任意大小写）
	Via(notifiable Notifiable) []string
}
