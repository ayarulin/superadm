package core

type CommandIO struct {
	UsersRepo
	YclientsSignValidator
  JwtCoder
}

type YclientsSignValidator interface {
	ValidateSign(data, sign string) bool
}

type JwtCoder interface {
	EncodeYclientsRegistration(*YclientsRegistration) (string, error)
	DecodeYclientsRegistration(string) (*YclientsRegistration, error)
}
