package auth

import (
	"math/rand"
	"time"

	"gopkg.in/mgo.v2/bson"
)

/*
Session identifies a logged in user by his token provided by the client.
*/
type Session struct {
	ID        bson.ObjectId `json:"id" structs:"id" bson:"_id,omitempty"`
	UserID    bson.ObjectId `json:"userID" structs:"userID" bson:"userID"`
	Token     string        `json:"token" structs:"token" bson:"token"`
	CreatedOn time.Time     `json:"createdOn" structs:"createdOn" bson:"createdOn"`
	ExpiresOn time.Time     `json:"expiresOn" structs:"expiresOn" bson:"expiresOn"`
	Active    bool          `json:"active" structs:"active" bson:"active"`
	IPAddress string        `json:"ipAddress" structs:"ipAddress" bson:"ipAddress"`
}

var tokenRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

/*
AssignToken generates a random token and adds it to he session.
*/
func (ses *Session) AssignToken() {
	b := make([]rune, 20)
	for i := range b {
		b[i] = tokenRunes[rand.Intn(len(tokenRunes))]
	}
	ses.Token = string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
