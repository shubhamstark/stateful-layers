package b

import (
	"fmt"

	"github.com/shubhamstark/stateful-layers.git/models/user"
)

type IUserGetter interface {
	GetUser(returnMemoized bool) user.User
}

type B struct {
	UserGetter IUserGetter
}

func (b B) PrintUser() {
	user := b.UserGetter.GetUser(true)

	fmt.Println(user)
}
