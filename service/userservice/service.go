package userservice

import (
	"fmt"
	"log"
	"tm-user/authentication"
	"tm-user/dao/userdao"
	"tm-user/global"
	"tm-user/model/user"

	"gorm.io/gorm"
)

const (
	loginError = "-1"
)

var (
	UserRepo UserRepoInterface = &userSQL{}
)

type UserRepoInterface interface {
	Initialize(*gorm.DB)
	Create(user.User) (user.UserInfo, error)
	Login(user.User) (user.UserInfo, error)
}

type userSQL struct {
	db *gorm.DB
}

// 初始化
func (um *userSQL) Initialize(db *gorm.DB) {
	um.db = db
}

// CreatUser 建立使用者
func (um *userSQL) Create(user user.User) (userInfo user.UserInfo, err error) {
	tx := um.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		if err != nil {
			log.Println(err)
		}
	}()

	// test
	fmt.Println("111")
	// 檢查使用者是否以創建
	userArr, err := userdao.GetUserByUsername(user, tx)
	if err != nil {

		panic(err)
	}
	// test
	fmt.Println("222")

	if len(userArr) != 0 {
		userInfo.JwtToken = global.LoginError
		return
	}
	// test
	fmt.Println("333")
	userInfo.UserID, err = userdao.InsertUser(user, tx)
	if err != nil {
		panic(err)
	}
	// test
	fmt.Println("4444", userInfo.UserID)
	tx.Commit()
	// 產生令牌
	userInfo.JwtToken, err = authentication.GenerateToken(user)
	if err != nil {
		userInfo.JwtToken = global.LoginError
		return
	}
	return
}

// LoginUser 登入使用者
func (um *userSQL) Login(user user.User) (userInfo user.UserInfo, err error) {
	tx := um.db.Begin()
	userTmp, err := userdao.GetUserbyNameAndPassword(user, tx)
	if err != nil {
		log.Println(err)
		return
	}

	if userTmp.UserName == user.UserName && userTmp.Password == user.Password {
		// 帳號密碼符合

		// 產生 jwt
		userInfo.JwtToken, err = authentication.GenerateToken(user)
		if err != nil {
			log.Println(err)
			return
		}

	} else {
		userInfo.JwtToken = global.LoginError
	}
	userInfo.UserID = userTmp.ID
	return
}
