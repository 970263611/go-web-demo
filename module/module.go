package module

type TokenUser struct {
	UserCode string
	Username string
	IsAdmin bool
	ExpirationTime int64
}
