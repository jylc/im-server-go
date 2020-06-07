package mysql_model

type UserFriend struct {
	Autokid       int   `gorm:"column:autokid;PRIMARY_KEY" json:"autokid"`
	RequestUserId int   `gorm:"column:request_user_id" json:"request_user_id"`
	AcceptUserId  int   `gorm:"column:accept_user_id" json:"accept_user_id"`
	Ctime         int64 `gorm:"column:c_time" json:"c_time"` //创建时间
	DTime         int64 `gorm:"column:d_time" json:"d_time"` //删除时间
	Status        int   `gorm:"column:status" json:"status"` //状态
}
