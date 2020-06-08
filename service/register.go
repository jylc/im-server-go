package service

import (
	dbConn "im-server-go/connections"
	apiResp "im-server-go/domain"
	"log"
)

func Register(telephone, pwd string) *apiResp.ApiResponse {
	conn := dbConn.GetSqlConn()
	response := &apiResp.ApiResponse{}
	if conn == nil {
		log.Println("[Service Register] Get Sql connection error")
		return response.Error("Get Sql connection error", nil)
	}
	if succ := conn.UserSignUp(telephone, pwd); succ {
		return response.Success("Register success!!!")
	} else {
		return response.Error("Register fail!!!", nil)
	}

}
