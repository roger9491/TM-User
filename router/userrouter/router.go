package userrouter

import (
	"fmt"
	"net/http"
	"tm-user/global"
	"tm-user/model/user"
	"tm-user/service/userservice"

	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func UserApi(e *gin.Engine) {
	UserGroup := e.Group("/api")
	{
		UserGroup.PUT("/user", PutUser)
		UserGroup.POST("/login", Login)
	}

}

// putUser 新增使用者
func PutUser(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}

	var user user.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// test
	fmt.Println("111")
	userInfo, err := userservice.UserRepo.Create(user)
	if err != nil || userInfo.JwtToken == global.LoginError {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// test
	fmt.Println("222", userInfo)
	c.JSON(http.StatusCreated, userInfo)

}

func Login(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}

	var user user.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userInfo, err := userservice.UserRepo.Login(user)
	if err != nil || userInfo.JwtToken == global.LoginError {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	c.JSON(http.StatusCreated, userInfo)
}
