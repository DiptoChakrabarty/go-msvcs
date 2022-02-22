package models

import (
	"fmt"
	"os"

	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Client = NewModelDB()

type Model struct {
	DBConn *gorm.DB
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func getDB() (db *gorm.DB, err error) {
	db_type := getEnv("DB_TYPE", "sqlite3")
	db_connection_string := getEnv("DB_CONNECTION_STRING", "./db/movie.db")
	return gorm.Open(db_type, db_connection_string)
}

func NewModelDB() *Model {
	db, err := getDB()
	if err != nil {
		fmt.Println(err.Error())
		panic("Unable to connect to DB")
	}
	db.AutoMigrate(&User{})
	return &Model{
		DBConn: db,
	}
}

func (db Model) Save(usr User) *resterrors.RestErr {
	db.DBConn.Model(&User{}).Save(&usr)
	return nil
}

func (db Model) Find(id uint64) (*User, *resterrors.RestErr) {
	return nil, nil
}
