package routerinit

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Option func(engine *gin.Engine)

var options []Option

func Include(optObjs ...Option) {

	options = append(options, optObjs...)
}

func InitRouters() *gin.Engine {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "X-Requested-With","Content-Type", "Accept", "Authorization", "UserID" },
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	  }))
	for _, opt := range options {
		opt(router)
	}

	return router
}
