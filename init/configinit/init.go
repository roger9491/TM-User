package configinit

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"tm-user/authentication"
	"tm-user/global"
	"tm-user/global/database"

	"github.com/joho/godotenv"
)

func init() {

	var (
	_, b, _, _ = runtime.Caller(0)

	projectRootPath = filepath.Join(filepath.Dir(b), "../../")
	)

	err := godotenv.Load(os.ExpandEnv(projectRootPath + "/.env"))
	if err != nil {
		log.Printf("Error getting env %v\n", err)
	}
	
	database.DBName = os.Getenv("DB_DBNAME")
	database.Host = os.Getenv("DB_HOST")
	database.Port = os.Getenv("DB_PORT")
	database.Username = os.Getenv("DB_USERNAME")
	database.Password = os.Getenv("DB_PASSWORD")

	authentication.Secret = []byte(os.Getenv("AUTH_SECRET"))

	global.IP = os.Getenv("HOST_IP")
	global.Port = os.Getenv("HOST_PORT")

}
