package mysql_model

type UserChatMessage struct {
	Autokid      int    `gorm:"column:autokid;PRIMARY_KEY" json:"autokid"`
	UserId       int    `gorm:"column:user_id" json:"user_id"`
	TargetUserId int    `gorm:"column:target_user_id" json:"target_user_id"`
	Message      string `gorm:"column:message" json:"message"`
	Ctime        int64  `gorm:"column:c_time" json:"c_time"` //创建时间
}
