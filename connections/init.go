package connections

var conn IUser

func init() {
	conn = InitMySql()
}

func GetSqlConn() IUser {
	return conn
}
