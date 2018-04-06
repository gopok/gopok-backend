package auth

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type user struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"-" bson:"password"`
}

func (u *user) hashPassword(password string) error {

	rawHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return errors.Wrap(err, "failed to hash user password:")
	}
	u.Password = string(rawHash)
	return nil
}
