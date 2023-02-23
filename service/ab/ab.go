package ab

type PrintableUser interface {
	PrintUser()
}

type AB struct {
	A PrintableUser
	B PrintableUser
}

func (t AB) Demo() {
	t.A.PrintUser()
	t.B.PrintUser()
}
