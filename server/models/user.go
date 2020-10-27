package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

//User a
type User struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	Email     string `gorm:"size:355"`
	Password  string `gorm:"size:355"`
	CreatedAt int64  `gorm:"autoCreateTime"`       // Use unix seconds as creating time
	UpdatedAt int64  `gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
}

//HashPassword hashes user password
func (user *User) HashPassword() {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(bytes)
}

//CheckPasswordHash compares hash with password
func (user *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
