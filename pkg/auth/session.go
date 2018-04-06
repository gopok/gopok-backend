package auth

import (
	"math/rand"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type session struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Token     string        `json:"token" bson:"token"`
	UserID    string        `json:"userID" bson:"userID"`
	CreatedOn time.Time     `json:"createdOn" bson:"createdOn"`
	ExpiresOn time.Time     `json:"expiresOn" bson:"expiresOn"`
	Active    bool          `json:"active" bson:"active"`
	IPAddress string        `json:"ipAddress" bson:"ipAddress"`
}

var tokenRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func (ses *session) AssignToken() {
	b := make([]rune, 20)
	for i := range b {
		b[i] = tokenRunes[rand.Intn(len(tokenRunes))]
	}
	ses.Token = string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
