package models

import (
	"fmt"
	"os"

	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors" // Sqlite driver based on GGO
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserModel interface {
	Save(usr User) *resterrors.RestErr
	Find(id uint64) User
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
	db_connection_string := getEnv("DB_CONNECTION_STRING", "./db/movie.db")
	fmt.Println(db_connection_string)
	return gorm.Open(sqlite.Open(db_connection_string), &gorm.Config{})
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
	fmt.Println("This is Model", usr)
	fmt.Println(usr.FirstName, usr.LastName, usr.Email, usr.Id)
	err := db.DBConn.Model(&User{}).Create(&usr)
	if err != nil {
		fmt.Println(err.Error)
		return resterrors.BadRequestError("Unable to save error")
	}
	return nil
}

func (db *Model) Find(id uint64) User {
	fmt.Println("This is Model find", id)
	var usr User
	db.DBConn.Model(&User{}).Set("gorm:auto_preload", true).Find(&usr, id)
	fmt.Println(usr)
	return usr
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
