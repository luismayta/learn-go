package repository

type UserRepository interface {
	GetUserByID(userID int) string
}
