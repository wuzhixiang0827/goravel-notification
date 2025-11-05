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