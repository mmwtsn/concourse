package present

import (
	"github.com/concourse/concourse/v7/atc"
	"github.com/concourse/concourse/v7/atc/db"
)

func User(user db.User) atc.User {
	return atc.User{
		ID:        user.ID(),
		Username:  user.Name(),
		Connector: user.Connector(),
		LastLogin: user.LastLogin().Unix(),
	}
}
