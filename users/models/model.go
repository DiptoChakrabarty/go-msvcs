package models

import (
	"fmt"
	"os"

	"github.com/DiptoChakrabarty/go-mvcs/logger"
	"github.com/DiptoChakrabarty/go-mvcs/resterrors" // Sqlite driver based on GGO
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserModel interface {
	Save(usr User) (*User, resterrors.RestErr)
	Find(id uint64) (*User, resterrors.RestErr)
	Update(usr User) resterrors.RestErr
	Delete(id uint64) resterrors.RestErr
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
		panic("unable to connect to DB")
	}
	db.AutoMigrate(&User{})
	return &Model{
		DBConn: db,
	}
}

func (db *Model) Save(usr User) (*User, resterrors.RestErr) {
	result := db.DBConn.Model(&User{}).Create(&usr)
	fmt.Println(result.RowsAffected, usr.ID)
	if result.Error != nil {
		logger.Error("Failed to save User", result.Error)
		return nil, resterrors.BadRequestError("unable to save error")
	}
	logger.Info("User Saved to DataBase")
	return &usr, nil
}

func (db *Model) Find(id uint64) (*User, resterrors.RestErr) {
	var usr User
	result := db.DBConn.Model(&User{}).Set("gorm:auto_preload", true).Find(&usr, id)
	if result.Error != nil {
		logger.Error("Failed to find User", result.Error)
		return nil, resterrors.NotFound("unable to find user")
	}
	logger.Info("User Retreived Successfully")
	return &usr, nil
}

func (db *Model) Update(usr User) resterrors.RestErr {
	result := db.DBConn.Model(&User{}).Save(&usr)
	if result.Error != nil {
		logger.Error("Failed to update User", result.Error)
		return resterrors.BadRequestError("unable to update user")
	}
	logger.Info("User Updated Successfully")
	return nil
}

func (db *Model) Delete(id uint64) resterrors.RestErr {
	result := db.DBConn.Model(&User{}).Delete(&User{}, id)
	if result.Error != nil {
		logger.Error("Failed to delete User", result.Error)
		return resterrors.BadRequestError("unable to delete user")
	}
	logger.Info("user deleted Successfully")
	return nil
}
