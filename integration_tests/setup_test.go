package integration__tests

import (
	"log"
	"os"
	"testing"
	"tm-user/init/sqlinit"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const (
	queryTruncateMessage = "TRUNCATE TABLE user;"
	queryInsertMessage   = "INSERT INTO messages(title, body, created_at) VALUES(?, ?, ?);"
	queryGetAllMessages  = "SELECT id, title, body, created_at FROM messages;"
)

var (
	dbConn *gorm.DB
)

// t *testing. 存在時自動執行
func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	os.Exit(m.Run())
}

func database() *gorm.DB {
	// dbDriver := os.Getenv("DBDRIVER_TEST")
	username := os.Getenv("USERNAME_TEST")
	password := os.Getenv("PASSWORD_TEST")
	host := os.Getenv("HOST_TEST")
	database := os.Getenv("DATABASE_TEST")
	port := os.Getenv("PORT_TEST")

	dbConn = sqlinit.TestInit(username, password, host, port, database)
	return dbConn
}

func refreshMessagesTable() {

	if err := dbConn.Exec(queryTruncateMessage).Error; err != nil {
		log.Fatalf("Error truncating messages table: %s", err)
	}
}

// func seedOneMessage() (domain.Message, error) {
// 	msg := domain.Message{
// 		Title:     "the title",
// 		Body:      "the body",
// 		CreatedAt: time.Now(),
// 	}
// 	stmt, err := dbConn.Prepare(queryInsertMessage)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	insertResult, createErr := stmt.Exec(msg.Title, msg.Body, msg.CreatedAt)
// 	if createErr != nil {
// 		log.Fatalf("Error creating message: %s", createErr)
// 	}
// 	msgId, err := insertResult.LastInsertId()
// 	if err != nil {
// 		log.Fatalf("Error creating message: %s", createErr)
// 	}
// 	msg.Id = msgId
// 	return msg, nil
// }

// func seedMessages() ([]domain.Message, error) {
// 	msgs := []domain.Message{
// 		{
// 			Title:     "first title",
// 			Body:      "first body",
// 			CreatedAt: time.Now(),
// 		},
// 		{
// 			Title:     "second title",
// 			Body:      "second body",
// 			CreatedAt: time.Now(),
// 		},
// 	}
// 	stmt, err := dbConn.Prepare(queryInsertMessage)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	for i, _ := range msgs {
// 		_, createErr := stmt.Exec(msgs[i].Title, msgs[i].Body, msgs[i].CreatedAt)
// 		if createErr != nil {
// 			return nil, createErr
// 		}
// 	}
// 	get_stmt, err := dbConn.Prepare(queryGetAllMessages)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	rows, err := get_stmt.Query()
// 	if err != nil {
// 		return nil,  err
// 	}
// 	defer rows.Close()

// 	results := make([]domain.Message, 0)

// 	for rows.Next() {
// 		var msg domain.Message
// 		if getError := rows.Scan(&msg.Id, &msg.Title, &msg.Body, &msg.CreatedAt); getError != nil {
// 			return nil, err
// 		}
// 		results = append(results, msg)
// 	}
// 	return results, nil
// }
