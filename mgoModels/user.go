package models

import (
	//"log"
	"gopkg.in/mgo.v2/bson"
	db "github.com/dangyanglim/go_cnode/database"
)

//"log"

type User struct {
	Name         string    `json:"name"`
	Loginname string `json:"loginname" `
	Pass    string `json:"pass" `
	Email string `json:"email" `
	Avatar string `json:"avatar" `
	AccessToken    string `json:"accessToken"`
	Score	uint `json:"score"`
}
type UserModel struct{}
func (p *UserModel) GetUserByName(name string)(user User,err error){
	mgodb := db.MogSession.DB("egg_cnode")
	err=mgodb.C("users").Find(bson.M{"name":name}).One(&user);
	return user,err 	
}


