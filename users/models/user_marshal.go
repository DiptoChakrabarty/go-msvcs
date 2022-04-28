package models

import "encoding/json"

type PublicUser struct {
	ID        uint64 `gorm:"primaryKey;auto_increment" json:"id"`
	FirstName string `json:"first" binding:"required" gorm:"type:varchar(20)"`
	LastName  string `json:"last" binding:"required" gorm:"type:varchar(20)"`
}

type PrivateUser struct {
	ID        uint64 `gorm:"primaryKey;auto_increment" json:"id"`
	FirstName string `json:"first" binding:"required" gorm:"type:varchar(20)"`
	LastName  string `json:"last" binding:"required" gorm:"type:varchar(20)"`
	Email     string `json:"email" binding:"required" gorm:"type:varchar(20)"`
	Created   string `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
	}

	userjson, _ := json.Marshal(user)
	var pvtUsr PrivateUser
	json.Unmarshal(userjson, &pvtUsr)
	return pvtUsr
}
