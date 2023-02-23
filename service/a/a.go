package a

import (
	"fmt"

	"github.com/shubhamstark/stateful-layers.git/models/user"
)

type IUserGetter interface {
	GetUser(returnMemoized bool) user.User
}

type A struct {
	UserGetter IUserGetter
}

func (a A) PrintUser() {

	user := a.UserGetter.GetUser(true)

	fmt.Println(user)
}
