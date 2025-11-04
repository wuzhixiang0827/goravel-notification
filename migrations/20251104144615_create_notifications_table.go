package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251104144615CreateNotificationsTable struct{}

// Signature The unique signature for the migration.
func (r *M20251104144615CreateNotificationsTable) Signature() string {
	return "20251104144615_create_notifications_table"
}

// Up Run the migrations.
func (r *M20251104144615CreateNotificationsTable) Up() error {
	if !facades.Schema().HasTable("notifications") {
		return facades.Schema().Create("notifications", func(table schema.Blueprint) {
			table.ID()
			table.String("type").Nullable().Comment("通知类型")
			table.String("notifiable_type").Nullable().Comment("通知对象类型")
			table.UnsignedBigInteger("notifiable_id").Nullable().Comment("通知对象ID")
			table.Text("data").Nullable().Comment("通知数据")
			table.Timestamp("read_at").Nullable().Comment("已读时间")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20251104144615CreateNotificationsTable) Down() error {
	return facades.Schema().DropIfExists("notifications")
}
