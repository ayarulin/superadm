package core

type UsersRepo interface {
  GetUser(string) (string, error)
}
