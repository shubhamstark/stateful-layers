package ab

import (
	userrepo "github.com/shubhamstark/stateful-layers.git/repo/userRepo"
	"github.com/shubhamstark/stateful-layers.git/service/a"
	"github.com/shubhamstark/stateful-layers.git/service/ab"
	"github.com/shubhamstark/stateful-layers.git/service/b"
)

func Build(userID string) ab.AB {
	userGetter := userrepo.NewUserRepo("user-1234")

	a := a.A{UserGetter: &userGetter}
	b := b.B{UserGetter: &userGetter}

	return ab.AB{
		A: a,
		B: b,
	}
}
