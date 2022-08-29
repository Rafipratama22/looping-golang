package service

type User struct {
	Nama string
}

type UserService struct {
	ListUser []User
}

type UserIface interface {
	Register(u *User) string
	Get() []User
}

func NewUserService(user []User) UserIface {
	return &UserService{
		ListUser: user,
	}
}

func (u *UserService) Register(user *User) string {
	u.ListUser = append(u.ListUser, *user)
	return user.Nama + "berhasil didaftarkan"
}

func (u *UserService) Get() []User {
	return u.ListUser
}