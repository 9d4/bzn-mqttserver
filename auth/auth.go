package auth

type Authenticator struct {
}

func (a *Authenticator) Authenticate(user, password []byte) bool {
	return true
}

func (a *Authenticator) ACL(user []byte, topic string, write bool) bool {
	return true
}
