package mysql_model

type UserBlackList struct {
	Autokid     int   `gorm:"column:autokid;PRIMARY_KEY" json:"autokid"`
	UserId      int   `gorm:"column:user_id" json:"user_id"`
	BlockUserId int   `gorm:"column:block_user_id" json:"block_user_id"`
	Ctime       int64 `gorm:"column:c_time" json:"c_time"` //创建时间
	DTime       int64 `gorm:"column:d_time" json:"d_time"` //删除时间
}
