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

type UserModel interface {
	Save(usr User) *resterrors.RestErr
	Find(id uint64) (*User, *resterrors.RestErr)
	Update(usr User) *resterrors.RestErr
	Delete(id uint64) *resterrors.RestErr
}

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

func NewModelDB() UserModel {
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

func (db *Model) Save(usr User) *resterrors.RestErr {
	err := db.DBConn.Model(&User{}).Save(&usr)
	if err != nil {
		return resterrors.BadRequestError("unable to save user")
	}
	fmt.Println(usr)
	return nil
}

func (db *Model) Find(id uint64) (*User, *resterrors.RestErr) {
	var usr User
	err := db.DBConn.Model(&User{}).Set("gorm:auto_preload", true).Find(&usr, id)
	if err != nil {
		return nil, resterrors.BadRequestError("user not found")
	}
	fmt.Println(usr)
	return &usr, nil
}

func (db *Model) Update(usr User) *resterrors.RestErr {
	err := db.DBConn.Model(&User{}).Save(&usr)
	if err != nil {
		return resterrors.BadRequestError("unable to update user")
	}
	return nil
}

func (db *Model) Delete(id uint64) *resterrors.RestErr {
	err := db.DBConn.Model(&User{}).Delete(&User{}, id)
	if err != nil {
		return resterrors.BadRequestError("unable to delete user")
	}
	return nil
}
