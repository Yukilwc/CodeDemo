package reflecttool

import (
	"database/sql"
	"time"
)

type EventNotification struct {
	Id          int64         `db:"id"`
	Name        string        `db:"name"`         // 角色名
	Calendar    int64         `db:"calendar"`     // 历法 1公历 2农历
	SolarDate   sql.NullTime  `db:"solar_date"`   // 公历日期
	ChineseDate sql.NullTime  `db:"chinese_date"` // 农历日期
	Time        string        `db:"time"`         // 时间
	Rule        int64         `db:"rule"`         // 规则
	Loop        int64         `db:"loop"`         // 是否循环
	Emails      string        `db:"emails"`
	CreateAt    time.Time     `db:"create_at"`
	UpdateAt    time.Time     `db:"update_at"`
	CreateBy    sql.NullInt64 `db:"create_by"`
	UpdateBy    sql.NullInt64 `db:"update_by"`
	DeleteAt    sql.NullTime  `db:"delete_at"`
}
type EventNotificationResponse struct {
	Id           int64  `db:"id"`
	Name         string `db:"name"`         // 角色名
	Calendar     int64  `db:"calendar"`     // 历法 1公历 2农历
	SolarDate    *int64 `db:"solar_date"`   // 公历日期
	ChineseDate  *int64 `db:"chinese_date"` // 农历日期
	Time         string `db:"time"`         // 时间
	Rule         int64  `db:"rule"`         // 规则
	Loop         int64  `db:"loop"`         // 是否循环
	Emails       string `db:"emails"`
	CreateAt     int64  `db:"create_at"`
	UpdateAt     int64  `db:"update_at"`
	CreateBy     int64  `db:"create_by"`
	CreateByName string `db:"create_by"`
	UpdateBy     int64  `db:"update_by"`
	UpdateByName string `db:"create_by"`
	DeleteAt     *int64 `db:"delete_at"`
}
