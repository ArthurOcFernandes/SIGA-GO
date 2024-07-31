package models

type SigaUser interface {
	Name() string
	SocialSecurityNumber() string
	Email() string
	Address() Address
}
