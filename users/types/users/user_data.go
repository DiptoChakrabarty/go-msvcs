package users

import (
	"github.com/DiptoChakrabarty/go-mvcs/users/models"
	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func (usr User) Save() *resterrors.RestErr {
	models.Client.DBConn.Model(&User{}).Save(&usr)
	return nil
}

func (usr User) Find(id uint64) (*User, *resterrors.RestErr) {
	return nil, nil
}
